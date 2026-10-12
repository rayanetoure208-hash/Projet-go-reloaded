// Package bin convertit un nombre binaire en décimal.
package bin

import "strconv"

// EnDecimal convertit mot (binaire) en décimal.
// Si mot n'est pas un nombre valide, il est retourné tel quel.
func EnDecimal(mot string) string {
	n, err := strconv.ParseInt(mot, 2, 64)
	if err != nil {
		return mot
	}
	return strconv.FormatInt(n, 10)
}
