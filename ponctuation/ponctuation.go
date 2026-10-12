// Package ponctuation corrige les espaces autour de . , ! ? : ;
package ponctuation

import (
	"strings"
	"unicode"
)

func estPonctuation(r rune) bool {
	return strings.ContainsRune(".,!?:;", r)
}

// entreChiffres vrai si r[i] est entouré de chiffres (ex : 3.14, 1,000, 10:30).
func entreChiffres(r []rune, i int) bool {
	return i > 0 && i+1 < len(r) && unicode.IsDigit(r[i-1]) && unicode.IsDigit(r[i+1])
}

// Corriger colle la ponctuation au mot précédent et la sépare du mot suivant.
// Les groupes comme "..." ou "!?" restent collés entre eux.
func Corriger(s string) string {
	r := []rune(s)
	out := make([]rune, 0, len(r))
	besoinEspace := false

	for i, c := range r {
		switch {
		case estPonctuation(c):
			if entreChiffres(r, i) {
				out = append(out, c)
				besoinEspace = false
				continue
			}
			for len(out) > 0 && out[len(out)-1] == ' ' {
				out = out[:len(out)-1]
			}
			out = append(out, c)
			besoinEspace = true
		case unicode.IsSpace(c):
			if len(out) > 0 && out[len(out)-1] != ' ' {
				out = append(out, ' ')
			}
			besoinEspace = false
		default:
			if besoinEspace {
				out = append(out, ' ')
			}
			besoinEspace = false
			out = append(out, c)
		}
	}
	return strings.TrimSpace(string(out))
}
