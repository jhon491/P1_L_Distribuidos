package main

import (
	"fmt"

	"estadisticas/capaFachadaServices/fachada"
	cola "estadisticas/componenteConexionCola"
)

func main() {
	f := fachada.NuevaFachadaEstadisticas()

	consumer, err := cola.NewRabbitConsumer(f)
	if err != nil {
		fmt.Println("Error iniciando consumidor:", err)
		return
	}
	defer consumer.Cerrar()

	if err := consumer.Escuchar(); err != nil {
		fmt.Println("Error escuchando la cola:", err)
	}
}
