package capaComunicacion

import (
	dto "cliente/capaComunicacion/DTO"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const urlBaseMetadata = "http://localhost:8080"

// ClienteMetadata encapsula todas las llamadas REST a ServidorMetadataDeAudios
type ClienteMetadata struct{}

func NuevoClienteMetadata() *ClienteMetadata {
	return &ClienteMetadata{}
}

// ObtenerTipos consulta GET /tipos
func (this *ClienteMetadata) ObtenerTipos() (dto.RespuestaTiposDTO, error) {
	fmt.Println("Consultando tipos de audio disponibles...")

	var respuesta dto.RespuestaTiposDTO
	body, err := this.realizarGet(urlBaseMetadata + "/tipos")
	if err != nil {
		return respuesta, err
	}

	if err := json.Unmarshal(body, &respuesta); err != nil {
		return respuesta, fmt.Errorf("error interpretando respuesta de tipos: %v", err)
	}
	return respuesta, nil
}

// ObtenerAudiosPorTipo consulta GET /tipos/:idTipo/audios
func (this *ClienteMetadata) ObtenerAudiosPorTipo(idTipo int) (dto.RespuestaAudiosDTO, error) {
	fmt.Printf("Consultando audios del tipo %d...\n", idTipo)

	var respuesta dto.RespuestaAudiosDTO
	url := fmt.Sprintf("%s/tipos/%d/audios", urlBaseMetadata, idTipo)
	body, err := this.realizarGet(url)
	if err != nil {
		return respuesta, err
	}

	if err := json.Unmarshal(body, &respuesta); err != nil {
		return respuesta, fmt.Errorf("error interpretando respuesta de audios: %v", err)
	}
	return respuesta, nil
}

// ObtenerDetalleAudio consulta GET /tipos/:idTipo/audios/:id
func (this *ClienteMetadata) ObtenerDetalleAudio(idTipo int, id int) (dto.RespuestaDetalleDTO, error) {
	fmt.Printf("Consultando detalle del audio %d (tipo %d)...\n", id, idTipo)

	var respuesta dto.RespuestaDetalleDTO
	url := fmt.Sprintf("%s/tipos/%d/audios/%d", urlBaseMetadata, idTipo, id)
	body, err := this.realizarGet(url)
	if err != nil {
		return respuesta, err
	}

	if err := json.Unmarshal(body, &respuesta); err != nil {
		return respuesta, fmt.Errorf("error interpretando respuesta de detalle: %v", err)
	}
	return respuesta, nil
}

// realizarGet centraliza la petición HTTP GET y la lectura del cuerpo
func (this *ClienteMetadata) realizarGet(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error conectando con ServidorMetadataDeAudios: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error leyendo respuesta: %v", err)
	}
	return body, nil
}
