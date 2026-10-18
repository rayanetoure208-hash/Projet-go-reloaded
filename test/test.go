package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type Test struct {
	nom     string
	entree  string
	attendu string
}

func main() {
	fmt.Println("===== TESTS GO RELOADED =====")

	dossier, err := os.MkdirTemp("", "go-reloaded-test")
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}
	defer os.RemoveAll(dossier)

	fichierEntree := filepath.Join(dossier, "entree.txt")
	fichierSortie := filepath.Join(dossier, "sortie.txt")

	tests := []Test{
		{"Majuscules", "hello (up)\n", "HELLO\n"},
		{"Minuscules", "HELLO (low)\n", "hello\n"},
		{"Capitales", "hello (cap)\n", "Hello\n"},
		{"Hexadécimal", "1E (hex)\n", "30\n"},
		{"Binaire", "101 (bin)\n", "5\n"},
		{"Ponctuation", "hello , world !\n", "hello, world!\n"},
		{"Apostrophes", "' hello world '\n", "'hello world'\n"},
		{"Article", "a apple\n", "an apple\n"},
		{"Transformations combinées", "hello (up), WORLD (low)\n", "HELLO, world\n"},
		{"Fichier vide", "", ""},
	}

	reussis := 0

	for _, test := range tests {
		err := os.WriteFile(fichierEntree, []byte(test.entree), 0644)
		if err != nil {
			fmt.Printf("[ÉCHEC] %s : préparation impossible\n", test.nom)
			continue
		}

		_ = os.Remove(fichierSortie)

		cmd := exec.Command("go", "run", ".", fichierEntree, fichierSortie)
		sortie, err := cmd.CombinedOutput()

		if err != nil {
			fmt.Printf("[ÉCHEC] %s : %v\n%s", test.nom, err, sortie)
			continue
		}

		resultat, err := os.ReadFile(fichierSortie)
		if err != nil {
			fmt.Printf("[ÉCHEC] %s : fichier de sortie absent\n", test.nom)
			continue
		}

		if string(resultat) == test.attendu {
			fmt.Printf("[RÉUSSI] %s\n", test.nom)
			reussis++
		} else {
			fmt.Printf("[ÉCHEC] %s\n", test.nom)
			fmt.Printf("  Attendu : %q\n", test.attendu)
			fmt.Printf("  Obtenu  : %q\n", string(resultat))
		}
	}

	fmt.Println("\n===== BILAN =====")
	fmt.Printf("%d tests réussis sur %d\n", reussis, len(tests))
}
