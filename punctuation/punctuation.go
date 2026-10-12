package punctuation

import (
	"regexp"
	"strings"
)

func corrigerPonctuation(texte string) string {
	// Supprimer les espaces avant la ponctuation.
	avant := regexp.MustCompile(`\s+([.,!?;:])`)
	texte = avant.ReplaceAllString(texte, "$1")

	// Ajouter un espace après un groupe de ponctuation
	// lorsqu'un autre mot suit immédiatement.
	apres := regexp.MustCompile(`([.,!?;:]+)(\S)`)
	texte = apres.ReplaceAllString(texte, "$1 $2")

	// Réduire les espaces multiples.
	espaces := regexp.MustCompile(`[ \t]+`)
	texte = espaces.ReplaceAllString(texte, " ")

	return strings.TrimSpace(texte)
}