package dto

// TipoAudioDTO representa un tipo de audio del catálogo
type TipoAudioDTO struct {
	IdTipo     int    `json:"idTipo"`
	NombreTipo string `json:"nombreTipo"`
}

// RespuestaTiposDTO es la respuesta de GET /tipos
type RespuestaTiposDTO struct {
	Codigo  int            `json:"codigo"`
	Mensaje string         `json:"mensaje"`
	Tipos   []TipoAudioDTO `json:"tipos"`
}

// ResumenAudioDTO representa un audio dentro de la lista de un tipo
type ResumenAudioDTO struct {
	Id     int    `json:"id"`
	Titulo string `json:"titulo"`
}

// RespuestaAudiosDTO es la respuesta de GET /tipos/:idTipo/audios
type RespuestaAudiosDTO struct {
	Codigo  int               `json:"codigo"`
	Mensaje string            `json:"mensaje"`
	Audios  []ResumenAudioDTO `json:"audios"`
}

// RespuestaDetalleDTO es la respuesta de GET /tipos/:idTipo/audios/:id
// ObjAudio se deja como mapa genérico porque cada tipo de audio tiene
// campos de metadata distintos (música, podcast, audiolibro, ruido blanco).
type RespuestaDetalleDTO struct {
	Codigo   int                    `json:"codigo"`
	Mensaje  string                 `json:"mensaje"`
	ObjAudio map[string]interface{} `json:"objAudio"`
}
