package fachada

import (
	"fmt"

	capaaccesoadatos "estadisticas/capaAccesoADatos"
	dtos "estadisticas/capaFachadaServices/DTOs"
)

type FachadaEstadisticas struct {
	repo *capaaccesoadatos.RepositorioEstadisticas
}

func NuevaFachadaEstadisticas() *FachadaEstadisticas {
	fmt.Println("Inicializando fachada de estadísticas...")
	return &FachadaEstadisticas{repo: capaaccesoadatos.GetRepositorioEstadisticas()}
}

// RegistrarReproduccion guarda el dato y muestra la confirmación por pantalla
func (f *FachadaEstadisticas) RegistrarReproduccion(msg dtos.ReproduccionAudioDTO) {
	entidad := capaaccesoadatos.EstadisticaReproduccion{
		TituloAudio: msg.TituloAudio,
		TipoAudio:   msg.TipoAudio,
		FechaHora:   msg.FechaHora,
	}

	f.repo.GuardarReproduccion(entidad)
	total := f.repo.TotalReproducciones()

	fmt.Printf("Estadística registrada -> Título: %s | Tipo: %s | Fecha: %s | Total acumulado: %d\n",
		entidad.TituloAudio, entidad.TipoAudio, entidad.FechaHora, total)
}
