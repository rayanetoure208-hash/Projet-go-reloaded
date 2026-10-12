package modificateur

import "testing"

func TestLire(t *testing.T) {
	m, n, ok := Lire([]string{"(up,", "2)"}, 0)
	if !ok || n != 2 || m.Nom != "up" || m.N != 2 {
		t.Errorf("got %+v %d %v", m, n, ok)
	}
	m, n, ok = Lire([]string{"(hex)"}, 0)
	if !ok || n != 1 || m.Nom != "hex" || m.N != 1 {
		t.Errorf("got %+v %d %v", m, n, ok)
	}
	if _, _, ok = Lire([]string{"bonjour"}, 0); ok {
		t.Error("bonjour ne doit pas être une balise")
	}
}
