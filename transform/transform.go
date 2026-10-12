package transform

import (
	"strings"

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


func Appliquer(texte string) string {
	lignes := strings.Split(texte, "\n")
	for i, l := range lignes {
		lignes[i] = appliquerLigne(l)
	}
	return strings.Join(lignes, "\n")
}

func appliquerLigne(ligne string) string {
	mots := appliquerModificateurs(strings.Fields(ligne))
	s := strings.Join(mots, " ")
	s = ponctuation.Corriger(s)
	s = apostrophe.Corriger(s)
	return article.Corriger(s)
}


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

func appliquerSurPrecedents(res []string, m modificateur.Modificateur) {
	debut := len(res) - m.N
	if debut < 0 {
		debut = 0
	}
	for j := debut; j < len(res); j++ {
		switch m.Nom {
		case "hex":
			res[j] = hex.EnDecimal(res[j])
		case "bin":
			res[j] = bin.EnDecimal(res[j])
		case "up":
			res[j] = up.EnMajuscules(res[j])
		case "low":
			res[j] = low.EnMinuscules(res[j])
		case "cap":
			res[j] = capitale.EnCapitale(res[j])
		}
	}
}
