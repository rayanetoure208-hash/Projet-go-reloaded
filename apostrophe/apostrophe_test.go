package apostrophe

import "testing"

func TestCorriger(t *testing.T) {
	tests := map[string]string{
		"me: ' awesome '":                    "me: 'awesome'",
		"said: ' I am the most well-known '": "said: 'I am the most well-known'",
		"don't stop":                         "don't stop",
		"' a ' and ' b '":                    "'a' and 'b'",
		"sans apostrophe":                    "sans apostrophe",
	}
	for entree, attendu := range tests {
		if got := Corriger(entree); got != attendu {
			t.Errorf("Corriger(%q) = %q, attendu %q", entree, got, attendu)
		}
	}
}
