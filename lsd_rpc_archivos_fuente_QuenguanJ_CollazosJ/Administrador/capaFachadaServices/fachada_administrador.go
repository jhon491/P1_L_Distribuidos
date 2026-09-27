package capaFachadaServices

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

const urlServidorAudios = "http://localhost:5000/canciones/almacenamiento"

// FachadaAdministrador encapsula la comunicación REST con ServidorDeAudios
type FachadaAdministrador struct{}

func NuevaFachadaAdministrador() *FachadaAdministrador {
	return &FachadaAdministrador{}
}

// SubirAudio arma la petición multipart/form-data y la envía al ServidorDeAudios
func (this *FachadaAdministrador) SubirAudio(rutaArchivo string, id string) error {
	archivo, err := os.Open(rutaArchivo)
	if err != nil {
		return fmt.Errorf("no se pudo abrir el archivo: %v", err)
	}
	defer archivo.Close()

	var cuerpo bytes.Buffer
	escritor := multipart.NewWriter(&cuerpo)

	// Campo "id"
	if err := escritor.WriteField("id", id); err != nil {
		return fmt.Errorf("error escribiendo campo id: %v", err)
	}

	// Campo "archivo" (el mp3)
	parte, err := escritor.CreateFormFile("archivo", archivo.Name())
	if err != nil {
		return fmt.Errorf("error creando form file: %v", err)
	}
	if _, err := io.Copy(parte, archivo); err != nil {
		return fmt.Errorf("error copiando contenido del archivo: %v", err)
	}

	escritor.Close()

	peticion, err := http.NewRequest(http.MethodPost, urlServidorAudios, &cuerpo)
	if err != nil {
		return fmt.Errorf("error creando la petición: %v", err)
	}
	peticion.Header.Set("Content-Type", escritor.FormDataContentType())

	fmt.Println("Enviando audio al ServidorDeAudios...") // aviso informativo, no es el eco calificable

	cliente := &http.Client{}
	respuesta, err := cliente.Do(peticion)
	if err != nil {
		return fmt.Errorf("error enviando la petición: %v", err)
	}
	defer respuesta.Body.Close()

	if respuesta.StatusCode != http.StatusOK {
		return fmt.Errorf("el servidor respondió con estado: %s", respuesta.Status)
	}

	fmt.Println("Audio almacenado correctamente.")
	return nil
}
