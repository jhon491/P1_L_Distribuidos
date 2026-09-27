package main

import (
	administrador "administrador/capaFachadaServices"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fachada := administrador.NuevaFachadaAdministrador()
	lector := bufio.NewReader(os.Stdin)

	for {
		mostrarMenu()
		opcion := leerOpcion(lector)

		switch opcion {
		case "1":
			procesarSubidaAudio(lector, fachada)
		case "2":
			fmt.Println("Saliendo...")
			return
		default:
			fmt.Println("Opción inválida, intenta de nuevo.")
		}
	}
}

func mostrarMenu() {
	fmt.Println("\n--- Administrador de Audios ---")
	fmt.Println("1. Subir un audio")
	fmt.Println("2. Salir")
	fmt.Print("Selecciona una opción: ")
}

func leerOpcion(lector *bufio.Reader) string {
	texto, _ := lector.ReadString('\n')
	return strings.TrimSpace(texto)
}

func procesarSubidaAudio(lector *bufio.Reader, fachada *administrador.FachadaAdministrador) {
	fmt.Print("Ruta del archivo mp3: ")
	ruta, _ := lector.ReadString('\n')
	ruta = strings.TrimSpace(ruta)

	fmt.Print("Id del audio: ")
	id, _ := lector.ReadString('\n')
	id = strings.TrimSpace(id)

	if err := fachada.SubirAudio(ruta, id); err != nil {
		fmt.Println("Error subiendo el audio:", err)
	}
}
