package ponctuation

import "testing"

func TestCorriger(t *testing.T) {
	tests := map[string]string{
		"there ,and then BAMM !!":     "there, and then BAMM!!",
		"thinking ... You were right": "thinking... You were right",
		"ok !?":                       "ok!?",
		"pi vaut 3.14":                "pi vaut 3.14",
		"a ; b : c":                   "a; b: c",
		"boring ,what do you think ?": "boring, what do you think?",
		"":                            "",
	}
	for entree, attendu := range tests {
		if got := Corriger(entree); got != attendu {
			t.Errorf("Corriger(%q) = %q, attendu %q", entree, got, attendu)
		}
	}
}
