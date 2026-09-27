package dto

// ResumenAudioDTO es la versión resumida de un audio, usada para listarlo
// dentro de un tipo antes de pedir el detalle completo.
type ResumenAudioDTO struct {
	Id     int    `json:"id"`
	Titulo string `json:"titulo"`
}
