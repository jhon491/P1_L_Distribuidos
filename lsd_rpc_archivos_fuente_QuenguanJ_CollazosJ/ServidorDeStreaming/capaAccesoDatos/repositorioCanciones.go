package capaaccesodatos

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// rutaAudios apunta a la carpeta compartida "audios", al mismo nivel que las
// carpetas de los programas. Se puede sobrescribir con RUTA_AUDIOS.
func rutaAudios() string {
	if r := os.Getenv("RUTA_AUDIOS"); r != "" {
		return r
	}
	return filepath.Join("..", "audios")
}

// idsTipos relaciona el nombre normalizado de cada tipo de audio (el que
// envía el cliente en AudioRequest.tipo_audio) con su idTipo en
// ServidorMetadataDeAudios.
var idsTipos = map[string]int{
	"musica":      1,
	"podcast":     2,
	"audiolibro":  3,
	"ruidoblanco": 4,
}

// normalizarTipo pasa a minúsculas y quita tildes y espacios:
// "Música" -> "musica", "Ruido Blanco" -> "ruidoblanco".
func normalizarTipo(nombre string) string {
	reemplazo := strings.NewReplacer("á", "a", "é", "e", "í", "i", "ó", "o", "ú", "u", "ü", "u", " ", "")
	return reemplazo.Replace(strings.ToLower(strings.TrimSpace(nombre)))
}

// IdTipoPorNombre devuelve el idTipo a partir del nombre del tipo. También
// acepta el id numérico como texto ("1").
func IdTipoPorNombre(nombre string) (int, error) {
	if n, err := strconv.Atoi(strings.TrimSpace(nombre)); err == nil && n > 0 {
		return n, nil
	}
	if id, ok := idsTipos[normalizarTipo(nombre)]; ok {
		return id, nil
	}
	return 0, fmt.Errorf("tipo de audio desconocido: %q", nombre)
}

// AbrirArchivo abre audios/<idTipo>_<idAudio>.mp3 y lo devuelve como *os.File.
// Los ids de audio se repiten entre tipos (cada tipo tiene su propio 1, 2...),
// por eso el nombre del archivo incluye también el id del tipo.
func AbrirArchivo(tipoAudio string, idAudio int32) (*os.File, error) {
	idTipo, err := IdTipoPorNombre(tipoAudio)
	if err != nil {
		return nil, err
	}

	ruta := filepath.Join(rutaAudios(), fmt.Sprintf("%d_%d.mp3", idTipo, idAudio))

	file, err := os.Open(ruta)
	if err != nil {
		fmt.Println("Error Audio abierto:", ruta)
		return nil, fmt.Errorf("no se pudo abrir el archivo: %w", err)
	}
	fmt.Println("Audio abierto:", ruta)
	return file, nil
}