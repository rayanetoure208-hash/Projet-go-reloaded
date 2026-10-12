// Package hex convertit un nombre hexadécimal en décimal.
package hex

import "strconv"

// EnDecimal convertit mot (hexadécimal) en décimal.
// Si mot n'est pas un nombre valide, il est retourné tel quel.
func EnDecimal(mot string) string {
	n, err := strconv.ParseInt(mot, 16, 64)
	if err != nil {
		return mot
	}
	return strconv.FormatInt(n, 10)
}
