package transform

import "testing"

func TestAppliquer(t *testing.T) {
	tests := []struct {
		nom, entree, attendu string
	}{
		// Exemples de l'énoncé
		{"hex", "1E (hex) files were added", "30 files were added"},
		{"bin", "It has been 10 (bin) years", "It has been 2 years"},
		{"up", "Ready, set, go (up) !", "Ready, set, GO!"},
		{"low", "I should stop SHOUTING (low)", "I should stop shouting"},
		{"cap", "Welcome to the Brooklyn bridge (cap)", "Welcome to the Brooklyn Bridge"},
		{"up avec nombre", "This is so exciting (up, 2)", "This is SO EXCITING"},
		{"ponctuation", "I was sitting over there ,and then BAMM !!", "I was sitting over there, and then BAMM!!"},
		{"groupe", "I was thinking ... You were right", "I was thinking... You were right"},
		{"apostrophe", "I am exactly how they describe me: ' awesome '", "I am exactly how they describe me: 'awesome'"},
		{"apostrophe longue", "As Elton John said: ' I am the most well-known homosexual in the world '", "As Elton John said: 'I am the most well-known homosexual in the world'"},
		{"article", "There it was. A amazing rock!", "There it was. An amazing rock!"},
		{"exemple 1", "it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it", "It was the best of times, it was the worst of TIMES, it was the age of wisdom, it"},
		{"exemple 2", "Simply add 42 (hex) and 10 (bin) and you will see the result is 68.", "Simply add 66 and 2 and you will see the result is 68."},
		{"exemple 3", "There is no greater agony than bearing a untold story inside you.", "There is no greater agony than bearing an untold story inside you."},
		{"exemple 4", "Punctuation tests are ... kinda boring ,what do you think ?", "Punctuation tests are... kinda boring, what do you think?"},

		// Cas supplémentaires
		{"cap avec nombre", "so very good day (cap, 3)", "so Very Good Day"},
		{"low avec nombre", "I AM VERY LOUD (low, 2)", "I AM very loud"},
		{"nombre trop grand", "hello (up, 5)", "HELLO"},
		{"ponctuation non comptée", "Hello , world (up, 2)", "HELLO, WORLD"},
		{"hex avec ponctuation", "1E, (hex) fichiers", "30, fichiers"},
		{"hex minuscule et bin", "ff (hex) et 1010 (bin)", "255 et 10"},
		{"apostrophe et ponctuation", "he said: ' wait ! '", "he said: 'wait!'"},
		{"apostrophe interne", "It's 'fine' (up) ok", "It's 'FINE' ok"},
		{"article h", "bearing a untold story and a hero", "bearing an untold story and an hero"},
		{"decimaux", "pi vaut 3.14 environ", "pi vaut 3.14 environ"},
		{"multi lignes", "a apple\nb (up)", "an apple\nB"},
		{"saut de ligne final", "hello (up)\n", "HELLO\n"},
		{"vide", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.nom, func(t *testing.T) {
			if got := Appliquer(tt.entree); got != tt.attendu {
				t.Errorf("Appliquer(%q)\n got  %q\n want %q", tt.entree, got, tt.attendu)
			}
		})
	}
}
