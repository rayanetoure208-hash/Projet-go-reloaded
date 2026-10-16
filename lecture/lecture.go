// Package lecture lit le fichier d'entrée.
package lecture

import (
	"fmt"
	"os"
)

// Lire retourne le contenu du fichier situé à chemin.
func Lire(chemin string) (string, error) {
	data, err := os.ReadFile(chemin)
	if err != nil {
		return "", fmt.Errorf("lecture de %q : %w", chemin, err)
	}
	return string(data), nil
}
