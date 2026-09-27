package service

import (
	"microservicio/capaAccesoADatos/repository"
	"microservicio/capaFachadaServices/dto"
)

// AudioService orquesta las consultas sobre los audios de cada tipo.
type AudioService struct {
	repository *repository.AudioRepository
}

func NewAudioService(repository *repository.AudioRepository) *AudioService {
	return &AudioService{repository: repository}
}

// ListarPorTipo devuelve el resumen (id + título) de los audios de un tipo.
func (this *AudioService) ListarPorTipo(idTipo int) []dto.ResumenAudioDTO {
	resumen := []dto.ResumenAudioDTO{}

	switch idTipo {
	case 1: // Música
		for _, m := range this.repository.ListarMusica() {
			resumen = append(resumen, dto.ResumenAudioDTO{Id: m.GetId(), Titulo: m.GetTitulo()})
		}
	case 2: // Podcast
		for _, p := range this.repository.ListarPodcasts() {
			resumen = append(resumen, dto.ResumenAudioDTO{Id: p.GetId(), Titulo: p.GetTituloEpisodio()})
		}
	case 3: // Audiolibro
		for _, a := range this.repository.ListarAudiolibros() {
			resumen = append(resumen, dto.ResumenAudioDTO{Id: a.GetId(), Titulo: a.GetTitulo()})
		}
	case 4: // Ruido Blanco
		for _, rb := range this.repository.ListarRuidosBlancos() {
			resumen = append(resumen, dto.ResumenAudioDTO{Id: rb.GetId(), Titulo: rb.GetTitulo()})
		}
	}

	return resumen
}

// ConsultarDetalle busca el audio específico dentro del tipo indicado.
func (this *AudioService) ConsultarDetalle(idTipo int, id int) (interface{}, bool) {
	switch idTipo {
	case 1:
		m, ok := this.repository.BuscarMusicaPorId(id)
		if !ok {
			return nil, false
		}
		return dto.MetadataMusicaDTO{
			Id: m.GetId(), ArtistaPrincipal: m.GetArtistaPrincipal(), Album: m.GetAlbum(),
			Genero: m.GetGenero(), Titulo: m.GetTitulo(),
			SelloDiscografico: m.GetSelloDiscografico(), AnioLanzamiento: m.GetAnioLanzamiento(),
		}, true

	case 2:
		p, ok := this.repository.BuscarPodcastPorId(id)
		if !ok {
			return nil, false
		}
		return dto.MetadataPodcastDTO{
			Id: p.GetId(), NombrePodcast: p.GetNombrePodcast(), TituloEpisodio: p.GetTituloEpisodio(),
			Anfitrion: p.GetAnfitrion(), EpisodioNumero: p.GetEpisodioNumero(),
			Notas: p.GetNotas(), Clasificacion: p.GetClasificacion(),
		}, true

	case 3:
		a, ok := this.repository.BuscarAudiolibroPorId(id)
		if !ok {
			return nil, false
		}
		return dto.MetadataAudiolibroDTO{
			Id: a.GetId(), Titulo: a.GetTitulo(), Autor: a.GetAutor(), Narrador: a.GetNarrador(),
			Editorial: a.GetEditorial(), ISBN: a.GetIsbn(), Capitulo: a.GetCapitulo(),
		}, true

	case 4:
		rb, ok := this.repository.BuscarRuidoBlancoPorId(id)
		if !ok {
			return nil, false
		}
		return dto.MetadataRuidoBlancoDTO{
			Id: rb.GetId(), Titulo: rb.GetTitulo(), TipoSonido: rb.GetTipoSonido(),
			FuenteAudio: rb.GetFuenteAudio(), UsoSugerido: rb.GetUsoSugerido(),
			Proveedor: rb.GetProveedor(), DuracionBucle: rb.GetDuracionBucle(),
			Frecuencia: rb.GetFrecuencia(),
		}, true
	}

	return nil, false
}
