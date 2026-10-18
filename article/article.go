package article

import (
	"strings"
	"unicode"
)

// commenceParVoyelle vérifie si le mot commence par a, e, i, o, u ou h.
func commenceParVoyelle(mot string) bool {
	for _, c := range mot {
		if unicode.IsLetter(c) {
			return strings.ContainsRune("aeiouh", unicode.ToLower(c))
		}
	}

	return false
}

// Corriger remplace "a" par "an" devant une voyelle ou un h.
func Corriger(s string) string {
	mots := strings.Fields(s)

	for i := 0; i+1 < len(mots); i++ {
		if mots[i] == "a" && commenceParVoyelle(mots[i+1]) {
			mots[i] = "an"
		} else if mots[i] == "A" && commenceParVoyelle(mots[i+1]) {
			mots[i] = "An"
		}
	}

	return strings.Join(mots, " ")
}
