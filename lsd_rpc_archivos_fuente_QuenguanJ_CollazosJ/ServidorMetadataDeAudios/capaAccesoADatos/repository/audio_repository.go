package repository

import "microservicio/capaAccesoADatos/entity"

// AudioRepository mantiene en memoria los audios de los 4 tipos soportados.
type AudioRepository struct {
	musicas       []entity.MetadataMusica
	podcasts      []entity.MetadataPotcast
	audiolibros   []entity.MetadataAudioLibro
	ruidosBlancos []entity.MetadataRuidoblanco
}

func NewAudioRepository() *AudioRepository {
	r := &AudioRepository{}
	r.CargarAudios()
	return r
}

// CargarAudios precarga al menos 2 audios de cada tipo (exigido por el PDF).
func (r *AudioRepository) CargarAudios() {
	var m1, m2 entity.MetadataMusica
	m1.SetId(1)
	m1.SetTitulo("Bohemian Rhapsody")
	m1.SetArtistaPrincipal("Queen")
	m1.SetAlbum("A Night at the Opera")
	m1.SetGenero("Rock")
	m1.SetSelloDiscografico("EMI")
	m1.SetAnioLanzamiento(1975)

	m2.SetId(2)
	m2.SetTitulo("Blinding Lights")
	m2.SetArtistaPrincipal("The Weeknd")
	m2.SetAlbum("After Hours")
	m2.SetGenero("Pop")
	m2.SetSelloDiscografico("XO/Republic")
	m2.SetAnioLanzamiento(2020)

	r.musicas = []entity.MetadataMusica{m1, m2}

	var p1, p2 entity.MetadataPotcast
	p1.SetId(1)
	p1.SetNombrePodcast("Radio Ambulante")
	p1.SetTituloEpisodio("La frontera")
	p1.SetAnfitrion("Daniel Alarcón")
	p1.SetEpisodioNumero(1)
	p1.SetNotas("Historias de Latinoamérica")
	p1.SetClasificacion("Para toda la familia")

	p2.SetId(2)
	p2.SetNombrePodcast("Cuonda Original")
	p2.SetTituloEpisodio("El caso")
	p2.SetAnfitrion("Equipo Cuonda")
	p2.SetEpisodioNumero(1)
	p2.SetNotas("Podcast de misterio")
	p2.SetClasificacion("Explícito")

	r.podcasts = []entity.MetadataPotcast{p1, p2}

	var a1, a2 entity.MetadataAudioLibro
	a1.SetId(1)
	a1.SetTitulo("Cien años de soledad")
	a1.SetAutor("Gabriel García Márquez")
	a1.SetNarrador("Gustavo Bonfigli")
	a1.SetEditorial("Penguin Random House")
	a1.SetIsbn("978-0060883287")
	a1.SetCapitulo(1)

	a2.SetId(2)
	a2.SetTitulo("Orgullo y prejuicio")
	a2.SetAutor("Jane Austen")
	a2.SetNarrador("Rosalinda Fernández")
	a2.SetEditorial("Alianza Editorial")
	a2.SetIsbn("978-8420674165")
	a2.SetCapitulo(1)

	r.audiolibros = []entity.MetadataAudioLibro{a1, a2}

	var rb1, rb2 entity.MetadataRuidoblanco
	rb1.SetId(1)
	rb1.SetTitulo("Lluvia suave")
	rb1.SetTipoSonido("Ruido Blanco")
	rb1.SetFuenteAudio("Lluvia")
	rb1.SetUsoSugerido("Dormir")
	rb1.SetProveedor("Calm")
	rb1.SetDuracionBucle(60)
	rb1.SetFrecuencia("Graves")

	rb2.SetId(2)
	rb2.SetTitulo("Bosque nocturno")
	rb2.SetTipoSonido("Ruido Marrón")
	rb2.SetFuenteAudio("Bosque")
	rb2.SetUsoSugerido("Concentración")
	rb2.SetProveedor("Noisli")
	rb2.SetDuracionBucle(45)
	rb2.SetFrecuencia("Agudos")

	r.ruidosBlancos = []entity.MetadataRuidoblanco{rb1, rb2}
}

func (r *AudioRepository) ListarMusica() []entity.MetadataMusica             { return r.musicas }
func (r *AudioRepository) ListarPodcasts() []entity.MetadataPotcast          { return r.podcasts }
func (r *AudioRepository) ListarAudiolibros() []entity.MetadataAudioLibro    { return r.audiolibros }
func (r *AudioRepository) ListarRuidosBlancos() []entity.MetadataRuidoblanco { return r.ruidosBlancos }

func (r *AudioRepository) BuscarMusicaPorId(id int) (entity.MetadataMusica, bool) {
	for _, m := range r.musicas {
		if m.GetId() == id {
			return m, true
		}
	}
	return entity.MetadataMusica{}, false
}

func (r *AudioRepository) BuscarPodcastPorId(id int) (entity.MetadataPotcast, bool) {
	for _, p := range r.podcasts {
		if p.GetId() == id {
			return p, true
		}
	}
	return entity.MetadataPotcast{}, false
}

func (r *AudioRepository) BuscarAudiolibroPorId(id int) (entity.MetadataAudioLibro, bool) {
	for _, a := range r.audiolibros {
		if a.GetId() == id {
			return a, true
		}
	}
	return entity.MetadataAudioLibro{}, false
}

func (r *AudioRepository) BuscarRuidoBlancoPorId(id int) (entity.MetadataRuidoblanco, bool) {
	for _, rb := range r.ruidosBlancos {
		if rb.GetId() == id {
			return rb, true
		}
	}
	return entity.MetadataRuidoblanco{}, false
}
