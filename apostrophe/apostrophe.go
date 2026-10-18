package apostrophe

import "unicode"

// dansMot vérifie si l'apostrophe est au milieu d'un mot,
// comme dans don't ou I'm.
func dansMot(r []rune, i int) bool {
	return i > 0 &&
		i+1 < len(r) &&
		unicode.IsLetter(r[i-1]) &&
		unicode.IsLetter(r[i+1])
}

// Corriger corrige les espaces autour des apostrophes.
func Corriger(s string) string {
	r := []rune(s)
	out := make([]rune, 0, len(r))
	positions := make([]int, 0)

	// Repère les apostrophes qui ne sont pas à l'intérieur d'un mot.
	for i, c := range r {
		if c == '\'' && !dansMot(r, i) {
			positions = append(positions, i)
		}
	}

	// Associe les apostrophes deux par deux.
	paires := make(map[int]bool)
	for i := 0; i+1 < len(positions); i += 2 {
		paires[positions[i]] = true
		paires[positions[i+1]] = true
	}

	ouverture := -1

	for i := 0; i < len(r); i++ {
		if r[i] != '\'' || dansMot(r, i) {
			out = append(out, r[i])
			continue
		}

		if !paires[i] {
			// Apostrophe isolée : recolle les morceaux d'un mot.
			for len(out) > 0 && unicode.IsSpace(out[len(out)-1]) {
				out = out[:len(out)-1]
			}
			out = append(out, '\'')
			for i+1 < len(r) && unicode.IsSpace(r[i+1]) {
				i++
			}
			continue
		}

		if ouverture == -1 {
			// Apostrophe ouvrante : conserve l'espace extérieur,
			// mais supprime les espaces à l'intérieur.
			out = append(out, '\'')
			ouverture = i
			for i+1 < len(r) && unicode.IsSpace(r[i+1]) {
				i++
			}
		} else {
			// Apostrophe fermante : retire les espaces intérieurs.
			for len(out) > 0 && unicode.IsSpace(out[len(out)-1]) {
				out = out[:len(out)-1]
			}
			out = append(out, '\'')
			ouverture = -1
		}
	}

	return string(out)
}
