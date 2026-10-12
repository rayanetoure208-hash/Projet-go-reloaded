// Package writefile écrit le fichier de sortie.
package writefile

import (
	"fmt"
	"os"
)

//  écrit contenu dans le fichier 

func Ecrire(chemin, contenu string) error {
	if err := os.WriteFile(chemin, []byte(contenu), 0o644); err != nil {
		return fmt.Errorf("écriture de %q : %w", chemin, err)
	}
	return nil
}
