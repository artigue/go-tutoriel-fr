🔝 Retour au [Sommaire](/SOMMAIRE.md)

# 8-4 : Documentation avec godoc

## Qu'est-ce que godoc ?

**Godoc** est l'outil officiel de Go pour générer automatiquement de la documentation à partir de votre code source. Il transforme vos commentaires en une belle documentation web, similaire à la documentation officielle de Go que vous consultez sur [pkg.go.dev](https://pkg.go.dev).

### Pourquoi documenter son code ?

```go
// ❌ Code sans documentation
func Calculate(x, y int, op string) int {
    switch op {
    case "+":
        return x + y
    case "*":
        return x * y
    default:
        return 0
    }
}

// ✅ Code bien documenté
// Calculate effectue une opération mathématique entre deux nombres.
//
// Paramètres :
//   - x : premier nombre
//   - y : deuxième nombre
//   - op : opération ("+", "*")
//
// Retourne le résultat de l'opération ou 0 si l'opération n'est pas supportée.
//
// Exemple :
//   result := Calculate(5, 3, "+") // result = 8
func Calculate(x, y int, op string) int {
    switch op {
    case "+":
        return x + y
    case "*":
        return x * y
    default:
        return 0
    }
}
```

## Installation et utilisation de godoc

### Installation

```bash
# Installer godoc (si pas déjà installé)
go install golang.org/x/tools/cmd/godoc@latest

# Vérifier l'installation
godoc -h
```

### Génération de documentation

```bash
# Serveur local de documentation
godoc -http=:6060

# Ouvrir http://localhost:6060 dans votre navigateur
# Votre package apparaîtra dans la liste !
```

### Alternative moderne : go doc

```bash
# Afficher la documentation d'un package dans le terminal
go doc package_name

# Exemples :
go doc fmt
go doc fmt.Println
go doc .        # Documentation du package actuel
```

## Règles de documentation en Go

### 1. Commentaires de package

Chaque package doit avoir un commentaire de package avant la déclaration `package`.

```go
// Package mathutils fournit des utilitaires mathématiques avancés.
//
// Ce package contient des fonctions pour des calculs qui ne sont pas
// disponibles dans le package math standard de Go. Il inclut des
// opérations sur les matrices, des fonctions statistiques et des
// utilitaires de géométrie.
//
// Exemple d'utilisation :
//   import "monprojet/mathutils"
//
//   moyenne := mathutils.Average([]float64{1, 2, 3, 4, 5})
//   fmt.Println(moyenne) // Affiche : 3
package mathutils
```

### 2. Commentaires de fonctions

Chaque fonction publique (qui commence par une majuscule) doit être documentée.

```go
// Add additionne deux nombres entiers et retourne le résultat.
//
// Cette fonction prend deux paramètres entiers et retourne leur somme.
// Elle gère correctement les débordements d'entiers.
//
// Exemple :
//   result := Add(5, 3)
//   fmt.Println(result) // Affiche : 8
func Add(a, b int) int {
    return a + b
}

// Divide divise deux nombres et retourne le résultat et une erreur.
//
// Cette fonction effectue une division sécurisée en gérant le cas
// de la division par zéro.
//
// Paramètres :
//   - a : le dividende
//   - b : le diviseur
//
// Retourne :
//   - float64 : le résultat de la division
//   - error : nil si succès, erreur si division par zéro
//
// Exemple :
//   result, err := Divide(10, 2)
//   if err != nil {
//       log.Fatal(err)
//   }
//   fmt.Println(result) // Affiche : 5
func Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division par zéro impossible")
    }
    return a / b, nil
}
```

### 3. Commentaires de types

Documentez vos structs, interfaces et types personnalisés.

```go
// User représente un utilisateur dans le système.
//
// Un utilisateur contient les informations de base nécessaires
// pour l'authentification et la gestion de profil.
type User struct {
    // ID est l'identifiant unique de l'utilisateur.
    ID int `json:"id"`

    // Name est le nom complet de l'utilisateur.
    // Il doit contenir au moins 2 caractères.
    Name string `json:"name"`

    // Email est l'adresse email de l'utilisateur.
    // Elle doit être unique dans le système.
    Email string `json:"email"`

    // CreatedAt est la date de création du compte.
    CreatedAt time.Time `json:"created_at"`
}

// UserService définit les opérations disponibles pour gérer les utilisateurs.
//
// Cette interface abstrait les opérations CRUD pour les utilisateurs,
// permettant différentes implémentations (base de données, mémoire, etc.).
type UserService interface {
    // CreateUser crée un nouvel utilisateur dans le système.
    CreateUser(user *User) error

    // GetUser récupère un utilisateur par son ID.
    GetUser(id int) (*User, error)

    // UpdateUser met à jour les informations d'un utilisateur.
    UpdateUser(user *User) error

    // DeleteUser supprime un utilisateur du système.
    DeleteUser(id int) error
}
```

### 4. Commentaires de méthodes

```go
// NewUser crée une nouvelle instance d'utilisateur avec validation.
//
// Cette fonction constructeur valide les paramètres d'entrée et
// retourne un utilisateur prêt à être utilisé ou une erreur si
// les paramètres sont invalides.
//
// Validation effectuée :
//   - Name : doit contenir au moins 2 caractères
//   - Email : doit être un format email valide
//
// Exemple :
//   user, err := NewUser("Alice Dupont", "alice@example.com")
//   if err != nil {
//       log.Fatal(err)
//   }
func NewUser(name, email string) (*User, error) {
    if len(name) < 2 {
        return nil, fmt.Errorf("le nom doit contenir au moins 2 caractères")
    }

    if !strings.Contains(email, "@") {
        return nil, fmt.Errorf("format email invalide")
    }

    return &User{
        Name:      name,
        Email:     email,
        CreatedAt: time.Now(),
    }, nil
}

// IsActive vérifie si l'utilisateur est considéré comme actif.
//
// Un utilisateur est actif si son compte a été créé il y a moins
// de 365 jours et si son email n'est pas vide.
//
// Retourne true si l'utilisateur est actif, false sinon.
func (u *User) IsActive() bool {
    if u.Email == "" {
        return false
    }

    oneYearAgo := time.Now().AddDate(-1, 0, 0)
    return u.CreatedAt.After(oneYearAgo)
}
```

## Exemple pratique : Package de calculatrice

Créons un package complètement documenté pour une calculatrice.

### Structure du projet

```text
calculator/
├── go.mod
├── doc.go              # Documentation du package
├── calculator.go       # Types et constructeurs
├── operations.go       # Opérations mathématiques
├── examples/
│   └── basic/
│       └── main.go     # Exemples d'utilisation
└── README.md
```

### 1. Fichier doc.go

```go
// Package calculator fournit une calculatrice programmable avec historique.
//
// Ce package implémente une calculatrice qui peut effectuer des opérations
// mathématiques de base et avancées tout en conservant un historique des
// calculs effectués.
//
// Fonctionnalités principales :
//   - Opérations de base : addition, soustraction, multiplication, division
//   - Opérations avancées : puissance, racine carrée
//   - Historique des calculs
//   - Gestion des erreurs (division par zéro, etc.)
//
// Utilisation de base :
//   calc := calculator.New()
//   result := calc.Add(10).Multiply(2).GetResult()
//   fmt.Println(result) // Affiche : 20
//
// Avec gestion d'erreurs :
//   calc := calculator.New()
//   result, err := calc.Add(10).Divide(0).GetResultWithError()
//   if err != nil {
//       log.Fatal(err)
//   }
//
// Historique :
//   calc := calculator.New()
//   calc.Add(5).Multiply(2)
//   history := calc.GetHistory()
//   for _, operation := range history {
//       fmt.Println(operation)
//   }
package calculator
```

### 2. Fichier calculator.go

```go
package calculator

import (
    "fmt"
    "time"
)

// Calculator représente une calculatrice avec historique.
//
// La calculatrice maintient un résultat courant et un historique
// de toutes les opérations effectuées. Elle supporte le chaînage
// d'opérations pour une utilisation fluide.
type Calculator struct {
    result  float64
    history []Operation
    hasError bool
    lastError error
}

// Operation représente une opération mathématique dans l'historique.
//
// Chaque opération enregistrée contient le type d'opération,
// les opérandes utilisées, le résultat et le moment de l'exécution.
type Operation struct {
    // Type de l'opération ("+", "-", "*", "/", "^", "√")
    Type string

    // Opérandes utilisées dans l'opération
    Operands []float64

    // Résultat de l'opération
    Result float64

    // Moment où l'opération a été effectuée
    Timestamp time.Time
}

// New crée une nouvelle instance de calculatrice.
//
// La calculatrice est initialisée avec un résultat de 0 et un
// historique vide. Elle est prête à effectuer des calculs.
//
// Exemple :
//   calc := New()
//   result := calc.Add(5).GetResult()
//   fmt.Println(result) // Affiche : 5
func New() *Calculator {
    return &Calculator{
        result:  0,
        history: make([]Operation, 0),
        hasError: false,
    }
}

// NewWithValue crée une nouvelle calculatrice avec une valeur initiale.
//
// Cette fonction est utile quand vous voulez commencer les calculs
// avec une valeur spécifique plutôt que zéro.
//
// Paramètres :
//   - value : la valeur initiale de la calculatrice
//
// Exemple :
//   calc := NewWithValue(10)
//   result := calc.Add(5).GetResult()
//   fmt.Println(result) // Affiche : 15
func NewWithValue(value float64) *Calculator {
    return &Calculator{
        result:  value,
        history: make([]Operation, 0),
        hasError: false,
    }
}

// GetResult retourne le résultat actuel de la calculatrice.
//
// Cette méthode ne vérifie pas les erreurs. Utilisez GetResultWithError
// si vous voulez une gestion d'erreurs complète.
//
// Retourne la valeur du résultat courant.
func (c *Calculator) GetResult() float64 {
    return c.result
}

// GetResultWithError retourne le résultat et l'erreur éventuelle.
//
// Cette méthode est recommandée pour une gestion robuste des erreurs,
// particulièrement lors d'opérations qui peuvent échouer (division par zéro).
//
// Retourne :
//   - float64 : le résultat actuel
//   - error : nil si aucune erreur, sinon l'erreur qui s'est produite
//
// Exemple :
//   calc := New()
//   result, err := calc.Add(10).Divide(0).GetResultWithError()
//   if err != nil {
//       fmt.Println("Erreur:", err)
//   }
func (c *Calculator) GetResultWithError() (float64, error) {
    if c.hasError {
        return 0, c.lastError
    }
    return c.result, nil
}

// Reset remet la calculatrice à zéro et efface l'historique.
//
// Cette méthode restaure la calculatrice à son état initial :
// résultat à 0, historique vide, aucune erreur.
//
// Exemple :
//   calc := New()
//   calc.Add(10).Multiply(2)
//   calc.Reset()
//   fmt.Println(calc.GetResult()) // Affiche : 0
func (c *Calculator) Reset() *Calculator {
    c.result = 0
    c.history = make([]Operation, 0)
    c.hasError = false
    c.lastError = nil
    return c
}

// GetHistory retourne l'historique complet des opérations.
//
// L'historique contient toutes les opérations effectuées depuis
// la création de la calculatrice ou le dernier Reset().
//
// Retourne une copie de l'historique pour éviter les modifications
// accidentelles.
func (c *Calculator) GetHistory() []Operation {
    // Retourner une copie pour éviter les modifications externes
    historyCopy := make([]Operation, len(c.history))
    copy(historyCopy, c.history)
    return historyCopy
}

// addToHistory ajoute une opération à l'historique (méthode privée).
func (c *Calculator) addToHistory(operationType string, operands []float64, result float64) {
    operation := Operation{
        Type:      operationType,
        Operands:  operands,
        Result:    result,
        Timestamp: time.Now(),
    }
    c.history = append(c.history, operation)
}

// setError définit une erreur dans la calculatrice (méthode privée).
func (c *Calculator) setError(err error) *Calculator {
    c.hasError = true
    c.lastError = err
    return c
}
```

### 3. Fichier operations.go

```go
package calculator

import (
    "fmt"
    "math"
)

// Add additionne une valeur au résultat actuel.
//
// Cette opération ajoute la valeur spécifiée au résultat courant
// de la calculatrice et met à jour l'historique.
//
// Paramètres :
//   - value : la valeur à ajouter
//
// Retourne la calculatrice pour permettre le chaînage d'opérations.
//
// Exemple :
//   calc := New()
//   result := calc.Add(5).Add(3).GetResult()
//   fmt.Println(result) // Affiche : 8
func (c *Calculator) Add(value float64) *Calculator {
    if c.hasError {
        return c
    }

    oldResult := c.result
    c.result += value
    c.addToHistory("+", []float64{oldResult, value}, c.result)
    return c
}

// Subtract soustrait une valeur du résultat actuel.
//
// Cette opération soustrait la valeur spécifiée du résultat courant
// de la calculatrice et met à jour l'historique.
//
// Paramètres :
//   - value : la valeur à soustraire
//
// Retourne la calculatrice pour permettre le chaînage d'opérations.
//
// Exemple :
//   calc := NewWithValue(10)
//   result := calc.Subtract(3).GetResult()
//   fmt.Println(result) // Affiche : 7
func (c *Calculator) Subtract(value float64) *Calculator {
    if c.hasError {
        return c
    }

    oldResult := c.result
    c.result -= value
    c.addToHistory("-", []float64{oldResult, value}, c.result)
    return c
}

// Multiply multiplie le résultat actuel par une valeur.
//
// Cette opération multiplie le résultat courant par la valeur
// spécifiée et met à jour l'historique.
//
// Paramètres :
//   - value : la valeur par laquelle multiplier
//
// Retourne la calculatrice pour permettre le chaînage d'opérations.
//
// Exemple :
//   calc := NewWithValue(5)
//   result := calc.Multiply(3).GetResult()
//   fmt.Println(result) // Affiche : 15
func (c *Calculator) Multiply(value float64) *Calculator {
    if c.hasError {
        return c
    }

    oldResult := c.result
    c.result *= value
    c.addToHistory("*", []float64{oldResult, value}, c.result)
    return c
}

// Divide divise le résultat actuel par une valeur.
//
// Cette opération divise le résultat courant par la valeur spécifiée.
// Si la valeur est zéro, une erreur est enregistrée et les opérations
// suivantes seront ignorées jusqu'à un Reset().
//
// Paramètres :
//   - value : la valeur par laquelle diviser
//
// Retourne la calculatrice pour permettre le chaînage d'opérations.
//
// Exemple avec succès :
//   calc := NewWithValue(10)
//   result := calc.Divide(2).GetResult()
//   fmt.Println(result) // Affiche : 5
//
// Exemple avec erreur :
//   calc := NewWithValue(10)
//   result, err := calc.Divide(0).GetResultWithError()
//   if err != nil {
//       fmt.Println("Erreur:", err) // Affiche : Erreur: division par zéro
//   }
func (c *Calculator) Divide(value float64) *Calculator {
    if c.hasError {
        return c
    }

    if value == 0 {
        return c.setError(fmt.Errorf("division par zéro"))
    }

    oldResult := c.result
    c.result /= value
    c.addToHistory("/", []float64{oldResult, value}, c.result)
    return c
}

// Power élève le résultat actuel à une puissance.
//
// Cette opération calcule le résultat courant élevé à la puissance
// spécifiée en utilisant math.Pow.
//
// Paramètres :
//   - exponent : l'exposant
//
// Retourne la calculatrice pour permettre le chaînage d'opérations.
//
// Exemple :
//   calc := NewWithValue(2)
//   result := calc.Power(3).GetResult()
//   fmt.Println(result) // Affiche : 8
func (c *Calculator) Power(exponent float64) *Calculator {
    if c.hasError {
        return c
    }

    oldResult := c.result
    c.result = math.Pow(c.result, exponent)
    c.addToHistory("^", []float64{oldResult, exponent}, c.result)
    return c
}

// SquareRoot calcule la racine carrée du résultat actuel.
//
// Cette opération remplace le résultat courant par sa racine carrée.
// Si le résultat actuel est négatif, une erreur est enregistrée.
//
// Retourne la calculatrice pour permettre le chaînage d'opérations.
//
// Exemple avec succès :
//   calc := NewWithValue(9)
//   result := calc.SquareRoot().GetResult()
//   fmt.Println(result) // Affiche : 3
//
// Exemple avec erreur :
//   calc := NewWithValue(-9)
//   result, err := calc.SquareRoot().GetResultWithError()
//   if err != nil {
//       fmt.Println("Erreur:", err) // Erreur: racine carrée d'un nombre négatif
//   }
func (c *Calculator) SquareRoot() *Calculator {
    if c.hasError {
        return c
    }

    if c.result < 0 {
        return c.setError(fmt.Errorf("racine carrée d'un nombre négatif"))
    }

    oldResult := c.result
    c.result = math.Sqrt(c.result)
    c.addToHistory("√", []float64{oldResult}, c.result)
    return c
}
```

### 4. Exemple d'utilisation (examples/basic/main.go)

```go
package main

import (
    "fmt"
    "log"

    "calculator"
)

func main() {
    fmt.Println("=== Exemples d'utilisation de la calculatrice ===\n")

    // Exemple 1 : Opérations de base avec chaînage
    fmt.Println("1. Opérations de base :")
    calc := calculator.New()
    result := calc.Add(10).Multiply(2).Subtract(5).GetResult()
    fmt.Printf("   (0 + 10) * 2 - 5 = %.2f\n", result)

    // Exemple 2 : Utilisation avec gestion d'erreurs
    fmt.Println("\n2. Gestion des erreurs :")
    calc2 := calculator.NewWithValue(10)
    result2, err := calc2.Divide(0).GetResultWithError()
    if err != nil {
        fmt.Printf("   Erreur attendue : %v\n", err)
    } else {
        fmt.Printf("   Résultat : %.2f\n", result2)
    }

    // Exemple 3 : Opérations avancées
    fmt.Println("\n3. Opérations avancées :")
    calc3 := calculator.NewWithValue(3)
    result3 := calc3.Power(2).SquareRoot().GetResult()
    fmt.Printf("   √(3²) = %.2f\n", result3)

    // Exemple 4 : Historique des opérations
    fmt.Println("\n4. Historique des opérations :")
    calc4 := calculator.New()
    calc4.Add(5).Multiply(3).Subtract(2)

    history := calc4.GetHistory()
    fmt.Printf("   Résultat final : %.2f\n", calc4.GetResult())
    fmt.Println("   Historique :")
    for i, op := range history {
        fmt.Printf("     %d. %s %v = %.2f\n", i+1, op.Type, op.Operands, op.Result)
    }
}
```

## Consulter la documentation

### 1. Serveur local

```bash
# Dans le dossier de votre projet
godoc -http=:6060

# Ouvrir http://localhost:6060
# Naviguer vers votre package dans la liste
```

### 2. Documentation en ligne de commande

```bash
# Documentation du package actuel
go doc

# Documentation d'une fonction spécifique
go doc Add

# Documentation complète
go doc -all
```

### 3. Publication en ligne

Une fois votre module publié sur GitHub, il apparaîtra automatiquement sur [pkg.go.dev](https://pkg.go.dev) !

## Conventions et bonnes pratiques

### 1. Structure des commentaires

```go
// NomFonction fait quelque chose de spécifique.
//
// Description plus détaillée si nécessaire.
// Peut s'étendre sur plusieurs lignes.
//
// Paramètres importants ou cas d'usage spéciaux.
//
// Exemple :
//   result := NomFonction(param)
//   fmt.Println(result)
func NomFonction(param string) string {
    return param
}
```

### 2. Utilisation d'exemples

```go
// Fibonacci calcule le nième nombre de Fibonacci.
//
// La séquence de Fibonacci commence par 0, 1 et chaque nombre
// suivant est la somme des deux précédents.
//
// Exemple :
//   fmt.Println(Fibonacci(0))  // 0
//   fmt.Println(Fibonacci(1))  // 1
//   fmt.Println(Fibonacci(5))  // 5
//   fmt.Println(Fibonacci(10)) // 55
func Fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return Fibonacci(n-1) + Fibonacci(n-2)
}
```

### 3. Documentation des erreurs

```go
// ParseConfig lit et parse un fichier de configuration.
//
// Cette fonction lit le fichier spécifié et le parse en tant que JSON.
// Elle retourne la configuration parsée ou une erreur.
//
// Erreurs possibles :
//   - Fichier introuvable
//   - Permissions insuffisantes
//   - JSON invalide
//   - Structure de configuration incorrecte
//
// Exemple :
//   config, err := ParseConfig("app.json")
//   if err != nil {
//       log.Fatal("Erreur de configuration:", err)
//   }
func ParseConfig(filename string) (*Config, error) {
    // Implementation...
}
```

## Tests de documentation

### Tests d'exemples

Go peut tester vos exemples de documentation !

```go
// calculator_test.go
package calculator

import "fmt"

// ExampleCalculator_Add démontre l'utilisation de la méthode Add.
func ExampleCalculator_Add() {
    calc := New()
    result := calc.Add(5).Add(3).GetResult()
    fmt.Println(result)
    // Output: 8
}

// ExampleNew démontre la création d'une calculatrice.
func ExampleNew() {
    calc := New()
    fmt.Println(calc.GetResult())
    // Output: 0
}
```

Exécuter les tests :

```bash
go test -v
```

## Résumé

La documentation avec godoc permet :

**📚 Documentation automatique :**

- Génération web à partir des commentaires
- Integration avec pkg.go.dev
- Recherche et navigation faciles

**✍️ Conventions simples :**

- Commentaires avant les déclarations
- Premier mot = nom de l'élément
- Exemples intégrés au code

**🔧 Outils intégrés :**

- `godoc -http` pour serveur local
- `go doc` pour consultation rapide
- Tests d'exemples automatiques

**💡 Bonnes pratiques :**

- Documenter tous les éléments publics
- Fournir des exemples concrets
- Expliquer les cas d'erreur
- Utiliser une structure cohérente

Une bonne documentation fait la différence entre un package utilisable et un package abandonné. Elle est investissement dans la durabilité et l'adoption de votre code !

---

**🎉 Félicitations !** Vous maîtrisez maintenant les packages et modules Go ! Dans le chapitre suivant, nous explorerons la concurrence avec les goroutines.

⏭️
