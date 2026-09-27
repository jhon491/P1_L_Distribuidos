package dto

type MetadataRuidoBlancoDTO struct {
	Id            int    `json:"id"`
	Titulo        string `json:"titulo"`
	TipoSonido    string `json:"tipoSonido"`
	FuenteAudio   string `json:"fuenteAudio"`
	UsoSugerido   string `json:"usoSugerido"`
	Proveedor     string `json:"proveedor"`
	DuracionBucle int    `json:"duracionBucle"`
	Frecuencia    string `json:"frecuencia"`
}

type RespuestaMetadataRuidoBlancoDTO struct {
	ObjRuidoBlanco MetadataRuidoBlancoDTO `json:"objRuidoBlanco"`
	Codigo         int                    `json:"codigo"`
	Mensaje        string                 `json:"mensaje"`
}
