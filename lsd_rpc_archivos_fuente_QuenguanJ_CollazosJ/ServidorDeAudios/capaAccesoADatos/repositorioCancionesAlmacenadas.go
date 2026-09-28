package capaaccesoadatos

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

type RepositorioCanciones struct {
	mu sync.Mutex
}

var (
	instancia *RepositorioCanciones
	once      sync.Once
)

// GetRepositorioCanciones aplica el patrón Singleton
func GetRepositorioCanciones() *RepositorioCanciones {
	once.Do(func() {
		instancia = &RepositorioCanciones{}
	})
	return instancia
}

// rutaAudios es la carpeta compartida con ServidorDeStreaming (misma lógica
// que allá): "../audios", o el valor de la variable RUTA_AUDIOS.
func rutaAudios() string {
	if r := os.Getenv("RUTA_AUDIOS"); r != "" {
		return r
	}
	return filepath.Join("..", "audios")
}

// GuardarCancion guarda el MP3 como audios/<idTipo>_<id>.mp3, el mismo nombre
// que ServidorDeStreaming busca al reproducir.
func (r *RepositorioCanciones) GuardarCancion(idTipo string, id string, data []byte) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Crear carpeta si no existe
	if err := os.MkdirAll(rutaAudios(), os.ModePerm); err != nil {
		return fmt.Errorf("error creando la carpeta de audios: %v", err)
	}

	fileName := fmt.Sprintf("%s_%s.mp3", idTipo, id)
	filePath := filepath.Join(rutaAudios(), fileName)

	// Guardar archivo físico
	err := os.WriteFile(filePath, data, 0644)
	if err != nil {
		return fmt.Errorf("error al guardar archivo: %v", err)
	}

	// crear registro en memoria
	return nil
}