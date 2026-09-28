package capaaccesodatos

import (
	"fmt"
	"os"
	"path/filepath"
)

// rutaAudios apunta a la carpeta compartida "audios", al mismo nivel que las
// carpetas de los programas. Se puede sobrescribir con RUTA_AUDIOS.
func rutaAudios() string {
	if r := os.Getenv("RUTA_AUDIOS"); r != "" {
		return r
	}
	return filepath.Join("..", "audios")
}

// AbrirArchivo abre audios/<id>.mp3 y lo devuelve como *os.File.
func AbrirArchivo(idAudio int32) (*os.File, error) {
	ruta := filepath.Join(rutaAudios(), fmt.Sprintf("%d.mp3", idAudio))

	file, err := os.Open(ruta)
	if err != nil {
		fmt.Println("Error Audio abierto:", ruta)
		return nil, fmt.Errorf("no se pudo abrir el archivo: %w", err)
	}
	fmt.Println("Audio abierto:", ruta)
	return file, nil
}
