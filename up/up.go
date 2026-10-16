// Package up met un mot en majuscules.
package up

import "strings"

// EnMajuscules retourne mot en majuscules.
func EnMajuscules(mot string) string {
	return strings.ToUpper(mot)
}
