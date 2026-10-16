package lecture

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLire(t *testing.T) {
	chemin := filepath.Join(t.TempDir(), "in.txt")
	if err := os.WriteFile(chemin, []byte("bonjour\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	got, err := Lire(chemin)
	if err != nil || got != "bonjour\n" {
		t.Errorf("Lire = %q, %v", got, err)
	}
	if _, err := Lire(filepath.Join(t.TempDir(), "absent.txt")); err == nil {
		t.Error("une erreur est attendue pour un fichier absent")
	}
}
