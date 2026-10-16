package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"

	"go-reloaded/apostrophe"
	"go-reloaded/article"
	"go-reloaded/bin"
	"go-reloaded/capitale"
	"go-reloaded/hex"
	"go-reloaded/lecture"
	"go-reloaded/low"
	"go-reloaded/modificateur"
	"go-reloaded/ponctuation"
	"go-reloaded/up"
	"go-reloaded/writefile"
)

var reMarqueur = regexp.MustCompile(`\(\s*(up|low|cap|hex|bin)\s*(?:,\s*(\d+)\s*)?\)`)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "usage : go run . <fichier_entree> <fichier_sortie>")
		os.Exit(1)
	}

	texte, err := lecture.Lire(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "erreur :", err)
		os.Exit(1)
	}

	lignes := strings.Split(texte, "\n")
	for i, l := range lignes {
		lignes[i] = traiterLigne(l)
	}

	if err := writefile.Ecrire(os.Args[2], strings.Join(lignes, "\n")); err != nil {
		fmt.Fprintln(os.Stderr, "erreur :", err)
		os.Exit(1)
	}
	}

func traiterLigne(ligne string) string {
	ligne = isolerMarqueurs(ligne)
	mots := appliquerModificateurs(strings.Fields(ligne))
	s := strings.Join(mots, " ")
	s = ponctuation.Corriger(s)
	s = apostrophe.Corriger(s)
	return article.Corriger(s)
}

// isolerMarqueurs entoure chaque balise d'espaces pour la séparer de la ponctuation.
func isolerMarqueurs(ligne string) string {
	return reMarqueur.ReplaceAllStringFunc(ligne, func(m string) string {
		sm := reMarqueur.FindStringSubmatch(m)
		if sm[2] == "" {
			return " (" + sm[1] + ") "
		}
		return " (" + sm[1] + ", " + sm[2] + ") "
	})
}

// appliquerModificateurs exécute les balises et les retire de la liste de mots.
func appliquerModificateurs(mots []string) []string {
	res := make([]string, 0, len(mots))
	for i := 0; i < len(mots); {
		m, consomme, ok := modificateur.Lire(mots, i)
		if !ok {
			res = append(res, mots[i])
			i++
			continue
		}
		i += consomme
		appliquerSurPrecedents(res, m)
	}
	return res
}

// appliquerSurPrecedents transforme les m.N mots précédents (la ponctuation seule ne compte pas).
func appliquerSurPrecedents(res []string, m modificateur.Modificateur) {
	restant := m.N
	for j := len(res) - 1; j >= 0 && restant > 0; j-- {
		if !contientAlnum(res[j]) {
			continue
		}
		res[j] = transformerMot(res[j], m.Nom)
		restant--
	}
}

func estAlnum(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

func contientAlnum(s string) bool {
	return strings.IndexFunc(s, estAlnum) >= 0
}

// bornes retourne le début et la fin (en octets) de la partie alphanumérique d'un mot.
func bornes(mot string) (int, int) {
	debut, fin := -1, 0
	for i, r := range mot {
		if estAlnum(r) {
			if debut < 0 {
				debut = i
			}
			fin = i + utf8.RuneLen(r)
		}
	}
	return debut, fin
}

func transformerMot(mot, nom string) string {
	debut, fin := bornes(mot)
	prefixe, coeur, suffixe := mot[:debut], mot[debut:fin], mot[fin:]
	switch nom {
	case "hex":
		coeur = hex.EnDecimal(coeur)
	case "bin":
		coeur = bin.EnDecimal(coeur)
	case "up":
		coeur = up.EnMajuscules(coeur)
	case "low":
		coeur = low.EnMinuscules(coeur)
	case "cap":
		coeur = capitale.EnCapitale(coeur)
	}
	return prefixe + coeur + suffixe

}