package capaFachadaServices

import (
	"bufio"
	"cliente/capaComunicacion"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type FachadaCliente struct {
	clienteMetadata *capaComunicacion.ClienteMetadata
	lector          *bufio.Reader
}

func NuevaFachadaCliente() *FachadaCliente {
	return &FachadaCliente{
		clienteMetadata: capaComunicacion.NuevoClienteMetadata(),
		lector:          bufio.NewReader(os.Stdin),
	}
}

// EjecutarMenuPrincipal es el punto de entrada del ciclo de menús
func (this *FachadaCliente) EjecutarMenuPrincipal() {
	for {
		fmt.Println("\n--- Spotify ---")
		fmt.Println("1. Ver tipos de audio")
		fmt.Println("2. Salir")
		fmt.Print("Selecciona una opción: ")

		opcion := this.leerLinea()
		switch opcion {
		case "1":
			this.menuTipos()
		case "2":
			fmt.Println("Saliendo...")
			return
		default:
			fmt.Println("Opción inválida.")
		}
	}
}

// menuTipos lista los tipos y permite entrar a uno
func (this *FachadaCliente) menuTipos() {
	respuesta, err := this.clienteMetadata.ObtenerTipos()
	if err != nil {
		fmt.Println("Error consultando tipos:", err)
		return
	}

	for {
		fmt.Println("\n--- Tipos de audio ---")
		for i, tipo := range respuesta.Tipos {
			fmt.Printf("%d. %s\n", i+1, tipo.NombreTipo)
		}
		fmt.Printf("%d. Atrás\n", len(respuesta.Tipos)+1)
		fmt.Print("Selecciona una opción: ")

		opcion := this.leerLinea()
		indice, err := strconv.Atoi(opcion)

		if err == nil && indice == len(respuesta.Tipos)+1 {
			return // Atrás
		}
		if err != nil || indice < 1 || indice > len(respuesta.Tipos) {
			fmt.Println("Opción inválida.")
			continue
		}

		tipoSeleccionado := respuesta.Tipos[indice-1]
		this.menuAudiosPorTipo(tipoSeleccionado.IdTipo, tipoSeleccionado.NombreTipo)
	}
}

// menuAudiosPorTipo lista los audios de un tipo y permite ver el detalle de uno
func (this *FachadaCliente) menuAudiosPorTipo(idTipo int, nombreTipo string) {
	respuesta, err := this.clienteMetadata.ObtenerAudiosPorTipo(idTipo)
	if err != nil {
		fmt.Println("Error consultando audios:", err)
		return
	}

	for {
		fmt.Printf("\n--- Tipo: %s ---\n", nombreTipo)
		for i, audio := range respuesta.Audios {
			fmt.Printf("%d. %s\n", i+1, audio.Titulo)
		}
		fmt.Printf("%d. Atrás\n", len(respuesta.Audios)+1)
		fmt.Print("Selecciona una opción: ")

		opcion := this.leerLinea()
		indice, err := strconv.Atoi(opcion)

		if err == nil && indice == len(respuesta.Audios)+1 {
			return // Atrás
		}
		if err != nil || indice < 1 || indice > len(respuesta.Audios) {
			fmt.Println("Opción inválida.")
			continue
		}

		audioSeleccionado := respuesta.Audios[indice-1]
		this.menuDetalleAudio(idTipo, audioSeleccionado.Id)
	}
}

// menuDetalleAudio muestra los metadatos completos y permite reproducir
func (this *FachadaCliente) menuDetalleAudio(idTipo int, id int) {
	respuesta, err := this.clienteMetadata.ObtenerDetalleAudio(idTipo, id)
	if err != nil {
		fmt.Println("Error consultando detalle:", err)
		return
	}

	for {
		fmt.Println("\n--- Detalle del audio ---")
		for campo, valor := range respuesta.ObjAudio {
			fmt.Printf("%s: %v\n", campo, valor)
		}
		fmt.Println("1. Reproducir")
		fmt.Println("2. Atrás")
		fmt.Print("Selecciona una opción: ")

		opcion := this.leerLinea()
		switch opcion {
		case "1":
			this.reproducirAudio(id)
		case "2":
			return
		default:
			fmt.Println("Opción inválida.")
		}
	}
}

// reproducirAudio es el punto donde más adelante se conecta el streaming por gRPC
func (this *FachadaCliente) reproducirAudio(id int) {
	fmt.Println("\n--- Reproduciendo audio ---")
	fmt.Println("(Streaming por gRPC pendiente de conectar aquí)")
	fmt.Println("1. Salir")
	this.leerLinea()
}

func (this *FachadaCliente) leerLinea() string {
	texto, _ := this.lector.ReadString('\n')
	return strings.TrimSpace(texto)
}
