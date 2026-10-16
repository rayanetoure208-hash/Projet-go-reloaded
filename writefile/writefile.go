// Package writefile écrit le fichier de sortie.
package writefile

import (
	"fmt"
	"os"
)

// Ecrire écrit contenu dans le fichier chemin (créé ou écrasé).
func Ecrire(chemin, contenu string) error {
	if err := os.WriteFile(chemin, []byte(contenu), 0o644); err != nil {
		return fmt.Errorf("écriture de %q : %w", chemin, err)
	}
	return nil
}
