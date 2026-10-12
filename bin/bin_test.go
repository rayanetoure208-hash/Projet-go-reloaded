package bin

import "testing"

func TestEnDecimal(t *testing.T) {
	tests := map[string]string{"10": "2", "1010": "10", "2": "2"}
	for entree, attendu := range tests {
		if got := EnDecimal(entree); got != attendu {
			t.Errorf("EnDecimal(%q) = %q, attendu %q", entree, got, attendu)
		}
	}
}
