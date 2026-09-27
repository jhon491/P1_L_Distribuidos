package entity

// MetadataRuidoblanco representa la información de un ruido blanco almacenada en el
// repositorio. Sus campos son privados (encapsulados) y se acceden mediante
// los métodos Get/Set correspondientes.

type MetadataRuidoblanco struct {
	id            int
	titulo        string
	tipoSonido    string
	fuenteAudio   string
	usoSugerido   string
	proveedor     string
	duracionBucle int
	frecuencia    string
}

func (this *MetadataRuidoblanco) GetId() int {
	return this.id
}
func (this *MetadataRuidoblanco) SetId(id int) {
	this.id = id
}

func (this *MetadataRuidoblanco) GetTitulo() string {
	return this.titulo
}
func (this *MetadataRuidoblanco) SetTitulo(v string) {
	this.titulo = v
}

func (this *MetadataRuidoblanco) GetTipoSonido() string {

	return this.tipoSonido
}

func (this *MetadataRuidoblanco) SetTipoSonido(tipoSonido string) {

	this.tipoSonido = tipoSonido
}

func (this *MetadataRuidoblanco) GetFuenteAudio() string {

	return this.fuenteAudio
}

func (this *MetadataRuidoblanco) SetFuenteAudio(fuenteAudio string) {

	this.fuenteAudio = fuenteAudio
}

func (this *MetadataRuidoblanco) GetUsoSugerido() string {

	return this.usoSugerido
}

func (this *MetadataRuidoblanco) SetUsoSugerido(usoSugerido string) {

	this.usoSugerido = usoSugerido
}

func (this *MetadataRuidoblanco) GetProveedor() string {

	return this.proveedor
}

func (this *MetadataRuidoblanco) SetProveedor(proveedor string) {

	this.proveedor = proveedor
}

func (this *MetadataRuidoblanco) GetDuracionBucle() int {

	return this.duracionBucle
}

func (this *MetadataRuidoblanco) SetDuracionBucle(duracionBucle int) {

	this.duracionBucle = duracionBucle
}

func (this *MetadataRuidoblanco) GetFrecuencia() string {

	return this.frecuencia
}

func (this *MetadataRuidoblanco) SetFrecuencia(frecuencia string) {

	this.frecuencia = frecuencia
}
