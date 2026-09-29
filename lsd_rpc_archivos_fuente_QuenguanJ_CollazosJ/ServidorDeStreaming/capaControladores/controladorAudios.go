package capacontroladores

import (
	"log"

	fachada "streaming/capaFachada"
	pb "streaming/serviciosAudio"
)

// ControladorServidor implementa el servicio gRPC AudioService.
type ControladorServidor struct {
	pb.UnimplementedAudioServiceServer
}

// AudioStream es la implementación del procedimiento remoto.
func (s *ControladorServidor) AudioStream(req *pb.AudioRequest, stream pb.AudioService_AudioStreamServer) error {
	// eco : llamada a procedimiento remoto gRPC
	log.Printf("[gRPC] AudioStream invocado: id=%d titulo=%q tipo=%q",
		req.IdAudio, req.TituloAudio, req.TipoAudio)

	// Delegar la lógica de lectura y envío de chunks (fragmentos) a la fachada.
	return fachada.EnviarFragmentosAudio(req, stream)
}
