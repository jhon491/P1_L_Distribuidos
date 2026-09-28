package capaComunicacion

import (
	"context"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "cliente/serviciosAudio"
)

const direccionStreaming = "localhost:50051"

// ClienteStreaming encapsula la llamada gRPC a ServidorDeStreaming
// y la reproducción en vivo del audio recibido.
type ClienteStreaming struct {
	conexion *grpc.ClientConn
	cliente  pb.AudioServiceClient
	// el speaker de beep solo puede inicializarse una vez por proceso;
	// se recuerda la frecuencia con la que se inicializó.
	speakerListo bool
	frecuencia   beep.SampleRate
}

func NuevoClienteStreaming() (*ClienteStreaming, error) {
	conexion, err := grpc.NewClient(direccionStreaming,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("error creando conexión gRPC con ServidorDeStreaming: %v", err)
	}
	return &ClienteStreaming{
		conexion: conexion,
		cliente:  pb.NewAudioServiceClient(conexion),
	}, nil
}

func (c *ClienteStreaming) Cerrar() {
	c.conexion.Close()
}

func (c *ClienteStreaming) ReproducirAudio(idAudio int, titulo, tipo string, detener <-chan struct{}) error {
	// El contexto se cancela al detener, lo que aborta el stream gRPC.
	ctx, cancelar := context.WithCancel(context.Background())
	defer cancelar()

	fmt.Printf("Solicitando streaming del audio %d por gRPC...\n", idAudio)
	stream, err := c.cliente.AudioStream(ctx, &pb.AudioRequest{
		IdAudio:     int32(idAudio),
		TituloAudio: titulo,
		TipoAudio:   tipo,
	})
	if err != nil {
		return fmt.Errorf("error invocando AudioStream: %v", err)
	}

	fmt.Println("Recibiendo y reproduciendo audio en vivo...")
	fin := make(chan struct{})
	defer close(fin) // libera a la goroutine vigilante en cualquier salida
	reader, writer := io.Pipe()
	canalSincronizacion := make(chan struct{})
	errDecodificacion := make(chan error, 1)
	var una sync.Once
	finalizar := func() { una.Do(func() { close(canalSincronizacion) }) }

	// Goroutine de decodificación y reproducción (lee bytes desde reader).
	go c.decodificarReproducir(reader, canalSincronizacion, finalizar, errDecodificacion)

	// Si el usuario detiene la reproducción: se cancela el stream y se
	// cierra el pipe para desbloquear al decodificador.
	go func() {
		select {
		case <-detener:
			cancelar()
			writer.CloseWithError(io.ErrClosedPipe)
			speaker.Clear()
			finalizar() // libera a la goroutine del decodificador
		case <-fin:
		}
	}()

	// Hilo principal: recibe los fragmentos y los escribe en el pipe.
	if errStream := c.recibirAudio(stream, writer); errStream != nil {
		select {
		case <-detener:
			// cancelado por el usuario: no es un error
		default:
			return fmt.Errorf("error en el streaming: %v", errStream)
		}
	}

	// Espera a que termine la reproducción o a que el usuario la detenga.
	select {
	case err := <-errDecodificacion:
		return err
	case <-canalSincronizacion:
	}
	select {
	case <-detener:
		fmt.Println("Reproducción detenida.")
	default:
		fmt.Println("Reproducción finalizada.")
	}
	return nil
}

// recibirAudio recibe cada fragmento del servidor y lo pasa al reproductor
// escribiéndolo en el pipe. Al terminar el stream cierra el writer (EOF).
func (c *ClienteStreaming) recibirAudio(stream pb.AudioService_AudioStreamClient, writer *io.PipeWriter) error {
	noFragmento := 0
	for {
		fragmento, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("\nAudio recibido completo.")
			writer.Close()
			return nil
		}
		if err != nil {
			writer.CloseWithError(err)
			return err
		}
		noFragmento++
		fmt.Printf("\rFragmento #%d recibido (%d bytes) reproduciendo ...", noFragmento, len(fragmento.Data))

		if _, err := writer.Write(fragmento.Data); err != nil {
			// el decodificador ya no lee (reproducción detenida o error)
			return nil
		}
	}
}

// decodificarReproducir toma los bytes mp3 que llegan por el pipe y los
// convierte en sonido real por el altavoz.
func (c *ClienteStreaming) decodificarReproducir(reader *io.PipeReader, canalSincronizacion chan struct{}, finalizar func(), errCh chan<- error) {
	streamer, formato, err := mp3.Decode(io.NopCloser(reader))
	if err != nil {
		reader.CloseWithError(err) // desbloquea a recibirAudio si estaba escribiendo
		errCh <- fmt.Errorf("error decodificando MP3: %v", err)
		return
	}
	defer streamer.Close()

	// Configura la salida de audio (solo la primera vez en el proceso).
	if !c.speakerListo {
		if err := speaker.Init(formato.SampleRate, formato.SampleRate.N(time.Second/2)); err != nil {
			reader.CloseWithError(err)
			errCh <- fmt.Errorf("error inicializando el dispositivo de audio: %v", err)
			return
		}
		c.speakerListo = true
		c.frecuencia = formato.SampleRate
	}

	// Si este audio tiene otra frecuencia de muestreo que la usada al
	// inicializar el speaker, se remuestrea para que no suene acelerado/lento.
	var fuente beep.Streamer = streamer
	if formato.SampleRate != c.frecuencia {
		fuente = beep.Resample(4, formato.SampleRate, c.frecuencia, streamer)
	}

	// Comienza a reproducir; el callback avisa al terminar.
	speaker.Play(beep.Seq(fuente, beep.Callback(finalizar)))

	// Se mantiene viva la goroutine hasta que termine la reproducción.
	<-canalSincronizacion
	reader.Close()
}