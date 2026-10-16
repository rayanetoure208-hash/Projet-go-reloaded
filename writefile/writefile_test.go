package writefile

import (
	"os"
	"path/filepath"
	"testing"
)

func TestEcrire(t *testing.T) {
	chemin := filepath.Join(t.TempDir(), "out.txt")
	if err := Ecrire(chemin, "salut\n"); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(chemin)
	if err != nil || string(data) != "salut\n" {
		t.Errorf("contenu = %q, %v", data, err)
	}
	if err := Ecrire(filepath.Join(t.TempDir(), "absent", "out.txt"), "x"); err == nil {
		t.Error("une erreur est attendue pour un dossier inexistant")
	}
}
