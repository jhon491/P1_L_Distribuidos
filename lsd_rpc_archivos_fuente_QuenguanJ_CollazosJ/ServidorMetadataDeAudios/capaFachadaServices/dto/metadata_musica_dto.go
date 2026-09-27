package dto

type MetadataMusicaDTO struct {
	Id                int    `json:"id"`
	ArtistaPrincipal  string `json:"artistaPrincipal"`
	Album             string `json:"album"`
	Genero            string `json:"genero"`
	Titulo            string `json:"titulo"`
	SelloDiscografico string `json:"selloDiscografico"`
	AnioLanzamiento   int    `json:"anioLanzamiento"`
}

type RespuestaMetadataMusicaDTO struct {
	ObjMusica MetadataMusicaDTO `json:"objMusica"`
	Codigo    int               `json:"codigo"`
	Mensaje   string            `json:"mensaje"`
}
