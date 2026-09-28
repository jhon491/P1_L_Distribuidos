package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	capacontroladores "streaming/capaControladores"
	fachada "streaming/capaFachada"
	cola "streaming/componenteConexionCola"
	pb "streaming/serviciosAudio"
)

func main() {
	// Conexión con la cola (RabbitMQ) para notificar las reproducciones.
	publisher, err := cola.NewRabbitPublisher()
	if err != nil {
		fmt.Println("Error iniciando publicador:", err)
		return
	}
	defer publisher.Cerrar()
	fachada.ConfigurarPublisher(publisher)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAudioServiceServer(grpcServer, &capacontroladores.ControladorServidor{})

	fmt.Println("Servidor gRPC escuchando en :50051...")
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
