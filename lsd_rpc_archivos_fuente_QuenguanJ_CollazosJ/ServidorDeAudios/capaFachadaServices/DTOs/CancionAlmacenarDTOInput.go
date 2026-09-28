package dtos

type CancionAlmacenarDTOInput struct {
	IdTipo string `json:"idTipo"` // 1=Música, 2=Podcast, 3=Audiolibro, 4=Ruido Blanco
	Id     string `json:"id"`     // id del audio dentro de su tipo
}