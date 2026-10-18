package transform

import (
	"strings"
	"unicode"
	"unicode/utf8"

	"go-reloaded/apostrophe"
	"go-reloaded/article"
	"go-reloaded/bin"
	"go-reloaded/capitale"
	"go-reloaded/hex"
	"go-reloaded/low"
	"go-reloaded/modificateur"
	"go-reloaded/ponctuation"
	"go-reloaded/up"
)

// Appliquer traite le texte ligne par ligne.
func Appliquer(texte string) string {
	lignes := strings.Split(texte, "\n")

	for i, ligne := range lignes {
		lignes[i] = appliquerLigne(ligne)
	}

	return strings.Join(lignes, "\n")
}

// appliquerLigne applique les transformations sur une ligne.
func appliquerLigne(ligne string) string {
	mots := appliquerModificateurs(strings.Fields(ligne))
	texte := strings.Join(mots, " ")

	texte = ponctuation.Corriger(texte)
	texte = apostrophe.Corriger(texte)
	texte = article.Corriger(texte)

	return texte
}

// appliquerModificateurs exécute les instructions et les retire du texte.
func appliquerModificateurs(mots []string) []string {
	resultat := make([]string, 0, len(mots))

	for i := 0; i < len(mots); {
		m, consomme, ok := modificateur.Lire(mots, i)

		if !ok {
			resultat = append(resultat, mots[i])
			i++
			continue
		}

		i += consomme
		appliquerSurPrecedents(resultat, m)
	}

	return resultat
}

// appliquerSurPrecedents transforme les mots précédents.
func appliquerSurPrecedents(resultat []string, m modificateur.Modificateur) {
	restant := m.N

	for j := len(resultat) - 1; j >= 0 && restant > 0; j-- {
		if !contientAlnum(resultat[j]) {
			continue
		}

		resultat[j] = transformerMot(resultat[j], m.Nom)
		restant--
	}
}

// estAlnum vérifie si le caractère est une lettre ou un chiffre.
func estAlnum(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

// contientAlnum vérifie si le texte contient une lettre ou un chiffre.
func contientAlnum(s string) bool {
	return strings.IndexFunc(s, estAlnum) >= 0
}

// bornes trouve les limites du mot en ignorant sa ponctuation extérieure.
func bornes(mot string) (int, int) {
	debut, fin := -1, 0

	for i, r := range mot {
		if estAlnum(r) {
			if debut == -1 {
				debut = i
			}

			fin = i + utf8.RuneLen(r)
		}
	}

	return debut, fin
}

// transformerMot applique la transformation demandée au mot.
func transformerMot(mot, nom string) string {
	debut, fin := bornes(mot)

	if debut == -1 {
		return mot
	}

	prefixe := mot[:debut]
	coeur := mot[debut:fin]
	suffixe := mot[fin:]

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