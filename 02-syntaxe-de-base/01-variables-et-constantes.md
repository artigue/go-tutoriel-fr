🔝 Retour au [Sommaire](/SOMMAIRE.md)

# 2-1 : Variables et constantes

## Introduction

Les variables et constantes sont les éléments fondamentaux de tout programme. Elles permettent de stocker et manipuler des données. Imaginez-les comme des **boîtes étiquetées** où vous rangez vos informations pour les retrouver et les utiliser plus tard.

## Qu'est-ce qu'une variable ?

### Définition simple

Une **variable** est un espace de stockage nommé qui peut contenir une valeur. Cette valeur peut changer au cours de l'exécution du programme (d'où le nom "variable").

### Analogie de la boîte

```html
┌─────────────────┐
│ Boîte étiquetée │
│ "nom"           │
│                 │
│ Contenu: "Alice"│
└─────────────────┘
```

En Go, cela s'écrit :

```go
var nom string = "Alice"
```

## Déclaration de variables en Go

### Méthode 1 : Déclaration explicite complète

**Syntaxe :**

```go
var nomVariable type = valeur
```

**Exemples :**

```go
var nom string = "Alice"
var age int = 25
var taille float64 = 1.75
var estEtudiant bool = true
```

**Décomposition :**

- `var` : mot-clé pour déclarer une variable
- `nom` : nom de la variable (identifiant)
- `string` : type de données
- `=` : opérateur d'assignation
- `"Alice"` : valeur initiale

### Méthode 2 : Déclaration avec inférence de type

Go peut deviner le type à partir de la valeur :

```go
var nom = "Alice"        // Go comprend que c'est un string
var age = 25             // Go comprend que c'est un int
var taille = 1.75        // Go comprend que c'est un float64
var estEtudiant = true   // Go comprend que c'est un bool
```

### Méthode 3 : Déclaration courte (recommandée)

**Syntaxe :**

```go
nomVariable := valeur
```

**Exemples :**

```go
nom := "Alice"
age := 25
taille := 1.75
estEtudiant := true
```

**Avantages :**

- Plus concise
- Type inféré automatiquement
- Syntaxe idiomatique Go

**Important :** `:=` ne fonctionne qu'à l'intérieur de fonctions !

### Méthode 4 : Déclaration sans valeur initiale

```go
var nom string       // Valeur par défaut : ""
var age int          // Valeur par défaut : 0
var taille float64   // Valeur par défaut : 0.0
var estEtudiant bool // Valeur par défaut : false
```

## Valeurs par défaut (valeurs zéro)

Go initialise automatiquement les variables avec des **valeurs zéro** :

| Type | Valeur zéro | Exemple |
|------|-------------|---------|
| `string` | `""` (chaîne vide) | `var nom string` → `""` |
| `int`, `int8`, `int16`, `int32`, `int64` | `0` | `var age int` → `0` |
| `uint`, `uint8`, `uint16`, `uint32`, `uint64` | `0` | `var compteur uint` → `0` |
| `float32`, `float64` | `0.0` | `var prix float64` → `0.0` |
| `bool` | `false` | `var actif bool` → `false` |
| `byte` | `0` | `var octet byte` → `0` |
| `rune` | `0` | `var caractere rune` → `0` |

**Exemple pratique :**

```go
package main

import "fmt"

func main() {
    var nom string
    var age int
    var actif bool

    fmt.Printf("nom: '%s'\n", nom)       // nom: ''
    fmt.Printf("age: %d\n", age)         // age: 0
    fmt.Printf("actif: %t\n", actif)     // actif: false
}
```

## Assignation et réassignation

### Première assignation

```go
var nom string    // Déclaration
nom = "Alice"     // Première assignation
```

### Réassignation

```go
nom = "Bob"       // Changement de valeur
nom = "Charlie"   // Nouveau changement
```

### Exemple complet

```go
package main

import "fmt"

func main() {
    // Déclaration et assignation
    var nom string = "Alice"
    fmt.Println("Nom initial:", nom)

    // Réassignation
    nom = "Bob"
    fmt.Println("Nouveau nom:", nom)

    // Avec la syntaxe courte
    age := 25
    fmt.Println("Age initial:", age)

    age = 26
    fmt.Println("Nouvel age:", age)
}
```

## Déclaration multiple

### Variables de même type

```go
var nom, prenom string
var x, y, z int

// Avec valeurs initiales
var nom, prenom string = "Dupont", "Alice"
var x, y, z int = 1, 2, 3
```

### Variables de types différents

```go
var (
    nom     string = "Alice"
    age     int    = 25
    taille  float64 = 1.75
    actif   bool   = true
)
```

### Avec la syntaxe courte

```go
nom, age := "Alice", 25
x, y, z := 1, 2, 3
```

## Portée des variables (Scope)

### Variables globales

Déclarées en dehors des fonctions, accessibles partout dans le package :

```go
package main

import "fmt"

var nomGlobal string = "Variable globale"  // Accessible partout

func main() {
    fmt.Println(nomGlobal)  // Fonctionne
    afficherNom()
}

func afficherNom() {
    fmt.Println(nomGlobal)  // Fonctionne aussi
}
```

### Variables locales

Déclarées à l'intérieur d'une fonction, accessibles uniquement dans cette fonction :

```go
func main() {
    var nomLocal string = "Variable locale"  // Accessible uniquement dans main()
    fmt.Println(nomLocal)
}

func autrefFonction() {
    // fmt.Println(nomLocal)  // ❌ ERREUR : nomLocal n'existe pas ici
}
```

### Variables de bloc

Déclarées dans un bloc (entre accolades), accessibles uniquement dans ce bloc :

```go
func main() {
    if true {
        var nomBloc string = "Variable de bloc"
        fmt.Println(nomBloc)  // ✅ Fonctionne
    }
    // fmt.Println(nomBloc)  // ❌ ERREUR : nomBloc n'existe plus
}
```

## Constantes

### Qu'est-ce qu'une constante ?

Une **constante** est une valeur qui ne peut **jamais changer** après sa déclaration. C'est comme une boîte scellée définitivement.

### Déclaration de constantes

**Syntaxe :**

```go
const nomConstante = valeur
```

**Exemples :**

```go
const pi = 3.14159
const nomSite = "MonSite.com"
const maxUtilisateurs = 1000
const estActif = true
```

### Constantes typées vs non-typées

**Constantes non-typées (recommandé) :**

```go
const pi = 3.14159          // Type flexible
const message = "Hello"     // Type flexible
```

**Constantes typées :**

```go
const pi float64 = 3.14159
const message string = "Hello"
```

### Avantages des constantes non-typées

```go
const nombre = 42

var entier8 int8 = nombre      // ✅ Fonctionne
var entier32 int32 = nombre    // ✅ Fonctionne
var entier64 int64 = nombre    // ✅ Fonctionne
var decimal float64 = nombre   // ✅ Fonctionne
```

### Déclaration multiple de constantes

```go
const (
    pi           = 3.14159
    e            = 2.71828
    vitesseLumiere = 299792458  // m/s
)
```

### Constantes avec iota

`iota` est un générateur de constantes qui s'incrémente automatiquement :

```go
const (
    Dimanche = iota  // 0
    Lundi           // 1
    Mardi           // 2
    Mercredi        // 3
    Jeudi           // 4
    Vendredi        // 5
    Samedi          // 6
)
```

**Utilisation pratique :**

```go
const (
    TaillePetite = iota  // 0
    TailleMoyenne       // 1
    TailleGrande        // 2
)

const (
    Rouge = iota + 1    // 1
    Vert               // 2
    Bleu               // 3
)
```

## Règles de nommage

### Conventions Go

**Variables et constantes locales (privées) :**

```go
var nom string              // ✅ minuscule
var ageUtilisateur int      // ✅ camelCase
var compteurVisites int     // ✅ descriptif
```

**Variables et constantes exportées (publiques) :**

```go
var Nom string              // ✅ Majuscule (exportée)
var AgeUtilisateur int      // ✅ PascalCase
var CompteurVisites int     // ✅ descriptif
```

### Bonnes pratiques

**✅ Noms descriptifs :**

```go
var utilisateurConnecte bool    // Clair
var nombreTentatives int        // Explicite
var messageErreur string        // Compréhensible
```

**❌ Noms cryptiques :**

```go
var u bool          // Que représente 'u' ?
var n int           // Trop vague
var msg string      // Abréviation peu claire
```

**✅ Constantes en majuscules (optionnel) :**

```go
const PI = 3.14159
const MAX_UTILISATEURS = 1000
const VERSION = "1.0.0"
```

## Erreurs courantes

### 1. Confusion entre := et =

```go
// ❌ Erreur : variable pas encore déclarée
nom = "Alice"

// ✅ Correct : déclaration + assignation
nom := "Alice"

// OU déclaration puis assignation
var nom string
nom = "Alice"
```

### 2. Redéclaration avec :=

```go
nom := "Alice"
// nom := "Bob"    // ❌ ERREUR : nom déjà déclaré

nom = "Bob"        // ✅ Correct : réassignation
```

### 3. Utilisation de := en dehors d'une fonction

```go
// ❌ ERREUR : := seulement dans les fonctions
// nom := "Alice"

// ✅ Correct : au niveau du package
var nom = "Alice"

func main() {
    // ✅ Correct : dans une fonction
    age := 25
}
```

### 4. Modification d'une constante

```go
const pi = 3.14159
// pi = 3.14    // ❌ ERREUR : impossible de modifier une constante
```

### 5. Variable déclarée mais non utilisée

```go
func main() {
    var nom string = "Alice"  // ❌ ERREUR : variable non utilisée
    // Go refuse de compiler si une variable n'est pas utilisée
}

func main() {
    var nom string = "Alice"
    fmt.Println(nom)          // ✅ Correct : variable utilisée
}
```

## Exemples pratiques

### Exemple 1 : Calculatrice simple

```go
package main

import "fmt"

func main() {
    // Variables pour les calculs
    a := 10.0
    b := 3.0

    // Calculs
    addition := a + b
    soustraction := a - b
    multiplication := a * b
    division := a / b

    // Affichage des résultats
    fmt.Printf("%.1f + %.1f = %.1f\n", a, b, addition)
    fmt.Printf("%.1f - %.1f = %.1f\n", a, b, soustraction)
    fmt.Printf("%.1f * %.1f = %.1f\n", a, b, multiplication)
    fmt.Printf("%.1f / %.1f = %.2f\n", a, b, division)
}
```

### Exemple 2 : Informations utilisateur

```go
package main

import "fmt"

// Constantes globales
const (
    NomApplication = "MonApp"
    Version        = "1.0.0"
    MaxUtilisateurs = 100
)

func main() {
    // Informations utilisateur
    nom := "Alice"
    prenom := "Dupont"
    age := 25
    estPremium := true

    // Variables calculées
    nomComplet := prenom + " " + nom
    peutUtiliserApp := age >= 18

    // Affichage
    fmt.Printf("=== %s v%s ===\n", NomApplication, Version)
    fmt.Printf("Utilisateur: %s\n", nomComplet)
    fmt.Printf("Age: %d ans\n", age)
    fmt.Printf("Statut Premium: %t\n", estPremium)
    fmt.Printf("Peut utiliser l'app: %t\n", peutUtiliserApp)
    fmt.Printf("Utilisateurs max: %d\n", MaxUtilisateurs)
}
```

### Exemple 3 : Gestion de stock

```go
package main

import "fmt"

func main() {
    // Stock initial
    var stockPommes int = 50
    var stockOranges int = 30
    var stockBananes int = 20

    // Ventes de la journée
    ventesPommes := 12
    ventesOranges := 8
    ventesBananes := 15

    // Calcul du stock restant
    stockPommes = stockPommes - ventesPommes
    stockOranges = stockOranges - ventesOranges
    stockBananes = stockBananes - ventesBananes

    // Stock total
    stockTotal := stockPommes + stockOranges + stockBananes

    // Affichage du rapport
    fmt.Println("=== RAPPORT DE STOCK ===")
    fmt.Printf("Pommes restantes: %d\n", stockPommes)
    fmt.Printf("Oranges restantes: %d\n", stockOranges)
    fmt.Printf("Bananes restantes: %d\n", stockBananes)
    fmt.Printf("Stock total: %d fruits\n", stockTotal)

    // Alertes
    const SEUIL_ALERTE = 10

    if stockPommes < SEUIL_ALERTE {
        fmt.Println("⚠️  ALERTE: Stock de pommes faible!")
    }
    if stockOranges < SEUIL_ALERTE {
        fmt.Println("⚠️  ALERTE: Stock d'oranges faible!")
    }
    if stockBananes < SEUIL_ALERTE {
        fmt.Println("⚠️  ALERTE: Stock de bananes faible!")
    }
}
```

## Exercices pratiques

### Exercice 1 : Informations personnelles

Créez un programme qui :

1. Déclare des variables pour votre nom, prénom, âge et ville
2. Affiche ces informations de manière formatée

### Exercice 2 : Calculateur d'âge

Créez un programme qui :

1. Déclare une constante pour l'année actuelle (2025)
2. Déclare une variable pour votre année de naissance
3. Calcule et affiche votre âge

### Exercice 3 : Conversion de température

Créez un programme qui :

1. Déclare une variable température en Celsius
2. Calcule l'équivalent en Fahrenheit (F = C × 9/5 + 32)
3. Affiche les deux températures

### Exercice 4 : Gestion de budget

Créez un programme qui :

1. Déclare des variables pour vos revenus et différentes dépenses
2. Calcule le total des dépenses
3. Calcule le solde restant
4. Affiche un résumé du budget

### Solutions

**Exercice 1 :**

```go
package main

import "fmt"

func main() {
    nom := "Dupont"
    prenom := "Alice"
    age := 25
    ville := "Paris"

    fmt.Printf("Nom: %s %s\n", prenom, nom)
    fmt.Printf("Age: %d ans\n", age)
    fmt.Printf("Ville: %s\n", ville)
}
```

**Exercice 2 :**

```go
package main

import "fmt"

func main() {
    const anneeActuelle = 2025
    var anneeNaissance int = 1998

    age := anneeActuelle - anneeNaissance

    fmt.Printf("Année de naissance: %d\n", anneeNaissance)
    fmt.Printf("Année actuelle: %d\n", anneeActuelle)
    fmt.Printf("Votre âge: %d ans\n", age)
}
```

**Exercice 3 :**

```go
package main

import "fmt"

func main() {
    celsius := 25.0
    fahrenheit := celsius*9/5 + 32

    fmt.Printf("Température en Celsius: %.1f°C\n", celsius)
    fmt.Printf("Température en Fahrenheit: %.1f°F\n", fahrenheit)
}
```

**Exercice 4 :**

```go
package main

import "fmt"

func main() {
    // Revenus
    revenus := 3000.0

    // Dépenses
    loyer := 800.0
    nourriture := 400.0
    transport := 150.0
    loisirs := 200.0

    // Calculs
    totalDepenses := loyer + nourriture + transport + loisirs
    solde := revenus - totalDepenses

    // Affichage
    fmt.Println("=== BUDGET MENSUEL ===")
    fmt.Printf("Revenus: %.2f€\n", revenus)
    fmt.Printf("Loyer: %.2f€\n", loyer)
    fmt.Printf("Nourriture: %.2f€\n", nourriture)
    fmt.Printf("Transport: %.2f€\n", transport)
    fmt.Printf("Loisirs: %.2f€\n", loisirs)
    fmt.Println("---")
    fmt.Printf("Total dépenses: %.2f€\n", totalDepenses)
    fmt.Printf("Solde restant: %.2f€\n", solde)
}
```

## Récapitulatif

**Ce que vous avez appris :**

- ✅ Différentes façons de déclarer des variables
- ✅ Valeurs par défaut (valeurs zéro) en Go
- ✅ Portée des variables (globale, locale, bloc)
- ✅ Constantes et leur utilisation
- ✅ Règles de nommage et bonnes pratiques
- ✅ Erreurs courantes à éviter

**Points clés à retenir :**

1. **`:=`** : déclaration courte (seulement dans les fonctions)
2. **`var`** : déclaration explicite (partout)
3. **`const`** : valeurs qui ne changent jamais
4. **Valeurs zéro** : Go initialise automatiquement
5. **Nommage** : descriptif et suivant les conventions

**Syntaxes importantes :**

```go
var nom string = "Alice"    // Déclaration complète
var nom = "Alice"           // Avec inférence de type
nom := "Alice"              // Déclaration courte (recommandée)
const PI = 3.14159          // Constante
```

---

*Excellent ! Vous maîtrisez maintenant les variables et constantes. Dans la prochaine section, nous explorerons les différents types de données primitifs disponibles en Go !*

⏭️
