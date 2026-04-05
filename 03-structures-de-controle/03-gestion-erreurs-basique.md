🔝 Retour au [Sommaire](/SOMMAIRE.md)

# 3-3 : Gestion des erreurs basique

## Introduction

La gestion des erreurs est une partie fondamentale de la programmation. Imaginez-vous en train de cuisiner : que se passe-t-il si vous n'avez plus de farine ? Vous devez gérer cette situation ! En programmation, c'est pareil : les erreurs arrivent (fichier introuvable, division par zéro, connexion internet coupée), et il faut savoir les gérer proprement.

Go a une approche unique et explicite pour gérer les erreurs qui peut sembler étrange au début, mais qui devient très naturelle avec la pratique.

## 1. Philosophie de Go concernant les erreurs

### Principe fondamental

En Go, **les erreurs sont des valeurs**, pas des exceptions. Cela signifie :

- Elles sont retournées explicitement par les fonctions
- Elles doivent être vérifiées manuellement
- Le code reste lisible et prévisible

### Comparaison avec d'autres langages

```go
// ❌ Autres langages (try/catch)
// try {
//     result = riskyOperation()
// } catch (Exception e) {
//     handleError(e)
// }

// ✅ Go style
result, err := riskyOperation()
if err != nil {
    // Gérer l'erreur
    handleError(err)
}
// Continuer avec result
```

## 2. Le type error

### Qu'est-ce qu'une error ?

En Go, `error` est une interface très simple :

```go
type error interface {
    Error() string
}
```

Toute structure qui a une méthode `Error() string` peut être une erreur.

### Valeur nil

Une erreur `nil` signifie "pas d'erreur" :

```go
package main

import "fmt"

func diviser(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division par zéro impossible")
    }
    return a / b, nil // nil = pas d'erreur
}

func main() {
    // Cas normal
    resultat, err := diviser(10, 2)
    if err != nil {
        fmt.Printf("Erreur : %v\n", err)
        return
    }
    fmt.Printf("10 ÷ 2 = %.2f\n", resultat)

    // Cas d'erreur
    resultat, err = diviser(10, 0)
    if err != nil {
        fmt.Printf("Erreur : %v\n", err)
        return
    }
    fmt.Printf("Résultat : %.2f\n", resultat)
}
```

## 3. Pattern de base : vérification d'erreur

### Le pattern classique

```go
result, err := functionThatCanFail()
if err != nil {
    // Gérer l'erreur
    return err // ou log, ou autre action
}
// Utiliser result
```

### Exemple concret

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    textes := []string{"123", "45.67", "abc", ""}

    for _, texte := range textes {
        fmt.Printf("\nConversion de '%s' : ", texte)

        // Tentative de conversion
        nombre, err := strconv.Atoi(texte)
        if err != nil {
            fmt.Printf("Erreur - %v", err)
        } else {
            fmt.Printf("Succès - %d", nombre)
        }
    }
}
```

## 4. Créer des erreurs

### Avec fmt.Errorf

```go
package main

import "fmt"

func verifierAge(age int) error {
    if age < 0 {
        return fmt.Errorf("âge invalide : %d (doit être positif)", age)
    }
    if age > 150 {
        return fmt.Errorf("âge invalide : %d (trop élevé)", age)
    }
    return nil
}

func main() {
    ages := []int{25, -5, 200, 0}

    for _, age := range ages {
        err := verifierAge(age)
        if err != nil {
            fmt.Printf("Âge %d : %v\n", age, err)
        } else {
            fmt.Printf("Âge %d : valide\n", age)
        }
    }
}
```

### Avec errors.New

```go
package main

import (
    "errors"
    "fmt"
)

func calculerRacine(x float64) (float64, error) {
    if x < 0 {
        return 0, errors.New("impossible de calculer la racine d'un nombre négatif")
    }
    // Simulation d'un calcul simple
    return x * 0.5, nil // Approximation grossière
}

func main() {
    nombres := []float64{9, -4, 16, 0}

    for _, nombre := range nombres {
        racine, err := calculerRacine(nombre)
        if err != nil {
            fmt.Printf("√%.1f : Erreur - %v\n", nombre, err)
        } else {
            fmt.Printf("√%.1f = %.2f\n", nombre, racine)
        }
    }
}
```

## 5. Gestion de multiples erreurs

### Early return pattern

```go
package main

import (
    "fmt"
    "strings"
)

func validerUtilisateur(nom, email string, age int) error {
    // Vérifier le nom
    if nom == "" {
        return fmt.Errorf("le nom ne peut pas être vide")
    }
    if len(nom) < 2 {
        return fmt.Errorf("le nom doit avoir au moins 2 caractères")
    }

    // Vérifier l'email
    if email == "" {
        return fmt.Errorf("l'email ne peut pas être vide")
    }
    if !strings.Contains(email, "@") {
        return fmt.Errorf("email invalide : doit contenir @")
    }

    // Vérifier l'âge
    if age < 0 {
        return fmt.Errorf("l'âge ne peut pas être négatif")
    }
    if age > 120 {
        return fmt.Errorf("âge invalide : %d ans", age)
    }

    return nil
}

func main() {
    utilisateurs := []struct {
        nom   string
        email string
        age   int
    }{
        {"Alice", "alice@email.com", 25},
        {"", "bob@email.com", 30},
        {"Charlie", "charlie-email", 35},
        {"Diana", "diana@email.com", -5},
    }

    for i, user := range utilisateurs {
        err := validerUtilisateur(user.nom, user.email, user.age)
        if err != nil {
            fmt.Printf("Utilisateur %d : ❌ %v\n", i+1, err)
        } else {
            fmt.Printf("Utilisateur %d : ✅ Valide\n", i+1)
        }
    }
}
```

## 6. Fonctions avec fichiers (exemple pratique)

```go
package main

import (
    "fmt"
    "os"
    "strings"
)

func lireFichier(nomFichier string) (string, error) {
    // Tentative de lecture du fichier
    contenu, err := os.ReadFile(nomFichier)
    if err != nil {
        return "", fmt.Errorf("impossible de lire le fichier '%s' : %w", nomFichier, err)
    }

    return string(contenu), nil
}

func compterMots(texte string) int {
    if texte == "" {
        return 0
    }
    mots := strings.Fields(texte)
    return len(mots)
}

func analyserFichier(nomFichier string) {
    fmt.Printf("Analyse du fichier '%s' :\n", nomFichier)

    contenu, err := lireFichier(nomFichier)
    if err != nil {
        fmt.Printf("❌ Erreur : %v\n\n", err)
        return
    }

    // Si on arrive ici, pas d'erreur
    nombreMots := compterMots(contenu)
    nombreLignes := len(strings.Split(contenu, "\n"))
    nombreCaracteres := len(contenu)

    fmt.Printf("✅ Fichier lu avec succès\n")
    fmt.Printf("   Caractères : %d\n", nombreCaracteres)
    fmt.Printf("   Mots : %d\n", nombreMots)
    fmt.Printf("   Lignes : %d\n\n", nombreLignes)
}

func main() {
    fichiers := []string{
        "test.txt",      // Fichier qui pourrait exister
        "inexistant.txt", // Fichier qui n'existe pas
        "data.txt",      // Autre fichier de test
    }

    for _, fichier := range fichiers {
        analyserFichier(fichier)
    }
}
```

## 7. Différentes stratégies de gestion

### 1. Propager l'erreur (return)

```go
func functionA() error {
    err := functionB()
    if err != nil {
        return err // Propager l'erreur vers le niveau supérieur
    }
    return nil
}
```

### 2. Logger et continuer

```go
package main

import (
    "fmt"
    "log"
)

func traiterElement(element string) error {
    if element == "" {
        return fmt.Errorf("élément vide")
    }
    // Traitement...
    return nil
}

func traiterListe(elements []string) {
    for i, element := range elements {
        err := traiterElement(element)
        if err != nil {
            log.Printf("Erreur sur l'élément %d : %v", i, err)
            continue // Continuer avec les autres éléments
        }
        fmt.Printf("Élément %d traité avec succès\n", i)
    }
}

func main() {
    elements := []string{"a", "", "c", "d"}
    traiterListe(elements)
}
```

### 3. Valeur par défaut

```go
package main

import (
    "fmt"
    "strconv"
)

func convertirAvecDefaut(texte string, defaut int) int {
    nombre, err := strconv.Atoi(texte)
    if err != nil {
        fmt.Printf("Conversion échouée pour '%s', utilisation de %d\n", texte, defaut)
        return defaut
    }
    return nombre
}

func main() {
    entrees := []string{"123", "abc", "456", "xyz"}

    for _, entree := range entrees {
        resultat := convertirAvecDefaut(entree, 0)
        fmt.Printf("'%s' → %d\n", entree, resultat)
    }
}
```

## 8. Wrapping d'erreurs (Go 1.13+)

### Ajouter du contexte

```go
package main

import (
    "fmt"
    "errors"
    "strconv"
)

func lireConfiguration(fichier string) (int, error) {
    // Simulation de lecture
    contenu := "abc" // Contenu invalide

    valeur, err := strconv.Atoi(contenu)
    if err != nil {
        // Wrapper l'erreur avec du contexte
        return 0, fmt.Errorf("erreur lors de la lecture de %s : %w", fichier, err)
    }

    return valeur, nil
}

func initialiser() error {
    _, err := lireConfiguration("config.txt")
    if err != nil {
        return fmt.Errorf("échec de l'initialisation : %w", err)
    }
    return nil
}

func main() {
    err := initialiser()
    if err != nil {
        fmt.Printf("Erreur : %v\n", err)

        // Décomposer l'erreur
        var numError *strconv.NumError
        if errors.As(err, &numError) {
            fmt.Printf("Erreur de conversion détectée : %s\n", numError.Num)
        }
    }
}
```

## 9. Exemples pratiques complets

### Calculatrice avec gestion d'erreurs

```go
package main

import (
    "fmt"
    "errors"
)

// Erreurs personnalisées
var (
    ErrDivisionParZero = errors.New("division par zéro")
    ErrOperationInconnue = errors.New("opération inconnue")
)

func calculer(a, b float64, operation string) (float64, error) {
    switch operation {
    case "+":
        return a + b, nil
    case "-":
        return a - b, nil
    case "*":
        return a * b, nil
    case "/":
        if b == 0 {
            return 0, ErrDivisionParZero
        }
        return a / b, nil
    default:
        return 0, fmt.Errorf("%w : %s", ErrOperationInconnue, operation)
    }
}

func main() {
    calculs := []struct {
        a, b float64
        op   string
    }{
        {10, 2, "+"},
        {10, 2, "/"},
        {10, 0, "/"},
        {5, 3, "%"},
    }

    for _, calc := range calculs {
        resultat, err := calculer(calc.a, calc.b, calc.op)

        if err != nil {
            fmt.Printf("%.1f %s %.1f = Erreur : %v\n",
                calc.a, calc.op, calc.b, err)

            // Gestion spécifique selon le type d'erreur
            if errors.Is(err, ErrDivisionParZero) {
                fmt.Println("  → Conseil : Vérifiez le dénominateur")
            } else if errors.Is(err, ErrOperationInconnue) {
                fmt.Println("  → Opérations disponibles : +, -, *, /")
            }
        } else {
            fmt.Printf("%.1f %s %.1f = %.2f\n",
                calc.a, calc.op, calc.b, resultat)
        }
    }
}
```

## 10. Bonnes pratiques

### 1. Messages d'erreur clairs

```go
// ✅ Bon
return fmt.Errorf("impossible de connecter à la base de données %s sur le port %d : %w",
    host, port, err)

// ❌ Moins bon
return errors.New("erreur de connexion")
```

### 2. Vérifier les erreurs immédiatement

```go
// ✅ Bon
file, err := os.Open("fichier.txt")
if err != nil {
    return err
}
defer file.Close()

// ❌ Mauvais
file, err := os.Open("fichier.txt")
defer file.Close()
if err != nil {
    return err
}
```

### 3. Ne pas ignorer les erreurs

```go
// ❌ Mauvais
data, _ := ioutil.ReadFile("config.txt")

// ✅ Bon
data, err := ioutil.ReadFile("config.txt")
if err != nil {
    log.Printf("Attention : impossible de lire config.txt : %v", err)
    // Utiliser une configuration par défaut
}
```

## 11. Exercices pratiques

### Exercice 1 : Validateur d'email

Créez une fonction qui valide un email et retourne des erreurs spécifiques pour chaque problème.

### Exercice 2 : Calculateur d'IMC

Créez un calculateur d'IMC qui gère les erreurs (poids/taille négatifs, division par zéro).

### Exercice 3 : Lecteur de fichier CSV

Créez une fonction qui lit un fichier CSV et gère toutes les erreurs possibles.

### Exercice 4 : Convertisseur d'unités

Créez un convertisseur température avec gestion d'erreurs pour les valeurs impossibles.

### Exercice 5 : Gestionnaire de mots de passe

Créez un validateur de mot de passe avec différents types d'erreurs.

# Solutions des Exercices - Gestion des erreurs basique

### Exercice 1 : Validateur d'email

#### Solution basique

```go
package main

import (
    "errors"
    "fmt"
    "strings"
)

// Erreurs spécifiques pour la validation d'email
var (
    ErrEmailVide        = errors.New("l'email ne peut pas être vide")
    ErrEmailSansArobase = errors.New("l'email doit contenir '@'")
    ErrEmailSansPoint   = errors.New("l'email doit contenir '.'")
    ErrEmailTropCourt   = errors.New("l'email est trop court")
    ErrPartieLocaleVide = errors.New("la partie avant '@' ne peut pas être vide")
    ErrDomaineVide      = errors.New("la partie après '@' ne peut pas être vide")
    ErrTropDArobase     = errors.New("l'email ne peut contenir qu'un seul '@'")
)

func validerEmail(email string) error {
    // Vérifier si l'email est vide
    if email == "" {
        return ErrEmailVide
    }

    // Vérifier la longueur minimale
    if len(email) < 5 { // minimum a@b.c
        return ErrEmailTropCourt
    }

    // Vérifier la présence de '@'
    if !strings.Contains(email, "@") {
        return ErrEmailSansArobase
    }

    // Vérifier qu'il n'y a qu'un seul '@'
    if strings.Count(email, "@") != 1 {
        return ErrTropDArobase
    }

    // Séparer en parties
    parties := strings.Split(email, "@")
    partieLocale := parties[0]
    domaine := parties[1]

    // Vérifier la partie locale
    if partieLocale == "" {
        return ErrPartieLocaleVide
    }

    // Vérifier le domaine
    if domaine == "" {
        return ErrDomaineVide
    }

    // Vérifier la présence de '.' dans le domaine
    if !strings.Contains(domaine, ".") {
        return ErrEmailSansPoint
    }

    return nil // Email valide
}

func main() {
    emails := []string{
        "john@example.com",
        "",
        "abc",
        "test@",
        "@example.com",
        "user@@domain.com",
        "a@b",
        "valid.email@domain.co.uk",
        "no-at-symbol.com",
    }

    fmt.Println("=== Validation d'emails ===")
    for _, email := range emails {
        err := validerEmail(email)
        if err != nil {
            fmt.Printf("❌ '%s' : %v\n", email, err)
        } else {
            fmt.Printf("✅ '%s' : Email valide\n", email)
        }
    }
}
```

#### Solution avancée avec structure de résultat

```go
package main

import (
    "errors"
    "fmt"
    "regexp"
    "strings"
)

type ResultatValidation struct {
    Valide  bool
    Erreurs []error
    Niveau  string // "critique", "avertissement", "info"
}

func (r *ResultatValidation) AjouterErreur(err error, niveau string) {
    r.Erreurs = append(r.Erreurs, err)
    r.Valide = false
    if niveau == "critique" {
        r.Niveau = "critique"
    }
}

func validerEmailAvance(email string) ResultatValidation {
    resultat := ResultatValidation{
        Valide:  true,
        Erreurs: []error{},
        Niveau:  "valide",
    }

    // Vérifications critiques
    if email == "" {
        resultat.AjouterErreur(errors.New("email vide"), "critique")
        return resultat
    }

    if len(email) < 3 {
        resultat.AjouterErreur(errors.New("email trop court"), "critique")
        return resultat
    }

    if !strings.Contains(email, "@") {
        resultat.AjouterErreur(errors.New("manque '@'"), "critique")
        return resultat
    }

    if strings.Count(email, "@") != 1 {
        resultat.AjouterErreur(errors.New("trop d'@"), "critique")
        return resultat
    }

    // Séparer en parties
    parties := strings.Split(email, "@")
    partieLocale, domaine := parties[0], parties[1]

    if partieLocale == "" {
        resultat.AjouterErreur(errors.New("partie locale vide"), "critique")
    }

    if domaine == "" {
        resultat.AjouterErreur(errors.New("domaine vide"), "critique")
    }

    if !strings.Contains(domaine, ".") {
        resultat.AjouterErreur(errors.New("domaine sans point"), "critique")
    }

    // Vérifications avec regex (optionnel)
    if resultat.Valide {
        regex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
        if !regex.MatchString(email) {
            resultat.AjouterErreur(errors.New("format non conforme RFC"), "avertissement")
            resultat.Niveau = "avertissement"
            resultat.Valide = false
        }
    }

    return resultat
}

func main() {
    emails := []string{
        "john.doe@example.com",
        "user@domain",
        "test@@double.com",
        "valid123@test-domain.co.uk",
        "",
        "no-at.com",
    }

    for _, email := range emails {
        resultat := validerEmailAvance(email)

        fmt.Printf("\nEmail : '%s'\n", email)
        if resultat.Valide {
            fmt.Println("✅ Valide")
        } else {
            fmt.Printf("❌ Invalide (%s)\n", resultat.Niveau)
            for _, err := range resultat.Erreurs {
                fmt.Printf("  - %v\n", err)
            }
        }
    }
}
```

### Exercice 2 : Calculateur d'IMC

#### Solution complète

```go
package main

import (
    "errors"
    "fmt"
)

// Erreurs spécifiques à l'IMC
var (
    ErrPoidsNegatif     = errors.New("le poids ne peut pas être négatif")
    ErrPoidsZero        = errors.New("le poids ne peut pas être zéro")
    ErrTailleNegative   = errors.New("la taille ne peut pas être négative")
    ErrTailleZero       = errors.New("la taille ne peut pas être zéro")
    ErrPoidsExcessif    = errors.New("poids irréaliste (>1000kg)")
    ErrTailleExcessive  = errors.New("taille irréaliste (>3m)")
    ErrTailleTropPetite = errors.New("taille irréaliste (<0.5m)")
)

type IMC struct {
    Valeur      float64
    Categorie   string
    Risque      string
}

func calculerIMC(poids, taille float64) (*IMC, error) {
    // Validation du poids
    if poids < 0 {
        return nil, ErrPoidsNegatif
    }
    if poids == 0 {
        return nil, ErrPoidsZero
    }
    if poids > 1000 {
        return nil, ErrPoidsExcessif
    }

    // Validation de la taille
    if taille < 0 {
        return nil, ErrTailleNegative
    }
    if taille == 0 {
        return nil, ErrTailleZero
    }
    if taille < 0.5 {
        return nil, ErrTailleTropPetite
    }
    if taille > 3.0 {
        return nil, ErrTailleExcessive
    }

    // Calcul de l'IMC
    valeurIMC := poids / (taille * taille)

    // Déterminer la catégorie
    var categorie, risque string
    switch {
    case valeurIMC < 16.5:
        categorie = "Dénutrition"
        risque = "Élevé"
    case valeurIMC < 18.5:
        categorie = "Maigreur"
        risque = "Modéré"
    case valeurIMC < 25:
        categorie = "Corpulence normale"
        risque = "Faible"
    case valeurIMC < 30:
        categorie = "Surpoids"
        risque = "Modéré"
    case valeurIMC < 35:
        categorie = "Obésité modérée"
        risque = "Élevé"
    case valeurIMC < 40:
        categorie = "Obésité sévère"
        risque = "Très élevé"
    default:
        categorie = "Obésité massive"
        risque = "Extrême"
    }

    return &IMC{
        Valeur:    valeurIMC,
        Categorie: categorie,
        Risque:    risque,
    }, nil
}

func main() {
    // Cas de test
    tests := []struct {
        nom    string
        poids  float64
        taille float64
    }{
        {"Personne normale", 70, 1.75},
        {"Poids négatif", -10, 1.70},
        {"Taille zéro", 70, 0},
        {"Personne très grande", 80, 2.10},
        {"Bébé", 5, 0.6},
        {"Cas extrême", 200, 1.60},
        {"Taille négative", 70, -1.70},
        {"Poids excessif", 1500, 1.80},
    }

    fmt.Println("=== Calculateur d'IMC ===")

    for _, test := range tests {
        fmt.Printf("\n%s (%.1fkg, %.2fm) :\n", test.nom, test.poids, test.taille)

        imc, err := calculerIMC(test.poids, test.taille)
        if err != nil {
            fmt.Printf("❌ Erreur : %v\n", err)
        } else {
            fmt.Printf("✅ IMC : %.1f\n", imc.Valeur)
            fmt.Printf("   Catégorie : %s\n", imc.Categorie)
            fmt.Printf("   Risque : %s\n", imc.Risque)
        }
    }

    // Version interactive
    fmt.Println("\n=== Mode interactif ===")
    var poids, taille float64

    fmt.Print("Entrez votre poids (kg) : ")
    fmt.Scanln(&poids)

    fmt.Print("Entrez votre taille (m) : ")
    fmt.Scanln(&taille)

    imc, err := calculerIMC(poids, taille)
    if err != nil {
        fmt.Printf("Erreur : %v\n", err)
    } else {
        fmt.Printf("\nVotre IMC : %.1f\n", imc.Valeur)
        fmt.Printf("Catégorie : %s\n", imc.Categorie)
        fmt.Printf("Niveau de risque : %s\n", imc.Risque)
    }
}
```

### Exercice 3 : Lecteur de fichier CSV

#### Solution avec gestion complète des erreurs

```go
package main

import (
    "encoding/csv"
    "errors"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Erreurs spécifiques CSV
var (
    ErrFichierInexistant = errors.New("fichier inexistant")
    ErrFichierVide       = errors.New("fichier vide")
    ErrFormatCSV         = errors.New("format CSV invalide")
    ErrColonnesManquantes = errors.New("colonnes manquantes")
    ErrDonneesInvalides  = errors.New("données invalides")
)

type DonneesCSV struct {
    Entetes []string
    Lignes  [][]string
    NbLignes int
    NbColonnes int
}

func lireCSV(nomFichier string) (*DonneesCSV, error) {
    // Vérifier l'existence du fichier
    if _, err := os.Stat(nomFichier); os.IsNotExist(err) {
        return nil, fmt.Errorf("%w : %s", ErrFichierInexistant, nomFichier)
    }

    // Ouvrir le fichier
    fichier, err := os.Open(nomFichier)
    if err != nil {
        return nil, fmt.Errorf("impossible d'ouvrir %s : %w", nomFichier, err)
    }
    defer fichier.Close()

    // Créer le lecteur CSV
    lecteur := csv.NewReader(fichier)
    lecteur.Comma = ','
    lecteur.TrimLeadingSpace = true

    // Lire toutes les lignes
    lignes, err := lecteur.ReadAll()
    if err != nil {
        return nil, fmt.Errorf("%w dans %s : %v", ErrFormatCSV, nomFichier, err)
    }

    // Vérifier que le fichier n'est pas vide
    if len(lignes) == 0 {
        return nil, fmt.Errorf("%w : %s", ErrFichierVide, nomFichier)
    }

    // Extraire les en-têtes
    entetes := lignes[0]
    donneesLignes := lignes[1:]

    // Vérifier la cohérence des colonnes
    nbColonnesAttendu := len(entetes)
    for i, ligne := range donneesLignes {
        if len(ligne) != nbColonnesAttendu {
            return nil, fmt.Errorf("%w : ligne %d a %d colonnes, %d attendues",
                ErrColonnesManquantes, i+2, len(ligne), nbColonnesAttendu)
        }
    }

    return &DonneesCSV{
        Entetes: entetes,
        Lignes: donneesLignes,
        NbLignes: len(donneesLignes),
        NbColonnes: nbColonnesAttendu,
    }, nil
}

// Fonction pour convertir une colonne en nombres
func extraireColonneNumerique(donnees *DonneesCSV, nomColonne string) ([]float64, error) {
    // Trouver l'index de la colonne
    indexColonne := -1
    for i, entete := range donnees.Entetes {
        if entete == nomColonne {
            indexColonne = i
            break
        }
    }

    if indexColonne == -1 {
        return nil, fmt.Errorf("colonne '%s' introuvable", nomColonne)
    }

    // Extraire et convertir les valeurs
    valeurs := make([]float64, 0, donnees.NbLignes)

    for i, ligne := range donnees.Lignes {
        valeurTexte := strings.TrimSpace(ligne[indexColonne])

        // Ignorer les cellules vides
        if valeurTexte == "" {
            continue
        }

        valeur, err := strconv.ParseFloat(valeurTexte, 64)
        if err != nil {
            return nil, fmt.Errorf("%w : ligne %d, colonne '%s' : '%s' n'est pas un nombre",
                ErrDonneesInvalides, i+2, nomColonne, valeurTexte)
        }

        valeurs = append(valeurs, valeur)
    }

    return valeurs, nil
}

// Créer un fichier CSV de test
func creerFichierTest() error {
    contenu := `nom,age,salaire
Alice,25,50000
Bob,30,60000
Charlie,35,abc
Diana,28,55000
Eve,,45000`

    return os.WriteFile("test.csv", []byte(contenu), 0644)
}

func main() {
    fmt.Println("=== Lecteur de fichier CSV ===")

    // Créer un fichier de test
    fmt.Println("Création du fichier de test...")
    if err := creerFichierTest(); err != nil {
        fmt.Printf("Erreur lors de la création du fichier test : %v\n", err)
        return
    }

    // Test avec différents fichiers
    fichiers := []string{
        "test.csv",
        "inexistant.csv",
        "",
    }

    for _, fichier := range fichiers {
        fmt.Printf("\n--- Lecture de '%s' ---\n", fichier)

        donnees, err := lireCSV(fichier)
        if err != nil {
            fmt.Printf("❌ Erreur : %v\n", err)

            // Gestion spécifique selon le type d'erreur
            if errors.Is(err, ErrFichierInexistant) {
                fmt.Println("   Conseil : Vérifiez le chemin du fichier")
            } else if errors.Is(err, ErrFormatCSV) {
                fmt.Println("   Conseil : Vérifiez le format CSV")
            }
            continue
        }

        // Afficher les informations du fichier
        fmt.Printf("✅ Fichier lu avec succès\n")
        fmt.Printf("   En-têtes : %v\n", donnees.Entetes)
        fmt.Printf("   Lignes de données : %d\n", donnees.NbLignes)
        fmt.Printf("   Colonnes : %d\n", donnees.NbColonnes)

        // Essayer d'extraire une colonne numérique
        if len(donnees.Entetes) > 0 {
            fmt.Printf("\n   Tentative d'extraction de la colonne 'age' :\n")
            ages, err := extraireColonneNumerique(donnees, "age")
            if err != nil {
                fmt.Printf("   ❌ %v\n", err)
            } else {
                fmt.Printf("   ✅ Ages extraits : %v\n", ages)
            }

            fmt.Printf("\n   Tentative d'extraction de la colonne 'salaire' :\n")
            salaires, err := extraireColonneNumerique(donnees, "salaire")
            if err != nil {
                fmt.Printf("   ❌ %v\n", err)
            } else {
                fmt.Printf("   ✅ Salaires extraits : %v\n", salaires)
            }
        }
    }
}
```

### Exercice 4 : Convertisseur d'unités

#### Solution complète pour les températures

```go
package main

import (
    "errors"
    "fmt"
    "math"
)

// Erreurs spécifiques aux conversions de température
var (
    ErrTempImpossible    = errors.New("température physiquement impossible")
    ErrUniteInconnue     = errors.New("unité de température inconnue")
    ErrConversionImpossible = errors.New("conversion impossible")
)

// Constantes physiques
const (
    ZeroAbsoluCelsius = -273.15
    ZeroAbsoluFahrenheit = -459.67
    ZeroAbsoluKelvin = 0
    TempMaxPratique = 10000 // Température maximale pratique
)

type Temperature struct {
    Valeur float64
    Unite  string
}

type ResultatConversion struct {
    Origine     Temperature
    Destination Temperature
    Valide      bool
    Erreur      error
}

// Validation des températures selon l'unité
func validerTemperature(temp float64, unite string) error {
    switch unite {
    case "C", "celsius":
        if temp < ZeroAbsoluCelsius {
            return fmt.Errorf("%w : %.2f°C (minimum: %.2f°C)",
                ErrTempImpossible, temp, ZeroAbsoluCelsius)
        }
    case "F", "fahrenheit":
        if temp < ZeroAbsoluFahrenheit {
            return fmt.Errorf("%w : %.2f°F (minimum: %.2f°F)",
                ErrTempImpossible, temp, ZeroAbsoluFahrenheit)
        }
    case "K", "kelvin":
        if temp < ZeroAbsoluKelvin {
            return fmt.Errorf("%w : %.2f K (minimum: %.2f K)",
                ErrTempImpossible, temp, ZeroAbsoluKelvin)
        }
    default:
        return fmt.Errorf("%w : %s (utilisez C, F, ou K)", ErrUniteInconnue, unite)
    }

    // Vérifier que la température n'est pas excessivement élevée
    if math.Abs(temp) > TempMaxPratique {
        return fmt.Errorf("température irréaliste : %.2f", temp)
    }

    return nil
}

// Normalisation des unités
func normaliserUnite(unite string) string {
    switch unite {
    case "C", "c", "celsius", "Celsius":
        return "C"
    case "F", "f", "fahrenheit", "Fahrenheit":
        return "F"
    case "K", "k", "kelvin", "Kelvin":
        return "K"
    default:
        return unite // Retourner tel quel pour générer une erreur
    }
}

// Conversion de température
func convertirTemperature(valeur float64, uniteOrigine, uniteDestination string) ResultatConversion {
    // Normaliser les unités
    uniteOrigine = normaliserUnite(uniteOrigine)
    uniteDestination = normaliserUnite(uniteDestination)

    origine := Temperature{Valeur: valeur, Unite: uniteOrigine}
    resultat := ResultatConversion{Origine: origine}

    // Valider la température d'origine
    if err := validerTemperature(valeur, uniteOrigine); err != nil {
        resultat.Erreur = err
        return resultat
    }

    // Valider l'unité de destination
    if err := validerTemperature(0, uniteDestination); err != nil {
        resultat.Erreur = fmt.Errorf("unité de destination invalide : %w", err)
        return resultat
    }

    // Si même unité, pas de conversion nécessaire
    if uniteOrigine == uniteDestination {
        resultat.Destination = origine
        resultat.Valide = true
        return resultat
    }

    // Convertir vers Celsius comme base
    var celsius float64
    switch uniteOrigine {
    case "C":
        celsius = valeur
    case "F":
        celsius = (valeur - 32) * 5 / 9
    case "K":
        celsius = valeur - 273.15
    default:
        resultat.Erreur = fmt.Errorf("%w : %s", ErrUniteInconnue, uniteOrigine)
        return resultat
    }

    // Convertir de Celsius vers l'unité de destination
    var valeurDestination float64
    switch uniteDestination {
    case "C":
        valeurDestination = celsius
    case "F":
        valeurDestination = celsius*9/5 + 32
    case "K":
        valeurDestination = celsius + 273.15
    default:
        resultat.Erreur = fmt.Errorf("%w : %s", ErrUniteInconnue, uniteDestination)
        return resultat
    }

    // Valider le résultat
    if err := validerTemperature(valeurDestination, uniteDestination); err != nil {
        resultat.Erreur = fmt.Errorf("résultat de conversion invalide : %w", err)
        return resultat
    }

    resultat.Destination = Temperature{Valeur: valeurDestination, Unite: uniteDestination}
    resultat.Valide = true
    return resultat
}

func main() {
    fmt.Println("=== Convertisseur de température ===")

    // Cas de test
    tests := []struct {
        nom                string
        valeur             float64
        uniteOrigine       string
        uniteDestination   string
    }{
        {"Eau bouillante", 100, "C", "F"},
        {"Eau gelée", 0, "C", "K"},
        {"Température corporelle", 98.6, "F", "C"},
        {"Zéro absolu (impossible)", -300, "C", "F"},
        {"Température négative valide", -50, "C", "F"},
        {"Unité inconnue", 25, "X", "C"},
        {"Même unité", 25, "C", "C"},
        {"Température très élevée", 5000, "C", "K"},
        {"Zéro Kelvin", 0, "K", "C"},
    }

    for _, test := range tests {
        fmt.Printf("\n--- %s ---\n", test.nom)
        fmt.Printf("Conversion : %.2f°%s → °%s\n",
            test.valeur, test.uniteOrigine, test.uniteDestination)

        resultat := convertirTemperature(test.valeur, test.uniteOrigine, test.uniteDestination)

        if resultat.Valide {
            fmt.Printf("✅ Résultat : %.2f°%s\n",
                resultat.Destination.Valeur, resultat.Destination.Unite)
        } else {
            fmt.Printf("❌ Erreur : %v\n", resultat.Erreur)

            // Suggestions selon le type d'erreur
            if errors.Is(resultat.Erreur, ErrTempImpossible) {
                fmt.Println("   💡 Conseil : Vérifiez que la température est au-dessus du zéro absolu")
            } else if errors.Is(resultat.Erreur, ErrUniteInconnue) {
                fmt.Println("   💡 Conseil : Utilisez C (Celsius), F (Fahrenheit) ou K (Kelvin)")
            }
        }
    }

    // Mode interactif
    fmt.Println("\n=== Mode interactif ===")
    for {
        var valeur float64
        var origine, destination string

        fmt.Print("\nEntrez la température (ou 0 pour quitter) : ")
        fmt.Scanln(&valeur)

        if valeur == 0 {
            break
        }

        fmt.Print("Unité d'origine (C/F/K) : ")
        fmt.Scanln(&origine)

        fmt.Print("Unité de destination (C/F/K) : ")
        fmt.Scanln(&destination)

        resultat := convertirTemperature(valeur, origine, destination)

        if resultat.Valide {
            fmt.Printf("✅ %.2f°%s = %.2f°%s\n",
                resultat.Origine.Valeur, resultat.Origine.Unite,
                resultat.Destination.Valeur, resultat.Destination.Unite)
        } else {
            fmt.Printf("❌ %v\n", resultat.Erreur)
        }
    }
}
```

### Exercice 5 : Gestionnaire de mots de passe

#### Solution complète avec différents niveaux de sécurité

```go
package main

import (
    "errors"
    "fmt"
    "regexp"
    "strings"
    "unicode"
)

// Erreurs spécifiques aux mots de passe
var (
    ErrMotDePasseVide        = errors.New("le mot de passe ne peut pas être vide")
    ErrTropCourt             = errors.New("mot de passe trop court")
    ErrTropLong              = errors.New("mot de passe trop long")
    ErrPasDeMinuscule        = errors.New("doit contenir au moins une minuscule")
    ErrPasDeMajuscule        = errors.New("doit contenir au moins une majuscule")
    ErrPasDeChiffre          = errors.New("doit contenir au moins un chiffre")
    ErrPasDeCaractereSpecial = errors.New("doit contenir au moins un caractère spécial")
    ErrCaractereInterdit     = errors.New("contient des caractères interdits")
    ErrMotDePasse Commun     = errors.New("mot de passe trop commun")
    ErrRepetitionExcessive   = errors.New("trop de caractères identiques consécutifs")
)

type NiveauSecurite int

const (
    Basique NiveauSecurite = iota
    Moyen
    Fort
    TresForte
)

type CriteresValidation struct {
    LongueurMin           int
    LongueurMax           int
    RequiertMinuscule     bool
    RequiertMajuscule     bool
    RequiertChiffre       bool
    RequiertCaractereSpecial bool
    MaxRepetitionConsecutive int
    VerifierMotsCourants  bool
}

type ResultatValidation struct {
    Valide            bool
    Erreurs           []error
    Score             int
    Niveau            NiveauSecurite
    Suggestions       []string
}

// Critères selon le niveau de sécurité
func obtenirCriteres(niveau NiveauSecurite) CriteresValidation {
    switch niveau {
    case Basique:
        return CriteresValidation{
            LongueurMin:              6,
            LongueurMax:              50,
            RequiertMinuscule:        false,
            RequiertMajuscule:        false,
            RequiertChiffre:          false,
            RequiertCaractereSpecial: false,
            MaxRepetitionConsecutive: 5,
            VerifierMotsCourants:     false,
        }
    case Moyen:
        return CriteresValidation{
            LongueurMin:              8,
            LongueurMax:              50,
            RequiertMinuscule:        true,
            RequiertMajuscule:        true,
            RequiertChiffre:          true,
            RequiertCaractereSpecial: false,
            MaxRepetitionConsecutive: 3,
            VerifierMotsCourants:     true,
        }
    case Fort:
        return CriteresValidation{
            LongueurMin:              10,
            LongueurMax:              100,
            RequiertMinuscule:        true,
            RequiertMajuscule:        true,
            RequiertChiffre:          true,
            RequiertCaractereSpecial: true,
            MaxRepetitionConsecutive: 2,
            VerifierMotsCourants:     true,
        }
    case TresForte:
        return CriteresValidation{
            LongueurMin:              12,
            LongueurMax:              100,
            RequiertMinuscule:        true,
            RequiertMajuscule:        true,
            RequiertChiffre:          true,
            RequiertCaractereSpecial: true,
            MaxRepetitionConsecutive: 2,
            VerifierMotsCourants:     true,
        }
    default:
        return obtenirCriteres(Moyen)
    }
}

// Liste de mots de passe courants (version simplifiée)
var motsDePasse Courants = []string{
    "password", "123456", "password123", "admin", "qwerty",
    "letmein", "welcome", "monkey", "dragon", "master",
    "azerty", "motdepasse", "bonjour", "salut", "secret",
}

// Vérifier si le mot de passe est dans la liste des mots courants
func estMotDePasseCourant(motDePasse string) bool {
    motDePasseMinuscule := strings.ToLower(motDePasse)
    for _, motCourant := range motsDePasse Courants {
        if motDePasseMinuscule == motCourant {
            return true
        }
        // Vérifier aussi avec des variantes simples
        if motDePasseMinuscule == motCourant+"123" ||
           motDePasseMinuscule == motCourant+"1" ||
           motDePasseMinuscule == "123"+motCourant {
            return true
        }
    }
    return false
}

// Analyser la composition du mot de passe
func analyserComposition(motDePasse string) (bool, bool, bool, bool) {
    hasMinuscule := false
    hasMajuscule := false
    hasChiffre := false
    hasSpecial := false

    for _, char := range motDePasse {
        switch {
        case unicode.IsLower(char):
            hasMinuscule = true
        case unicode.IsUpper(char):
            hasMajuscule = true
        case unicode.IsDigit(char):
            hasChiffre = true
        case unicode.IsPunct(char) || unicode.IsSymbol(char):
            hasSpecial = true
        }
    }

    return hasMinuscule, hasMajuscule, hasChiffre, hasSpecial
}

// Vérifier les répétitions excessives
func verifierRepetitions(motDePasse string, maxRepetition int) bool {
    if len(motDePasse) < 2 {
        return true
    }

    compteur := 1
    for i := 1; i < len(motDePasse); i++ {
        if motDePasse[i] == motDePasse[i-1] {
            compteur++
            if compteur > maxRepetition {
                return false
            }
        } else {
            compteur = 1
        }
    }
    return true
}

// Calculer le score de force du mot de passe
func calculerScore(motDePasse string) int {
    score := 0

    // Points pour la longueur
    longueur := len(motDePasse)
    switch {
    case longueur >= 12:
        score += 25
    case longueur >= 8:
        score += 15
    case longueur >= 6:
        score += 5
    }

    // Points pour la diversité des caractères
    hasMin, hasMaj, hasNum, hasSpec := analyserComposition(motDePasse)

    if hasMin {
        score += 10
    }
    if hasMaj {
        score += 10
    }
    if hasNum {
        score += 10
    }
    if hasSpec {
        score += 15
    }

    // Points bonus pour la complexité
    if hasMin && hasMaj && hasNum && hasSpec {
        score += 20
    }

    // Malus pour les répétitions
    if !verifierRepetitions(motDePasse, 2) {
        score -= 10
    }

    // Malus pour les mots courants
    if estMotDePasseCourant(motDePasse) {
        score -= 30
    }

    // Assurer que le score reste dans les limites
    if score < 0 {
        score = 0
    }
    if score > 100 {
        score = 100
    }

    return score
}

// Déterminer le niveau selon le score
func determinerNiveau(score int) NiveauSecurite {
    switch {
    case score >= 80:
        return TresForte
    case score >= 60:
        return Fort
    case score >= 40:
        return Moyen
    default:
        return Basique
    }
}

// Générer des suggestions d'amélioration
func genererSuggestions(motDePasse string, criteres CriteresValidation) []string {
    var suggestions []string

    hasMin, hasMaj, hasNum, hasSpec := analyserComposition(motDePasse)

    if len(motDePasse) < criteres.LongueurMin {
        suggestions = append(suggestions,
            fmt.Sprintf("Augmentez la longueur à au moins %d caractères", criteres.LongueurMin))
    }

    if criteres.RequiertMinuscule && !hasMin {
        suggestions = append(suggestions, "Ajoutez des lettres minuscules (a-z)")
    }

    if criteres.RequiertMajuscule && !hasMaj {
        suggestions = append(suggestions, "Ajoutez des lettres majuscules (A-Z)")
    }

    if criteres.RequiertChiffre && !hasNum {
        suggestions = append(suggestions, "Ajoutez des chiffres (0-9)")
    }

    if criteres.RequiertCaractereSpecial && !hasSpec {
        suggestions = append(suggestions, "Ajoutez des caractères spéciaux (!@#$%^&*)")
    }

    if !verifierRepetitions(motDePasse, criteres.MaxRepetitionConsecutive) {
        suggestions = append(suggestions,
            fmt.Sprintf("Évitez plus de %d caractères identiques consécutifs",
                criteres.MaxRepetitionConsecutive))
    }

    if criteres.VerifierMotsCourants && estMotDePasseCourant(motDePasse) {
        suggestions = append(suggestions, "Évitez les mots de passe courants")
    }

    if len(suggestions) == 0 {
        suggestions = append(suggestions, "Mot de passe conforme aux critères")
    }

    return suggestions
}

// Fonction principale de validation
func validerMotDePasse(motDePasse string, niveau NiveauSecurite) ResultatValidation {
    resultat := ResultatValidation{
        Valide: true,
        Erreurs: []error{},
    }

    criteres := obtenirCriteres(niveau)

    // Vérification de base
    if motDePasse == "" {
        resultat.Erreurs = append(resultat.Erreurs, ErrMotDePasseVide)
        resultat.Valide = false
    }

    // Vérifier la longueur
    longueur := len(motDePasse)
    if longueur < criteres.LongueurMin {
        resultat.Erreurs = append(resultat.Erreurs,
            fmt.Errorf("%w : %d caractères (minimum: %d)",
                ErrTropCourt, longueur, criteres.LongueurMin))
        resultat.Valide = false
    }

    if longueur > criteres.LongueurMax {
        resultat.Erreurs = append(resultat.Erreurs,
            fmt.Errorf("%w : %d caractères (maximum: %d)",
                ErrTropLong, longueur, criteres.LongueurMax))
        resultat.Valide = false
    }

    // Analyser la composition
    hasMin, hasMaj, hasNum, hasSpec := analyserComposition(motDePasse)

    if criteres.RequiertMinuscule && !hasMin {
        resultat.Erreurs = append(resultat.Erreurs, ErrPasDeMinuscule)
        resultat.Valide = false
    }

    if criteres.RequiertMajuscule && !hasMaj {
        resultat.Erreurs = append(resultat.Erreurs, ErrPasDeMajuscule)
        resultat.Valide = false
    }

    if criteres.RequiertChiffre && !hasNum {
        resultat.Erreurs = append(resultat.Erreurs, ErrPasDeChiffre)
        resultat.Valide = false
    }

    if criteres.RequiertCaractereSpecial && !hasSpec {
        resultat.Erreurs = append(resultat.Erreurs, ErrPasDeCaractereSpecial)
        resultat.Valide = false
    }

    // Vérifier les répétitions
    if !verifierRepetitions(motDePasse, criteres.MaxRepetitionConsecutive) {
        resultat.Erreurs = append(resultat.Erreurs, ErrRepetitionExcessive)
        resultat.Valide = false
    }

    // Vérifier les mots courants
    if criteres.VerifierMotsCourants && estMotDePasseCourant(motDePasse) {
        resultat.Erreurs = append(resultat.Erreurs, ErrMotDePasseCourant)
        resultat.Valide = false
    }

    // Calculer le score et le niveau
    resultat.Score = calculerScore(motDePasse)
    resultat.Niveau = determinerNiveau(resultat.Score)

    // Générer les suggestions
    resultat.Suggestions = genererSuggestions(motDePasse, criteres)

    return resultat
}

// Fonction d'affichage du résultat
func afficherResultat(motDePasse string, niveau NiveauSecurite, resultat ResultatValidation) {
    niveaux := []string{"Basique", "Moyen", "Fort", "Très Fort"}

    fmt.Printf("\n=== Validation du mot de passe ===\n")
    fmt.Printf("Mot de passe : %s\n", strings.Repeat("*", len(motDePasse)))
    fmt.Printf("Niveau requis : %s\n", niveaux[niveau])
    fmt.Printf("Score : %d/100\n", resultat.Score)
    fmt.Printf("Niveau atteint : %s\n", niveaux[resultat.Niveau])

    if resultat.Valide {
        fmt.Println("✅ Mot de passe VALIDE")
    } else {
        fmt.Println("❌ Mot de passe INVALIDE")
        fmt.Println("\nErreurs :")
        for _, err := range resultat.Erreurs {
            fmt.Printf("  - %v\n", err)
        }
    }

    fmt.Println("\nSuggestions :")
    for _, suggestion := range resultat.Suggestions {
        fmt.Printf("  💡 %s\n", suggestion)
    }
}

func main() {
    fmt.Println("=== Gestionnaire de mots de passe ===")

    // Tests avec différents mots de passe et niveaux
    tests := []struct {
        nom        string
        motDePasse string
        niveau     NiveauSecurite
    }{
        {"Mot de passe faible", "123456", Moyen},
        {"Mot de passe moyen", "Password123", Moyen},
        {"Mot de passe fort", "MyS3cur3P@ssw0rd!", Fort},
        {"Mot de passe très fort", "Tr3s$ecur3Motd3Pass3#2024!", TresForte},
        {"Mot de passe avec répétitions", "aaabbbccc", Basique},
        {"Mot de passe courant", "password123", Moyen},
        {"Mot de passe court", "Ab1!", Fort},
        {"Mot de passe long et complexe", "Ceci-Est-Un-Mot-De-Passe-Tres-Long-Et-Complexe123!", TresForte},
    }

    for _, test := range tests {
        fmt.Printf("\n" + strings.Repeat("=", 50))
        fmt.Printf("\nTest : %s", test.nom)

        resultat := validerMotDePasse(test.motDePasse, test.niveau)
        afficherResultat(test.motDePasse, test.niveau, resultat)
    }

    // Mode interactif
    fmt.Printf("\n" + strings.Repeat("=", 50))
    fmt.Println("\n=== Mode interactif ===")

    for {
        var motDePasse string
        var choixNiveau int

        fmt.Print("\nEntrez un mot de passe (ou 'quit' pour quitter) : ")
        fmt.Scanln(&motDePasse)

        if motDePasse == "quit" {
            break
        }

        fmt.Println("Choisissez le niveau de sécurité :")
        fmt.Println("0 - Basique")
        fmt.Println("1 - Moyen")
        fmt.Println("2 - Fort")
        fmt.Println("3 - Très Fort")
        fmt.Print("Niveau (0-3) : ")
        fmt.Scanln(&choixNiveau)

        if choixNiveau < 0 || choixNiveau > 3 {
            fmt.Println("❌ Niveau invalide, utilisation du niveau Moyen")
            choixNiveau = 1
        }

        niveau := NiveauSecurite(choixNiveau)
        resultat := validerMotDePasse(motDePasse, niveau)
        afficherResultat(motDePasse, niveau, resultat)
    }

    fmt.Println("\nMerci d'avoir utilisé le gestionnaire de mots de passe !")
}
```

### Fonctionnalités clés du gestionnaire de mots de passe

#### 🔒 **Niveaux de sécurité**

- **Basique** : Longueur minimale seulement
- **Moyen** : Minuscules, majuscules, chiffres
- **Fort** : + caractères spéciaux, longueur 10+
- **Très Fort** : Critères renforcés, longueur 12+

#### 🎯 **Validations implémentées**

- Longueur (min/max selon niveau)
- Composition (maj/min/chiffres/spéciaux)
- Répétitions excessives
- Mots de passe courants
- Score de force (0-100)

#### 💡 **Fonctionnalités avancées**

- Système de scoring intelligent
- Suggestions personnalisées
- Mode interactif
- Gestion d'erreurs spécifiques
- Détection de patterns faibles

#### 🛡️ **Sécurité**

- Masquage du mot de passe à l'affichage
- Vérification contre liste de mots courants
- Détection de variantes simples (password123, etc.)
- Validation de caractères interdits

Ce gestionnaire de mots de passe démontre une gestion d'erreurs complète avec différents types d'erreurs, validation par niveaux, et retour utilisateur constructif !

## 12. Comparaison avec try/catch

Pour ceux qui viennent d'autres langages :

| Try/Catch | Go |
|-----------|-------|
| Exceptions cachées | Erreurs explicites |
| Flux complexe | Flux linéaire |
| Performance imprévisible | Performance prévisible |
| Peut être ignoré | Doit être géré |

```go
// Autres langages
// try {
//     result = operation()
//     use(result)
// } catch (Exception e) {
//     handle(e)
// }

// Go
result, err := operation()
if err != nil {
    handle(err)
    return
}
use(result)
```

## Résumé

La gestion d'erreurs en Go suit ces principes :

- **Explicit is better than implicit** : Les erreurs sont visibles
- **Errors are values** : Les erreurs se manipulent comme des valeurs
- **Handle errors at the right level** : Gérez les erreurs au bon endroit
- **Don't ignore errors** : Ne jamais ignorer une erreur
- **Fail fast** : Échouez rapidement et proprement

**Points clés à retenir :**

- Toujours vérifier `if err != nil`
- Utiliser `fmt.Errorf` pour ajouter du contexte
- Propager, logger ou gérer selon le contexte
- Messages d'erreur clairs et utiles
- Early return pattern pour éviter l'imbrication

La gestion d'erreurs en Go peut sembler verbeuse au début, mais elle rend le code plus robuste et prévisible !

Dans la prochaine section, nous verrons les fonctions en détail !

⏭️
