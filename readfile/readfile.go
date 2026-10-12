package readfile

import "os"

func lireFichier(nom string) (string, error) {
	// Lire le contenu du fichier.
	contenu, err := os.ReadFile(nom)
	if err != nil {
		return "", err
	}

	// Convertir les octets en texte.
	return string(contenu), nil
}