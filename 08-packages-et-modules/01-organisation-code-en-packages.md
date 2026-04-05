🔝 Retour au [Sommaire](/SOMMAIRE.md)

# 8-1 : Organisation du code en packages

## Qu'est-ce qu'un package ?

Un **package** en Go est simplement un dossier qui contient des fichiers `.go`. Tous les fichiers dans le même dossier appartiennent au même package et peuvent partager leurs fonctions, variables et types entre eux.

### Exemple simple

```text
mon-projet/
├── main.go          # package main
└── utils/
    ├── math.go      # package utils
    └── string.go    # package utils
```

## Règles de base des packages

### 1. Déclaration du package

Chaque fichier Go doit commencer par une déclaration de package :

```go
package nom_du_package
```

**Exemple :**

```go
// fichier: utils/math.go
package utils

func Add(a, b int) int {
    return a + b
}
```

### 2. Le package main

Le package `main` est spécial : c'est le point d'entrée de votre programme. Il doit contenir une fonction `main()`.

```go
// fichier: main.go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

### 3. Un dossier = un package

Tous les fichiers dans un même dossier doivent avoir la même déclaration de package.

## Importer des packages

### Import local

Pour utiliser un package de votre projet :

```go
package main

import (
    "fmt"
    "./utils"  // Import local (ancienne méthode)
)

func main() {
    result := utils.Add(5, 3)
    fmt.Println(result)
}
```

### Import avec modules (recommandé)

```go
package main

import (
    "fmt"
    "mon-projet/utils"
)

func main() {
    result := utils.Add(5, 3)
    fmt.Println(result)
}
```

## Exemple pratique : Calculatrice

Créons une calculatrice simple pour illustrer l'organisation en packages.

### Structure du projet

```text
calculatrice/
├── go.mod
├── main.go
├── math/
│   ├── basic.go
│   └── advanced.go
└── utils/
    └── display.go
```

### 1. Initialisation du module

```bash
go mod init calculatrice
```

### 2. Package math/basic.go

```go
package math

// Add additionne deux nombres
func Add(a, b float64) float64 {
    return a + b
}

// Subtract soustrait deux nombres
func Subtract(a, b float64) float64 {
    return a - b
}

// Multiply multiplie deux nombres
func Multiply(a, b float64) float64 {
    return a * b
}

// Divide divise deux nombres
func Divide(a, b float64) float64 {
    if b == 0 {
        return 0 // Gestion simple de la division par zéro
    }
    return a / b
}
```

### 3. Package math/advanced.go

```go
package math

import "math"

// Power élève un nombre à une puissance
func Power(base, exponent float64) float64 {
    return math.Pow(base, exponent)
}

// SquareRoot calcule la racine carrée
func SquareRoot(n float64) float64 {
    return math.Sqrt(n)
}
```

### 4. Package utils/display.go

```go
package utils

import "fmt"

// PrintResult affiche le résultat d'une opération
func PrintResult(operation string, a, b, result float64) {
    fmt.Printf("%s: %.2f et %.2f = %.2f\n", operation, a, b, result)
}

// PrintSingleResult affiche le résultat d'une opération à un seul paramètre
func PrintSingleResult(operation string, a, result float64) {
    fmt.Printf("%s: %.2f = %.2f\n", operation, a, result)
}
```

### 5. Fichier main.go

```go
package main

import (
    "calculatrice/math"
    "calculatrice/utils"
)

func main() {
    // Opérations basiques
    a, b := 10.0, 3.0

    result := math.Add(a, b)
    utils.PrintResult("Addition", a, b, result)

    result = math.Subtract(a, b)
    utils.PrintResult("Soustraction", a, b, result)

    result = math.Multiply(a, b)
    utils.PrintResult("Multiplication", a, b, result)

    result = math.Divide(a, b)
    utils.PrintResult("Division", a, b, result)

    // Opérations avancées
    result = math.Power(a, b)
    utils.PrintResult("Puissance", a, b, result)

    result = math.SquareRoot(a)
    utils.PrintSingleResult("Racine carrée", a, result)
}
```

## Bonnes pratiques d'organisation

### 1. Nommage des packages

- Utilisez des noms courts et descriptifs
- Évitez les underscores (utilisez `mathutils` plutôt que `math_utils`)
- Préférez le singulier au pluriel (`user` plutôt que `users`)

### 2. Structure par fonctionnalité

```text
projet/
├── user/          # Tout ce qui concerne les utilisateurs
│   ├── model.go
│   ├── service.go
│   └── handler.go
├── product/       # Tout ce qui concerne les produits
│   ├── model.go
│   └── service.go
└── database/      # Gestion de la base de données
    └── connection.go
```

### 3. Éviter les packages trop gros

Si un package contient plus de 5-7 fichiers, considérez le diviser :

```text
// Au lieu de :
utils/
├── math.go
├── string.go
├── file.go
├── network.go
├── crypto.go
└── json.go

// Préférez :
mathutils/
└── math.go
stringutils/
└── string.go
fileutils/
└── file.go
```

### 4. Package internal

Pour du code privé à votre projet, utilisez un dossier `internal/` :

```text
projet/
├── main.go
├── public/        # Code utilisable par d'autres projets
│   └── api.go
└── internal/      # Code privé, non accessible depuis l'extérieur
    └── config.go
```

## Erreurs communes à éviter

### 1. Imports circulaires

```go
// Package A importe B
// Package B importe A
// ❌ Ceci causera une erreur !
```

### 2. Packages trop petits

```go
// ❌ Évitez d'avoir un package avec une seule fonction simple
package add

func Add(a, b int) int {
    return a + b
}
```

### 3. Mauvais nommage

```go
// ❌ Évitez les noms génériques
package utils
package helpers
package common

// ✅ Préférez des noms spécifiques
package mathutils
package stringhelpers
package configloader
```

## Exercice pratique

Créez un projet "gestion-bibliotheque" avec la structure suivante :

```text
gestion-bibliotheque/
├── go.mod
├── main.go
├── book/
│   └── book.go      # Struct Book et ses méthodes
├── library/
│   └── library.go   # Gestion de la collection de livres
└── display/
    └── display.go   # Fonctions d'affichage
```

**Fonctionnalités à implémenter :**

- Struct `Book` avec titre, auteur, année
- Fonctions pour ajouter/supprimer des livres
- Fonction pour lister tous les livres
- Affichage formaté des résultats

## Résumé

Les packages en Go permettent :

- ✅ D'organiser le code de manière logique
- ✅ De réutiliser du code facilement
- ✅ De collaborer efficacement en équipe
- ✅ De maintenir des projets de grande taille

**Points clés à retenir :**

1. Un dossier = un package
2. Tous les fichiers d'un dossier partagent le même package
3. Utilisez des noms courts et descriptifs
4. Organisez par fonctionnalité, pas par type de fichier
5. Évitez les imports circulaires

Dans la section suivante, nous verrons comment contrôler la visibilité des éléments avec les concepts public/private en Go.

⏭️
