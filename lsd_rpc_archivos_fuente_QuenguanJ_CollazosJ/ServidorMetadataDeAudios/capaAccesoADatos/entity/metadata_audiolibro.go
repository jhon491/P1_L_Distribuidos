package entity

// MetadataAudioLibro representa la información de un audio almacenada en el
// repositorio. Sus campos son privados (encapsulados) y se acceden mediante
// los métodos Get/Set correspondientes.

type MetadataAudioLibro struct {
	id        int
	titulo    string
	autor     string
	narrador  string
	editorial string
	isbn      string
	capitulo  int
}

func (this *MetadataAudioLibro) GetId() int {
	return this.id
}

func (this *MetadataAudioLibro) SetId(id int) {
	this.id = id
}

func (this *MetadataAudioLibro) GetTitulo() string {

	return this.titulo
}

func (this *MetadataAudioLibro) SetTitulo(titulo string) {

	this.titulo = titulo
}

func (this *MetadataAudioLibro) GetAutor() string {

	return this.autor
}

func (this *MetadataAudioLibro) SetAutor(autor string) {

	this.autor = autor
}

func (this *MetadataAudioLibro) GetNarrador() string {

	return this.narrador
}

func (this *MetadataAudioLibro) SetNarrador(narrador string) {

	this.narrador = narrador
}

func (this *MetadataAudioLibro) GetEditorial() string {

	return this.editorial
}

func (this *MetadataAudioLibro) SetEditorial(editorial string) {

	this.editorial = editorial
}

func (this *MetadataAudioLibro) GetIsbn() string {

	return this.isbn
}

func (this *MetadataAudioLibro) SetIsbn(isbn string) {

	this.isbn = isbn
}

func (this *MetadataAudioLibro) GetCapitulo() int {

	return this.capitulo
}

func (this *MetadataAudioLibro) SetCapitulo(capitulo int) {

	this.capitulo = capitulo
}
