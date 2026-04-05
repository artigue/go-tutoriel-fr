🔝 Retour au [Sommaire](/SOMMAIRE.md)

# 2-2 : Types de données primitifs

## Introduction

Les types de données primitifs sont les briques de base pour stocker l'information dans vos programmes. Imaginez-les comme différents types de **contenants** : une bouteille pour les liquides, une boîte pour les objets solides, un tiroir pour les documents. Chaque type est optimisé pour un usage spécifique.

## Qu'est-ce qu'un type de données ?

### Définition simple

Un **type de données** définit :

- **Quelle sorte de valeur** peut être stockée
- **Combien de mémoire** est nécessaire
- **Quelles opérations** sont possibles

### Analogie des contenants

```text
┌─────────────┐  ┌─────────────┐  ┌─────────────┐
│    int      │  │   string    │  │    bool     │
│  (nombre)   │  │   (texte)   │  │ (vrai/faux) │
│             │  │             │  │             │
│     42      │  │  "Hello"    │  │    true     │
└─────────────┘  └─────────────┘  └─────────────┘
```

## Vue d'ensemble des types primitifs

Go propose plusieurs catégories de types primitifs :

| Catégorie | Types | Usage |
|-----------|--------|--------|
| **Entiers** | `int`, `int8`, `int16`, `int32`, `int64` | Nombres entiers |
| **Entiers non signés** | `uint`, `uint8`, `uint16`, `uint32`, `uint64` | Nombres positifs uniquement |
| **Décimaux** | `float32`, `float64` | Nombres à virgule |
| **Texte** | `string` | Chaînes de caractères |
| **Booléens** | `bool` | Vrai ou faux |
| **Caractères** | `byte`, `rune` | Caractères individuels |

## Types d'entiers (int)

### Types d'entiers signés

**Peuvent stocker des nombres positifs ET négatifs :**

| Type | Taille | Plage de valeurs | Exemple d'usage |
|------|--------|------------------|-----------------|
| `int8` | 8 bits | -128 à 127 | Température en Celsius |
| `int16` | 16 bits | -32,768 à 32,767 | Altitude en mètres |
| `int32` | 32 bits | -2,147,483,648 à 2,147,483,647 | Population d'une ville |
| `int64` | 64 bits | -9,223,372,036,854,775,808 à 9,223,372,036,854,775,807 | Timestamp Unix |
| `int` | 32 ou 64 bits* | Dépend de l'architecture | Usage général |

*`int` fait 32 bits sur les systèmes 32 bits, 64 bits sur les systèmes 64 bits

### Types d'entiers non signés

**Peuvent stocker uniquement des nombres positifs (0 et plus) :**

| Type | Taille | Plage de valeurs | Exemple d'usage |
|------|--------|------------------|-----------------|
| `uint8` | 8 bits | 0 à 255 | Valeur RGB d'une couleur |
| `uint16` | 16 bits | 0 à 65,535 | Port réseau |
| `uint32` | 32 bits | 0 à 4,294,967,295 | ID d'utilisateur |
| `uint64` | 64 bits | 0 à 18,446,744,073,709,551,615 | Taille de fichier en octets |
| `uint` | 32 ou 64 bits | Dépend de l'architecture | Index de tableau |

### Exemples pratiques

```go
package main

import "fmt"

func main() {
    // Entiers signés
    var temperature int8 = -15        // Température en hiver
    var altitude int16 = 8848         // Mont Everest en mètres
    var population int32 = 2161000    // Population de Paris
    var timestamp int64 = 1640995200  // 1er janvier 2022

    // Entiers non signés
    var rouge uint8 = 255             // Composant rouge RGB
    var port uint16 = 8080            // Port web
    var userId uint32 = 123456        // ID utilisateur
    var taillefichier uint64 = 1073741824  // 1 GB en octets

    // Type int général (recommandé pour l'usage courant)
    var age int = 25
    var compteur int = 0

    fmt.Printf("Température: %d°C\n", temperature)
    fmt.Printf("Altitude: %d m\n", altitude)
    fmt.Printf("Population: %d habitants\n", population)
    fmt.Printf("Port: %d\n", port)
    fmt.Printf("Age: %d ans\n", age)
}
```

### Débordement d'entiers (Overflow)

**Attention !** Les entiers ont des limites :

```go
package main

import "fmt"

func main() {
    var petit int8 = 127      // Valeur maximale pour int8
    fmt.Printf("Avant: %d\n", petit)

    petit = petit + 1         // Débordement !
    fmt.Printf("Après: %d\n", petit)  // Résultat: -128 (inattendu!)
}
```

**Conseil :** Utilisez `int` par défaut, sauf si vous avez une raison spécifique d'utiliser un autre type.

## Types de nombres décimaux (float)

### Types disponibles

| Type | Taille | Précision | Plage approximative |
|------|--------|-----------|-------------------|
| `float32` | 32 bits | ~7 chiffres décimaux | ±1.18e-38 à ±3.4e+38 |
| `float64` | 64 bits | ~15 chiffres décimaux | ±2.23e-308 à ±1.80e+308 |

### Exemples pratiques

```go
package main

import "fmt"

func main() {
    // float32 : précision simple
    var prix float32 = 19.99
    var taux float32 = 0.075

    // float64 : précision double (recommandé)
    var pi float64 = 3.141592653589793
    var vitesse float64 = 299792458.0  // Vitesse de la lumière en m/s

    // Calculs
    total := prix * (1 + taux)
    circonference := 2 * pi * 10  // Rayon = 10

    fmt.Printf("Prix TTC: %.2f€\n", total)
    fmt.Printf("Circonférence: %.2f\n", circonference)

    // Notation scientifique
    var tresPetit float64 = 1.23e-10
    var tresGrand float64 = 6.02e23   // Nombre d'Avogadro

    fmt.Printf("Très petit: %e\n", tresPetit)
    fmt.Printf("Très grand: %e\n", tresGrand)
}
```

### Précision des nombres flottants

**Important :** Les nombres flottants ne sont pas toujours exacts :

```go
package main

import "fmt"

func main() {
    var a float64 = 0.1
    var b float64 = 0.2
    var c float64 = a + b

    fmt.Printf("0.1 + 0.2 = %.17f\n", c)  // Résultat: 0.30000000000000004

    // Pour les calculs monétaires, utilisez des entiers !
    var prixCentimes int = 10 + 20  // 30 centimes = 0.30€
    fmt.Printf("Prix exact: %.2f€\n", float64(prixCentimes)/100)
}
```

## Type chaîne de caractères (string)

### Définition et caractéristiques

Le type `string` stocke du **texte** (séquence de caractères).

**Caractéristiques importantes :**

- Les strings sont **immutables** (ne peuvent pas être modifiées)
- Encodage **UTF-8** par défaut
- Délimitées par des **guillemets doubles** `"`

### Déclaration et utilisation

```go
package main

import "fmt"

func main() {
    // Déclarations simples
    var nom string = "Alice"
    var message string = "Bonjour tout le monde!"

    // Avec la syntaxe courte
    prenom := "Bob"
    ville := "Paris"

    // String vide
    var description string  // Valeur par défaut: ""

    // String multiligne
    var poeme string = `
    Les sanglots longs
    Des violons
    De l'automne
    `

    fmt.Printf("Nom: %s\n", nom)
    fmt.Printf("Message: %s\n", message)
    fmt.Printf("Ville: %s\n", ville)
    fmt.Printf("Poème:%s\n", poeme)
}
```

### Opérations sur les strings

**Concaténation (assemblage) :**

```go
package main

import "fmt"

func main() {
    prenom := "Alice"
    nom := "Dupont"

    // Méthode 1 : avec +
    nomComplet := prenom + " " + nom

    // Méthode 2 : avec fmt.Sprintf
    presentation := fmt.Sprintf("Je m'appelle %s %s", prenom, nom)

    fmt.Println(nomComplet)      // Alice Dupont
    fmt.Println(presentation)    // Je m'appelle Alice Dupont
}
```

**Longueur d'une string :**

```go
package main

import "fmt"

func main() {
    message := "Hello, 世界"  // Mélange anglais/chinois

    fmt.Printf("Longueur en octets: %d\n", len(message))        // 13
    fmt.Printf("Nombre de caractères: %d\n", len([]rune(message)))  // 9
}
```

### Caractères d'échappement

**Caractères spéciaux dans les strings :**

| Séquence | Signification | Exemple |
|----------|---------------|---------|
| `\n` | Nouvelle ligne | `"Ligne 1\nLigne 2"` |
| `\t` | Tabulation | `"Nom:\tAlice"` |
| `\"` | Guillemet double | `"Il a dit \"Bonjour\""` |
| `\\` | Antislash | `"Chemin: C:\\Users"` |
| `\r` | Retour chariot | `"Windows\r\n"` |

```go
package main

import "fmt"

func main() {
    // Avec échappement
    citation := "Einstein a dit: \"L'imagination est plus importante que la connaissance.\""

    // Avec backticks (pas d'échappement nécessaire)
    chemin := `C:\Users\Alice\Documents\fichier.txt`

    // Multiligne avec backticks
    html := `
    <html>
        <body>
            <h1>Titre</h1>
        </body>
    </html>
    `

    fmt.Println(citation)
    fmt.Println(chemin)
    fmt.Println(html)
}
```

## Type booléen (bool)

### Définition

Le type `bool` ne peut avoir que **deux valeurs** :

- `true` (vrai)
- `false` (faux)

### Utilisation pratique

```go
package main

import "fmt"

func main() {
    // Déclarations simples
    var estConnecte bool = true
    var aPaye bool = false

    // Avec la syntaxe courte
    estAdmin := false
    estActif := true

    // Valeur par défaut
    var estVerifie bool  // false par défaut

    // Utilisation dans les conditions
    if estConnecte {
        fmt.Println("Utilisateur connecté")
    }

    if !aPaye {  // ! signifie "non" ou "pas"
        fmt.Println("Paiement en attente")
    }

    // Combinaisons logiques
    peutAcceder := estConnecte && estActif  // ET logique
    doitVerifier := !estVerifie || !aPaye   // OU logique

    fmt.Printf("Peut accéder: %t\n", peutAcceder)
    fmt.Printf("Doit vérifier: %t\n", doitVerifier)
}
```

### Conversion vers bool

```go
package main

import "fmt"

func main() {
    // Les valeurs "zéro" sont considérées comme false
    age := 0
    nom := ""

    estMajeur := age >= 18
    aNom := nom != ""

    fmt.Printf("Est majeur: %t\n", estMajeur)    // false
    fmt.Printf("A un nom: %t\n", aNom)           // false
}
```

## Types de caractères

### byte (uint8)

**Usage :** Stocke un octet (0-255), souvent utilisé pour les caractères ASCII.

```go
package main

import "fmt"

func main() {
    var lettre byte = 'A'        // Caractère ASCII
    var valeur byte = 65         // Même chose (65 = code ASCII de 'A')

    fmt.Printf("Lettre: %c\n", lettre)     // A
    fmt.Printf("Valeur: %d\n", lettre)     // 65
    fmt.Printf("Identique: %t\n", lettre == valeur)  // true
}
```

### rune (int32)

**Usage :** Stocke un caractère Unicode (peut représenter n'importe quel caractère mondial).

```go
package main

import "fmt"

func main() {
    var emoji rune = '😀'        // Emoji
    var chinois rune = '中'       // Caractère chinois
    var francais rune = 'é'      // Caractère français avec accent

    fmt.Printf("Emoji: %c (code: %d)\n", emoji, emoji)
    fmt.Printf("Chinois: %c (code: %d)\n", chinois, chinois)
    fmt.Printf("Français: %c (code: %d)\n", francais, francais)

    // Parcourir une string caractère par caractère
    texte := "Héllo 世界 😀"
    for _, char := range texte {
        fmt.Printf("%c ", char)
    }
    fmt.Println()
}
```

## Conversion entre types

### Conversions explicites

Go **ne fait jamais** de conversions automatiques. Vous devez être explicite :

```go
package main

import "fmt"

func main() {
    // Conversions entre entiers
    var a int = 42
    var b int64 = int64(a)      // Conversion explicite requise
    var c int8 = int8(a)        // Attention au débordement !

    // Conversions entre entiers et flottants
    var entier int = 42
    var decimal float64 = float64(entier)
    var retourEntier int = int(decimal)

    // Conversions avec strings
    var nombre int = 123
    var texte string = fmt.Sprintf("%d", nombre)  // int vers string

    fmt.Printf("a: %d, b: %d, c: %d\n", a, b, c)
    fmt.Printf("Entier: %d, Décimal: %.1f\n", entier, decimal)
    fmt.Printf("Nombre en texte: %s\n", texte)
}
```

### Conversions avancées avec le package strconv

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    // String vers nombre
    texteNombre := "123"
    nombre, err := strconv.Atoi(texteNombre)
    if err != nil {
        fmt.Println("Erreur de conversion:", err)
    } else {
        fmt.Printf("Nombre: %d\n", nombre)
    }

    // Nombre vers string
    valeur := 456
    texte := strconv.Itoa(valeur)
    fmt.Printf("Texte: %s\n", texte)

    // Float vers string
    pi := 3.14159
    piTexte := strconv.FormatFloat(pi, 'f', 2, 64)
    fmt.Printf("Pi: %s\n", piTexte)  // 3.14
}
```

## Comment choisir le bon type ?

### Guide de sélection

**Pour les nombres entiers :**

```go
var age int              // ✅ Usage général
var temperature int8     // ✅ Si vous savez que les valeurs sont petites
var population int64     // ✅ Si vous savez que les valeurs sont grandes
var couleurRouge uint8   // ✅ Si vous voulez seulement des positifs (0-255)
```

**Pour les nombres décimaux :**

```go
var prix float64         // ✅ Recommandé par défaut
var coordonnee float32   // ✅ Si la précision réduite suffit et vous voulez économiser la mémoire
```

**Pour le texte :**

```go
var nom string           // ✅ Pour tout texte
var initiale rune        // ✅ Pour un seul caractère Unicode
var ascii byte           // ✅ Pour un seul caractère ASCII
```

**Pour les valeurs vrai/faux :**

```go
var estValide bool       // ✅ Toujours bool pour les booléens
```

### Exemples d'application

```go
package main

import "fmt"

func main() {
    // Profil utilisateur
    var (
        nom          string  = "Alice Dupont"
        age          int     = 25
        taille       float64 = 1.68
        estEtudiant  bool    = true
        noteMoyenne  float64 = 15.7
        nbCours      uint8   = 6
    )

    // Affichage du profil
    fmt.Println("=== PROFIL UTILISATEUR ===")
    fmt.Printf("Nom: %s\n", nom)
    fmt.Printf("Age: %d ans\n", age)
    fmt.Printf("Taille: %.2f m\n", taille)
    fmt.Printf("Étudiant: %t\n", estEtudiant)
    fmt.Printf("Note moyenne: %.1f/20\n", noteMoyenne)
    fmt.Printf("Nombre de cours: %d\n", nbCours)

    // Calculs
    tailleEnCm := taille * 100
    estMajeur := age >= 18

    fmt.Printf("Taille en cm: %.0f cm\n", tailleEnCm)
    fmt.Printf("Est majeur: %t\n", estMajeur)
}
```

## Erreurs courantes

### 1. Confusion entre types numériques

```go
// ❌ ERREUR : types incompatibles
var a int = 42
var b int64 = 42
// var c int = a + b  // ERREUR !

// ✅ CORRECT : conversion explicite
var c int = a + int(b)
```

### 2. Débordement d'entiers

```go
// ❌ ATTENTION : débordement
var petit int8 = 200  // ERREUR : 200 > 127 (max pour int8)

// ✅ CORRECT : choisir le bon type
var nombre int = 200
```

### 3. Précision des flottants

```go
// ❌ PROBLÈME : imprécision
var prix float64 = 0.1 + 0.2  // 0.30000000000000004

// ✅ SOLUTION : arithmétique entière pour l'argent
var prixCentimes int = 10 + 20  // 30 centimes
var prixEuros float64 = float64(prixCentimes) / 100
```

### 4. Modification de string

```go
// ❌ ERREUR : strings sont immutables
var nom string = "Alice"
// nom[0] = 'B'  // ERREUR !

// ✅ CORRECT : créer une nouvelle string
nom = "B" + nom[1:]  // Résultat: "Blice"
```

## Exercices pratiques

### Exercice 1 : Calculateur d'IMC

Créez un programme qui :

1. Déclare des variables pour le poids (kg) et la taille (m)
2. Calcule l'IMC (poids / taille²)
3. Affiche le résultat avec 2 décimales

### Exercice 2 : Profil complet

Créez un programme avec :

1. Informations personnelles (nom, âge, etc.)
2. Préférences (couleur favorite, sport, etc.)
3. Statistiques (nombre d'amis, note moyenne, etc.)
4. Affichez tout de manière formatée

### Exercice 3 : Conversions

Créez un programme qui :

1. Déclare une température en Celsius (float64)
2. Convertit en Fahrenheit et Kelvin
3. Affiche les trois températures avec des précisions appropriées

### Exercice 4 : Analyse de texte

Créez un programme qui :

1. Déclare une phrase
2. Compte le nombre de caractères
3. Extrait et affiche le premier et dernier caractère
4. Vérifie si la phrase contient certains mots

### Solutions

**Exercice 1 :**

```go
package main

import "fmt"

func main() {
    var poids float64 = 70.5  // kg
    var taille float64 = 1.75 // m

    imc := poids / (taille * taille)

    fmt.Printf("Poids: %.1f kg\n", poids)
    fmt.Printf("Taille: %.2f m\n", taille)
    fmt.Printf("IMC: %.2f\n", imc)

    // Interprétation
    var interpretation string
    if imc < 18.5 {
        interpretation = "Insuffisance pondérale"
    } else if imc < 25 {
        interpretation = "Poids normal"
    } else {
        interpretation = "Surpoids"
    }

    fmt.Printf("Interprétation: %s\n", interpretation)
}
```

**Exercice 2 :**

```go
package main

import "fmt"

func main() {
    // Informations personnelles
    var nom string = "Dupont"
    var prenom string = "Alice"
    var age uint8 = 25
    var taille float32 = 1.68

    // Préférences
    var couleurFavorite string = "Bleu"
    var sportPrefere string = "Tennis"
    var aimeChocolat bool = true

    // Statistiques
    var nbAmis uint16 = 150
    var noteMoyenne float64 = 15.7
    var nbLivresLus uint8 = 24

    // Affichage
    fmt.Println("=== PROFIL COMPLET ===")
    fmt.Printf("Nom: %s %s\n", prenom, nom)
    fmt.Printf("Age: %d ans\n", age)
    fmt.Printf("Taille: %.2f m\n", taille)
    fmt.Println()
    fmt.Println("=== PRÉFÉRENCES ===")
    fmt.Printf("Couleur favorite: %s\n", couleurFavorite)
    fmt.Printf("Sport préféré: %s\n", sportPrefere)
    fmt.Printf("Aime le chocolat: %t\n", aimeChocolat)
    fmt.Println()
    fmt.Println("=== STATISTIQUES ===")
    fmt.Printf("Nombre d'amis: %d\n", nbAmis)
    fmt.Printf("Note moyenne: %.1f/20\n", noteMoyenne)
    fmt.Printf("Livres lus cette année: %d\n", nbLivresLus)
}
```

**Exercice 3 :**

```go
package main

import "fmt"

func main() {
    var celsius float64 = 25.5

    // Conversions
    fahrenheit := celsius*9/5 + 32
    kelvin := celsius + 273.15

    // Affichage
    fmt.Printf("Température en Celsius: %.1f°C\n", celsius)
    fmt.Printf("Température en Fahrenheit: %.1f°F\n", fahrenheit)
    fmt.Printf("Température en Kelvin: %.2f K\n", kelvin)
}
```

**Exercice 4 :**

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    phrase := "Go est un langage fantastique pour apprendre la programmation"

    // Analyse
    nbCaracteres := len(phrase)
    runes := []rune(phrase)
    premierChar := runes[0]
    dernierChar := runes[len(runes)-1]

    // Recherches
    contientGo := strings.Contains(phrase, "Go")
    contientPython := strings.Contains(phrase, "Python")

    // Affichage
    fmt.Printf("Phrase: %s\n", phrase)
    fmt.Printf("Nombre de caractères: %d\n", nbCaracteres)
    fmt.Printf("Premier caractère: %c\n", premierChar)
    fmt.Printf("Dernier caractère: %c\n", dernierChar)
    fmt.Printf("Contient 'Go': %t\n", contientGo)
    fmt.Printf("Contient 'Python': %t\n", contientPython)
}
```

## Récapitulatif

**Ce que vous avez appris :**

- ✅ Types d'entiers (signés et non signés)
- ✅ Types de nombres décimaux (float32, float64)
- ✅ Type chaîne de caractères (string)
- ✅ Type booléen (bool)
- ✅ Types de caractères (byte, rune)
- ✅ Conversions entre types
- ✅ Guide pour choisir le bon type

**Points clés à retenir :**

1. **Go est strict** : pas de conversions automatiques
2. **int et float64** : choix par défaut pour la plupart des cas
3. **string sont immutables** : ne peuvent pas être modifiées
4. **Attention aux débordements** : choisir la bonne taille
5. **UTF-8 par défaut** : support international excellent

**Types recommandés pour débuter :**

```go
var nombre int           // Entiers
var decimal float64      // Nombres à virgule
var texte string         // Texte
var vrai bool           // Vrai/faux
```

---

*Parfait ! Vous maîtrisez maintenant les types de données primitifs. Dans la prochaine section, nous découvrirons les opérateurs pour manipuler et comparer ces données !*

⏭️
