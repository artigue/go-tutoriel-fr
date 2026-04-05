🔝 Retour au [Sommaire](/SOMMAIRE.md)

# 4-3 : Fonctions variadiques

## Introduction

Imaginez que vous voulez créer une fonction pour calculer la somme de plusieurs nombres, mais vous ne savez pas à l'avance combien de nombres vous aurez. Parfois 2, parfois 5, parfois 10... Les **fonctions variadiques** résolvent exactement ce problème ! Elles permettent de passer un nombre variable d'arguments du même type à une fonction.

Le terme "variadique" vient du latin et signifie "qui varie". Ces fonctions sont très pratiques et vous les avez déjà utilisées sans le savoir : `fmt.Println()` est une fonction variadique !

## Syntaxe des fonctions variadiques

Une fonction variadique se déclare en ajoutant `...` devant le type du dernier paramètre :

```go
func nomFonction(parametres normaux, variadic ...type) {
    // Le paramètre variadique devient un slice à l'intérieur de la fonction
}
```

### Exemple simple : calculer une somme

```go
package main

import "fmt"

// Fonction variadique pour calculer la somme
func somme(nombres ...int) int {
    total := 0
    fmt.Printf("J'ai reçu %d nombre(s): %v\n", len(nombres), nombres)

    for _, nombre := range nombres {
        total += nombre
    }
    return total
}

func main() {
    // Appels avec différents nombres d'arguments
    fmt.Println("Somme de rien:", somme())
    fmt.Println("Somme de 1 nombre:", somme(5))
    fmt.Println("Somme de 3 nombres:", somme(1, 2, 3))
    fmt.Println("Somme de 5 nombres:", somme(10, 20, 30, 40, 50))
}
```

**Résultat :**

```
J'ai reçu 0 nombre(s): []
Somme de rien: 0
J'ai reçu 1 nombre(s): [5]
Somme de 1 nombre: 5
J'ai reçu 3 nombre(s): [1 2 3]
Somme de 3 nombres: 6
J'ai reçu 5 nombre(s): [10 20 30 40 50]
Somme de 5 nombres: 150
```

**Points importants :**

- `nombres ...int` signifie "zéro ou plusieurs entiers"
- À l'intérieur de la fonction, `nombres` est un slice `[]int`
- On peut appeler la fonction avec 0, 1, ou plusieurs arguments

## Mélanger paramètres normaux et variadiques

Les paramètres variadiques doivent **toujours être en dernier** :

```go
package main

import "fmt"

// Fonction avec paramètres normaux ET variadiques
func saluerPlusieurs(salutation string, noms ...string) {
    if len(noms) == 0 {
        fmt.Println("Personne à saluer !")
        return
    }

    fmt.Printf("%s ", salutation)
    for i, nom := range noms {
        if i == len(noms)-1 {
            fmt.Printf("%s !\n", nom)
        } else if i == len(noms)-2 {
            fmt.Printf("%s et ", nom)
        } else {
            fmt.Printf("%s, ", nom)
        }
    }
}

func main() {
    saluerPlusieurs("Bonjour")
    saluerPlusieurs("Bonjour", "Alice")
    saluerPlusieurs("Bonjour", "Alice", "Bob")
    saluerPlusieurs("Bonjour", "Alice", "Bob", "Charlie", "Diana")
}
```

**Résultat :**

```
Personne à saluer !
Bonjour Alice !
Bonjour Alice et Bob !
Bonjour Alice, Bob, Charlie et Diana !
```

## Passer un slice à une fonction variadique

Vous pouvez passer un slice existant à une fonction variadique en utilisant `...` :

```go
package main

import "fmt"

func maximum(nombres ...int) int {
    if len(nombres) == 0 {
        return 0 // ou on pourrait retourner une erreur
    }

    max := nombres[0]
    for _, nombre := range nombres {
        if nombre > max {
            max = nombre
        }
    }
    return max
}

func main() {
    // Appel normal avec arguments individuels
    fmt.Println("Maximum de 3, 7, 2:", maximum(3, 7, 2))

    // Appel avec un slice en utilisant ...
    mesNombres := []int{15, 3, 9, 27, 1}
    fmt.Println("Maximum du slice:", maximum(mesNombres...))

    // On peut aussi mélanger (mais tous les arguments doivent être du même type)
    autresNombres := []int{100, 200}
    fmt.Println("Maximum mixte:", maximum(50, autresNombres...)) // Équivaut à maximum(50, 100, 200)
}
```

**Résultat :**

```
Maximum de 3, 7, 2: 7
Maximum du slice: 27
Maximum mixte: 200
```

**Explication :**

- `maximum(mesNombres...)` "décompresse" le slice et passe chaque élément comme argument séparé
- C'est équivalent à `maximum(15, 3, 9, 27, 1)`

## Exemples pratiques avec différents types

### Fonction variadique avec des strings

```go
package main

import (
    "fmt"
    "strings"
)

func creerPhrase(mots ...string) string {
    if len(mots) == 0 {
        return "Phrase vide"
    }

    // Joindre tous les mots avec des espaces
    phrase := strings.Join(mots, " ")

    // Ajouter un point à la fin
    return phrase + "."
}

func main() {
    fmt.Println(creerPhrase())
    fmt.Println(creerPhrase("Bonjour"))
    fmt.Println(creerPhrase("Go", "est", "un", "super", "langage"))

    // Avec un slice de strings
    mots := []string{"J'aime", "programmer", "en", "Go"}
    fmt.Println(creerPhrase(mots...))
}
```

**Résultat :**

```
Phrase vide
Bonjour.
Go est un super langage.
J'aime programmer en Go.
```

### Fonction variadique avec des structures

```go
package main

import "fmt"

type Produit struct {
    Nom  string
    Prix float64
}

func calculerTotal(produits ...Produit) (float64, int) {
    total := 0.0
    nombre := len(produits)

    fmt.Println("Produits dans le panier:")
    for _, produit := range produits {
        fmt.Printf("- %s: %.2f€\n", produit.Nom, produit.Prix)
        total += produit.Prix
    }

    return total, nombre
}

func main() {
    // Créer quelques produits
    pain := Produit{Nom: "Pain", Prix: 1.20}
    lait := Produit{Nom: "Lait", Prix: 0.85}
    fromage := Produit{Nom: "Fromage", Prix: 4.50}

    // Calculer le total de différentes façons
    fmt.Println("=== Achat 1 ===")
    total1, nb1 := calculerTotal(pain)
    fmt.Printf("Total: %.2f€ (%d produit)\n\n", total1, nb1)

    fmt.Println("=== Achat 2 ===")
    total2, nb2 := calculerTotal(pain, lait, fromage)
    fmt.Printf("Total: %.2f€ (%d produits)\n\n", total2, nb2)

    // Avec un slice
    courses := []Produit{pain, lait}
    fmt.Println("=== Achat 3 ===")
    total3, nb3 := calculerTotal(courses...)
    fmt.Printf("Total: %.2f€ (%d produits)\n", total3, nb3)
}
```

**Résultat :**

```
=== Achat 1 ===
Produits dans le panier:
- Pain: 1.20€
Total: 1.20€ (1 produit)

=== Achat 2 ===
Produits dans le panier:
- Pain: 1.20€
- Lait: 0.85€
- Fromage: 4.50€
Total: 6.55€ (3 produits)

=== Achat 3 ===
Produits dans le panier:
- Pain: 1.20€
- Lait: 0.85€
Total: 2.05€ (2 produits)
```

## Fonctions variadiques avec interface{} (tous types)

Parfois, vous voulez accepter des arguments de types différents. Utilisez `interface{}` :

```go
package main

import "fmt"

func afficherTout(elements ...interface{}) {
    fmt.Printf("J'ai reçu %d élément(s):\n", len(elements))

    for i, element := range elements {
        fmt.Printf("  [%d] %v (type: %T)\n", i, element, element)
    }
}

func main() {
    afficherTout("Hello", 42, 3.14, true, []int{1, 2, 3})
}
```

**Résultat :**

```
J'ai reçu 5 élément(s):
  [0] Hello (type: string)
  [1] 42 (type: int)
  [2] 3.14 (type: float64)
  [3] true (type: bool)
  [4] [1 2 3] (type: []int)
```

**Note :** C'est exactement comme ça que `fmt.Println()` fonctionne !

## Comparaison avec fmt.Println

Maintenant vous comprenez pourquoi `fmt.Println` peut prendre autant d'arguments que vous voulez :

```go
package main

import "fmt"

func main() {
    // fmt.Println est définie comme : func Println(a ...interface{}) (n int, err error)

    fmt.Println() // 0 argument
    fmt.Println("Un seul argument") // 1 argument
    fmt.Println("Plusieurs", "arguments", "de", "types", 42, true) // 6 arguments

    // Vous pouvez même passer un slice
    valeurs := []interface{}{"Depuis", "un", "slice", 123}
    fmt.Println(valeurs...)
}
```

## Créer sa propre fonction de logging

Voici un exemple pratique d'utilisation des fonctions variadiques :

```go
package main

import (
    "fmt"
    "time"
)

func log(niveau string, messages ...interface{}) {
    // Créer un timestamp
    timestamp := time.Now().Format("2006-01-02 15:04:05")

    // Afficher le niveau et le timestamp
    fmt.Printf("[%s] %s: ", timestamp, niveau)

    // Afficher tous les messages
    fmt.Println(messages...)
}

func main() {
    log("INFO", "Application démarrée")
    log("DEBUG", "Valeur de x:", 42)
    log("ERROR", "Erreur dans", "la fonction", "calculer", "avec code", 500)

    // Avec des variables
    utilisateur := "Alice"
    score := 95.5
    log("INFO", "Utilisateur", utilisateur, "a obtenu le score", score)
}
```

**Résultat :**

```
[2024-01-15 14:30:45] INFO: Application démarrée
[2024-01-15 14:30:45] DEBUG: Valeur de x: 42
[2024-01-15 14:30:45] ERROR: Erreur dans la fonction calculer avec code 500
[2024-01-15 14:30:45] INFO: Utilisateur Alice a obtenu le score 95.5
```

## Exercices pratiques

### Exercice 1 : Fonction de moyenne

Créez une fonction variadique `moyenne` qui calcule la moyenne de nombres flottants.

<details>
<summary>Solution</summary>

```go
package main

import "fmt"

func moyenne(nombres ...float64) (float64, bool) {
    if len(nombres) == 0 {
        return 0, false // Pas de moyenne possible
    }

    somme := 0.0
    for _, nombre := range nombres {
        somme += nombre
    }

    return somme / float64(len(nombres)), true
}

func main() {
    // Test avec différents nombres d'arguments
    testCases := [][]float64{
        {},
        {5.0},
        {2.0, 4.0, 6.0},
        {1.5, 2.5, 3.5, 4.5, 5.5},
    }

    for i, test := range testCases {
        fmt.Printf("Test %d avec %v:\n", i+1, test)
        if moy, valide := moyenne(test...); valide {
            fmt.Printf("  Moyenne: %.2f\n", moy)
        } else {
            fmt.Println("  Impossible de calculer (aucun nombre)")
        }
        fmt.Println()
    }
}
```

</details>

### Exercice 2 : Constructeur flexible

Créez une fonction variadique pour créer des utilisateurs avec des informations optionnelles.

<details>
<summary>Solution</summary>

```go
package main

import "fmt"

type Utilisateur struct {
    Nom    string
    Age    int
    Email  string
    Actif  bool
}

func creerUtilisateur(nom string, infos ...interface{}) Utilisateur {
    utilisateur := Utilisateur{
        Nom:   nom,
        Age:   0,
        Email: "",
        Actif: true, // Par défaut actif
    }

    // Traiter les informations optionnelles
    for _, info := range infos {
        switch v := info.(type) {
        case int:
            utilisateur.Age = v
        case string:
            utilisateur.Email = v
        case bool:
            utilisateur.Actif = v
        }
    }

    return utilisateur
}

func (u Utilisateur) String() string {
    return fmt.Sprintf("Utilisateur{Nom: %s, Age: %d, Email: %s, Actif: %v}",
        u.Nom, u.Age, u.Email, u.Actif)
}

func main() {
    // Différentes façons de créer des utilisateurs
    user1 := creerUtilisateur("Alice")
    user2 := creerUtilisateur("Bob", 25)
    user3 := creerUtilisateur("Charlie", 30, "charlie@example.com")
    user4 := creerUtilisateur("Diana", 28, "diana@example.com", false)
    user5 := creerUtilisateur("Eve", false, 35, "eve@example.com")

    fmt.Println(user1)
    fmt.Println(user2)
    fmt.Println(user3)
    fmt.Println(user4)
    fmt.Println(user5)
}
```

</details>

### Exercice 3 : Fonction de validation multiple

Créez une fonction variadique qui vérifie si tous les arguments passés respectent une condition.

<details>
<summary>Solution</summary>

```go
package main

import "fmt"

func tousPositifs(nombres ...int) (bool, []int) {
    var negatifs []int

    for _, nombre := range nombres {
        if nombre < 0 {
            negatifs = append(negatifs, nombre)
        }
    }

    return len(negatifs) == 0, negatifs
}

func tousSuperieurA(limite int, nombres ...int) (bool, []int) {
    var inferieurs []int

    for _, nombre := range nombres {
        if nombre <= limite {
            inferieurs = append(inferieurs, nombre)
        }
    }

    return len(inferieurs) == 0, inferieurs
}

func main() {
    // Test tous positifs
    fmt.Println("=== Test tous positifs ===")
    tests := [][]int{
        {1, 2, 3, 4, 5},
        {-1, 2, 3},
        {0, 1, 2},
        {},
    }

    for _, test := range tests {
        valide, invalides := tousPositifs(test...)
        fmt.Printf("Nombres: %v\n", test)
        if valide {
            fmt.Println("  ✅ Tous positifs")
        } else {
            fmt.Printf("  ❌ Nombres négatifs trouvés: %v\n", invalides)
        }
        fmt.Println()
    }

    // Test supérieurs à 10
    fmt.Println("=== Test tous supérieurs à 10 ===")
    nombres := []int{15, 20, 5, 25, 8, 30}
    valide, invalides := tousSuperieurA(10, nombres...)
    fmt.Printf("Nombres: %v\n", nombres)
    if valide {
        fmt.Println("  ✅ Tous supérieurs à 10")
    } else {
        fmt.Printf("  ❌ Nombres ≤ 10 trouvés: %v\n", invalides)
    }
}
```

</details>

## Patterns avancés avec les fonctions variadiques

### Composition de fonctions

```go
package main

import "fmt"

// Fonction qui applique plusieurs transformations
func appliquerTransformations(nombre int, transformations ...func(int) int) int {
    resultat := nombre

    fmt.Printf("Nombre initial: %d\n", resultat)

    for i, transformation := range transformations {
        ancien := resultat
        resultat = transformation(resultat)
        fmt.Printf("  Étape %d: %d -> %d\n", i+1, ancien, resultat)
    }

    return resultat
}

func doubler(x int) int    { return x * 2 }
func ajouter10(x int) int  { return x + 10 }
func carre(x int) int      { return x * x }

func main() {
    resultat := appliquerTransformations(
        5,
        doubler,
        ajouter10,
        carre,
    )

    fmt.Printf("Résultat final: %d\n", resultat)
}
```

**Résultat :**

```
Nombre initial: 5
  Étape 1: 5 -> 10
  Étape 2: 10 -> 20
  Étape 3: 20 -> 400
Résultat final: 400
```

## Erreurs courantes à éviter

### 1. Paramètre variadique pas en dernier

```go
// ❌ Incorrect - le paramètre variadique doit être en dernier
func mauvaise(params ...string, autreParam int) { }

// ✅ Correct
func bonne(autreParam int, params ...string) { }
```

### 2. Oublier les ... pour passer un slice

```go
func somme(nombres ...int) int {
    total := 0
    for _, n := range nombres {
        total += n
    }
    return total
}

func main() {
    mesNombres := []int{1, 2, 3, 4, 5}

    // ❌ Ceci ne compile pas
    // resultat := somme(mesNombres)

    // ✅ Correct
    resultat := somme(mesNombres...)
    fmt.Println(resultat)
}
```

### 3. Confusion entre slice et arguments variadiques

```go
// Ces deux approches sont différentes :

// Approche 1 : fonction variadique
func methode1(nombres ...int) {
    // nombres est un []int à l'intérieur de la fonction
}

// Approche 2 : fonction qui prend un slice
func methode2(nombres []int) {
    // nombres est un []int passé explicitement
}

func main() {
    // Pour methode1 (variadique)
    methode1(1, 2, 3)           // ✅ OK
    methode1([]int{1, 2, 3}...) // ✅ OK avec ...

    // Pour methode2 (slice normal)
    methode2([]int{1, 2, 3})    // ✅ OK
    // methode2(1, 2, 3)        // ❌ Ne compile pas
}
```

## Résumé

Les fonctions variadiques sont un outil puissant de Go qui permet de :

**Créer des APIs flexibles** : Vos fonctions peuvent accepter un nombre variable d'arguments, rendant leur utilisation plus naturelle.

**Simplifier le code** : Plus besoin de créer des slices manuellement pour passer plusieurs éléments.

**Améliorer la lisibilité** : Le code devient plus expressif et proche du langage naturel.

**Points clés à retenir :**

- Syntaxe : `param ...type`
- Le paramètre variadique devient un slice dans la fonction
- Doit toujours être le dernier paramètre
- Utilisez `slice...` pour passer un slice existant
- `interface{}` permet d'accepter différents types

**Cas d'usage courants :**

- Fonctions de calcul (somme, moyenne, maximum)
- Logging et debugging
- Constructeurs flexibles
- Validation de données
- Composition de fonctions

Dans la section suivante, nous explorerons les fonctions anonymes et les closures, des concepts avancés qui donnent encore plus de flexibilité à vos programmes Go.

⏭️
