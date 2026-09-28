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
	clienteMetadata  *capaComunicacion.ClienteMetadata
	clienteStreaming *capaComunicacion.ClienteStreaming
	// lineas recibe TODO lo que se escribe por teclado. Un único goroutine
	// lee stdin, así el menú y la reproducción nunca compiten por la entrada.
	lineas chan string
}

func NuevaFachadaCliente() *FachadaCliente {
	streaming, err := capaComunicacion.NuevoClienteStreaming()
	if err != nil {
		fmt.Println("Error iniciando cliente de streaming:", err)
	}
	fachada := &FachadaCliente{
		clienteMetadata:  capaComunicacion.NuevoClienteMetadata(),
		clienteStreaming: streaming,
		lineas:           make(chan string),
	}
	go fachada.leerTeclado()
	return fachada
}

// leerTeclado es el único lector de stdin: envía cada línea al canal y lo
// cierra si la entrada termina (Ctrl+D o stdin cerrado).
func (this *FachadaCliente) leerTeclado() {
	lector := bufio.NewReader(os.Stdin)
	for {
		texto, err := lector.ReadString('\n')
		if texto != "" || err == nil {
			this.lineas <- strings.TrimSpace(texto)
		}
		if err != nil {
			close(this.lineas)
			return
		}
	}
}

// EjecutarMenuPrincipal es el punto de entrada del ciclo de menús
func (this *FachadaCliente) EjecutarMenuPrincipal() {
	if this.clienteStreaming != nil {
		defer this.clienteStreaming.Cerrar()
	}
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
		this.menuDetalleAudio(idTipo, nombreTipo, audioSeleccionado.Id, audioSeleccionado.Titulo)
	}
}

// menuDetalleAudio muestra los metadatos completos y permite reproducir
func (this *FachadaCliente) menuDetalleAudio(idTipo int, nombreTipo string, id int, titulo string) {
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
			this.reproducirAudio(id, titulo, nombreTipo)
		case "2":
			return
		default:
			fmt.Println("Opción inválida.")
		}
	}
}

// reproducirAudio solicita el streaming por gRPC y reproduce el audio.
// Mientras suena, el usuario puede presionar Enter para detenerlo.
func (this *FachadaCliente) reproducirAudio(id int, titulo string, tipo string) {
	if this.clienteStreaming == nil {
		fmt.Println("El cliente de streaming no está disponible.")
		return
	}

	fmt.Printf("\n--- Reproduciendo: %s ---\n", titulo)
	fmt.Println("(presiona Enter para detener)")

	detener := make(chan struct{})
	terminado := make(chan struct{})

	go func() {
		defer close(terminado)
		if err := this.clienteStreaming.ReproducirAudio(id, titulo, tipo, detener); err != nil {
			fmt.Println("\nError reproduciendo audio:", err)
		}
	}()

	// Se espera lo primero que ocurra: que termine la reproducción o que el
	// usuario escriba algo. No queda ninguna lectura pendiente al terminar,
	// por lo que la siguiente opción del menú no se pierde.
	select {
	case <-terminado:
	case <-this.lineas:
		close(detener)
		<-terminado
	}
}

func (this *FachadaCliente) leerLinea() string {
	texto, abierto := <-this.lineas
	if !abierto {
		fmt.Println("\nEntrada cerrada. Saliendo...")
		if this.clienteStreaming != nil {
			this.clienteStreaming.Cerrar()
		}
		os.Exit(0)
	}
	return texto
}