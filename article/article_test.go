package article

import "testing"

func TestCorriger(t *testing.T) {
	tests := map[string]string{
		"There it was. A amazing rock!": "There it was. An amazing rock!",
		"a untold story":                "an untold story",
		"a hour":                        "an hour",
		"a dog":                         "a dog",
		"a":                             "a",
		"banana apple":                  "banana apple",
		"a 'apple'":                     "an 'apple'",
	}
	for entree, attendu := range tests {
		if got := Corriger(entree); got != attendu {
			t.Errorf("Corriger(%q) = %q, attendu %q", entree, got, attendu)
		}
	}
}
