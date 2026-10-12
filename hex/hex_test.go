package hex

import "testing"

func TestEnDecimal(t *testing.T) {
	tests := map[string]string{"1E": "30", "42": "66", "ff": "255", "zz": "zz"}
	for entree, attendu := range tests {
		if got := EnDecimal(entree); got != attendu {
			t.Errorf("EnDecimal(%q) = %q, attendu %q", entree, got, attendu)
		}
	}
}
