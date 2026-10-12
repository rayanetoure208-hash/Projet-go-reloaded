package cap

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// mettreEnCapitale met la première lettre en majuscule.
func mettreEnCapitale(mot string) string {
	if mot == "" {
		return mot
	}

	premiere, taille := utf8.DecodeRuneInString(mot)

	debut := string(unicode.ToUpper(premiere))
	reste := strings.ToLower(mot[taille:])

	return debut + reste
}