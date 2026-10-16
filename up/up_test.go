package up

import "testing"

func TestEnMajuscules(t *testing.T) {
	tests := map[string]string{"go": "GO", "GO": "GO", "": "", "été": "ÉTÉ"}
	for entree, attendu := range tests {
		if got := EnMajuscules(entree); got != attendu {
			t.Errorf("EnMajuscules(%q) = %q, attendu %q", entree, got, attendu)
		}
	}
}
