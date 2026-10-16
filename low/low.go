// Package low met un mot en minuscules.
package low

import "strings"

// EnMinuscules retourne mot en minuscules.
func EnMinuscules(mot string) string {
	return strings.ToLower(mot)
}
