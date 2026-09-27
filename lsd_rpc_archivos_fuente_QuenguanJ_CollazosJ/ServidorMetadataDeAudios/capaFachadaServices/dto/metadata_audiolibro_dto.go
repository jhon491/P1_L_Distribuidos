package dto

type MetadataAudiolibroDTO struct {
	Id        int    `json:"id"`
	Titulo    string `json:"titulo"`
	Autor     string `json:"autor"`
	Narrador  string `json:"narrador"`
	Editorial string `json:"editorial"`
	ISBN      string `json:"isbn"`
	Capitulo  int    `json:"capitulo"`
}

type RespuestaMetadataAudiolibroDTO struct {
	ObjAudiolibro MetadataAudiolibroDTO `json:"objAudiolibro"`
	Codigo        int                   `json:"codigo"`
	Mensaje       string                `json:"mensaje"`
}
