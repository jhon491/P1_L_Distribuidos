package repository

import "microservicio/capaAccesoADatos/entity"

// TipoAudioRepository gestiona en memoria los tipos de audio disponibles.
type TipoAudioRepository struct {
	vectorTipos []entity.TipoAudio
}

func NewTipoAudioRepository() *TipoAudioRepository {
	repo := &TipoAudioRepository{}
	repo.CargarTipos()
	return repo
}

func (r *TipoAudioRepository) CargarTipos() {
	var tipo1, tipo2, tipo3, tipo4 entity.TipoAudio

	tipo1.SetIdTipo(1)
	tipo1.SetNombreTipo("Música")

	tipo2.SetIdTipo(2)
	tipo2.SetNombreTipo("Podcast")

	tipo3.SetIdTipo(3)
	tipo3.SetNombreTipo("Audiolibro")

	tipo4.SetIdTipo(4)
	tipo4.SetNombreTipo("Ruido Blanco")

	r.vectorTipos = []entity.TipoAudio{tipo1, tipo2, tipo3, tipo4}
}

func (r *TipoAudioRepository) RegistrarTipo(tipo entity.TipoAudio) {
	r.vectorTipos = append(r.vectorTipos, tipo)
}

func (r *TipoAudioRepository) BuscarPorId(id int) (entity.TipoAudio, bool) {
	for _, tipo := range r.vectorTipos {
		if tipo.GetIdTipo() == id {
			return tipo, true
		}
	}
	return entity.TipoAudio{}, false
}

func (r *TipoAudioRepository) BuscarPorNombre(nombre string) (entity.TipoAudio, bool) {
	for _, tipo := range r.vectorTipos {
		if tipo.GetNombreTipo() == nombre {
			return tipo, true
		}
	}
	return entity.TipoAudio{}, false
}

func (r *TipoAudioRepository) ListarTipos() []entity.TipoAudio {
	return r.vectorTipos
}
