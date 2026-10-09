package main

import "os"

func readFile(filename string) (string, error) {
	// Lire le contenu du fichier
	data, err := os.ReadFile(filename)

	// Vérifier si la lecture a échoué
	if err != nil {
		return "", err
	}

	// Transformer les données en texte
	return string(data), nil
}