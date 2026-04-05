🔝 Retour au [Sommaire](/SOMMAIRE.md)

# 4-4 : Fonctions anonymes et closures

## Introduction

Jusqu'à présent, nous avons créé des fonctions avec des noms, comme `calculer()` ou `saluer()`. Mais Go permet aussi de créer des fonctions **sans nom**, appelées **fonctions anonymes**. Et quand ces fonctions "capturent" des variables de leur environnement, elles deviennent des **closures** (fermetures en français).

Imaginez les fonctions anonymes comme des "assistants temporaires" que vous créez pour une tâche spécifique, sans avoir besoin de leur donner un nom permanent. Les closures sont comme des assistants qui "se souviennent" du contexte dans lequel ils ont été créés.

Ces concepts ouvrent la porte à des patterns de programmation très puissants et élégants !

## Fonctions anonymes : les bases

### Syntaxe de base

Une fonction anonyme se déclare comme une fonction normale, mais sans nom :

```go
func(paramètres) typeRetour {
    // corps de la fonction
}
```

### Premier exemple simple

```go
package main

import "fmt"

func main() {
    // Déclarer ET appeler immédiatement une fonction anonyme
    func() {
        fmt.Println("Je suis une fonction anonyme !")
    }() // Les () à la fin l'exécutent immédiatement

    // Fonction anonyme avec paramètres
    func(nom string) {
        fmt.Printf("Bonjour %s depuis une fonction anonyme !\n", nom)
    }("Alice")

    // Fonction anonyme avec retour
    resultat := func(a, b int) int {
        return a + b
    }(5, 3)

    fmt.Printf("5 + 3 = %d\n", resultat)
}
```

**Résultat :**

```
Je suis une fonction anonyme !
Bonjour Alice depuis une fonction anonyme !
5 + 3 = 8
```

## Assigner des fonctions anonymes à des variables

Les fonctions sont des "citoyens de première classe" en Go, vous pouvez les stocker dans des variables :

```go
package main

import "fmt"

func main() {
    // Assigner une fonction anonyme à une variable
    saluer := func(nom string) {
        fmt.Printf("Salut %s !\n", nom)
    }

    // Maintenant on peut l'appeler comme une fonction normale
    saluer("Bob")
    saluer("Charlie")

    // Fonction qui calcule et retourne
    calculer := func(x, y int) (int, int) {
        somme := x + y
        produit := x * y
        return somme, produit
    }

    s, p := calculer(4, 5)
    fmt.Printf("4 + 5 = %d, 4 × 5 = %d\n", s, p)

    // On peut même changer la fonction assignée à la variable !
    calculer = func(x, y int) (int, int) {
        difference := x - y
        quotient := x / y
        return difference, quotient
    }

    d, q := calculer(10, 2)
    fmt.Printf("10 - 2 = %d, 10 ÷ 2 = %d\n", d, q)
}
```

**Résultat :**

```
Salut Bob !
Salut Charlie !
4 + 5 = 9, 4 × 5 = 20
10 - 2 = 8, 10 ÷ 2 = 5
```

## Introduction aux closures

Une **closure** (fermeture) est une fonction anonyme qui "capture" et utilise des variables de son environnement extérieur. C'est comme si la fonction "se souvenait" des variables qui existaient quand elle a été créée.

### Exemple simple de closure

```go
package main

import "fmt"

func main() {
    message := "Hello depuis l'extérieur"

    // Cette fonction anonyme "capture" la variable message
    afficher := func() {
        fmt.Println(message) // Utilise la variable de l'environnement externe
    }

    afficher() // Affiche: Hello depuis l'extérieur

    // Même si on change la variable externe, la closure la voit !
    message = "Message modifié"
    afficher() // Affiche: Message modifié
}
```

**Résultat :**

```
Hello depuis l'extérieur
Message modifié
```

**Points importants :**

- La fonction anonyme peut accéder à `message` même si elle n'est pas passée en paramètre
- Quand `message` change, la closure voit le changement
- C'est ça une closure : une fonction + son environnement capturé

### Exemple plus pratique : compteur

```go
package main

import "fmt"

func creerCompteur() func() int {
    compte := 0 // Variable locale à creerCompteur

    // Retourner une fonction anonyme qui "capture" compte
    return func() int {
        compte++ // Modifie la variable capturée
        return compte
    }
}

func main() {
    // Créer un compteur
    compteur1 := creerCompteur()

    // Chaque appel incrémente et retourne la valeur
    fmt.Println("Compteur1:", compteur1()) // 1
    fmt.Println("Compteur1:", compteur1()) // 2
    fmt.Println("Compteur1:", compteur1()) // 3

    // Créer un second compteur indépendant
    compteur2 := creerCompteur()
    fmt.Println("Compteur2:", compteur2()) // 1 (indépendant du premier)
    fmt.Println("Compteur1:", compteur1()) // 4 (continue sa séquence)
    fmt.Println("Compteur2:", compteur2()) // 2
}
```

**Résultat :**

```
Compteur1: 1
Compteur1: 2
Compteur1: 3
Compteur2: 1
Compteur1: 4
Compteur2: 2
```

**Explication magique :**

- Chaque appel à `creerCompteur()` crée une nouvelle variable `compte`
- La fonction retournée "se souvient" de SA propre variable `compte`
- Même après que `creerCompteur()` termine, la variable `compte` reste accessible à la closure
- C'est pourquoi les deux compteurs sont indépendants !

## Closures avec paramètres

Les closures peuvent à la fois capturer des variables ET prendre des paramètres :

```go
package main

import "fmt"

func creerMultiplicateur(facteur int) func(int) int {
    // facteur est capturé par la closure
    return func(nombre int) int {
        return nombre * facteur
    }
}

func main() {
    // Créer différents multiplicateurs
    doubler := creerMultiplicateur(2)
    tripler := creerMultiplicateur(3)
    parDix := creerMultiplicateur(10)

    // Utiliser les closures
    fmt.Printf("Doubler 5: %d\n", doubler(5))       // 10
    fmt.Printf("Tripler 4: %d\n", tripler(4))       // 12
    fmt.Printf("5 × 10: %d\n", parDix(5))           // 50

    // On peut les utiliser avec des variables
    nombres := []int{1, 2, 3, 4, 5}
    fmt.Print("Doublés: ")
    for _, n := range nombres {
        fmt.Printf("%d ", doubler(n))
    }
    fmt.Println()
}
```

**Résultat :**

```
Doubler 5: 10
Tripler 4: 12
5 × 10: 50
Doublés: 2 4 6 8 10
```

## Utilisations pratiques des fonctions anonymes

### 1. Callbacks et gestionnaires d'événements

```go
package main

import "fmt"

// Fonction qui prend une autre fonction en paramètre
func traiterNombres(nombres []int, traitement func(int) int) []int {
    resultat := make([]int, len(nombres))
    for i, nombre := range nombres {
        resultat[i] = traitement(nombre)
    }
    return resultat
}

func main() {
    nombres := []int{1, 2, 3, 4, 5}

    // Utiliser des fonctions anonymes comme callbacks
    doubles := traiterNombres(nombres, func(x int) int {
        return x * 2
    })

    carres := traiterNombres(nombres, func(x int) int {
        return x * x
    })

    negatifs := traiterNombres(nombres, func(x int) int {
        return -x
    })

    fmt.Printf("Originaux: %v\n", nombres)
    fmt.Printf("Doublés:   %v\n", doubles)
    fmt.Printf("Carrés:    %v\n", carres)
    fmt.Printf("Négatifs:  %v\n", negatifs)
}
```

**Résultat :**

```
Originaux: [1 2 3 4 5]
Doublés:   [2 4 6 8 10]
Carrés:    [1 4 9 16 25]
Négatifs:  [-1 -2 -3 -4 -5]
```

### 2. Filtrage avec des closures

```go
package main

import "fmt"

func filtrer(nombres []int, condition func(int) bool) []int {
    var resultat []int
    for _, nombre := range nombres {
        if condition(nombre) {
            resultat = append(resultat, nombre)
        }
    }
    return resultat
}

func main() {
    nombres := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    // Filtrer les nombres pairs
    pairs := filtrer(nombres, func(x int) bool {
        return x%2 == 0
    })

    // Filtrer les nombres > 5
    grands := filtrer(nombres, func(x int) bool {
        return x > 5
    })

    // Utiliser une closure qui capture une variable
    limite := 7
    petits := filtrer(nombres, func(x int) bool {
        return x < limite // Capture la variable limite
    })

    fmt.Printf("Tous:    %v\n", nombres)
    fmt.Printf("Pairs:   %v\n", pairs)
    fmt.Printf("> 5:     %v\n", grands)
    fmt.Printf("< %d:     %v\n", limite, petits)
}
```

**Résultat :**

```
Tous:    [1 2 3 4 5 6 7 8 9 10]
Pairs:   [2 4 6 8 10]
> 5:     [6 7 8 9 10]
< 7:     [1 2 3 4 5 6]
```

## Closures avancées : état persistant

### Créateur de générateurs

```go
package main

import "fmt"

func creerGenerateurSuite(debut, increment int) func() int {
    valeur := debut - increment // On commence avant pour que le premier appel donne 'debut'

    return func() int {
        valeur += increment
        return valeur
    }
}

func main() {
    // Générateur de nombres pairs
    pairs := creerGenerateurSuite(0, 2)
    fmt.Print("Nombres pairs: ")
    for i := 0; i < 5; i++ {
        fmt.Printf("%d ", pairs())
    }
    fmt.Println()

    // Générateur de multiples de 5
    multiples5 := creerGenerateurSuite(5, 5)
    fmt.Print("Multiples de 5: ")
    for i := 0; i < 4; i++ {
        fmt.Printf("%d ", multiples5())
    }
    fmt.Println()

    // Générateur de nombres impairs en partant de 1
    impairs := creerGenerateurSuite(1, 2)
    fmt.Print("Nombres impairs: ")
    for i := 0; i < 6; i++ {
        fmt.Printf("%d ", impairs())
    }
    fmt.Println()
}
```

**Résultat :**

```
Nombres pairs: 0 2 4 6 8
Multiples de 5: 5 10 15 20
Nombres impairs: 1 3 5 7 9 11
```

### Cache avec closure

```go
package main

import "fmt"

func creerCache() (func(string, string), func(string) (string, bool)) {
    // Map privée capturée par les closures
    cache := make(map[string]string)

    // Fonction pour ajouter au cache
    ajouter := func(cle, valeur string) {
        cache[cle] = valeur
        fmt.Printf("Ajouté au cache: %s -> %s\n", cle, valeur)
    }

    // Fonction pour récupérer du cache
    recuperer := func(cle string) (string, bool) {
        valeur, existe := cache[cle]
        if existe {
            fmt.Printf("Trouvé dans le cache: %s -> %s\n", cle, valeur)
        } else {
            fmt.Printf("Pas trouvé dans le cache: %s\n", cle)
        }
        return valeur, existe
    }

    return ajouter, recuperer
}

func main() {
    // Créer un système de cache
    ajouter, recuperer := creerCache()

    // Ajouter quelques éléments
    ajouter("nom", "Alice")
    ajouter("age", "25")
    ajouter("ville", "Paris")

    fmt.Println("\n--- Recherches ---")

    // Rechercher des éléments
    if nom, trouve := recuperer("nom"); trouve {
        fmt.Printf("Le nom est: %s\n", nom)
    }

    recuperer("profession") // N'existe pas

    if age, trouve := recuperer("age"); trouve {
        fmt.Printf("L'âge est: %s ans\n", age)
    }
}
```

**Résultat :**

```
Ajouté au cache: nom -> Alice
Ajouté au cache: age -> 25
Ajouté au cache: ville -> Paris

--- Recherches ---
Trouvé dans le cache: nom -> Alice
Le nom est: Alice
Pas trouvé dans le cache: profession
Trouvé dans le cache: age -> 25
L'âge est: 25 ans
```

## Fonctions qui retournent des fonctions

### Fabrique de validateurs

```go
package main

import (
    "fmt"
    "strings"
)

func creerValidateur(regle string) func(string) (bool, string) {
    switch regle {
    case "email":
        return func(valeur string) (bool, string) {
            if strings.Contains(valeur, "@") && strings.Contains(valeur, ".") {
                return true, "Email valide"
            }
            return false, "Format d'email invalide"
        }

    case "longueur":
        return func(valeur string) (bool, string) {
            if len(valeur) >= 3 {
                return true, "Longueur suffisante"
            }
            return false, "Trop court (minimum 3 caractères)"
        }

    case "majuscule":
        return func(valeur string) (bool, string) {
            if strings.ToUpper(valeur) == valeur {
                return true, "Tout en majuscules"
            }
            return false, "Doit être en majuscules"
        }

    default:
        return func(valeur string) (bool, string) {
            return true, "Aucune validation"
        }
    }
}

func main() {
    // Créer différents validateurs
    validEmail := creerValidateur("email")
    validLongueur := creerValidateur("longueur")
    validMajuscule := creerValidateur("majuscule")

    // Tests
    tests := []struct {
        valeur     string
        validateur func(string) (bool, string)
        nom        string
    }{
        {"alice@example.com", validEmail, "Email"},
        {"alice", validEmail, "Email"},
        {"Bob", validLongueur, "Longueur"},
        {"Al", validLongueur, "Longueur"},
        {"HELLO", validMajuscule, "Majuscule"},
        {"Hello", validMajuscule, "Majuscule"},
    }

    for _, test := range tests {
        valide, message := test.validateur(test.valeur)
        status := "✅"
        if !valide {
            status = "❌"
        }
        fmt.Printf("%s [%s] '%s': %s\n", status, test.nom, test.valeur, message)
    }
}
```

**Résultat :**

```
✅ [Email] 'alice@example.com': Email valide
❌ [Email] 'alice': Format d'email invalide
✅ [Longueur] 'Bob': Longueur suffisante
❌ [Longueur] 'Al': Trop court (minimum 3 caractères)
✅ [Majuscule] 'HELLO': Tout en majuscules
❌ [Majuscule] 'Hello': Doit être en majuscules
```

## Exercices pratiques

### Exercice 1 : Accumulateur personnalisé

Créez une fonction qui retourne une closure permettant d'accumuler des valeurs avec une opération personnalisée (addition, multiplication, etc.).

<details>
<summary>Solution</summary>

```go
package main

import "fmt"

func creerAccumulateur(valeurInitiale int, operation string) func(int) int {
    total := valeurInitiale

    switch operation {
    case "+":
        return func(valeur int) int {
            total += valeur
            return total
        }
    case "*":
        return func(valeur int) int {
            total *= valeur
            return total
        }
    case "-":
        return func(valeur int) int {
            total -= valeur
            return total
        }
    default:
        return func(valeur int) int {
            return total
        }
    }
}

func main() {
    // Accumulateur d'addition commençant à 0
    somme := creerAccumulateur(0, "+")
    fmt.Printf("Somme après ajout de 5: %d\n", somme(5))   // 5
    fmt.Printf("Somme après ajout de 3: %d\n", somme(3))   // 8
    fmt.Printf("Somme après ajout de 7: %d\n", somme(7))   // 15

    // Accumulateur de produit commençant à 1
    produit := creerAccumulateur(1, "*")
    fmt.Printf("Produit après × 2: %d\n", produit(2))      // 2
    fmt.Printf("Produit après × 3: %d\n", produit(3))      // 6
    fmt.Printf("Produit après × 4: %d\n", produit(4))      // 24

    // Accumulateur de soustraction commençant à 100
    difference := creerAccumulateur(100, "-")
    fmt.Printf("Différence après - 10: %d\n", difference(10))  // 90
    fmt.Printf("Différence après - 5: %d\n", difference(5))    // 85
}
```

</details>

### Exercice 2 : Système de rappels (callbacks)

Créez un système qui permet d'enregistrer des fonctions de rappel et de les exécuter toutes ensemble.

<details>
<summary>Solution</summary>

```go
package main

import "fmt"

func creerGestionnaireCallbacks() (func(func()), func()) {
    var callbacks []func()

    // Fonction pour ajouter un callback
    ajouter := func(callback func()) {
        callbacks = append(callbacks, callback)
        fmt.Printf("Callback ajouté (total: %d)\n", len(callbacks))
    }

    // Fonction pour exécuter tous les callbacks
    executer := func() {
        fmt.Printf("Exécution de %d callback(s):\n", len(callbacks))
        for i, callback := range callbacks {
            fmt.Printf("  Callback %d: ", i+1)
            callback()
        }
        fmt.Println("Tous les callbacks exécutés.")
    }

    return ajouter, executer
}

func main() {
    ajouter, executer := creerGestionnaireCallbacks()

    // Ajouter des callbacks avec des closures
    nom := "Alice"
    ajouter(func() {
        fmt.Printf("Bonjour %s!\n", nom)
    })

    count := 0
    ajouter(func() {
        count++
        fmt.Printf("Compteur: %d\n", count)
    })

    ajouter(func() {
        fmt.Println("Callback statique")
    })

    // Variables capturées par les closures
    message := "Message initial"
    ajouter(func() {
        fmt.Printf("Message: %s\n", message)
    })

    fmt.Println("\n--- Première exécution ---")
    executer()

    // Modifier les variables capturées
    nom = "Bob"
    message = "Message modifié"

    fmt.Println("\n--- Deuxième exécution ---")
    executer()
}
```

</details>

### Exercice 3 : Builder pattern avec closures

Créez un système pour construire des objets étape par étape en utilisant des closures.

<details>
<summary>Solution</summary>

```go
package main

import "fmt"

type Personne struct {
    Nom        string
    Age        int
    Email      string
    Telephone  string
    Adresse    string
}

func (p Personne) String() string {
    return fmt.Sprintf("Personne{Nom: %s, Age: %d, Email: %s, Tel: %s, Adresse: %s}",
        p.Nom, p.Age, p.Email, p.Telephone, p.Adresse)
}

func nouveauBuilder() func() (func(string) func() interface{}, func() Personne) {
    p := Personne{}

    return func() (func(string) func() interface{}, func() Personne) {
        // Fonction pour définir des propriétés
        definir := func(propriete string) func() interface{} {
            switch propriete {
            case "nom":
                return func() interface{} {
                    return func(nom string) {
                        p.Nom = nom
                        fmt.Printf("Nom défini: %s\n", nom)
                    }
                }
            case "age":
                return func() interface{} {
                    return func(age int) {
                        p.Age = age
                        fmt.Printf("Âge défini: %d\n", age)
                    }
                }
            case "email":
                return func() interface{} {
                    return func(email string) {
                        p.Email = email
                        fmt.Printf("Email défini: %s\n", email)
                    }
                }
            case "telephone":
                return func() interface{} {
                    return func(tel string) {
                        p.Telephone = tel
                        fmt.Printf("Téléphone défini: %s\n", tel)
                    }
                }
            case "adresse":
                return func() interface{} {
                    return func(adresse string) {
                        p.Adresse = adresse
                        fmt.Printf("Adresse définie: %s\n", adresse)
                    }
                }
            default:
                return func() interface{} {
                    return func() {
                        fmt.Printf("Propriété inconnue: %s\n", propriete)
                    }
                }
            }
        }

        // Fonction pour construire l'objet final
        construire := func() Personne {
            fmt.Println("Construction terminée!")
            return p
        }

        return definir, construire
    }
}

// Version simplifiée plus pratique
func creerPersonneBuilder() (func(string), func(int), func(string), func(string), func(string), func() Personne) {
    p := Personne{}

    nom := func(n string) { p.Nom = n }
    age := func(a int) { p.Age = a }
    email := func(e string) { p.Email = e }
    telephone := func(t string) { p.Telephone = t }
    adresse := func(a string) { p.Adresse = a }
    construire := func() Personne { return p }

    return nom, age, email, telephone, adresse, construire
}

func main() {
    fmt.Println("=== Builder pattern avec closures ===")

    // Utilisation du builder simplifié
    nom, age, email, tel, adresse, construire := creerPersonneBuilder()

    // Construction étape par étape
    nom("Alice Dupont")
    age(28)
    email("alice@example.com")
    tel("06 12 34 56 78")
    adresse("123 Rue de la Paix, Paris")

    // Construire l'objet final
    personne := construire()
    fmt.Printf("\nPersonne créée: %s\n", personne)

    // Créer une autre personne avec le même builder
    fmt.Println("\n=== Deuxième personne ===")
    nom2, age2, email2, _, _, construire2 := creerPersonneBuilder()

    nom2("Bob Martin")
    age2(35)
    email2("bob@example.com")
    // On ne définit pas téléphone et adresse

    personne2 := construire2()
    fmt.Printf("Deuxième personne: %s\n", personne2)
}
```

</details>

## Patterns avancés et cas d'usage

### Middleware pattern

```go
package main

import (
    "fmt"
    "time"
)

type HandlerFunc func(string) string

// Middleware de logging
func withLogging(handler HandlerFunc) HandlerFunc {
    return func(input string) string {
        fmt.Printf("[LOG] Traitement de: %s\n", input)
        start := time.Now()

        result := handler(input)

        duration := time.Since(start)
        fmt.Printf("[LOG] Terminé en %v, résultat: %s\n", duration, result)
        return result
    }
}

// Middleware de validation
func withValidation(handler HandlerFunc) HandlerFunc {
    return func(input string) string {
        if len(input) == 0 {
            return "ERREUR: entrée vide"
        }
        if len(input) > 100 {
            return "ERREUR: entrée trop longue"
        }

        return handler(input)
    }
}

// Handler de base
func traiterTexte(texte string) string {
    return fmt.Sprintf("Traité: %s", texte)
}

func main() {
    // Handler simple
    fmt.Println("=== Handler simple ===")
    resultat1 := traiterTexte("Hello")
    fmt.Printf("Résultat: %s\n\n", resultat1)

    // Handler avec logging
    fmt.Println("=== Avec logging ===")
    handlerAvecLog := withLogging(traiterTexte)
    resultat2 := handlerAvecLog("Hello avec log")
    fmt.Printf("Résultat final: %s\n\n", resultat2)

    // Handler avec validation ET logging
    fmt.Println("=== Avec validation et logging ===")
    handlerComplet := withLogging(withValidation(traiterTexte))

    resultat3 := handlerComplet("Texte valide")
    fmt.Printf("Résultat final: %s\n\n", resultat3)

    resultat4 := handlerComplet("") // Texte vide
    fmt.Printf("Résultat final: %s\n", resultat4)
}
```

**Résultat :**

```
=== Handler simple ===
Résultat: Traité: Hello

=== Avec logging ===
[LOG] Traitement de: Hello avec log
[LOG] Terminé en 167ns, résultat: Traité: Hello avec log
Résultat final: Traité: Hello avec log

=== Avec validation et logging ===
[LOG] Traitement de: Texte valide
[LOG] Terminé en 209ns, résultat: Traité: Texte valide
Résultat final: Traité: Texte valide

[LOG] Traitement de:
[LOG] Terminé en 42ns, résultat: ERREUR: entrée vide
Résultat final: ERREUR: entrée vide
```

## Erreurs courantes à éviter

### 1. Confusion entre déclaration et appel

```go
// ❌ Incorrect - ne fait qu'assigner la fonction
maFonction := func() {
    fmt.Println("Hello")
}
// La fonction n'est pas exécutée !

// ✅ Correct - assigne ET exécute
func() {
    fmt.Println("Hello")
}() // Les () à la fin exécutent la fonction

// ✅ Ou assigner puis appeler
maFonction := func() {
    fmt.Println("Hello")
}
maFonction() // Appel explicit
```

### 2. Capture de variables de boucle

```go
package main

import "fmt"

func main() {
    var fonctions []func()

    // ❌ Problème classique
    for i := 0; i < 3; i++ {
        fonctions = append(fonctions, func() {
            fmt.Printf("Valeur: %d\n", i) // Capture i par référence !
        })
    }

    fmt.Println("Résultat incorrect:")
    for _, f := range fonctions {
        f() // Affiche toujours 3 !
    }

    fmt.Println("\nRésultat correct - Solution 1 (copie locale):")
    var bonnesFonctions1 []func()
    for i := 0; i < 3; i++ {
        valeur := i // Copie locale
        bonnesFonctions1 = append(bonnesFonctions1, func() {
            fmt.Printf("Valeur: %d\n", valeur)
        })
    }

    for _, f := range bonnesFonctions1 {
        f() // Affiche 0, 1, 2
    }

    fmt.Println("\nRésultat correct - Solution 2 (paramètre):")
    var bonnesFonctions2 []func()
    for i := 0; i < 3; i++ {
        bonnesFonctions2 = append(bonnesFonctions2, func(val int) func() {
            return func() {
                fmt.Printf("Valeur: %d\n", val)
            }
        }(i)) // Passer i comme paramètre immédiatement
    }

    for _, f := range bonnesFonctions2 {
        f()
    }
}
```

**Résultat :**

```
Résultat incorrect:
Valeur: 3
Valeur: 3
Valeur: 3

Résultat correct - Solution 1 (copie locale):
Valeur: 0
Valeur: 1
Valeur: 2

Résultat correct - Solution 2 (paramètre):
Valeur: 0
Valeur: 1
Valeur: 2
```

**Explication du problème :**

- Toutes les closures capturent la **même** variable `i`
- Quand la boucle finit, `i` vaut 3
- Toutes les closures voient cette valeur finale

### 3. Fuite mémoire avec les closures

```go
package main

import (
    "fmt"
    "runtime"
)

func creerGrosseClosure() func() {
    // Gros tableau qui sera capturé par la closure
    grosTableau := make([]int, 1000000)
    for i := range grosTableau {
        grosTableau[i] = i
    }

    // Cette closure capture TOUT le tableau, même si elle n'utilise qu'un élément
    return func() {
        fmt.Printf("Premier élément: %d\n", grosTableau[0])
    }
}

func creerBonneClosure() func() {
    grosTableau := make([]int, 1000000)
    for i := range grosTableau {
        grosTableau[i] = i
    }

    // Copier seulement ce dont on a besoin
    premierElement := grosTableau[0]

    return func() {
        fmt.Printf("Premier élément: %d\n", premierElement)
    }
    // grosTableau peut maintenant être collecté par le GC
}

func main() {
    runtime.GC() // Force le garbage collection
    var m1 runtime.MemStats
    runtime.ReadMemStats(&m1)
    fmt.Printf("Mémoire avant: %d KB\n", m1.Alloc/1024)

    // ❌ Mauvaise pratique - garde tout le tableau en mémoire
    mauvaiseClosure := creerGrosseClosure()

    runtime.GC()
    var m2 runtime.MemStats
    runtime.ReadMemStats(&m2)
    fmt.Printf("Mémoire après mauvaise closure: %d KB\n", m2.Alloc/1024)

    mauvaiseClosure()

    // ✅ Bonne pratique - ne garde que ce qui est nécessaire
    bonneClosure := creerBonneClosure()

    runtime.GC()
    var m3 runtime.MemStats
    runtime.ReadMemStats(&m3)
    fmt.Printf("Mémoire après bonne closure: %d KB\n", m3.Alloc/1024)

    bonneClosure()
}
```

## Bonnes pratiques avec les closures

### 1. Documentez les captures

```go
// ✅ Bon : documentation claire de ce qui est capturé
func creerValidateur(longueurMin int) func(string) bool {
    // Cette closure capture longueurMin
    return func(texte string) bool {
        return len(texte) >= longueurMin
    }
}
```

### 2. Minimisez les captures

```go
type Config struct {
    Host     string
    Port     int
    Database string
    Secret   string
}

// ❌ Capture toute la config
func mauvaisConnecteur(config Config) func() {
    return func() {
        // N'utilise que Host et Port, mais capture tout
        fmt.Printf("Connexion à %s:%d\n", config.Host, config.Port)
    }
}

// ✅ Capture seulement ce qui est nécessaire
func bonConnecteur(config Config) func() {
    host := config.Host
    port := config.Port
    return func() {
        fmt.Printf("Connexion à %s:%d\n", host, port)
    }
}
```

### 3. Attention aux références circulaires

```go
package main

import "fmt"

type Node struct {
    Value int
    Next  *Node
    Print func() // ⚠️ Risque de référence circulaire
}

func creerNode(value int) *Node {
    node := &Node{Value: value}

    // ❌ Référence circulaire : node -> Print -> node
    node.Print = func() {
        fmt.Printf("Node value: %d\n", node.Value)
    }

    return node
}

// ✅ Solution sans référence circulaire
func creerNodeSafe(value int) *Node {
    node := &Node{Value: value}

    // Capture la valeur, pas la référence au node
    nodeValue := node.Value
    node.Print = func() {
        fmt.Printf("Node value: %d\n", nodeValue)
    }

    return node
}

func main() {
    node1 := creerNode(42)
    node1.Print()

    node2 := creerNodeSafe(24)
    node2.Print()
}
```

## Cas d'usage avancés

### Configuration dynamique

```go
package main

import (
    "fmt"
    "strings"
)

type Logger struct {
    prefix string
    level  string
}

func (l Logger) log(message string) {
    fmt.Printf("[%s][%s] %s\n", l.level, l.prefix, message)
}

func creerLoggerBuilder() func(string) func(string) Logger {
    return func(prefix string) func(string) Logger {
        return func(level string) Logger {
            return Logger{
                prefix: prefix,
                level:  strings.ToUpper(level),
            }
        }
    }
}

// Alternative avec closure qui maintient l'état
func creerLoggerAvecConfig() (func(string), func(string), func() Logger) {
    var prefix, level string

    setPrefix := func(p string) {
        prefix = p
        fmt.Printf("Prefix configuré: %s\n", prefix)
    }

    setLevel := func(l string) {
        level = strings.ToUpper(l)
        fmt.Printf("Level configuré: %s\n", level)
    }

    build := func() Logger {
        if prefix == "" {
            prefix = "APP"
        }
        if level == "" {
            level = "INFO"
        }
        return Logger{prefix: prefix, level: level}
    }

    return setPrefix, setLevel, build
}

func main() {
    // Méthode 1 : Builder fonctionnel
    fmt.Println("=== Builder fonctionnel ===")
    builder := creerLoggerBuilder()
    logger1 := builder("WEB")("debug")
    logger1.log("Serveur démarré")

    logger2 := builder("DB")("error")
    logger2.log("Connexion échouée")

    // Méthode 2 : Configuration par étapes
    fmt.Println("\n=== Configuration par étapes ===")
    setPrefix, setLevel, build := creerLoggerAvecConfig()

    setPrefix("API")
    setLevel("warning")
    logger3 := build()
    logger3.log("Limite de taux atteinte")

    // Reconfigurer
    setPrefix("AUTH")
    setLevel("info")
    logger4 := build()
    logger4.log("Utilisateur connecté")
}
```

**Résultat :**

```
=== Builder fonctionnel ===
[DEBUG][WEB] Serveur démarré
[ERROR][DB] Connexion échouée

=== Configuration par étapes ===
Prefix configuré: API
Level configuré: WARNING
[WARNING][API] Limite de taux atteinte
Prefix configuré: AUTH
Level configuré: INFO
[INFO][AUTH] Utilisateur connecté
```

### Plugin system avec closures

```go
package main

import "fmt"

type Plugin func(string) string

func creerRegistrePlugins() (func(string, Plugin), func(string, string) string, func()) {
    plugins := make(map[string]Plugin)

    // Enregistrer un plugin
    enregistrer := func(nom string, plugin Plugin) {
        plugins[nom] = plugin
        fmt.Printf("Plugin '%s' enregistré\n", nom)
    }

    // Exécuter un plugin
    executer := func(nomPlugin, input string) string {
        if plugin, existe := plugins[nomPlugin]; existe {
            return plugin(input)
        }
        return fmt.Sprintf("Plugin '%s' non trouvé", nomPlugin)
    }

    // Lister les plugins
    lister := func() {
        fmt.Printf("Plugins disponibles (%d):\n", len(plugins))
        for nom := range plugins {
            fmt.Printf("  - %s\n", nom)
        }
    }

    return enregistrer, executer, lister
}

func main() {
    enregistrer, executer, lister := creerRegistrePlugins()

    // Créer et enregistrer des plugins
    enregistrer("majuscules", func(texte string) string {
        return strings.ToUpper(texte)
    })

    enregistrer("reverse", func(texte string) string {
        runes := []rune(texte)
        for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
            runes[i], runes[j] = runes[j], runes[i]
        }
        return string(runes)
    })

    // Plugin avec état capturé
    compteur := 0
    enregistrer("compteur", func(texte string) string {
        compteur++
        return fmt.Sprintf("[%d] %s", compteur, texte)
    })

    lister()
    fmt.Println()

    // Utiliser les plugins
    fmt.Println("Tests des plugins:")
    fmt.Printf("Majuscules: %s\n", executer("majuscules", "hello world"))
    fmt.Printf("Reverse: %s\n", executer("reverse", "hello"))
    fmt.Printf("Compteur 1: %s\n", executer("compteur", "premier"))
    fmt.Printf("Compteur 2: %s\n", executer("compteur", "deuxième"))
    fmt.Printf("Plugin inexistant: %s\n", executer("inconnu", "test"))
}
```

## Performance et optimisations

### Comparaison fonction normale vs closure

```go
package main

import (
    "fmt"
    "time"
)

// Fonction normale
func calculNormal(x int) int {
    return x * x + 2*x + 1
}

// Closure simple
func creerCalculClosure() func(int) int {
    return func(x int) int {
        return x * x + 2*x + 1
    }
}

// Closure avec capture
func creerCalculAvecCapture(a, b, c int) func(int) int {
    return func(x int) int {
        return a*x*x + b*x + c
    }
}

func benchmarkFonction(nom string, fn func(int) int, iterations int) {
    start := time.Now()

    result := 0
    for i := 0; i < iterations; i++ {
        result += fn(i % 1000)
    }

    duration := time.Since(start)
    fmt.Printf("%s: %v (résultat: %d)\n", nom, duration, result)
}

func main() {
    iterations := 10000000

    // Préparer les fonctions
    closureSimple := creerCalculClosure()
    closureAvecCapture := creerCalculAvecCapture(1, 2, 1)

    fmt.Printf("Benchmark avec %d itérations:\n", iterations)

    // Comparer les performances
    benchmarkFonction("Fonction normale", calculNormal, iterations)
    benchmarkFonction("Closure simple", closureSimple, iterations)
    benchmarkFonction("Closure avec capture", closureAvecCapture, iterations)
}
```

### Optimisation mémoire

```go
package main

import (
    "fmt"
    "runtime"
    "time"
)

// ❌ Closure lourde - capture toute la structure
type ConfigLourde struct {
    DatabaseURL string
    SecretKey   [1000]byte
    Cache       map[string][]byte
    Settings    map[string]interface{}
}

func creerHandlerLourd(config ConfigLourde) func(string) {
    return func(message string) {
        // N'utilise que DatabaseURL mais capture tout
        fmt.Printf("DB: %s, Message: %s\n", config.DatabaseURL, message)
    }
}

// ✅ Closure optimisée - capture seulement ce qui est nécessaire
func creerHandlerOptimise(config ConfigLourde) func(string) {
    dbURL := config.DatabaseURL // Copie seulement ce qui est utilisé
    return func(message string) {
        fmt.Printf("DB: %s, Message: %s\n", dbURL, message)
    }
}

func mesureMemoire(nom string, fn func()) {
    runtime.GC()
    var avant runtime.MemStats
    runtime.ReadMemStats(&avant)

    start := time.Now()
    fn()
    duration := time.Since(start)

    runtime.GC()
    var apres runtime.MemStats
    runtime.ReadMemStats(&apres)

    fmt.Printf("%s:\n", nom)
    fmt.Printf("  Temps: %v\n", duration)
    fmt.Printf("  Mémoire utilisée: %d KB\n", (apres.Alloc-avant.Alloc)/1024)
    fmt.Printf("  Objets alloués: %d\n", apres.Mallocs-avant.Mallocs)
    fmt.Println()
}

func main() {
    // Créer une config lourde
    config := ConfigLourde{
        DatabaseURL: "postgres://localhost/db",
        Cache:       make(map[string][]byte),
        Settings:    make(map[string]interface{}),
    }

    // Remplir avec des données
    for i := 0; i < 1000; i++ {
        key := fmt.Sprintf("key_%d", i)
        config.Cache[key] = make([]byte, 1024) // 1KB par entrée
        config.Settings[key] = fmt.Sprintf("value_%d", i)
    }

    fmt.Println("Comparaison utilisation mémoire:")

    // Test handler lourd
    mesureMemoire("Handler lourd (capture tout)", func() {
        handlers := make([]func(string), 100)
        for i := 0; i < 100; i++ {
            handlers[i] = creerHandlerLourd(config)
        }
        // Utiliser les handlers pour éviter l'optimisation du compilateur
        for _, h := range handlers {
            h("test")
        }
    })

    // Test handler optimisé
    mesureMemoire("Handler optimisé (capture minimum)", func() {
        handlers := make([]func(string), 100)
        for i := 0; i < 100; i++ {
            handlers[i] = creerHandlerOptimise(config)
        }
        for _, h := range handlers {
            h("test")
        }
    })
}
```

## Patterns idiomatiques Go

### Options pattern avec closures

```go
package main

import "fmt"

type Server struct {
    Host    string
    Port    int
    Timeout int
    Debug   bool
    MaxConn int
}

type ServerOption func(*Server)

func WithHost(host string) ServerOption {
    return func(s *Server) {
        s.Host = host
    }
}

func WithPort(port int) ServerOption {
    return func(s *Server) {
        s.Port = port
    }
}

func WithTimeout(timeout int) ServerOption {
    return func(s *Server) {
        s.Timeout = timeout
    }
}

func WithDebug(debug bool) ServerOption {
    return func(s *Server) {
        s.Debug = debug
    }
}

func WithMaxConnections(max int) ServerOption {
    return func(s *Server) {
        s.MaxConn = max
    }
}

func NewServer(options ...ServerOption) *Server {
    // Configuration par défaut
    server := &Server{
        Host:    "localhost",
        Port:    8080,
        Timeout: 30,
        Debug:   false,
        MaxConn: 100,
    }

    // Appliquer les options
    for _, option := range options {
        option(server)
    }

    return server
}

func (s *Server) String() string {
    return fmt.Sprintf("Server{Host: %s, Port: %d, Timeout: %ds, Debug: %v, MaxConn: %d}",
        s.Host, s.Port, s.Timeout, s.Debug, s.MaxConn)
}

func main() {
    // Serveur avec configuration par défaut
    server1 := NewServer()
    fmt.Printf("Serveur par défaut: %s\n", server1)

    // Serveur avec quelques options
    server2 := NewServer(
        WithHost("0.0.0.0"),
        WithPort(3000),
        WithDebug(true),
    )
    fmt.Printf("Serveur personnalisé: %s\n", server2)

    // Serveur avec toutes les options
    server3 := NewServer(
        WithHost("api.example.com"),
        WithPort(443),
        WithTimeout(60),
        WithDebug(false),
        WithMaxConnections(500),
    )
    fmt.Printf("Serveur complet: %s\n", server3)
}
```

**Résultat :**

```
Serveur par défaut: Server{Host: localhost, Port: 8080, Timeout: 30s, Debug: false, MaxConn: 100}
Serveur personnalisé: Server{Host: 0.0.0.0, Port: 3000, Timeout: 30s, Debug: true, MaxConn: 100}
Serveur complet: Server{Host: api.example.com, Port: 443, Timeout: 60s, Debug: false, MaxConn: 500}
```

## Résumé et conseils finaux

### Ce que nous avons appris

**Fonctions anonymes :**

- Syntaxe : `func(params) retour { ... }`
- Peuvent être appelées immédiatement ou stockées dans des variables
- Utiles pour les callbacks et les transformations ponctuelles

**Closures :**

- Fonctions qui capturent des variables de leur environnement
- Permettent de créer des fonctions avec état persistant
- Très puissantes pour les patterns comme les factories et les builders

**Patterns courants :**

- **Callbacks** : passer des fonctions comme paramètres
- **Factories** : fonctions qui créent et retournent d'autres fonctions
- **Options pattern** : configuration flexible avec des closures
- **Middleware** : composition de fonctions pour ajouter des comportements

### Quand utiliser quoi ?

**Utilisez des fonctions anonymes quand :**

- Vous avez besoin d'une fonction simple et ponctuelle
- Vous implémentez des callbacks
- Vous voulez éviter de polluer l'espace de noms global

**Utilisez des closures quand :**

- Vous avez besoin d'état persistant entre les appels
- Vous créez des configurations ou des builders flexibles
- Vous implémentez des patterns comme le middleware
- Vous voulez encapsuler des données privées

### Bonnes pratiques finales

1. **Simplicité** : commencez simple, ajoutez la complexité seulement quand nécessaire
2. **Documentation** : expliquez ce qui est capturé et pourquoi
3. **Performance** : attention aux captures lourdes et aux fuites mémoire
4. **Lisibilité** : préférez des fonctions nommées si la logique est complexe
5. **Test** : les closures peuvent être plus difficiles à tester, planifiez vos tests

### Exemple récapitulatif

```go
package main

import "fmt"

// Démonstration des concepts principaux
func demonstrationComplete() {
    fmt.Println("=== Démonstration complète ===")

    // 1. Fonction anonyme simple
    func() {
        fmt.Println("1. Fonction anonyme exécutée immédiatement")
    }()

    // 2. Fonction stockée dans une variable
    saluer := func(nom string) {
        fmt.Printf("2. Bonjour %s !\n", nom)
    }
    saluer("Go")

    // 3. Closure avec capture
    compteur := 0
    incrementer := func() int {
        compteur++
        return compteur
    }
    fmt.Printf("3. Compteur: %d, %d, %d\n",
        incrementer(), incrementer(), incrementer())

    // 4. Factory de fonctions
    creerMultiplicateur := func(facteur int) func(int) int {
        return func(x int) int {
            return x * facteur
        }
    }
    doubler := creerMultiplicateur(2)
    fmt.Printf("4. 5 × 2 = %d\n", doubler(5))

    // 5. Callback avec closure
    nombres := []int{1, 2, 3, 4, 5}
    base := 10
    transformer := func(slice []int, fn func(int) int) []int {
        result := make([]int, len(slice))
        for i, v := range slice {
            result[i] = fn(v)
        }
        return result
    }

    ajouterBase := transformer(nombres, func(x int) int {
        return x + base // Capture 'base'
    })
    fmt.Printf("5. Ajout de %d: %v\n", base, ajouterBase)
}

func main() {
    demonstrationComplete()
}
```

**Résultat :**

```
=== Démonstration complète ===
1. Fonction anonyme exécutée immédiatement
2. Bonjour Go !
3. Compteur: 1, 2, 3
4. 5 × 2 = 10
5. Ajout de 10: [11 12 13 14 15]
```

Les fonctions anonymes et les closures sont des outils puissants qui, une fois maîtrisés, vous permettront d'écrire du code Go plus expressif et flexible. Elles sont particulièrement utiles dans la programmation fonctionnelle et pour créer des APIs élégantes.

Dans le prochain chapitre, nous explorerons les structures de données Go et comment elles interagissent avec les fonctions que nous venons d'apprendre !

⏭️
