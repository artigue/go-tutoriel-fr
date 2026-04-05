🔝 Retour au [Sommaire](/SOMMAIRE.md)

# 3-1 : Conditions (if/else, switch)

## Introduction

Les conditions sont des structures qui permettent à votre programme de prendre des décisions. Imaginez-les comme des carrefours dans votre code : selon la situation, votre programme va emprunter une route ou une autre.

## 1. La structure if/else

### Syntaxe de base

La structure `if` est la plus simple pour tester une condition :

```go
if condition {
    // Code exécuté si la condition est vraie
}
```

**Exemple concret :**

```go
package main

import "fmt"

func main() {
    age := 18

    if age >= 18 {
        fmt.Println("Vous êtes majeur")
    }
}
```

### Ajouter une alternative avec else

```go
if condition {
    // Code si la condition est vraie
} else {
    // Code si la condition est fausse
}
```

**Exemple :**

```go
package main

import "fmt"

func main() {
    temperature := 25

    if temperature > 30 {
        fmt.Println("Il fait chaud")
    } else {
        fmt.Println("Il fait bon")
    }
}
```

### Plusieurs conditions avec else if

```go
if condition1 {
    // Code pour condition1
} else if condition2 {
    // Code pour condition2
} else {
    // Code si aucune condition n'est vraie
}
```

**Exemple pratique :**

```go
package main

import "fmt"

func main() {
    note := 85

    if note >= 90 {
        fmt.Println("Excellent !")
    } else if note >= 80 {
        fmt.Println("Très bien")
    } else if note >= 70 {
        fmt.Println("Bien")
    } else if note >= 60 {
        fmt.Println("Passable")
    } else {
        fmt.Println("Insuffisant")
    }
}
```

## 2. Déclaration de variables dans le if

Go permet de déclarer une variable directement dans la condition. C'est très pratique !

```go
if variable := expression; condition {
    // La variable n'existe que dans ce bloc
}
```

**Exemple :**

```go
package main

import "fmt"

func main() {
    // La variable 'resultat' n'existe que dans le bloc if
    if resultat := 10 * 5; resultat > 40 {
        fmt.Printf("Le résultat %d est supérieur à 40\n", resultat)
    }

    // fmt.Println(resultat) // Erreur ! resultat n'existe plus ici
}
```

**Cas d'usage courant avec les erreurs :**

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    texte := "123"

    // Conversion et vérification d'erreur en une ligne
    if nombre, err := strconv.Atoi(texte); err == nil {
        fmt.Printf("Nombre converti : %d\n", nombre)
    } else {
        fmt.Printf("Erreur de conversion : %v\n", err)
    }
}
```

## 3. Opérateurs de comparaison

Pour créer des conditions, vous devez connaître les opérateurs :

| Opérateur | Signification |
|-----------|---------------|
| `==` | Égal à |
| `!=` | Différent de |
| `<` | Inférieur à |
| `<=` | Inférieur ou égal à |
| `>` | Supérieur à |
| `>=` | Supérieur ou égal à |

**Exemples :**

```go
package main

import "fmt"

func main() {
    a := 10
    b := 20

    if a == b {
        fmt.Println("a est égal à b")
    }

    if a != b {
        fmt.Println("a est différent de b")
    }

    if a < b {
        fmt.Println("a est inférieur à b")
    }
}
```

## 4. Opérateurs logiques

Vous pouvez combiner plusieurs conditions :

| Opérateur | Signification |
|-----------|---------------|
| `&&` | ET logique (AND) |
| `\|\|` | OU logique (OR) |
| `!` | NON logique (NOT) |

**Exemples :**

```go
package main

import "fmt"

func main() {
    age := 25
    permis := true

    // ET logique - les deux conditions doivent être vraies
    if age >= 18 && permis {
        fmt.Println("Vous pouvez conduire")
    }

    // OU logique - au moins une condition doit être vraie
    if age < 18 || !permis {
        fmt.Println("Vous ne pouvez pas conduire")
    }

    // NON logique - inverse la condition
    if !permis {
        fmt.Println("Vous n'avez pas le permis")
    }
}
```

## 5. La structure switch

Le `switch` est utile quand vous avez beaucoup de conditions à tester sur la même variable.

### Syntaxe de base

```go
switch variable {
case valeur1:
    // Code pour valeur1
case valeur2:
    // Code pour valeur2
default:
    // Code si aucune valeur ne correspond
}
```

**Exemple simple :**

```go
package main

import "fmt"

func main() {
    jour := "lundi"

    switch jour {
    case "lundi":
        fmt.Println("Début de semaine")
    case "mardi", "mercredi", "jeudi":
        fmt.Println("Milieu de semaine")
    case "vendredi":
        fmt.Println("Fin de semaine")
    case "samedi", "dimanche":
        fmt.Println("Week-end")
    default:
        fmt.Println("Jour inconnu")
    }
}
```

### Switch avec des conditions

```go
package main

import "fmt"

func main() {
    note := 85

    switch {
    case note >= 90:
        fmt.Println("Excellent")
    case note >= 80:
        fmt.Println("Très bien")
    case note >= 70:
        fmt.Println("Bien")
    case note >= 60:
        fmt.Println("Passable")
    default:
        fmt.Println("Insuffisant")
    }
}
```

### Switch avec initialisation

```go
package main

import "fmt"

func main() {
    switch heure := 14; {
    case heure < 12:
        fmt.Println("Bonjour")
    case heure < 18:
        fmt.Println("Bon après-midi")
    default:
        fmt.Println("Bonsoir")
    }
}
```

## 6. Différences importantes avec d'autres langages

### Pas de break nécessaire

En Go, chaque `case` se termine automatiquement. Pas besoin de `break` :

```go
switch x {
case 1:
    fmt.Println("Un")
    // Pas de break nécessaire
case 2:
    fmt.Println("Deux")
}
```

### Utiliser fallthrough pour continuer

Si vous voulez vraiment continuer vers le case suivant :

```go
package main

import "fmt"

func main() {
    note := 85

    switch {
    case note >= 80:
        fmt.Println("Très bien")
        fallthrough
    case note >= 60:
        fmt.Println("Vous avez réussi")
    default:
        fmt.Println("Essayez encore")
    }
}
```

## 7. Exemples pratiques

### Exemple 1 : Calculateur simple

```go
package main

import "fmt"

func main() {
    a := 10
    b := 5
    operation := "+"

    switch operation {
    case "+":
        fmt.Printf("%d + %d = %d\n", a, b, a+b)
    case "-":
        fmt.Printf("%d - %d = %d\n", a, b, a-b)
    case "*":
        fmt.Printf("%d * %d = %d\n", a, b, a*b)
    case "/":
        if b != 0 {
            fmt.Printf("%d / %d = %d\n", a, b, a/b)
        } else {
            fmt.Println("Division par zéro impossible")
        }
    default:
        fmt.Println("Opération inconnue")
    }
}
```

### Exemple 2 : Vérification d'âge

```go
package main

import "fmt"

func main() {
    age := 16

    if age < 0 {
        fmt.Println("Âge invalide")
    } else if age < 13 {
        fmt.Println("Enfant")
    } else if age < 18 {
        fmt.Println("Adolescent")
    } else if age < 65 {
        fmt.Println("Adulte")
    } else {
        fmt.Println("Senior")
    }
}
```

### Exemple 3 : Validation de données

```go
package main

import "fmt"

func main() {
    username := "john_doe"
    password := "secret123"

    if len(username) < 3 {
        fmt.Println("Le nom d'utilisateur est trop court")
    } else if len(password) < 8 {
        fmt.Println("Le mot de passe est trop court")
    } else {
        fmt.Println("Identifiants valides")
    }
}
```

## 8. Bonnes pratiques

### 1. Utilisez le "early return" pattern

```go
func validateUser(age int, name string) error {
    if age < 0 {
        return fmt.Errorf("âge invalide")
    }

    if name == "" {
        return fmt.Errorf("nom requis")
    }

    // Traitement principal
    return nil
}
```

### 2. Préférez switch pour de multiples conditions

```go
// Préférez ceci
switch status {
case "active", "pending":
    processUser()
case "inactive":
    archiveUser()
default:
    handleError()
}

// Plutôt que ceci
if status == "active" || status == "pending" {
    processUser()
} else if status == "inactive" {
    archiveUser()
} else {
    handleError()
}
```

### 3. Évitez l'imbrication excessive

```go
// Difficile à lire
if condition1 {
    if condition2 {
        if condition3 {
            // code
        }
    }
}

// Plus lisible
if !condition1 {
    return
}
if !condition2 {
    return
}
if !condition3 {
    return
}
// code
```

## 9. Exercices pratiques

### Exercice 1 : Calculateur de TVA

Créez un programme qui calcule le prix TTC selon le pays :

- France : 20%
- Belgique : 21%
- Suisse : 7.7%
- Autre : 0%

### Exercice 2 : Système de notes

Créez un programme qui convertit une note numérique en lettre :

- 90-100 : A
- 80-89 : B
- 70-79 : C
- 60-69 : D
- 0-59 : F

### Exercice 3 : Validation d'email

Créez un programme qui valide basiquement un email :

- Doit contenir "@"
- Doit contenir "."
- Ne doit pas être vide

## Solutions des Exercices - Conditions (if/else, switch)

### Exercice 1 : Calculateur de TVA

#### Solution avec switch

```go
package main

import "fmt"

func main() {
    // Exemple d'utilisation
    prixHT := 100.0
    pays := "France"

    // Calcul du prix TTC selon le pays
    var tauxTVA float64

    switch pays {
    case "France":
        tauxTVA = 0.20 // 20%
    case "Belgique":
        tauxTVA = 0.21 // 21%
    case "Suisse":
        tauxTVA = 0.077 // 7.7%
    default:
        tauxTVA = 0.0 // 0% pour les autres pays
    }

    prixTTC := prixHT * (1 + tauxTVA)

    fmt.Printf("Prix HT : %.2f€\n", prixHT)
    fmt.Printf("Pays : %s\n", pays)
    fmt.Printf("Taux TVA : %.1f%%\n", tauxTVA*100)
    fmt.Printf("Prix TTC : %.2f€\n", prixTTC)
}
```

#### Solution avec fonction réutilisable

```go
package main

import "fmt"

// Fonction pour calculer le prix TTC
func calculerPrixTTC(prixHT float64, pays string) (float64, float64) {
    var tauxTVA float64

    switch pays {
    case "France":
        tauxTVA = 0.20
    case "Belgique":
        tauxTVA = 0.21
    case "Suisse":
        tauxTVA = 0.077
    default:
        tauxTVA = 0.0
    }

    prixTTC := prixHT * (1 + tauxTVA)
    return prixTTC, tauxTVA
}

func main() {
    // Tests avec différents pays
    tests := []struct {
        prix float64
        pays string
    }{
        {100.0, "France"},
        {50.0, "Belgique"},
        {200.0, "Suisse"},
        {75.0, "Canada"},
    }

    for _, test := range tests {
        prixTTC, tauxTVA := calculerPrixTTC(test.prix, test.pays)
        fmt.Printf("--- %s ---\n", test.pays)
        fmt.Printf("Prix HT : %.2f€\n", test.prix)
        fmt.Printf("TVA : %.1f%%\n", tauxTVA*100)
        fmt.Printf("Prix TTC : %.2f€\n\n", prixTTC)
    }
}
```

#### Solution interactive

```go
package main

import "fmt"

func main() {
    var prixHT float64
    var pays string

    // Demander les informations à l'utilisateur
    fmt.Print("Entrez le prix HT : ")
    fmt.Scanln(&prixHT)

    fmt.Print("Entrez le pays (France/Belgique/Suisse) : ")
    fmt.Scanln(&pays)

    // Validation du prix
    if prixHT < 0 {
        fmt.Println("Erreur : le prix ne peut pas être négatif")
        return
    }

    // Calcul de la TVA
    var tauxTVA float64
    var nomTVA string

    switch pays {
    case "France":
        tauxTVA = 0.20
        nomTVA = "20%"
    case "Belgique":
        tauxTVA = 0.21
        nomTVA = "21%"
    case "Suisse":
        tauxTVA = 0.077
        nomTVA = "7.7%"
    default:
        tauxTVA = 0.0
        nomTVA = "0%"
        fmt.Printf("Pays '%s' non reconnu, TVA par défaut : 0%%\n", pays)
    }

    montantTVA := prixHT * tauxTVA
    prixTTC := prixHT + montantTVA

    // Affichage du résultat
    fmt.Println("\n--- FACTURE ---")
    fmt.Printf("Prix HT : %.2f€\n", prixHT)
    fmt.Printf("TVA (%s) : %.2f€\n", nomTVA, montantTVA)
    fmt.Printf("Prix TTC : %.2f€\n", prixTTC)
}
```

### Exercice 2 : Système de notes

#### Solution avec if/else

```go
package main

import "fmt"

func main() {
    // Exemple de notes à convertir
    notes := []int{95, 87, 73, 62, 45, 101, -5}

    for _, note := range notes {
        var lettre string

        // Validation de la note
        if note < 0 || note > 100 {
            fmt.Printf("Note %d : INVALIDE (doit être entre 0 et 100)\n", note)
            continue
        }

        // Conversion en lettre
        if note >= 90 {
            lettre = "A"
        } else if note >= 80 {
            lettre = "B"
        } else if note >= 70 {
            lettre = "C"
        } else if note >= 60 {
            lettre = "D"
        } else {
            lettre = "F"
        }

        fmt.Printf("Note %d : %s\n", note, lettre)
    }
}
```

#### Solution avec switch

```go
package main

import "fmt"

func obtenirLettre(note int) string {
    // Validation
    if note < 0 || note > 100 {
        return "INVALIDE"
    }

    // Utilisation du switch avec des tranches
    switch {
    case note >= 90:
        return "A"
    case note >= 80:
        return "B"
    case note >= 70:
        return "C"
    case note >= 60:
        return "D"
    default:
        return "F"
    }
}

func main() {
    // Test avec différentes notes
    notes := []int{100, 95, 89, 85, 79, 75, 69, 65, 59, 0, 105, -10}

    fmt.Println("Conversion des notes :")
    for _, note := range notes {
        lettre := obtenirLettre(note)
        if lettre == "INVALIDE" {
            fmt.Printf("Note %d : %s\n", note, lettre)
        } else {
            fmt.Printf("Note %d : %s\n", note, lettre)
        }
    }
}
```

#### Solution avec commentaires détaillés

```go
package main

import "fmt"

// Fonction pour convertir une note numérique en lettre
func convertirNote(note int) (string, string) {
    // Validation de la note
    if note < 0 || note > 100 {
        return "INVALIDE", "La note doit être entre 0 et 100"
    }

    // Conversion avec commentaires
    switch {
    case note >= 90:
        return "A", "Excellent"
    case note >= 80:
        return "B", "Très bien"
    case note >= 70:
        return "C", "Bien"
    case note >= 60:
        return "D", "Passable"
    default:
        return "F", "Insuffisant"
    }
}

func main() {
    // Demander une note à l'utilisateur
    var note int
    fmt.Print("Entrez une note (0-100) : ")
    fmt.Scanln(&note)

    lettre, commentaire := convertirNote(note)

    if lettre == "INVALIDE" {
        fmt.Printf("Erreur : %s\n", commentaire)
    } else {
        fmt.Printf("Note : %d\n", note)
        fmt.Printf("Lettre : %s\n", lettre)
        fmt.Printf("Commentaire : %s\n", commentaire)
    }
}
```

### Exercice 3 : Validation d'email

#### Solution basique

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    // Emails de test
    emails := []string{
        "john@example.com",
        "marie.dupont@gmail.com",
        "invalide@",
        "@example.com",
        "sans-arobase.com",
        "",
        "test@domain",
        "user@domain.co.uk",
    }

    for _, email := range emails {
        if estEmailValide(email) {
            fmt.Printf("✓ '%s' est valide\n", email)
        } else {
            fmt.Printf("✗ '%s' est invalide\n", email)
        }
    }
}

func estEmailValide(email string) bool {
    // Vérifier si l'email n'est pas vide
    if email == "" {
        return false
    }

    // Vérifier la présence de "@"
    if !strings.Contains(email, "@") {
        return false
    }

    // Vérifier la présence de "."
    if !strings.Contains(email, ".") {
        return false
    }

    return true
}
```

#### Solution avec validation détaillée

```go
package main

import (
    "fmt"
    "strings"
)

// Fonction de validation avec messages d'erreur détaillés
func validerEmail(email string) (bool, string) {
    // Vérifier si l'email n'est pas vide
    if email == "" {
        return false, "L'email ne peut pas être vide"
    }

    // Vérifier la présence de "@"
    if !strings.Contains(email, "@") {
        return false, "L'email doit contenir '@'"
    }

    // Vérifier la présence de "."
    if !strings.Contains(email, ".") {
        return false, "L'email doit contenir '.'"
    }

    // Vérifications supplémentaires
    if strings.HasPrefix(email, "@") {
        return false, "L'email ne peut pas commencer par '@'"
    }

    if strings.HasSuffix(email, "@") {
        return false, "L'email ne peut pas finir par '@'"
    }

    // Vérifier qu'il n'y a qu'un seul "@"
    if strings.Count(email, "@") != 1 {
        return false, "L'email doit contenir exactement un '@'"
    }

    // Séparer la partie locale et le domaine
    parties := strings.Split(email, "@")
    partieLocale := parties[0]
    domaine := parties[1]

    // Vérifier que la partie locale n'est pas vide
    if partieLocale == "" {
        return false, "La partie avant '@' ne peut pas être vide"
    }

    // Vérifier que le domaine n'est pas vide
    if domaine == "" {
        return false, "La partie après '@' ne peut pas être vide"
    }

    // Vérifier que le domaine contient un point
    if !strings.Contains(domaine, ".") {
        return false, "Le domaine doit contenir un point"
    }

    return true, "Email valide"
}

func main() {
    // Test interactif
    var email string
    fmt.Print("Entrez une adresse email : ")
    fmt.Scanln(&email)

    valide, message := validerEmail(email)

    if valide {
        fmt.Printf("✓ %s\n", message)
    } else {
        fmt.Printf("✗ %s\n", message)
    }
}
```

#### Solution avec structure de données

```go
package main

import (
    "fmt"
    "strings"
)

// Structure pour représenter un résultat de validation
type ResultatValidation struct {
    Valide   bool
    Message  string
    Erreurs  []string
}

// Fonction complète de validation
func validerEmailComplet(email string) ResultatValidation {
    resultat := ResultatValidation{
        Valide:  true,
        Message: "",
        Erreurs: []string{},
    }

    // Vérifications une par une
    if email == "" {
        resultat.Valide = false
        resultat.Erreurs = append(resultat.Erreurs, "L'email ne peut pas être vide")
    }

    if !strings.Contains(email, "@") {
        resultat.Valide = false
        resultat.Erreurs = append(resultat.Erreurs, "L'email doit contenir '@'")
    }

    if !strings.Contains(email, ".") {
        resultat.Valide = false
        resultat.Erreurs = append(resultat.Erreurs, "L'email doit contenir '.'")
    }

    if strings.Count(email, "@") > 1 {
        resultat.Valide = false
        resultat.Erreurs = append(resultat.Erreurs, "L'email ne peut contenir qu'un seul '@'")
    }

    if strings.HasPrefix(email, "@") || strings.HasSuffix(email, "@") {
        resultat.Valide = false
        resultat.Erreurs = append(resultat.Erreurs, "L'email ne peut pas commencer ou finir par '@'")
    }

    // Message final
    if resultat.Valide {
        resultat.Message = "Email valide"
    } else {
        resultat.Message = "Email invalide"
    }

    return resultat
}

func main() {
    // Test avec plusieurs emails
    emails := []string{
        "john@example.com",
        "marie.dupont@gmail.com",
        "invalide@",
        "@example.com",
        "sans-arobase.com",
        "",
        "multiple@@signs.com",
        "user@domain.co.uk",
    }

    fmt.Println("Validation des emails :")
    fmt.Println("=====================")

    for _, email := range emails {
        resultat := validerEmailComplet(email)

        if resultat.Valide {
            fmt.Printf("✓ '%s' : %s\n", email, resultat.Message)
        } else {
            fmt.Printf("✗ '%s' : %s\n", email, resultat.Message)
            for _, erreur := range resultat.Erreurs {
                fmt.Printf("  - %s\n", erreur)
            }
        }
        fmt.Println()
    }
}
```

### Points clés des solutions

#### Exercice 1 - TVA

- **Switch** pour les différents pays
- **Validation** des données d'entrée
- **Fonctions réutilisables** pour un code propre
- **Interface utilisateur** simple

#### Exercice 2 - Notes

- **Validation** des notes (0-100)
- **Switch avec conditions** pour les tranches
- **Messages descriptifs** pour chaque note
- **Gestion des cas limites**

#### Exercice 3 - Email

- **Validation progressive** avec strings.Contains
- **Gestion des erreurs** multiples
- **Structure de données** pour les résultats
- **Messages d'erreur** détaillés

Ces solutions montrent différentes approches : du simple au complexe.

## Résumé

Les conditions sont essentielles pour créer des programmes intelligents. Retenez :

- `if/else` pour des conditions simples
- `switch` pour de multiples choix
- Possibilité de déclarer des variables dans les conditions
- Pas de parenthèses obligatoires autour des conditions
- Pas de `break` nécessaire dans `switch`
- Utilisez les bonnes pratiques pour un code lisible

Dans la prochaine section, nous verrons les boucles qui permettent de répéter des actions !

⏭️
