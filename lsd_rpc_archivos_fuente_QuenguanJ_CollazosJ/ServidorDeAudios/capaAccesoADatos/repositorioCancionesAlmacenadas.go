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

// GuardarCancion guarda el archivo MP3 en la carpeta 'audios'
func (r *RepositorioCanciones) GuardarCancion(id string, data []byte) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Crear carpeta si no existe
	os.MkdirAll("../audios", os.ModePerm)

	// Construir nombre del archivo: titulo_genero_artista.mp3
	fileName := fmt.Sprintf("cancion_%s.mp3", id)
	filePath := filepath.Join("..", "audios", fileName)

	// Guardar archivo físico
	err := os.WriteFile(filePath, data, 0644)
	if err != nil {
		return fmt.Errorf("error al guardar archivo: %v", err)
	}

	// crear registro en memoria
	return nil
}
