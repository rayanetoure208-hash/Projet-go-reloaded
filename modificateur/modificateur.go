// Package modificateur reconnaît les balises (hex), (bin), (up), (low), (cap)
// ainsi que leurs variantes avec nombre : (up, 2).
package modificateur

import (
	"regexp"
	"strconv"
	"strings"
)

// Modificateur décrit une balise : son nom et le nombre de mots concernés.
type Modificateur struct {
	Nom string
	N   int
}

var motif = regexp.MustCompile(`^\((hex|bin|up|low|cap)(?:,(\d+))?\)$`)

// Lire tente de lire une balise à partir de mots[i].
// Elle retourne la balise, le nombre de mots consommés et true si une balise a été trouvée.
// "(up, 2)" est découpé par strings.Fields en "(up," et "2)" : on recolle les deux.
func Lire(mots []string, i int) (Modificateur, int, bool) {
	candidat := mots[i]
	consomme := 1
	if strings.HasSuffix(candidat, ",") && i+1 < len(mots) {
		candidat += mots[i+1]
		consomme = 2
	}

	m := motif.FindStringSubmatch(candidat)
	if m == nil {
		return Modificateur{}, 0, false
	}

	n := 1
	if m[2] != "" {
		v, err := strconv.Atoi(m[2])
		if err != nil || v < 1 {
			return Modificateur{}, 0, false
		}
		n = v
	}
	return Modificateur{Nom: m[1], N: n}, consomme, true
}
