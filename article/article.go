// Package article transforme "a" en "an" devant une voyelle ou un h.
package article

import (
	"strings"
	"unicode"
)

// commenceParVoyelle vrai si la première lettre du mot est a, e, i, o, u ou h.
// Les caractères non alphabétiques initiaux (ex : une apostrophe) sont ignorés.
func commenceParVoyelle(mot string) bool {
	for _, c := range mot {
		if !unicode.IsLetter(c) {
			continue
		}
		return strings.ContainsRune("aeiouh", unicode.ToLower(c))
	}
	return false
}

// Corriger remplace chaque "a"/"A" par "an"/"An" quand le mot suivant le demande.
func Corriger(s string) string {
	mots := strings.Fields(s)
	for i := 0; i+1 < len(mots); i++ {
		if (mots[i] == "a" || mots[i] == "A") && commenceParVoyelle(mots[i+1]) {
			mots[i] += "n"
		}
	}
	return strings.Join(mots, " ")
}
