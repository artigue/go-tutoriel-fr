🔝 Retour au [Sommaire](/SOMMAIRE.md)

# 5-3 : Structs

## Introduction

Une **struct** (structure) est un type de données personnalisé qui permet de regrouper plusieurs valeurs de types différents sous une même entité. C'est l'équivalent des classes dans d'autres langages, mais en plus simple. Les structs sont la base de la programmation orientée données en Go.

## Concept de base

Imaginez que vous voulez représenter une personne dans votre programme. Une personne a plusieurs attributs :

- Un nom (string)
- Un âge (int)
- Une taille (float64)
- Une adresse email (string)

Au lieu de gérer ces informations séparément, une struct permet de les regrouper :

```go
type Personne struct {
    Nom    string
    Age    int
    Taille float64
    Email  string
}
```

## Déclaration de Structs

### Syntaxe de base

```go
// Déclaration d'une struct
type NomDeLaStruct struct {
    champ1 type1
    champ2 type2
    champ3 type3
}
```

### Exemples concrets

```go
// Struct simple pour un produit
type Produit struct {
    Nom         string
    Prix        float64
    EnStock     bool
    Quantite    int
}

// Struct pour représenter un point en 2D
type Point struct {
    X float64
    Y float64
}

// Struct pour un compte bancaire
type CompteBancaire struct {
    Numero    string
    Titulaire string
    Solde     float64
    Actif     bool
}
```

## Création et initialisation

### Méthode 1 : Initialisation avec valeurs par défaut

```go
type Personne struct {
    Nom string
    Age int
}

func main() {
    // Création avec zero values
    var p Personne
    fmt.Printf("%+v\n", p) // {Nom: Age:0}

    // Les zero values sont :
    // string: ""
    // int: 0
    // float64: 0.0
    // bool: false
}
```

### Méthode 2 : Initialisation avec valeurs spécifiques

```go
// Syntaxe avec noms de champs (recommandée)
personne1 := Personne{
    Nom: "Alice",
    Age: 25,
}

// Syntaxe positionnelle (à éviter généralement)
personne2 := Personne{"Bob", 30}

// Initialisation partielle (autres champs = zero value)
personne3 := Personne{
    Nom: "Charlie", // Age sera 0
}
```

### Méthode 3 : Avec new()

```go
// new() retourne un pointeur vers la struct
p := new(Personne)
p.Nom = "Diana"
p.Age = 28

fmt.Printf("%+v\n", *p) // {Nom:Diana Age:28}
```

### Méthode 4 : Avec l'opérateur &

```go
// Création et obtention d'un pointeur
p := &Personne{
    Nom: "Eve",
    Age: 22,
}

fmt.Printf("%+v\n", *p) // {Nom:Eve Age:22}
```

## Accès aux champs

```go
type Voiture struct {
    Marque string
    Modele string
    Annee  int
}

func main() {
    voiture := Voiture{
        Marque: "Toyota",
        Modele: "Corolla",
        Annee:  2023,
    }

    // Lecture des champs
    fmt.Println("Marque:", voiture.Marque)
    fmt.Println("Modèle:", voiture.Modele)
    fmt.Println("Année:", voiture.Annee)

    // Modification des champs
    voiture.Annee = 2024
    voiture.Marque = "Honda"

    fmt.Printf("Voiture mise à jour: %+v\n", voiture)
}
```

## Structs avec différents types de champs

```go
type Etudiant struct {
    ID          int
    Nom         string
    Age         int
    Notes       []float64    // Slice de notes
    Actif       bool
    Informations map[string]string // Map d'informations
}

func main() {
    etudiant := Etudiant{
        ID:    12345,
        Nom:   "Alice Dupont",
        Age:   20,
        Notes: []float64{15.5, 18.0, 16.5},
        Actif: true,
        Informations: map[string]string{
            "email":     "alice@email.com",
            "telephone": "01-23-45-67-89",
        },
    }

    // Accès aux champs complexes
    fmt.Printf("Nom: %s\n", etudiant.Nom)
    fmt.Printf("Première note: %.1f\n", etudiant.Notes[0])
    fmt.Printf("Email: %s\n", etudiant.Informations["email"])

    // Ajouter une note
    etudiant.Notes = append(etudiant.Notes, 17.0)

    // Ajouter une information
    etudiant.Informations["adresse"] = "123 Rue de la Paix"
}
```

## Structs imbriquées (Composition)

```go
type Adresse struct {
    Rue        string
    Ville      string
    CodePostal string
    Pays       string
}

type Personne struct {
    Nom        string
    Age        int
    Adresse    Adresse // Struct imbriquée
}

func main() {
    personne := Personne{
        Nom: "Bob Martin",
        Age: 35,
        Adresse: Adresse{
            Rue:        "456 Avenue des Champs",
            Ville:      "Paris",
            CodePostal: "75001",
            Pays:       "France",
        },
    }

    // Accès aux champs imbriqués
    fmt.Printf("Nom: %s\n", personne.Nom)
    fmt.Printf("Ville: %s\n", personne.Adresse.Ville)
    fmt.Printf("Adresse complète: %s, %s %s\n",
        personne.Adresse.Rue,
        personne.Adresse.CodePostal,
        personne.Adresse.Ville)
}
```

## Structs anonymes et embedding

### Structs anonymes

```go
// Struct anonyme pour usage ponctuel
point := struct {
    X, Y int
}{10, 20}

fmt.Printf("Point: (%d, %d)\n", point.X, point.Y)

// Utilisation dans une fonction
func traiterDonnees() {
    config := struct {
        Host string
        Port int
        SSL  bool
    }{
        Host: "localhost",
        Port: 8080,
        SSL:  true,
    }

    fmt.Printf("Configuration: %+v\n", config)
}
```

### Embedding (Composition anonyme)

```go
type Adresse struct {
    Rue   string
    Ville string
}

type Personne struct {
    Nom string
    Age int
    Adresse // Embedding anonyme
}

func main() {
    p := Personne{
        Nom: "Alice",
        Age: 25,
        Adresse: Adresse{
            Rue:   "123 Rue de la Liberté",
            Ville: "Lyon",
        },
    }

    // Accès direct aux champs de la struct embedée
    fmt.Printf("Nom: %s\n", p.Nom)
    fmt.Printf("Rue: %s\n", p.Rue)   // Accès direct !
    fmt.Printf("Ville: %s\n", p.Ville) // Accès direct !

    // Ou accès explicite
    fmt.Printf("Rue (explicite): %s\n", p.Adresse.Rue)
}
```

## Méthodes sur les Structs

Les méthodes permettent d'associer des fonctions à vos structs :

```go
type Rectangle struct {
    Largeur float64
    Hauteur float64
}

// Méthode avec receiver par valeur
func (r Rectangle) Aire() float64 {
    return r.Largeur * r.Hauteur
}

// Méthode avec receiver par pointeur
func (r *Rectangle) Redimensionner(facteur float64) {
    r.Largeur *= facteur
    r.Hauteur *= facteur
}

// Méthode pour affichage
func (r Rectangle) String() string {
    return fmt.Sprintf("Rectangle(%.2f x %.2f)", r.Largeur, r.Hauteur)
}

func main() {
    rect := Rectangle{Largeur: 10, Hauteur: 5}

    fmt.Printf("Rectangle: %s\n", rect)
    fmt.Printf("Aire: %.2f\n", rect.Aire())

    // Modification via pointeur
    rect.Redimensionner(2.0)
    fmt.Printf("Après redimensionnement: %s\n", rect)
    fmt.Printf("Nouvelle aire: %.2f\n", rect.Aire())
}
```

## Constructeurs (Fonctions de création)

Go n'a pas de constructeurs intégrés, mais on utilise des fonctions par convention :

```go
type CompteBancaire struct {
    numero    string  // privé (minuscule)
    titulaire string  // privé
    solde     float64 // privé
}

// Constructeur
func NouveauCompte(numero, titulaire string, soldeInitial float64) *CompteBancaire {
    return &CompteBancaire{
        numero:    numero,
        titulaire: titulaire,
        solde:     soldeInitial,
    }
}

// Méthodes publiques pour accéder aux champs privés
func (c *CompteBancaire) Numero() string {
    return c.numero
}

func (c *CompteBancaire) Titulaire() string {
    return c.titulaire
}

func (c *CompteBancaire) Solde() float64 {
    return c.solde
}

func (c *CompteBancaire) Deposer(montant float64) {
    if montant > 0 {
        c.solde += montant
    }
}

func (c *CompteBancaire) Retirer(montant float64) bool {
    if montant > 0 && c.solde >= montant {
        c.solde -= montant
        return true
    }
    return false
}

func main() {
    compte := NouveauCompte("123456789", "Alice Dupont", 1000.0)

    fmt.Printf("Compte: %s - %s\n", compte.Numero(), compte.Titulaire())
    fmt.Printf("Solde initial: %.2f€\n", compte.Solde())

    compte.Deposer(500.0)
    fmt.Printf("Après dépôt: %.2f€\n", compte.Solde())

    if compte.Retirer(200.0) {
        fmt.Printf("Après retrait: %.2f€\n", compte.Solde())
    }
}
```

## Comparaison de Structs

```go
type Point struct {
    X, Y int
}

func main() {
    p1 := Point{X: 10, Y: 20}
    p2 := Point{X: 10, Y: 20}
    p3 := Point{X: 15, Y: 25}

    // Comparaison directe (tous les champs doivent être comparables)
    fmt.Printf("p1 == p2: %t\n", p1 == p2) // true
    fmt.Printf("p1 == p3: %t\n", p1 == p3) // false
}
```

**Attention :** Les structs contenant des slices, maps ou fonctions ne peuvent pas être comparées avec `==`.

```go
type Personne struct {
    Nom   string
    Notes []int // Rend la struct non-comparable
}

func main() {
    p1 := Personne{Nom: "Alice", Notes: []int{15, 18}}
    p2 := Personne{Nom: "Alice", Notes: []int{15, 18}}

    // Ceci ne compile pas !
    // fmt.Println(p1 == p2) // Erreur de compilation

    // Solution : comparer champ par champ
    fmt.Printf("Même nom: %t\n", p1.Nom == p2.Nom)
}
```

## Tags de Structs

Les tags permettent d'ajouter des métadonnées aux champs :

```go
import (
    "encoding/json"
    "fmt"
)

type Utilisateur struct {
    ID       int    `json:"id"`
    Nom      string `json:"name"`
    Email    string `json:"email"`
    MotDePasse string `json:"-"` // Exclure du JSON
}

func main() {
    user := Utilisateur{
        ID:       1,
        Nom:      "Alice",
        Email:    "alice@example.com",
        MotDePasse: "secret123",
    }

    // Conversion en JSON
    jsonData, err := json.Marshal(user)
    if err != nil {
        fmt.Println("Erreur:", err)
        return
    }

    fmt.Printf("JSON: %s\n", jsonData)
    // Résultat: {"id":1,"name":"Alice","email":"alice@example.com"}
    // Le mot de passe n'est pas inclus !
}
```

## Exemples pratiques

### Exemple 1 : Système de gestion d'étudiants

```go
type Etudiant struct {
    ID        int
    Nom       string
    Prenom    string
    Email     string
    Notes     []float64
    Moyenne   float64
}

func (e *Etudiant) AjouterNote(note float64) {
    e.Notes = append(e.Notes, note)
    e.CalculerMoyenne()
}

func (e *Etudiant) CalculerMoyenne() {
    if len(e.Notes) == 0 {
        e.Moyenne = 0
        return
    }

    var somme float64
    for _, note := range e.Notes {
        somme += note
    }
    e.Moyenne = somme / float64(len(e.Notes))
}

func (e Etudiant) String() string {
    return fmt.Sprintf("%s %s (ID: %d) - Moyenne: %.2f",
        e.Prenom, e.Nom, e.ID, e.Moyenne)
}

func main() {
    etudiant := &Etudiant{
        ID:     12345,
        Nom:    "Dupont",
        Prenom: "Alice",
        Email:  "alice.dupont@email.com",
    }

    etudiant.AjouterNote(15.5)
    etudiant.AjouterNote(18.0)
    etudiant.AjouterNote(16.5)

    fmt.Println(etudiant)
    fmt.Printf("Notes: %v\n", etudiant.Notes)
}
```

### Exemple 2 : E-commerce simple

```go
type Produit struct {
    ID          int
    Nom         string
    Prix        float64
    Stock       int
    Description string
}

type Panier struct {
    Items []ItemPanier
    Total float64
}

type ItemPanier struct {
    Produit  Produit
    Quantite int
}

func (p *Panier) AjouterProduit(produit Produit, quantite int) error {
    if produit.Stock < quantite {
        return fmt.Errorf("stock insuffisant pour %s", produit.Nom)
    }

    // Vérifier si le produit est déjà dans le panier
    for i, item := range p.Items {
        if item.Produit.ID == produit.ID {
            p.Items[i].Quantite += quantite
            p.CalculerTotal()
            return nil
        }
    }

    // Ajouter nouveau produit
    p.Items = append(p.Items, ItemPanier{
        Produit:  produit,
        Quantite: quantite,
    })

    p.CalculerTotal()
    return nil
}

func (p *Panier) CalculerTotal() {
    p.Total = 0
    for _, item := range p.Items {
        p.Total += item.Produit.Prix * float64(item.Quantite)
    }
}

func (p Panier) Afficher() {
    fmt.Println("=== PANIER ===")
    for _, item := range p.Items {
        fmt.Printf("%s x%d - %.2f€ = %.2f€\n",
            item.Produit.Nom,
            item.Quantite,
            item.Produit.Prix,
            item.Produit.Prix*float64(item.Quantite))
    }
    fmt.Printf("TOTAL: %.2f€\n", p.Total)
}

func main() {
    // Créer des produits
    produit1 := Produit{
        ID:          1,
        Nom:         "MacBook Pro",
        Prix:        2499.99,
        Stock:       5,
        Description: "Ordinateur portable haut de gamme",
    }

    produit2 := Produit{
        ID:          2,
        Nom:         "iPhone 15",
        Prix:        1199.99,
        Stock:       10,
        Description: "Smartphone dernière génération",
    }

    // Créer un panier
    panier := &Panier{}

    // Ajouter des produits
    if err := panier.AjouterProduit(produit1, 1); err != nil {
        fmt.Printf("Erreur: %v\n", err)
    }

    if err := panier.AjouterProduit(produit2, 2); err != nil {
        fmt.Printf("Erreur: %v\n", err)
    }

    panier.Afficher()
}
```

### Exemple 3 : Jeu simple (RPG)

```go
type Personnage struct {
    Nom       string
    Vie       int
    VieMax    int
    Attaque   int
    Defense   int
    Niveau    int
    Experience int
}

func NouveauPersonnage(nom string) *Personnage {
    return &Personnage{
        Nom:        nom,
        Vie:        100,
        VieMax:     100,
        Attaque:    10,
        Defense:    5,
        Niveau:     1,
        Experience: 0,
    }
}

func (p *Personnage) AttaquerPersonnage(cible *Personnage) {
    degats := p.Attaque - cible.Defense
    if degats < 1 {
        degats = 1
    }

    cible.Vie -= degats
    if cible.Vie < 0 {
        cible.Vie = 0
    }

    fmt.Printf("%s attaque %s pour %d dégâts!\n", p.Nom, cible.Nom, degats)
    fmt.Printf("%s a maintenant %d PV\n", cible.Nom, cible.Vie)
}

func (p *Personnage) EstVivant() bool {
    return p.Vie > 0
}

func (p *Personnage) Guerir(montant int) {
    p.Vie += montant
    if p.Vie > p.VieMax {
        p.Vie = p.VieMax
    }
    fmt.Printf("%s récupère %d PV (PV actuels: %d/%d)\n",
        p.Nom, montant, p.Vie, p.VieMax)
}

func (p *Personnage) GagnerExperience(exp int) {
    p.Experience += exp
    fmt.Printf("%s gagne %d XP (Total: %d XP)\n", p.Nom, exp, p.Experience)

    // Niveau suivant tous les 100 XP
    if p.Experience >= p.Niveau*100 {
        p.MonterNiveau()
    }
}

func (p *Personnage) MonterNiveau() {
    p.Niveau++
    p.VieMax += 20
    p.Vie = p.VieMax // Guérison complète
    p.Attaque += 5
    p.Defense += 2

    fmt.Printf("🎉 %s monte au niveau %d!\n", p.Nom, p.Niveau)
    fmt.Printf("   Nouvelles stats: Vie=%d, Attaque=%d, Défense=%d\n",
        p.VieMax, p.Attaque, p.Defense)
}

func (p Personnage) String() string {
    return fmt.Sprintf("%s (Niv.%d) - PV: %d/%d, ATT: %d, DEF: %d, XP: %d",
        p.Nom, p.Niveau, p.Vie, p.VieMax, p.Attaque, p.Defense, p.Experience)
}

func main() {
    // Créer des personnages
    hero := NouveauPersonnage("Aragorn")
    monstre := NouveauPersonnage("Orc")

    fmt.Println("=== COMBAT ===")
    fmt.Printf("Héros: %s\n", hero)
    fmt.Printf("Monstre: %s\n", monstre)

    // Combat simple
    tour := 1
    for hero.EstVivant() && monstre.EstVivant() {
        fmt.Printf("\n--- Tour %d ---\n", tour)

        hero.AttaquerPersonnage(monstre)
        if monstre.EstVivant() {
            monstre.AttaquerPersonnage(hero)
        }

        tour++
        if tour > 10 { // Éviter une boucle infinie
            break
        }
    }

    fmt.Println("\n=== RÉSULTAT ===")
    if hero.EstVivant() {
        fmt.Printf("🏆 %s gagne le combat!\n", hero.Nom)
        hero.GagnerExperience(50)
    } else {
        fmt.Printf("💀 %s est vaincu...\n", hero.Nom)
    }
}
```

## Bonnes pratiques

### 1. Nommage des structs

```go
// ✅ Bon : noms descriptifs et en PascalCase
type Utilisateur struct {
    ID   int
    Nom  string
}

// ✅ Bon : prefixe pour éviter les conflits
type HTTPClient struct {
    Timeout time.Duration
}

// ❌ Évitez : noms trop génériques
type Data struct {
    Value interface{}
}
```

### 2. Organisation des champs

```go
// ✅ Bon : grouper les champs logiquement
type Produit struct {
    // Identification
    ID  int
    SKU string

    // Informations
    Nom         string
    Description string

    // Financier
    Prix     float64
    TVA      float64

    // Stock
    Quantite    int
    SeuilAlerte int
}
```

### 3. Utilisation des pointeurs

```go
// ✅ Pointeur quand la struct est grosse ou modifiée
func (p *GrosseStruct) Modifier() {
    // Modification des champs
}

// ✅ Valeur pour les petites structs en lecture seule
func (p Point) Distance() float64 {
    return math.Sqrt(p.X*p.X + p.Y*p.Y)
}
```

### 4. Validation dans les constructeurs

```go
func NouvelUtilisateur(nom, email string, age int) (*Utilisateur, error) {
    if nom == "" {
        return nil, fmt.Errorf("le nom ne peut pas être vide")
    }
    if age < 0 {
        return nil, fmt.Errorf("l'âge ne peut pas être négatif")
    }
    if !strings.Contains(email, "@") {
        return nil, fmt.Errorf("email invalide")
    }

    return &Utilisateur{
        Nom:   nom,
        Email: email,
        Age:   age,
    }, nil
}
```

## Exercices pratiques

### Exercice 1 : Basique

Créez une struct `Livre` avec les champs : titre, auteur, pages, prix. Implémentez :

1. Un constructeur `NouveauLivre()`
2. Une méthode `String()` pour l'affichage
3. Une méthode `AppliquerRemise(pourcentage float64)`

### Exercice 2 : Intermédiaire

Créez un système de bibliothèque avec :

1. Struct `Livre` (titre, auteur, ISBN, disponible)
2. Struct `Bibliotheque` avec une slice de livres
3. Méthodes : `AjouterLivre()`, `EmprunterLivre()`, `RetournerLivre()`, `RechercherParAuteur()`

### Exercice 3 : Avancé

Implémentez un système de gestion d'employés :

1. Struct `Adresse` (rue, ville, code postal)
2. Struct `Employe` (nom, prénom, adresse, salaire, département)
3. Struct `Entreprise` avec slice d'employés
4. Méthodes pour calculer la masse salariale par département, trouver les employés par ville, etc.

```go
package main

import (
 "fmt"
 "strings"
 "time"
)

// ==========================================
// EXERCICE 1 : BASIQUE
// ==========================================

// Struct Livre avec les champs requis
type Livre struct {
 Titre  string
 Auteur string
 Pages  int
 Prix   float64
}

// Constructeur avec validation
func NouveauLivre(titre, auteur string, pages int, prix float64) (*Livre, error) {
 // Validation des paramètres
 if titre == "" {
  return nil, fmt.Errorf("le titre ne peut pas être vide")
 }
 if auteur == "" {
  return nil, fmt.Errorf("l'auteur ne peut pas être vide")
 }
 if pages <= 0 {
  return nil, fmt.Errorf("le nombre de pages doit être positif")
 }
 if prix < 0 {
  return nil, fmt.Errorf("le prix ne peut pas être négatif")
 }

 return &Livre{
  Titre:  titre,
  Auteur: auteur,
  Pages:  pages,
  Prix:   prix,
 }, nil
}

// Méthode String pour l'affichage
func (l Livre) String() string {
 return fmt.Sprintf("📖 \"%s\" par %s (%d pages) - %.2f€",
  l.Titre, l.Auteur, l.Pages, l.Prix)
}

// Méthode pour appliquer une remise
func (l *Livre) AppliquerRemise(pourcentage float64) error {
 if pourcentage < 0 || pourcentage > 100 {
  return fmt.Errorf("le pourcentage doit être entre 0 et 100")
 }

 ancienPrix := l.Prix
 l.Prix = l.Prix * (1.0 - pourcentage/100.0)

 fmt.Printf("Remise de %.1f%% appliquée: %.2f€ → %.2f€\n",
  pourcentage, ancienPrix, l.Prix)
 return nil
}

// Méthode bonus pour calculer le prix par page
func (l Livre) PrixParPage() float64 {
 if l.Pages == 0 {
  return 0
 }
 return l.Prix / float64(l.Pages)
}

func exercice1() {
 fmt.Println("=== EXERCICE 1 : BASIQUE ===")

 // Créer des livres avec validation
 livre1, err := NouveauLivre("Le Petit Prince", "Antoine de Saint-Exupéry", 96, 8.50)
 if err != nil {
  fmt.Printf("Erreur: %v\n", err)
  return
 }

 livre2, err := NouveauLivre("1984", "George Orwell", 328, 12.99)
 if err != nil {
  fmt.Printf("Erreur: %v\n", err)
  return
 }

 livre3, err := NouveauLivre("Go Programming", "Alan Donovan", 380, 45.00)
 if err != nil {
  fmt.Printf("Erreur: %v\n", err)
  return
 }

 // Afficher les livres
 fmt.Println("Livres créés:")
 fmt.Printf("1. %s\n", livre1)
 fmt.Printf("2. %s\n", livre2)
 fmt.Printf("3. %s\n", livre3)

 // Appliquer des remises
 fmt.Println("\nApplication de remises:")
 livre1.AppliquerRemise(10.0) // 10% de remise
 livre2.AppliquerRemise(15.0) // 15% de remise
 livre3.AppliquerRemise(20.0) // 20% de remise

 // Afficher après remises
 fmt.Println("\nAprès remises:")
 fmt.Printf("1. %s (%.4f€/page)\n", livre1, livre1.PrixParPage())
 fmt.Printf("2. %s (%.4f€/page)\n", livre2, livre2.PrixParPage())
 fmt.Printf("3. %s (%.4f€/page)\n", livre3, livre3.PrixParPage())

 // Test de validation
 fmt.Println("\nTest de validation:")
 if _, err := NouveauLivre("", "Auteur", 100, 10.0); err != nil {
  fmt.Printf("Erreur attendue: %v\n", err)
 }

 if err := livre1.AppliquerRemise(150.0); err != nil {
  fmt.Printf("Erreur attendue: %v\n", err)
 }

 fmt.Println()
}

// ==========================================
// EXERCICE 2 : INTERMÉDIAIRE
// ==========================================

// Struct Livre pour la bibliothèque (différente de l'exercice 1)
type LivreBiblio struct {
 Titre      string
 Auteur     string
 ISBN       string
 Disponible bool
}

// Constructeur pour LivreBiblio
func NouveauLivreBiblio(titre, auteur, isbn string) *LivreBiblio {
 return &LivreBiblio{
  Titre:      titre,
  Auteur:     auteur,
  ISBN:       isbn,
  Disponible: true, // Disponible par défaut
 }
}

func (l LivreBiblio) String() string {
 statut := "📗 Disponible"
 if !l.Disponible {
  statut = "📕 Emprunté"
 }
 return fmt.Sprintf("%s - \"%s\" par %s (ISBN: %s)",
  statut, l.Titre, l.Auteur, l.ISBN)
}

// Struct Bibliothèque
type Bibliotheque struct {
 Nom    string
 Livres []*LivreBiblio
}

// Constructeur pour Bibliothèque
func NouvelleBibliotheque(nom string) *Bibliotheque {
 return &Bibliotheque{
  Nom:    nom,
  Livres: make([]*LivreBiblio, 0),
 }
}

// Ajouter un livre
func (b *Bibliotheque) AjouterLivre(livre *LivreBiblio) {
 // Vérifier si le livre existe déjà (par ISBN)
 for _, l := range b.Livres {
  if l.ISBN == livre.ISBN {
   fmt.Printf("⚠️  Livre avec ISBN %s existe déjà\n", livre.ISBN)
   return
  }
 }

 b.Livres = append(b.Livres, livre)
 fmt.Printf("✅ Livre ajouté: %s\n", livre.Titre)
}

// Emprunter un livre par ISBN
func (b *Bibliotheque) EmprunterLivre(isbn string) error {
 for _, livre := range b.Livres {
  if livre.ISBN == isbn {
   if !livre.Disponible {
    return fmt.Errorf("le livre \"%s\" est déjà emprunté", livre.Titre)
   }
   livre.Disponible = false
   fmt.Printf("📚 Livre emprunté: %s\n", livre.Titre)
   return nil
  }
 }
 return fmt.Errorf("livre avec ISBN %s non trouvé", isbn)
}

// Retourner un livre par ISBN
func (b *Bibliotheque) RetournerLivre(isbn string) error {
 for _, livre := range b.Livres {
  if livre.ISBN == isbn {
   if livre.Disponible {
    return fmt.Errorf("le livre \"%s\" n'était pas emprunté", livre.Titre)
   }
   livre.Disponible = true
   fmt.Printf("📖 Livre retourné: %s\n", livre.Titre)
   return nil
  }
 }
 return fmt.Errorf("livre avec ISBN %s non trouvé", isbn)
}

// Rechercher par auteur
func (b *Bibliotheque) RechercherParAuteur(auteur string) []*LivreBiblio {
 var resultats []*LivreBiblio
 auteurLower := strings.ToLower(auteur)

 for _, livre := range b.Livres {
  if strings.Contains(strings.ToLower(livre.Auteur), auteurLower) {
   resultats = append(resultats, livre)
  }
 }

 return resultats
}

// Afficher tous les livres
func (b *Bibliotheque) AfficherCatalogue() {
 fmt.Printf("\n=== CATALOGUE DE %s ===\n", strings.ToUpper(b.Nom))
 if len(b.Livres) == 0 {
  fmt.Println("Aucun livre dans la bibliothèque.")
  return
 }

 for i, livre := range b.Livres {
  fmt.Printf("%d. %s\n", i+1, livre)
 }
}

// Statistiques de la bibliothèque
func (b *Bibliotheque) Statistiques() {
 total := len(b.Livres)
 disponibles := 0
 empruntes := 0

 for _, livre := range b.Livres {
  if livre.Disponible {
   disponibles++
  } else {
   empruntes++
  }
 }

 fmt.Printf("\n=== STATISTIQUES %s ===\n", strings.ToUpper(b.Nom))
 fmt.Printf("Total des livres: %d\n", total)
 fmt.Printf("Disponibles: %d\n", disponibles)
 fmt.Printf("Empruntés: %d\n", empruntes)
 if total > 0 {
  fmt.Printf("Taux d'emprunt: %.1f%%\n", float64(empruntes)/float64(total)*100)
 }
}

func exercice2() {
 fmt.Println("=== EXERCICE 2 : INTERMÉDIAIRE ===")

 // Créer une bibliothèque
 biblio := NouvelleBibliotheque("Bibliothèque Municipale")

 // Créer des livres
 livre1 := NouveauLivreBiblio("Le Petit Prince", "Antoine de Saint-Exupéry", "978-0156012195")
 livre2 := NouveauLivreBiblio("1984", "George Orwell", "978-0452284234")
 livre3 := NouveauLivreBiblio("Dune", "Frank Herbert", "978-0441172719")
 livre4 := NouveauLivreBiblio("La Ferme des Animaux", "George Orwell", "978-0451526342")
 livre5 := NouveauLivreBiblio("Le Guide du Routard Galactique", "Douglas Adams", "978-0345391803")

 // Ajouter les livres
 fmt.Println("\n--- Ajout de livres ---")
 biblio.AjouterLivre(livre1)
 biblio.AjouterLivre(livre2)
 biblio.AjouterLivre(livre3)
 biblio.AjouterLivre(livre4)
 biblio.AjouterLivre(livre5)

 // Tentative d'ajout d'un doublon
 livre1bis := NouveauLivreBiblio("Le Petit Prince (édition illustrée)", "Antoine de Saint-Exupéry", "978-0156012195")
 biblio.AjouterLivre(livre1bis)

 // Afficher le catalogue
 biblio.AfficherCatalogue()

 // Emprunter des livres
 fmt.Println("\n--- Emprunts ---")
 biblio.EmprunterLivre("978-0156012195") // Le Petit Prince
 biblio.EmprunterLivre("978-0441172719") // Dune

 // Tentative d'emprunt d'un livre déjà emprunté
 if err := biblio.EmprunterLivre("978-0156012195"); err != nil {
  fmt.Printf("❌ %v\n", err)
 }

 // Tentative d'emprunt d'un livre inexistant
 if err := biblio.EmprunterLivre("978-9999999999"); err != nil {
  fmt.Printf("❌ %v\n", err)
 }

 // Afficher le catalogue après emprunts
 biblio.AfficherCatalogue()

 // Retourner un livre
 fmt.Println("\n--- Retours ---")
 biblio.RetournerLivre("978-0156012195") // Le Petit Prince

 // Recherche par auteur
 fmt.Println("\n--- Recherche par auteur ---")
 livresOrwell := biblio.RechercherParAuteur("George Orwell")
 fmt.Printf("Livres de George Orwell (%d trouvé(s)):\n", len(livresOrwell))
 for _, livre := range livresOrwell {
  fmt.Printf("  - %s\n", livre)
 }

 livresSaint := biblio.RechercherParAuteur("Saint-Exupéry")
 fmt.Printf("Livres de Saint-Exupéry (%d trouvé(s)):\n", len(livresSaint))
 for _, livre := range livresSaint {
  fmt.Printf("  - %s\n", livre)
 }

 // Statistiques
 biblio.Statistiques()

 fmt.Println()
}

// ==========================================
// EXERCICE 3 : AVANCÉ
// ==========================================

// Struct Adresse
type Adresse struct {
 Rue        string
 Ville      string
 CodePostal string
}

func (a Adresse) String() string {
 return fmt.Sprintf("%s, %s %s", a.Rue, a.CodePostal, a.Ville)
}

// Struct Employé
type Employe struct {
 ID          int
 Nom         string
 Prenom      string
 Adresse     Adresse
 Salaire     float64
 Departement string
 DateEmbauche time.Time
}

func (e Employe) String() string {
 return fmt.Sprintf("%s %s (ID: %d) - %s - %.2f€/mois",
  e.Prenom, e.Nom, e.ID, e.Departement, e.Salaire)
}

func (e Employe) NomComplet() string {
 return fmt.Sprintf("%s %s", e.Prenom, e.Nom)
}

func (e Employe) AncienneteAnnees() int {
 return int(time.Since(e.DateEmbauche).Hours() / 24 / 365)
}

// Constructeur pour Employé
func NouvelEmploye(id int, nom, prenom string, adresse Adresse,
 salaire float64, departement string) *Employe {
 return &Employe{
  ID:           id,
  Nom:          nom,
  Prenom:       prenom,
  Adresse:      adresse,
  Salaire:      salaire,
  Departement:  departement,
  DateEmbauche: time.Now(),
 }
}

// Struct Entreprise
type Entreprise struct {
 Nom       string
 Employes  []*Employe
 prochainID int
}

// Constructeur pour Entreprise
func NouvelleEntreprise(nom string) *Entreprise {
 return &Entreprise{
  Nom:        nom,
  Employes:   make([]*Employe, 0),
  prochainID: 1,
 }
}

// Ajouter un employé
func (ent *Entreprise) AjouterEmploye(nom, prenom string, adresse Adresse,
 salaire float64, departement string) *Employe {
 employe := NouvelEmploye(ent.prochainID, nom, prenom, adresse, salaire, departement)
 ent.prochainID++
 ent.Employes = append(ent.Employes, employe)

 fmt.Printf("✅ Employé ajouté: %s\n", employe.NomComplet())
 return employe
}

// Supprimer un employé par ID
func (ent *Entreprise) SupprimerEmploye(id int) error {
 for i, employe := range ent.Employes {
  if employe.ID == id {
   // Supprimer l'employé du slice
   ent.Employes = append(ent.Employes[:i], ent.Employes[i+1:]...)
   fmt.Printf("🗑️  Employé supprimé: %s\n", employe.NomComplet())
   return nil
  }
 }
 return fmt.Errorf("employé avec ID %d non trouvé", id)
}

// Calculer la masse salariale par département
func (ent *Entreprise) MasseSalarialeParDepartement() map[string]float64 {
 masseSalariale := make(map[string]float64)

 for _, employe := range ent.Employes {
  masseSalariale[employe.Departement] += employe.Salaire
 }

 return masseSalariale
}

// Trouver les employés par ville
func (ent *Entreprise) EmployesParVille(ville string) []*Employe {
 var resultats []*Employe
 villeLower := strings.ToLower(ville)

 for _, employe := range ent.Employes {
  if strings.ToLower(employe.Adresse.Ville) == villeLower {
   resultats = append(resultats, employe)
  }
 }

 return resultats
}

// Trouver les employés par département
func (ent *Entreprise) EmployesParDepartement(departement string) []*Employe {
 var resultats []*Employe
 deptLower := strings.ToLower(departement)

 for _, employe := range ent.Employes {
  if strings.ToLower(employe.Departement) == deptLower {
   resultats = append(resultats, employe)
  }
 }

 return resultats
}

// Calculer le salaire moyen
func (ent *Entreprise) SalaireMoyen() float64 {
 if len(ent.Employes) == 0 {
  return 0
 }

 var total float64
 for _, employe := range ent.Employes {
  total += employe.Salaire
 }

 return total / float64(len(ent.Employes))
}

// Trouver l'employé avec le salaire le plus élevé
func (ent *Entreprise) EmployeAvecSalaireMax() *Employe {
 if len(ent.Employes) == 0 {
  return nil
 }

 maxEmploye := ent.Employes[0]
 for _, employe := range ent.Employes[1:] {
  if employe.Salaire > maxEmploye.Salaire {
   maxEmploye = employe
  }
 }

 return maxEmploye
}

// Augmenter les salaires d'un département
func (ent *Entreprise) AugmenterSalaires(departement string, pourcentage float64) {
 employesAffectes := 0

 for _, employe := range ent.Employes {
  if strings.ToLower(employe.Departement) == strings.ToLower(departement) {
   ancienSalaire := employe.Salaire
   employe.Salaire *= (1.0 + pourcentage/100.0)
   fmt.Printf("💰 %s: %.2f€ → %.2f€ (+%.1f%%)\n",
    employe.NomComplet(), ancienSalaire, employe.Salaire, pourcentage)
   employesAffectes++
  }
 }

 if employesAffectes == 0 {
  fmt.Printf("Aucun employé trouvé dans le département: %s\n", departement)
 } else {
  fmt.Printf("Augmentation appliquée à %d employé(s)\n", employesAffectes)
 }
}

// Afficher tous les employés
func (ent *Entreprise) AfficherEmployes() {
 fmt.Printf("\n=== EMPLOYÉS DE %s ===\n", strings.ToUpper(ent.Nom))
 if len(ent.Employes) == 0 {
  fmt.Println("Aucun employé.")
  return
 }

 for i, employe := range ent.Employes {
  fmt.Printf("%d. %s\n", i+1, employe)
  fmt.Printf("   📍 %s\n", employe.Adresse)
  fmt.Printf("   📅 Embauché depuis %d an(s)\n", employe.AncienneteAnnees())
 }
}

// Statistiques complètes
func (ent *Entreprise) StatistiquesCompletes() {
 fmt.Printf("\n=== STATISTIQUES %s ===\n", strings.ToUpper(ent.Nom))

 if len(ent.Employes) == 0 {
  fmt.Println("Aucun employé.")
  return
 }

 // Statistiques générales
 fmt.Printf("Nombre total d'employés: %d\n", len(ent.Employes))
 fmt.Printf("Salaire moyen: %.2f€\n", ent.SalaireMoyen())

 employeMax := ent.EmployeAvecSalaireMax()
 if employeMax != nil {
  fmt.Printf("Salaire le plus élevé: %.2f€ (%s)\n",
   employeMax.Salaire, employeMax.NomComplet())
 }

 // Masse salariale par département
 fmt.Println("\nMasse salariale par département:")
 masseSalariale := ent.MasseSalarialeParDepartement()
 var totalMasse float64
 for dept, masse := range masseSalariale {
  fmt.Printf("  %s: %.2f€/mois\n", dept, masse)
  totalMasse += masse
 }
 fmt.Printf("  TOTAL: %.2f€/mois (%.2f€/an)\n", totalMasse, totalMasse*12)

 // Répartition par département
 fmt.Println("\nRépartition par département:")
 departements := make(map[string]int)
 for _, employe := range ent.Employes {
  departements[employe.Departement]++
 }
 for dept, count := range departements {
  fmt.Printf("  %s: %d employé(s)\n", dept, count)
 }

 // Répartition par ville
 fmt.Println("\nRépartition par ville:")
 villes := make(map[string]int)
 for _, employe := range ent.Employes {
  villes[employe.Adresse.Ville]++
 }
 for ville, count := range villes {
  fmt.Printf("  %s: %d employé(s)\n", ville, count)
 }
}

func exercice3() {
 fmt.Println("=== EXERCICE 3 : AVANCÉ ===")

 // Créer une entreprise
 entreprise := NouvelleEntreprise("TechCorp Solutions")

 // Créer des adresses
 adresse1 := Adresse{"123 Rue de la République", "Lyon", "69001"}
 adresse2 := Adresse{"456 Avenue des Champs", "Paris", "75001"}
 adresse3 := Adresse{"789 Boulevard Central", "Lyon", "69002"}
 adresse4 := Adresse{"321 Rue Victor Hugo", "Marseille", "13001"}
 adresse5 := Adresse{"654 Place de la Liberté", "Paris", "75002"}

 // Ajouter des employés
 fmt.Println("\n--- Ajout d'employés ---")
 entreprise.AjouterEmploye("Dupont", "Alice", adresse1, 4500.0, "Développement")
 entreprise.AjouterEmploye("Martin", "Bob", adresse2, 5200.0, "Développement")
 entreprise.AjouterEmploye("Durand", "Charlie", adresse3, 3800.0, "Marketing")
 entreprise.AjouterEmploye("Bernard", "Diana", adresse4, 4100.0, "Marketing")
 entreprise.AjouterEmploye("Petit", "Eve", adresse5, 6000.0, "Direction")
 entreprise.AjouterEmploye("Moreau", "Frank", adresse1, 3500.0, "RH")
 entreprise.AjouterEmploye("Simon", "Grace", adresse2, 4800.0, "Développement")

 // Afficher tous les employés
 entreprise.AfficherEmployes()

 // Recherches par ville
 fmt.Println("\n--- Recherche par ville ---")
 employesLyon := entreprise.EmployesParVille("Lyon")
 fmt.Printf("Employés à Lyon (%d):\n", len(employesLyon))
 for _, emp := range employesLyon {
  fmt.Printf("  - %s (%s)\n", emp.NomComplet(), emp.Departement)
 }

 employesParis := entreprise.EmployesParVille("Paris")
 fmt.Printf("Employés à Paris (%d):\n", len(employesParis))
 for _, emp := range employesParis {
  fmt.Printf("  - %s (%s)\n", emp.NomComplet(), emp.Departement)
 }

 // Recherche par département
 fmt.Println("\n--- Recherche par département ---")
 employesDev := entreprise.EmployesParDepartement("Développement")
 fmt.Printf("Employés en Développement (%d):\n", len(employesDev))
 for _, emp := range employesDev {
  fmt.Printf("  - %s (%.2f€)\n", emp.NomComplet(), emp.Salaire)
 }

 // Augmentation de salaire
 fmt.Println("\n--- Augmentation de salaire ---")
 entreprise.AugmenterSalaires("Développement", 8.0) // 8% d'augmentation
 entreprise.AugmenterSalaires("Marketing", 5.0)     // 5% d'augmentation

 // Statistiques complètes
 entreprise.StatistiquesCompletes()

 // Test de suppression
 fmt.Println("\n--- Suppression d'employé ---")
 if err := entreprise.SupprimerEmploye(3); err != nil {
  fmt.Printf("Erreur: %v\n", err)
 }

 // Statistiques après suppression
 fmt.Println("\nStatistiques après suppression:")
 fmt.Printf("Nombre d'employés: %d\n", len(entreprise.Employes))
 fmt.Printf("Nouvelle masse salariale totale: %.2f€/mois\n",
  func() float64 {
   var total float64
   for _, masse := range entreprise.MasseSalarialeParDepartement() {
    total += masse
   }
   return total
  }())

 fmt.Println()
}

// ==========================================
// FONCTIONS BONUS
// ==========================================

// Fonction pour créer des données de test rapidement
func creerDonneesTest() (*Bibliotheque, *Entreprise) {
 // Bibliothèque de test
 biblio := NouvelleBibliotheque("Bibliothèque de Test")
 biblio.AjouterLivre(NouveauLivreBiblio("Test Livre 1", "Auteur Test", "111"))
 biblio.AjouterLivre(NouveauLivreBiblio("Test Livre 2", "Auteur Test", "222"))

 // Entreprise de test
 entreprise := NouvelleEntreprise("Test Corp")
 adresse := Adresse{"Test Rue", "Test Ville", "12345"}
 entreprise.AjouterEmploye("Test", "Employé", adresse, 3000.0, "Test")

 return biblio, entreprise
}

// Démonstration des patterns avancés
func demonstrationPatterns() {
 fmt.Println("=== DÉMONSTRATION PATTERNS AVANCÉS ===")

 // Pattern Builder pour Employé
 type EmployeBuilder struct {
  employe *Employe
 }

 func NouvelEmployeBuilder() *EmployeBuilder {
  return &EmployeBuilder{
   employe: &Employe{},
  }
 }

 func (eb *EmployeBuilder) Nom(nom string) *EmployeBuilder {
  eb.employe.Nom = nom
  return eb
 }

 func (eb *EmployeBuilder) Prenom(prenom string) *EmployeBuilder {
  eb.employe.Prenom = prenom
  return eb
 }

 func (eb *EmployeBuilder) Salaire(salaire float64) *EmployeBuilder {
  eb.employe.Salaire = salaire
  return eb
 }

 func (eb *EmployeBuilder) Build() *Employe {
  eb.employe.DateEmbauche = time.Now()
  return eb.employe
 }

 // Utilisation du Builder
 employe := NouvelEmployeBuilder().
  Nom("Builder").
  Prenom("Pattern").
  Salaire(5000.0).
  Build()

 fmt.Printf("Employé créé avec Builder: %s\n", employe.NomComplet())
}

// ==========================================
// FONCTION PRINCIPALE
// ==========================================

func main() {
 fmt.Println("SOLUTIONS DES EXERCICES - STRUCTS")
 fmt.Println("==================================")

 exercice1()
 exercice2()
 exercice3()

 // Bonus: démonstration de patterns avancés
 demonstrationPatterns()

 // Bonus: test de performance
 demonstrationPerformance()
}

// ==========================================
// DÉMONSTRATIONS BONUS
// ==========================================

func demonstrationPerformance() {
 fmt.Println("\n=== DÉMONSTRATION PERFORMANCE ===")

 // Comparaison passage par valeur vs pointeur
 type GrosseStruct struct {
  Donnees [1000]int
  Nom     string
  Actif   bool
 }

 // Fonction avec passage par valeur (copie)
 func traiterParValeur(gs GrosseStruct) {
  // Toute la struct est copiée !
  gs.Nom = "Modifié"
 }

 // Fonction avec passage par pointeur (référence)
 func traiterParPointeur(gs *GrosseStruct) {
  // Seul le pointeur est passé
  gs.Nom = "Modifié"
 }

 gs := GrosseStruct{Nom: "Original"}

 fmt.Printf("Avant traitement par valeur: %s\n", gs.Nom)
 traiterParValeur(gs)
 fmt.Printf("Après traitement par valeur: %s\n", gs.Nom) // Pas modifié

 fmt.Printf("Avant traitement par pointeur: %s\n", gs.Nom)
 traiterParPointeur(&gs)
 fmt.Printf("Après traitement par pointeur: %s\n", gs.Nom) // Modifié

 // Démonstration avec benchmark simulé
 fmt.Println("\n💡 Pour les grosses structs, utilisez des pointeurs!")
 fmt.Println("   - Évite la copie de données")
 fmt.Println("   - Permet la modification")
 fmt.Println("   - Plus efficace en mémoire")
}
```

## Résumé

Les structs sont le fondement de l'organisation des données en Go :

**Points clés à retenir :**

- Les structs regroupent des données de types différents
- Elles supportent la composition via l'embedding
- Les méthodes peuvent être attachées aux structs
- Utilisez des constructeurs pour la validation
- Les champs privés commencent par une minuscule
- Préférez les pointeurs pour les grosses structs ou les modifications

**Patterns courants :**

- Constructeurs avec validation
- Méthodes receiver par pointeur pour modification
- Embedding pour la composition
- Tags pour les métadonnées

Les structs permettent de créer des abstractions puissantes et de structurer votre code de manière claire et maintenable.

⏭️
