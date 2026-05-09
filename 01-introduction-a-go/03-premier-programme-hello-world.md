🔝 Retour au [Sommaire](/SOMMAIRE.md)

# 1-3 : Premier programme "Hello World"

## Introduction

Le traditionnel programme "Hello World" est souvent le premier programme qu'on écrit dans un nouveau langage. C'est une excellente façon de vérifier que votre environnement fonctionne et de découvrir la syntaxe de base de Go.

## Création de votre premier programme

### Étape 1 : Préparer l'environnement

**Créer un nouveau répertoire :**

```bash
mkdir hello-world
cd hello-world
```

**Initialiser un module Go :**

```bash
go mod init hello-world
```

Cette commande crée un fichier `go.mod` qui identifie votre projet comme un module Go.

### Étape 2 : Écrire le code

**Créer le fichier `main.go` :**

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

### Étape 3 : Exécuter le programme

```bash
go run main.go
```

**Résultat attendu :**

```html
Hello, World!
```

\***\*Félicitations ! Vous venez d'exécuter votre premier programme Go ! 🎉\*\***

## Analyse ligne par ligne

Maintenant, décortiquons chaque ligne de ce programme pour comprendre ce qui se passe :

### Ligne 1 : `package main`

```go
package main
```

**Qu'est-ce qu'un package ?**

- Un package est un moyen d'organiser le code Go
- Chaque fichier Go doit commencer par une déclaration de package
- `main` est un package spécial qui indique que ce fichier contient un programme exécutable

**Pourquoi `main` ?**

- Le package `main` dit à Go : "Ce fichier contient un programme qui peut être exécuté"
- Sans `package main`, vous ne pourriez pas créer un programme exécutable

### Ligne 2 : `import "fmt"`

```go
import "fmt"
```

**Qu'est-ce qu'un import ?**

- `import` permet d'utiliser du code d'autres packages
- `fmt` est un package de la bibliothèque standard de Go
- `fmt` signifie "format" et contient des fonctions pour afficher du texte

**Analogie :**
C'est comme dire "J'ai besoin d'emprunter des outils de la boîte à outils 'fmt' pour mon projet"

### Ligne 3 : Ligne vide

```go

```

**Pourquoi une ligne vide ?**

- Améliore la lisibilité du code
- Sépare les déclarations des instructions
- Bonne pratique en Go

### Ligne 4 : `func main() {`

```go
func main() {
```

**Qu'est-ce qu'une fonction ?**

- Une fonction est un bloc de code qui effectue une tâche spécifique
- `func` est le mot-clé pour déclarer une fonction
- `main` est le nom de la fonction

**Pourquoi `main` ?**

- La fonction `main` est le point d'entrée de votre programme
- C'est la première fonction qui s'exécute quand vous lancez le programme
- Obligatoire dans le package `main`

**Les accolades `{` :**

- Indiquent le début du bloc de code de la fonction
- Tout ce qui se trouve entre `{` et `}` appartient à la fonction

### Ligne 5 : `fmt.Println("Hello, World!")`

```go
    fmt.Println("Hello, World!")
```

**Décomposition :**

- `fmt` : Le package qu'on a importé
- `.` : L'opérateur d'accès (comme dire "dans le package fmt")
- `Println` : Une fonction du package fmt qui affiche du texte
- `("Hello, World!")` : Les parenthèses contiennent ce qu'on veut afficher
- `"Hello, World!"` : Une chaîne de caractères (string)

**Que fait `Println` ?**

- Affiche le texte suivi d'un retour à la ligne
- "Print line" = Afficher une ligne

### Ligne 6 : `}`

```go
}
```

**L'accolade fermante :**

- Marque la fin de la fonction `main`
- Doit toujours fermer chaque `{` ouvert

## Variantes du programme Hello World

### Version avec une variable

```go
package main

import "fmt"

func main() {
    message := "Hello, World!"
    fmt.Println(message)
}
```

**Nouveautés :**

- `message :=` : Déclaration et assignation d'une variable
- `fmt.Println(message)` : Affiche le contenu de la variable

### Version avec plusieurs messages

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
    fmt.Println("Bienvenue dans Go!")
    fmt.Println("C'est votre premier programme!")
}
```

### Version avec formatage

```go
package main

import "fmt"

func main() {
    nom := "Alice"
    age := 25
    fmt.Printf("Bonjour, je m'appelle %s et j'ai %d ans.\n", nom, age)
}
```

**Nouveautés :**

- `Printf` : Fonction de formatage (comme `printf` en C)
- `%s` : Placeholder pour une chaîne de caractères
- `%d` : Placeholder pour un nombre entier
- `\n` : Retour à la ligne

## Compilation vs Exécution

### Exécution directe

```bash
go run main.go
```

**Avantages :**

- Rapide pour tester
- Pas de fichier temporaire créé
- Idéal pour le développement

### Compilation puis exécution

```bash
# Compiler le programme
go build main.go

# Exécuter le fichier compilé
./main          # Sur Linux/macOS
main.exe        # Sur Windows
```

**Avantages :**

- Crée un fichier exécutable
- Plus rapide à l'exécution (déjà compilé)
- Peut être distribué sans Go installé

### Compilation avec nom personnalisé

```bash
# Compiler avec un nom spécifique
go build -o mon-programme main.go

# Exécuter
./mon-programme
```

## Erreurs courantes et solutions

### Erreur : "package main is not in GOROOT"

**Cause :** Vous n'avez pas initialisé un module Go

**Solution :**

```bash
go mod init nom-du-projet
```

### Erreur : "undefined: fmt"

**Cause :** Vous avez oublié d'importer le package fmt

**Solution :**

```go
import "fmt"  // Ajouter cette ligne
```

### Erreur : "expected 'package', found 'import'"

**Cause :** La déclaration `package main` doit être la première ligne

**Solution :**

```go
package main  // Doit être en première ligne

import "fmt"
```

### Erreur : "main function not found"

**Cause :** Pas de fonction `main` dans le package `main`

**Solution :**

```go
func main() {
    // Votre code ici
}
```

## Bonnes pratiques

### Formatage automatique

Toujours formater votre code avec :

```bash
go fmt main.go
```

**Avant le formatage :**

```go
package main
import"fmt"
func main(){
fmt.Println("Hello, World!")
}
```

**Après le formatage :**

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

### Vérification du code

```bash
go vet main.go
```

Cette commande vérifie les erreurs potentielles dans votre code.

### Structure recommandée

```go
package main

import (
    "fmt"
    // Autres imports si nécessaire
)

func main() {
    // Votre code ici
}
```

## Exercices pratiques

### Exercice 1 : Message personnalisé

Modifiez le programme pour qu'il affiche votre nom :

```text
Hello, [Votre nom]!
```

### Exercice 2 : Plusieurs lignes

Créez un programme qui affiche :

```text
Bonjour !
Je m'appelle [Votre nom]
J'apprends Go
```

### Exercice 3 : Avec variables

Utilisez des variables pour stocker votre nom et votre âge, puis affichez-les.

### Exercice 4 : Compilation

Compilez votre programme et créez un fichier exécutable nommé `salut`.

### Solutions

**Exercice 1 :**

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Alice!")
}
```

**Exercice 2 :**

```go
package main

import "fmt"

func main() {
    fmt.Println("Bonjour !")
    fmt.Println("Je m'appelle Alice")
    fmt.Println("J'apprends Go")
}
```

**Exercice 3 :**

```go
package main

import "fmt"

func main() {
    nom := "Alice"
    age := 25
    fmt.Printf("Je m'appelle %s et j'ai %d ans.\n", nom, age)
}
```

**Exercice 4 :**

```bash
go build -o salut main.go
./salut
```

## Récapitulatif

**Ce que vous avez appris :**

- ✅ Structure de base d'un programme Go
- ✅ Rôle du package `main` et de la fonction `main`
- ✅ Comment importer et utiliser des packages
- ✅ Différence entre `go run` et `go build`
- ✅ Bonnes pratiques de formatage

**Points clés à retenir :**

1. Tout programme Go exécutable doit avoir `package main`
2. La fonction `main()` est le point d'entrée
3. `import` permet d'utiliser des packages externes
4. `fmt.Println()` affiche du texte
5. `go fmt` formate automatiquement votre code

**Commandes importantes :**

```bash
go mod init projet    # Initialiser un module
go run main.go        # Exécuter directement
go build main.go      # Compiler
go fmt main.go        # Formater
go vet main.go        # Vérifier
```

---

_Dans la prochaine section, nous explorerons la structure d'un projet Go et découvrirons comment organiser votre code de manière professionnelle !_

⏭️
