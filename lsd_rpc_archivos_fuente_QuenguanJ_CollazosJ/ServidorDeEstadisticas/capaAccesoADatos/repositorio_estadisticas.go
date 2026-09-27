package capaaccesoadatos

import "sync"

// EstadisticaReproduccion es el registro que se guarda internamente
type EstadisticaReproduccion struct {
	TituloAudio string
	TipoAudio   string
	FechaHora   string
}

type RepositorioEstadisticas struct {
	mu        sync.Mutex
	registros []EstadisticaReproduccion
}

var instancia *RepositorioEstadisticas
var once sync.Once

// GetRepositorioEstadisticas devuelve siempre la misma instancia (singleton)
func GetRepositorioEstadisticas() *RepositorioEstadisticas {
	once.Do(func() {
		instancia = &RepositorioEstadisticas{registros: []EstadisticaReproduccion{}}
	})
	return instancia
}

func (r *RepositorioEstadisticas) GuardarReproduccion(e EstadisticaReproduccion) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.registros = append(r.registros, e)
}

func (r *RepositorioEstadisticas) TotalReproducciones() int {
	r.mu.Lock()
	defer r.mu.Unlock()
	return len(r.registros)
}

func (r *RepositorioEstadisticas) ObtenerTodas() []EstadisticaReproduccion {
	r.mu.Lock()
	defer r.mu.Unlock()
	copia := make([]EstadisticaReproduccion, len(r.registros))
	copy(copia, r.registros)
	return copia
}
