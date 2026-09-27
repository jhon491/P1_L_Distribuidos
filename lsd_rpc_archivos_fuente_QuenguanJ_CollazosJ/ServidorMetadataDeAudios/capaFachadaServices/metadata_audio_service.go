package service

import (
	"microservicio/capaAccesoADatos/entity"
	"microservicio/capaAccesoADatos/repository"
	"microservicio/capaFachadaServices/dto"
)

// MetadataAudioService expone la lógica de negocio para manejar tipos de audio.
type MetadataAudioService struct {
	repository *repository.TipoAudioRepository
}

func NewMetadataAudioService(repository *repository.TipoAudioRepository) *MetadataAudioService {
	return &MetadataAudioService{repository: repository}
}

// RegistrarTipo recibe un DTO y lo guarda en el catálogo de tipos.
func (this *MetadataAudioService) RegistrarTipo(tipoDTO dto.TipoAudioDTO) {
	var tipo entity.TipoAudio
	tipo.SetIdTipo(tipoDTO.IdTipo)
	tipo.SetNombreTipo(tipoDTO.NombreTipo)

	this.repository.RegistrarTipo(tipo)
}

// ConsultarTipo busca un tipo por su id y lo convierte a DTO de respuesta.
func (this *MetadataAudioService) ConsultarTipo(id int) dto.RespuestaTipoAudioDTO {
	var respuesta dto.RespuestaTipoAudioDTO

	tipo, encontrado := this.repository.BuscarPorId(id)
	if encontrado {
		respuesta.ObjTipo = dto.TipoAudioDTO{
			IdTipo:     tipo.GetIdTipo(),
			NombreTipo: tipo.GetNombreTipo(),
		}
		respuesta.Codigo = 200
		respuesta.Mensaje = "Tipo de audio encontrado"
		return respuesta
	}

	respuesta.Codigo = 400
	respuesta.Mensaje = "El tipo de audio no se encontró"
	return respuesta
}

// ListarTipos devuelve todos los tipos registrados.
func (this *MetadataAudioService) ListarTipos() []dto.TipoAudioDTO {
	tipos := this.repository.ListarTipos()
	result := make([]dto.TipoAudioDTO, 0, len(tipos))

	for _, tipo := range tipos {
		result = append(result, dto.TipoAudioDTO{
			IdTipo:     tipo.GetIdTipo(),
			NombreTipo: tipo.GetNombreTipo(),
		})
	}

	return result
}
