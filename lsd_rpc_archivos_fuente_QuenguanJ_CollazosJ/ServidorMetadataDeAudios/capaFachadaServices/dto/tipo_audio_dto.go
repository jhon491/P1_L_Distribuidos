package dto

// TipoAudioDTO representa el tipo de audio que se expone por REST.
type TipoAudioDTO struct {
	IdTipo     int    `json:"idTipo"`
	NombreTipo string `json:"nombreTipo"`
}

// RespuestaTipoAudioDTO es la respuesta para consultas de tipos de audio.
type RespuestaTipoAudioDTO struct {
	ObjTipo TipoAudioDTO `json:"objTipo"`
	Codigo  int          `json:"codigo"`
	Mensaje string       `json:"mensaje"`
}
