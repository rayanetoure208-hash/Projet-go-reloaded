package modificateur

import (
	"regexp"
	"strconv"
	"strings"
)

type Modificateur struct {
	Nom string
	N   int
}

var motif = regexp.MustCompile(`^\((hex|bin|up|low|cap)(?:,(\d+))?\)$`)

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
