package fachada

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	capaaccesodatos "streaming/capaAccesoDatos"
	cola "streaming/componenteConexionCola"
	pb "streaming/serviciosAudio"
)

// publisher se inyecta desde main para poder notificar a la cola.
var publisher *cola.RabbitPublisher

// ConfigurarPublisher recibe el publicador de RabbitMQ creado en main.
func ConfigurarPublisher(p *cola.RabbitPublisher) {
	publisher = p
}

// abrirArchivo actúa como fachada para obtener el archivo de la canción
// usando la capa de acceso a datos.
func abrirArchivo(tipoAudio string, idAudio int32) (*os.File, error) {
	log.Printf("GetAudioFile llamado con tipo=%q id=%d", tipoAudio, idAudio)
	return capaaccesodatos.AbrirArchivo(tipoAudio, idAudio)
}

// notificarReproduccion publica de forma asíncrona (goroutine) el evento de
// reproducción en la cola.
func notificarReproduccion(titulo, tipo string) {
	if publisher == nil {
		return
	}
	go func() {
		err := publisher.PublicarReproduccion(cola.ReproduccionAudioDTO{
			TituloAudio: titulo,
			TipoAudio:   tipo,
			FechaHora:   time.Now().Format("2006-01-02 15:04:05"),
		})
		if err != nil {
			log.Println("Error notificando la reproducción:", err)
		}
	}()
}

// EnviarFragmentosAudio lee el archivo de la canción en chunks y los envía al
// stream gRPC.
func EnviarFragmentosAudio(req *pb.AudioRequest, stream pb.AudioService_AudioStreamServer) error {
	log.Printf("Enviando fragmentos de audio para id=%d", req.IdAudio)
	file, err := abrirArchivo(req.TipoAudio, req.IdAudio)
	if err != nil {
		return err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("error leyendo información del archivo: %w", err)
	}

	// Nueva reproducción: se notifica a la cola (asíncrono).
	notificarReproduccion(req.TituloAudio, req.TipoAudio)

	buf := make([]byte, 32*1024) // 32KB por chunk (fragmento)
	chunkNum := 0

	for {
		// n es la cantidad de bytes que devolvió file.Read(buf).
		// en buf se guarda el fragmento leído
		n, err := file.Read(buf)
		if err == io.EOF {
			log.Println("Canción enviada completa.")
			break
		}
		if err != nil {
			return fmt.Errorf("error leyendo archivo: %w", err)
		}

		chunkNum++

		if n > 0 {
			// se crea un objeto de tipo AudioChunk
			objChunk := &pb.AudioChunk{
				Data:       buf[:n],
				Numero:     int32(chunkNum),
				TotalBytes: info.Size(),
			}
			// con la función Send se envía el objeto de tipo AudioChunk
			if err := stream.Send(objChunk); err != nil {
				return fmt.Errorf("Error enviando chunk #%d: %w", chunkNum, err)
			}
			log.Printf("Chunk #%d enviado (%d bytes)\n", chunkNum, n)
		}
	}
	return nil
}