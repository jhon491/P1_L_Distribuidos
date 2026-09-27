package main

import (
	"fmt"

	cola "streaming/componenteConexionCola"
)

func main() {
	publisher, err := cola.NewRabbitPublisher()
	if err != nil {
		fmt.Println("Error iniciando publicador:", err)
		return
	}
	defer publisher.Cerrar()

	publisher.PublicarReproduccion(cola.ReproduccionAudioDTO{
		TituloAudio: "Cien años de soledad",
		TipoAudio:   "Audiolibro",
		FechaHora:   "2026-09-26 10:00",
	})
}
