// Package apostrophe place les guillemets simples contre les mots qu'ils entourent.
package apostrophe

import "unicode"

// dansMot vrai pour une apostrophe interne comme dans "don't".
func dansMot(r []rune, i int) bool {
	return i > 0 && i+1 < len(r) && unicode.IsLetter(r[i-1]) && unicode.IsLetter(r[i+1])
}

// Corriger retire les espaces après l'apostrophe ouvrante et avant la fermante.
func Corriger(s string) string {
	r := []rune(s)
	out := make([]rune, 0, len(r))
	ouvert := false

	for i := 0; i < len(r); i++ {
		c := r[i]
		if c != '\'' || dansMot(r, i) {
			out = append(out, c)
			continue
		}
		if !ouvert {
			out = append(out, c)
			for i+1 < len(r) && r[i+1] == ' ' {
				i++
			}
			ouvert = true
		} else {
			for len(out) > 0 && out[len(out)-1] == ' ' {
				out = out[:len(out)-1]
			}
			out = append(out, c)
			ouvert = false
		}
	}
	return string(out)
}
