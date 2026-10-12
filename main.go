package main

import (
	"fmt"
	"os"

	"go-reloaded/lecture"
	"go-reloaded/transform"
	"go-reloaded/writefile"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "usage : go run . <fichier_entree> <fichier_sortie>")
		os.Exit(1)
	}

	texte, err := lecture.Lire(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "erreur :", err)
		os.Exit(1)
	}

	if err := writefile.Ecrire(os.Args[2], transform.Appliquer(texte)); err != nil {
		fmt.Fprintln(os.Stderr, "erreur :", err)
		os.Exit(1)
	}
}
