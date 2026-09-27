package entity

// MetadataPotcast representa la información de un podcast almacenada en el
// repositorio. Sus campos son privados (encapsulados) y se acceden mediante
// los métodos Get/Set correspondientes.

type MetadataPotcast struct {
	id             int
	nombrePodcast  string
	tituloEpisodio string
	anfitrion      string
	episodioNumero int
	notas          string
	clasificacion  string
}

func (this *MetadataPotcast) GetId() int {
	return this.id
}
func (this *MetadataPotcast) SetId(id int) {
	this.id = id
}

func (this *MetadataPotcast) GetNombrePodcast() string {

	return this.nombrePodcast
}

func (this *MetadataPotcast) SetNombrePodcast(nombrePodcast string) {

	this.nombrePodcast = nombrePodcast
}

func (this *MetadataPotcast) GetTituloEpisodio() string {

	return this.tituloEpisodio
}

func (this *MetadataPotcast) SetTituloEpisodio(tituloEpisodio string) {

	this.tituloEpisodio = tituloEpisodio
}

func (this *MetadataPotcast) GetAnfitrion() string {

	return this.anfitrion
}

func (this *MetadataPotcast) SetAnfitrion(anfitrion string) {

	this.anfitrion = anfitrion
}

func (this *MetadataPotcast) GetEpisodioNumero() int {

	return this.episodioNumero
}

func (this *MetadataPotcast) SetEpisodioNumero(episodioNumero int) {

	this.episodioNumero = episodioNumero
}

func (this *MetadataPotcast) GetNotas() string {

	return this.notas
}

func (this *MetadataPotcast) SetNotas(notas string) {

	this.notas = notas
}

func (this *MetadataPotcast) GetClasificacion() string {

	return this.clasificacion
}

func (this *MetadataPotcast) SetClasificacion(clasificacion string) {

	this.clasificacion = clasificacion
}
