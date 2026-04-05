🔝 Retour au [Sommaire](/SOMMAIRE.md)

# 6-3 : Interface vide et assertions de type

## Introduction

L'**interface vide** (`interface{}` ou `any` depuis Go 1.18) est un concept unique en Go qui peut contenir n'importe quel type de valeur. Les **assertions de type** permettent de récupérer le type concret d'une interface et de travailler avec la valeur originale.

Ces concepts sont essentiels pour comprendre comment Go gère les types dynamiques tout en conservant la sécurité des types.

## L'interface vide

### Qu'est-ce que l'interface vide ?

L'interface vide ne définit aucune méthode, ce qui signifie que **tous les types** l'implémentent automatiquement :

```go
// Anciennes versions de Go
var v interface{}

// Go 1.18+ (recommandé)
var v any
```

### Exemples de base

```go
package main

import "fmt"

func main() {
    // L'interface vide peut contenir n'importe quel type
    var valeur any

    valeur = 42
    fmt.Printf("Valeur: %v, Type: %T\n", valeur, valeur)

    valeur = "Hello, Go!"
    fmt.Printf("Valeur: %v, Type: %T\n", valeur, valeur)

    valeur = []int{1, 2, 3}
    fmt.Printf("Valeur: %v, Type: %T\n", valeur, valeur)

    valeur = map[string]int{"a": 1, "b": 2}
    fmt.Printf("Valeur: %v, Type: %T\n", valeur, valeur)
}
```

**Sortie :**

```text
Valeur: 42, Type: int
Valeur: Hello, Go!, Type: string
Valeur: [1 2 3], Type: []int
Valeur: map[a:1 b:2], Type: map[string]int
```

## Fonctions acceptant n'importe quel type

```go
package main

import "fmt"

// Fonction qui accepte n'importe quel type
func Afficher(valeur any) {
    fmt.Printf("J'ai reçu: %v (type: %T)\n", valeur, valeur)
}

// Fonction avec plusieurs paramètres de types différents
func AfficherPlusieurs(valeurs ...any) {
    fmt.Println("=== Liste des valeurs ===")
    for i, valeur := range valeurs {
        fmt.Printf("[%d] %v (%T)\n", i, valeur, valeur)
    }
}

func main() {
    Afficher(123)
    Afficher("Bonjour")
    Afficher(3.14)

    fmt.Println()
    AfficherPlusieurs(42, "test", true, []string{"a", "b"}, map[string]int{"score": 100})
}
```

## Assertions de type

### Assertion simple

L'assertion de type permet de récupérer le type concret d'une interface :

```go
valeur.(Type)
```

```go
package main

import "fmt"

func main() {
    var valeur any = "Hello, World!"

    // Assertion de type simple
    texte := valeur.(string)
    fmt.Printf("Texte: %s\n", texte)

    // Ceci provoquera un panic car valeur n'est pas un int
    // nombre := valeur.(int) // PANIC!
}
```

### Assertion sécurisée

Pour éviter les panics, utilisez la forme à deux valeurs de retour :

```go
package main

import "fmt"

func TraiterValeur(valeur any) {
    // Assertion sécurisée
    if texte, ok := valeur.(string); ok {
        fmt.Printf("C'est un string: %s\n", texte)
    } else if nombre, ok := valeur.(int); ok {
        fmt.Printf("C'est un int: %d\n", nombre)
    } else if decimal, ok := valeur.(float64); ok {
        fmt.Printf("C'est un float64: %.2f\n", decimal)
    } else {
        fmt.Printf("Type non géré: %T\n", valeur)
    }
}

func main() {
    TraiterValeur("Bonjour")
    TraiterValeur(42)
    TraiterValeur(3.14)
    TraiterValeur([]int{1, 2, 3})
}
```

**Sortie :**

```text
C'est un string: Bonjour
C'est un int: 42
C'est un float64: 3.14
Type non géré: []int
```

## Switch de type

Le `switch` de type est une manière élégante de gérer plusieurs types :

```go
package main

import "fmt"

func AnalyserType(valeur any) {
    switch v := valeur.(type) {
    case string:
        fmt.Printf("String de %d caractères: %q\n", len(v), v)
    case int:
        if v >= 0 {
            fmt.Printf("Entier positif: %d\n", v)
        } else {
            fmt.Printf("Entier négatif: %d\n", v)
        }
    case float64:
        fmt.Printf("Nombre décimal: %.2f\n", v)
    case bool:
        if v {
            fmt.Println("Valeur booléenne: vraie")
        } else {
            fmt.Println("Valeur booléenne: fausse")
        }
    case []int:
        fmt.Printf("Slice d'entiers avec %d éléments: %v\n", len(v), v)
    case map[string]int:
        fmt.Printf("Map avec %d clés: %v\n", len(v), v)
    case nil:
        fmt.Println("Valeur nulle")
    default:
        fmt.Printf("Type non géré: %T\n", v)
    }
}

func main() {
    valeurs := []any{
        "Hello Go",
        42,
        -10,
        3.14159,
        true,
        false,
        []int{1, 2, 3, 4, 5},
        map[string]int{"score": 100, "niveau": 5},
        nil,
        complex(1, 2), // Type non géré
    }

    for i, valeur := range valeurs {
        fmt.Printf("[%d] ", i)
        AnalyserType(valeur)
    }
}
```

## Exemple pratique : Calculatrice universelle

```go
package main

import (
    "fmt"
    "errors"
)

type Calculatrice struct{}

// Addition qui accepte différents types numériques
func (c *Calculatrice) Additionner(a, b any) (any, error) {
    switch va := a.(type) {
    case int:
        switch vb := b.(type) {
        case int:
            return va + vb, nil
        case float64:
            return float64(va) + vb, nil
        default:
            return nil, errors.New("type non supporté pour b")
        }
    case float64:
        switch vb := b.(type) {
        case int:
            return va + float64(vb), nil
        case float64:
            return va + vb, nil
        default:
            return nil, errors.New("type non supporté pour b")
        }
    case string:
        if vb, ok := b.(string); ok {
            return va + vb, nil // Concaténation
        }
        return nil, errors.New("impossible de concaténer avec un non-string")
    default:
        return nil, errors.New("type non supporté pour a")
    }
}

// Fonction pour afficher le résultat de manière lisible
func (c *Calculatrice) AfficherResultat(a, b any) {
    resultat, err := c.Additionner(a, b)
    if err != nil {
        fmt.Printf("Erreur: %s + %s = %v\n", fmt.Sprintf("%v", a), fmt.Sprintf("%v", b), err)
        return
    }

    fmt.Printf("%v + %v = %v\n", a, b, resultat)
}

func main() {
    calc := &Calculatrice{}

    fmt.Println("=== Calculatrice universelle ===")

    // Tests avec différents types
    calc.AfficherResultat(5, 3)              // int + int
    calc.AfficherResultat(5, 3.14)           // int + float64
    calc.AfficherResultat(2.5, 7)            // float64 + int
    calc.AfficherResultat(1.5, 2.5)          // float64 + float64
    calc.AfficherResultat("Hello", " World") // string + string
    calc.AfficherResultat(5, "test")         // Erreur
}
```

## Détection d'interfaces

Vous pouvez aussi vérifier si un type implémente une interface spécifique :

```go
package main

import "fmt"

// Interface personnalisée
type Parleur interface {
    Parler() string
}

// Types qui implémentent Parleur
type Chien struct {
    nom string
}

func (c Chien) Parler() string {
    return "Woof!"
}

type Robot struct {
    id string
}

func (r Robot) Parler() string {
    return "Beep boop"
}

// Fonction qui vérifie si un objet peut parler
func FaireParler(objet any) {
    // Vérifier si l'objet implémente l'interface Parleur
    if parleur, ok := objet.(Parleur); ok {
        fmt.Printf("L'objet dit: %s\n", parleur.Parler())
    } else {
        fmt.Printf("L'objet de type %T ne peut pas parler\n", objet)
    }
}

func main() {
    objets := []any{
        Chien{nom: "Rex"},
        Robot{id: "R2D2"},
        "Une simple string", // Ne peut pas parler
        42,                  // Ne peut pas parler
    }

    for _, objet := range objets {
        FaireParler(objet)
    }
}
```

## Exemple complet : Système de configuration

```go
package main

import (
    "fmt"
    "encoding/json"
)

// Structure de configuration qui peut contenir différents types
type Configuration struct {
    donnees map[string]any
}

func NouvelleConfiguration() *Configuration {
    return &Configuration{
        donnees: make(map[string]any),
    }
}

func (c *Configuration) Definir(cle string, valeur any) {
    c.donnees[cle] = valeur
}

// Méthodes spécialisées pour récupérer des types spécifiques
func (c *Configuration) ObtenirString(cle string) (string, bool) {
    valeur, existe := c.donnees[cle]
    if !existe {
        return "", false
    }

    if str, ok := valeur.(string); ok {
        return str, true
    }
    return "", false
}

func (c *Configuration) ObtenirInt(cle string) (int, bool) {
    valeur, existe := c.donnees[cle]
    if !existe {
        return 0, false
    }

    switch v := valeur.(type) {
    case int:
        return v, true
    case float64: // JSON decode souvent en float64
        return int(v), true
    default:
        return 0, false
    }
}

func (c *Configuration) ObtenirBool(cle string) (bool, bool) {
    valeur, existe := c.donnees[cle]
    if !existe {
        return false, false
    }

    if b, ok := valeur.(bool); ok {
        return b, true
    }
    return false, false
}

func (c *Configuration) ObtenirSliceString(cle string) ([]string, bool) {
    valeur, existe := c.donnees[cle]
    if !existe {
        return nil, false
    }

    // Gérer []interface{} (souvent le cas avec JSON)
    if slice, ok := valeur.([]interface{}); ok {
        result := make([]string, len(slice))
        for i, item := range slice {
            if str, ok := item.(string); ok {
                result[i] = str
            } else {
                return nil, false // Un élément n'est pas un string
            }
        }
        return result, true
    }

    // Gérer []string directement
    if slice, ok := valeur.([]string); ok {
        return slice, true
    }

    return nil, false
}

func (c *Configuration) AfficherTout() {
    fmt.Println("=== Configuration ===")
    for cle, valeur := range c.donnees {
        fmt.Printf("  %s: %v (%T)\n", cle, valeur, valeur)
    }
}

// Charger depuis JSON
func (c *Configuration) ChargerJSON(jsonData string) error {
    var donnees map[string]interface{}
    err := json.Unmarshal([]byte(jsonData), &donnees)
    if err != nil {
        return err
    }

    for cle, valeur := range donnees {
        c.donnees[cle] = valeur
    }
    return nil
}

func main() {
    config := NouvelleConfiguration()

    // Définir des valeurs de différents types
    config.Definir("nom_app", "MonApplication")
    config.Definir("version", "1.2.3")
    config.Definir("port", 8080)
    config.Definir("debug", true)
    config.Definir("features", []string{"auth", "cache", "logging"})

    // Afficher toute la configuration
    config.AfficherTout()

    // Récupérer des valeurs avec vérification de type
    if nom, ok := config.ObtenirString("nom_app"); ok {
        fmt.Printf("\nNom de l'application: %s\n", nom)
    }

    if port, ok := config.ObtenirInt("port"); ok {
        fmt.Printf("Port: %d\n", port)
    }

    if debug, ok := config.ObtenirBool("debug"); ok {
        fmt.Printf("Mode debug: %t\n", debug)
    }

    if features, ok := config.ObtenirSliceString("features"); ok {
        fmt.Printf("Fonctionnalités: %v\n", features)
    }

    // Test avec du JSON
    fmt.Println("\n=== Chargement depuis JSON ===")
    jsonConfig := `{
        "database_url": "postgresql://localhost:5432/mydb",
        "max_connections": 100,
        "timeout": 30.5,
        "ssl_enabled": true,
        "allowed_hosts": ["localhost", "example.com", "api.example.com"]
    }`

    config2 := NouvelleConfiguration()
    err := config2.ChargerJSON(jsonConfig)
    if err != nil {
        fmt.Printf("Erreur de chargement JSON: %v\n", err)
        return
    }

    config2.AfficherTout()

    // Vérifier les types après chargement JSON
    if url, ok := config2.ObtenirString("database_url"); ok {
        fmt.Printf("\nURL de base de données: %s\n", url)
    }

    if maxConn, ok := config2.ObtenirInt("max_connections"); ok {
        fmt.Printf("Connexions max: %d\n", maxConn)
    }

    if hosts, ok := config2.ObtenirSliceString("allowed_hosts"); ok {
        fmt.Printf("Hôtes autorisés: %v\n", hosts)
    }
}
```

## Cas d'usage courants

### 1. Fonctions de logging

```go
func Log(level string, message any) {
    switch v := message.(type) {
    case string:
        fmt.Printf("[%s] %s\n", level, v)
    case error:
        fmt.Printf("[%s] Erreur: %s\n", level, v.Error())
    default:
        fmt.Printf("[%s] %v\n", level, v)
    }
}
```

### 2. Cache générique

```go
type Cache struct {
    donnees map[string]any
}

func (c *Cache) Stocker(cle string, valeur any) {
    c.donnees[cle] = valeur
}

func (c *Cache) Recuperer(cle string) (any, bool) {
    valeur, existe := c.donnees[cle]
    return valeur, existe
}
```

## Pièges à éviter

1. **Surutilisation** : N'utilisez `any` que quand c'est vraiment nécessaire
2. **Perte de sécurité des types** : Vérifiez toujours les types avant utilisation
3. **Performance** : Les assertions de type ont un coût en performance
4. **Lisibilité** : Le code devient moins lisible avec trop d'interfaces vides

## Bonnes pratiques

1. **Préférez les interfaces spécifiques** à `any` quand possible
2. **Utilisez la forme sécurisée** des assertions de type (`valeur, ok := x.(Type)`)
3. **Documentez** quand vous utilisez `any` et pourquoi
4. **Validez toujours** les types avant utilisation
5. **Considérez les génériques** (Go 1.18+) comme alternative à `any`

## Résumé

L'interface vide et les assertions de type sont des outils puissants mais à utiliser avec précaution :

- **Interface vide (`any`)** : Peut contenir n'importe quel type
- **Assertions de type** : Permettent de récupérer le type concret
- **Switch de type** : Méthode élégante pour gérer plusieurs types
- **Forme sécurisée** : Toujours préférer `valeur, ok := x.(Type)`

Ces concepts sont essentiels pour comprendre comment Go gère les types dynamiques tout en préservant la sécurité des types. Dans la prochaine section, nous verrons comment utiliser la composition plutôt que l'héritage en Go.

⏭️
