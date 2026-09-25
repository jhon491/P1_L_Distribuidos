package entity

// MetadataMusica representa la información de una canción almacenada en el
// repositorio. Sus campos son privados (encapsulados) y se acceden mediante
// los métodos Get/Set correspondientes.
type MetadataMusica struct {
	artistaPrincipal  string
	album             string
	genero            string
	titulo            string
	selloDiscografico string
	anioLanzamiento   int
}

func (this *MetadataMusica) GetArtistaPrincipal() string {
	return this.artistaPrincipal
}

func (this *MetadataMusica) SetArtistaPrincipal(artistaPrincipal string) {
	this.artistaPrincipal = artistaPrincipal
}

func (this *MetadataMusica) GetAlbum() string {
	return this.album
}

func (this *MetadataMusica) SetAlbum(album string) {
	this.album = album
}

func (this *MetadataMusica) GetGenero() string {
	return this.genero
}

func (this *MetadataMusica) SetGenero(genero string) {
	this.genero = genero
}

func (this *MetadataMusica) GetTitulo() string {
	return this.titulo
}

func (this *MetadataMusica) SetTitulo(titulo string) {
	this.titulo = titulo
}

func (this *MetadataMusica) GetSelloDiscografico() string {
	return this.selloDiscografico
}

func (this *MetadataMusica) SetSelloDiscografico(selloDiscografico string) {
	this.selloDiscografico = selloDiscografico
}

func (this *MetadataMusica) GetAnioLanzamiento() int {
	return this.anioLanzamiento
}

func (this *MetadataMusica) SetAnioLanzamiento(anioLanzamiento int) {
	this.anioLanzamiento = anioLanzamiento
}
