// Package capitale met la première lettre d'un mot en majuscule.
package capitale

import (
	"strings"
	"unicode"
)

// EnCapitale retourne mot avec une majuscule initiale et le reste en minuscules.
func EnCapitale(mot string) string {
	r := []rune(strings.ToLower(mot))
	if len(r) == 0 {
		return mot
	}
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}
