package low

import "testing"

func TestEnMinuscules(t *testing.T) {
	tests := map[string]string{"SHOUTING": "shouting", "shouting": "shouting", "": "", "ÉTÉ": "été"}
	for entree, attendu := range tests {
		if got := EnMinuscules(entree); got != attendu {
			t.Errorf("EnMinuscules(%q) = %q, attendu %q", entree, got, attendu)
		}
	}
}
