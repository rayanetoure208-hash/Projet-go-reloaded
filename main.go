package main  

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	// pour verifier si on a bien deux fichiers 
	if len(os.arg) !=3 {
		return
	}
	//pour lire le fichier sample 
	data, err := os.Readfile(os.args[1])
	if err != nil {
      return 
	}
	text := string(data)

	// pour separer le texte en mots 
	words := strings.fields(text)

	// Parcourir les mots 
	for i := 0; i < len(words); i++ {
		
	}
}