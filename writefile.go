package main

import "os"

func writeFile(filename, text string) error {
	// Transformer le texte en tableau d'octets et écrire le fichier
	err := os.WriteFile(filename, []byte(text), 0644)

	// Retourner l'erreur éventuelle
	return err
}