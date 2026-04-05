🔝 Retour au [Sommaire](/SOMMAIRE.md)

# 6-1 : Définition de méthodes

## Introduction

En Go, les **méthodes** sont des fonctions qui sont associées à un type spécifique. Contrairement aux langages orientés objet traditionnels, Go n'a pas de classes, mais il permet d'attacher des méthodes à n'importe quel type que vous définissez.

Une méthode est essentiellement une fonction avec un **récepteur** (receiver) spécial qui détermine sur quel type la méthode peut être appelée.

## Syntaxe de base

```go
func (récepteur TypeRécepteur) nomMéthode(paramètres) typeRetour {
    // Corps de la méthode
}
```

Le récepteur est placé entre `func` et le nom de la méthode, entre parenthèses.

## Exemple simple avec une struct

Commençons par un exemple concret avec une struct `Rectangle` :

```go
package main

import "fmt"

// Définition de la struct Rectangle
type Rectangle struct {
    largeur  float64
    hauteur  float64
}

// Méthode pour calculer l'aire du rectangle
func (r Rectangle) Aire() float64 {
    return r.largeur * r.hauteur
}

// Méthode pour calculer le périmètre du rectangle
func (r Rectangle) Perimetre() float64 {
    return 2 * (r.largeur + r.hauteur)
}

func main() {
    // Création d'un rectangle
    rect := Rectangle{largeur: 5.0, hauteur: 3.0}

    // Appel des méthodes
    fmt.Printf("Aire: %.2f\n", rect.Aire())
    fmt.Printf("Périmètre: %.2f\n", rect.Perimetre())
}
```

**Sortie :**

```
Aire: 15.00
Périmètre: 16.00
```

## Récepteur par valeur vs par pointeur

### Récepteur par valeur

Quand vous utilisez un récepteur par valeur, Go travaille avec une **copie** de la valeur :

```go
type Compteur struct {
    valeur int
}

// Récepteur par valeur - ne modifie PAS l'original
func (c Compteur) Incrementer() {
    c.valeur++ // Modifie seulement la copie
}

// Méthode pour obtenir la valeur
func (c Compteur) Valeur() int {
    return c.valeur
}

func main() {
    compteur := Compteur{valeur: 0}
    compteur.Incrementer()
    fmt.Println(compteur.Valeur()) // Affiche: 0 (pas modifié!)
}
```

### Récepteur par pointeur

Avec un récepteur par pointeur, vous travaillez avec l'**original** :

```go
type Compteur struct {
    valeur int
}

// Récepteur par pointeur - modifie l'original
func (c *Compteur) Incrementer() {
    c.valeur++ // Modifie l'original
}

// Méthode pour obtenir la valeur
func (c *Compteur) Valeur() int {
    return c.valeur
}

func main() {
    compteur := Compteur{valeur: 0}
    compteur.Incrementer() // Go convertit automatiquement en (&compteur).Incrementer()
    fmt.Println(compteur.Valeur()) // Affiche: 1 (modifié!)
}
```

## Quand utiliser un récepteur par pointeur ?

Utilisez un récepteur par pointeur quand :

1. **Vous devez modifier la valeur** du récepteur
2. **Le récepteur est volumineux** (pour éviter la copie)
3. **Vous voulez la cohérence** (si certaines méthodes utilisent des pointeurs)

### Exemple pratique : Compte bancaire

```go
type CompteBancaire struct {
    solde   float64
    titulaire string
}

// Déposer de l'argent (modifie le solde)
func (c *CompteBancaire) Deposer(montant float64) {
    if montant > 0 {
        c.solde += montant
    }
}

// Retirer de l'argent (modifie le solde)
func (c *CompteBancaire) Retirer(montant float64) bool {
    if montant > 0 && c.solde >= montant {
        c.solde -= montant
        return true
    }
    return false
}

// Consulter le solde (lecture seule)
func (c *CompteBancaire) Solde() float64 {
    return c.solde
}

// Obtenir le titulaire (lecture seule)
func (c *CompteBancaire) Titulaire() string {
    return c.titulaire
}

func main() {
    compte := CompteBancaire{
        solde:     1000.0,
        titulaire: "Alice Martin",
    }

    fmt.Printf("Solde initial: %.2f€\n", compte.Solde())

    compte.Deposer(500.0)
    fmt.Printf("Après dépôt: %.2f€\n", compte.Solde())

    if compte.Retirer(200.0) {
        fmt.Printf("Après retrait: %.2f€\n", compte.Solde())
    }
}
```

## Méthodes sur des types personnalisés

Vous pouvez définir des méthodes sur n'importe quel type que vous créez, pas seulement les structs :

```go
// Type personnalisé basé sur int
type Temperature int

// Méthode pour convertir en Celsius
func (t Temperature) Celsius() float64 {
    return float64(t)
}

// Méthode pour convertir en Fahrenheit
func (t Temperature) Fahrenheit() float64 {
    return float64(t)*9/5 + 32
}

// Méthode pour convertir en Kelvin
func (t Temperature) Kelvin() float64 {
    return float64(t) + 273.15
}

func main() {
    temp := Temperature(25) // 25°C

    fmt.Printf("Température: %.1f°C\n", temp.Celsius())
    fmt.Printf("En Fahrenheit: %.1f°F\n", temp.Fahrenheit())
    fmt.Printf("En Kelvin: %.1fK\n", temp.Kelvin())
}
```

## Exemple complet : Gestionnaire de tâches

```go
package main

import (
    "fmt"
    "time"
)

type Tache struct {
    id          int
    description string
    terminee    bool
    dateCreation time.Time
}

type GestionnaireTaches struct {
    taches []Tache
    prochainId int
}

// Créer une nouvelle tâche
func (g *GestionnaireTaches) AjouterTache(description string) {
    tache := Tache{
        id:           g.prochainId,
        description:  description,
        terminee:     false,
        dateCreation: time.Now(),
    }
    g.taches = append(g.taches, tache)
    g.prochainId++
}

// Marquer une tâche comme terminée
func (g *GestionnaireTaches) TerminerTache(id int) bool {
    for i := range g.taches {
        if g.taches[i].id == id {
            g.taches[i].terminee = true
            return true
        }
    }
    return false
}

// Afficher toutes les tâches
func (g *GestionnaireTaches) AfficherTaches() {
    fmt.Println("=== Liste des tâches ===")
    for _, tache := range g.taches {
        statut := "❌"
        if tache.terminee {
            statut = "✅"
        }
        fmt.Printf("%s [%d] %s\n", statut, tache.id, tache.description)
    }
}

// Compter les tâches terminées
func (g *GestionnaireTaches) NombreTachesTerminees() int {
    count := 0
    for _, tache := range g.taches {
        if tache.terminee {
            count++
        }
    }
    return count
}

func main() {
    // Créer un gestionnaire de tâches
    gestionnaire := GestionnaireTaches{
        taches:     make([]Tache, 0),
        prochainId: 1,
    }

    // Ajouter des tâches
    gestionnaire.AjouterTache("Apprendre Go")
    gestionnaire.AjouterTache("Faire les courses")
    gestionnaire.AjouterTache("Réviser pour l'examen")

    // Afficher les tâches
    gestionnaire.AfficherTaches()

    // Terminer une tâche
    gestionnaire.TerminerTache(1)

    fmt.Println("\nAprès avoir terminé la tâche 1:")
    gestionnaire.AfficherTaches()

    fmt.Printf("\nTâches terminées: %d\n", gestionnaire.NombreTachesTerminees())
}
```

## Bonnes pratiques

1. **Cohérence des récepteurs** : Si une méthode d'un type utilise un récepteur par pointeur, utilisez des pointeurs pour toutes les méthodes de ce type.

2. **Noms de récepteurs courts** : Utilisez des noms courts (généralement la première lettre du type) : `func (r Rectangle)`, `func (c Compteur)`.

3. **Récepteurs par pointeur pour les modifications** : Utilisez toujours des pointeurs quand vous devez modifier l'état du récepteur.

4. **Méthodes sur des types intégrés** : Vous ne pouvez pas définir de méthodes sur des types intégrés comme `int`, `string`, etc. Créez d'abord un type personnalisé.

## Résumé

Les méthodes en Go permettent d'associer des fonctions à des types spécifiques, créant ainsi une forme d'encapsulation. Les points clés à retenir :

- Les méthodes ont un récepteur qui détermine le type sur lequel elles peuvent être appelées
- Utilisez des récepteurs par valeur pour la lecture, par pointeur pour la modification
- Les méthodes peuvent être définies sur n'importe quel type personnalisé
- Go convertit automatiquement entre valeurs et pointeurs lors de l'appel de méthodes

Dans la prochaine section, nous verrons comment utiliser les interfaces pour créer des abstractions puissantes avec ces méthodes.

⏭️
