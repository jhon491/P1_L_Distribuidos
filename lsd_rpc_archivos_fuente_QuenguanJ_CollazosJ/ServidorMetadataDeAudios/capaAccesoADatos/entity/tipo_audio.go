package entity

// TipoAudio representa el catálogo de tipos de audio soportados por el servidor.
type TipoAudio struct {
	idTipo     int
	nombreTipo string
}

func (this *TipoAudio) GetIdTipo() int {
	return this.idTipo
}

func (this *TipoAudio) SetIdTipo(idTipo int) {
	this.idTipo = idTipo
}

func (this *TipoAudio) GetNombreTipo() string {
	return this.nombreTipo
}

func (this *TipoAudio) SetNombreTipo(nombreTipo string) {
	this.nombreTipo = nombreTipo
}
