package capitale

import "testing"

func TestEnCapitale(t *testing.T) {
	tests := map[string]string{"bridge": "Bridge", "BRIDGE": "Bridge", "": "", "é": "É"}
	for entree, attendu := range tests {
		if got := EnCapitale(entree); got != attendu {
			t.Errorf("EnCapitale(%q) = %q, attendu %q", entree, got, attendu)
		}
	}
}
