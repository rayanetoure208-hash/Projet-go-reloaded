# go-reloaded

Outil en Go qui complète, édite et corrige automatiquement un texte.

Le programme lit un fichier d'entrée, applique des transformations,
puis écrit le texte modifié dans un fichier de sortie.

## Utilisation

```bash
go run . sample.txt result.txt
```

- `sample.txt` : fichier d'entrée (le texte à modifier)
- `result.txt` : fichier de sortie (créé ou écrasé par le programme)

## Exemple

Entrée :

```
Simply add 42 (hex) and 10 (bin) and you will see the result is 68.
```

Sortie :

```
Simply add 66 and 2 and you will see the result is 68.
```

## Transformations

| Règle | Effet | Exemple |
|---|---|---|
| `(hex)` | mot précédent converti de l'hexadécimal en décimal | `1E (hex)` → `30` |
| `(bin)` | mot précédent converti du binaire en décimal | `10 (bin)` → `2` |
| `(up)` | mot précédent en majuscules | `go (up)` → `GO` |
| `(low)` | mot précédent en minuscules | `SHOUTING (low)` → `shouting` |
| `(cap)` | mot précédent avec une majuscule initiale | `bridge (cap)` → `Bridge` |
| `(up, n)` `(low, n)` `(cap, n)` | s'applique aux `n` mots précédents | `so exciting (up, 2)` → `SO EXCITING` |
| Ponctuation `. , ! ? : ;` | collée au mot précédent, espace après | `there ,and` → `there, and` |
| Groupes `...` `!?` | restent collés ensemble | `thinking ...` → `thinking...` |
| Apostrophes `' '` | collées au texte qu'elles entourent | `' awesome '` → `'awesome'` |
| `a` | devient `an` devant une voyelle ou un `h` | `a amazing` → `an amazing` |

## Structure du projet

Un dossier par fonctionnalité :

| Dossier | Rôle |
|---|---|
| `lecture` | lecture du fichier d'entrée |
| `writefile` | écriture du fichier de sortie |
| `hex`, `bin` | conversions en décimal |
| `up`, `low`, `capitale` | changements de casse |
| `modificateur` | détection des balises `(up, 2)`, etc. |
| `ponctuation` | correction de la ponctuation |
| `apostrophe` | correction des apostrophes |
| `article` | transformation de `a` en `an` |
| `transform` | enchaîne toutes les étapes |

## Tests

```bash
go test ./...
```

## Auteur

Rayane Toure 