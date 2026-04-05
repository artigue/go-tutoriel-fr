🔝 Retour au [Sommaire](/SOMMAIRE.md)

# 5-2 : Maps

## Introduction

Une **map** (carte ou dictionnaire en français) est une structure de données qui stocke des paires **clé-valeur**. C'est l'équivalent des dictionnaires en Python, des objets en JavaScript, ou des HashMap en Java. Les maps permettent de retrouver rapidement une valeur en utilisant sa clé unique.

## Concept de base

Imaginez un dictionnaire de traduction :

- **Clé** : "hello" (mot en anglais)
- **Valeur** : "bonjour" (traduction en français)

En Go, cela se représente ainsi :

```go
traduction := map[string]string{
    "hello":   "bonjour",
    "goodbye": "au revoir",
    "thank you": "merci",
}
```

## Déclaration de Maps

### Méthode 1 : Déclaration avec valeurs initiales

```go
// Map avec des valeurs initiales
ages := map[string]int{
    "Alice": 25,
    "Bob":   30,
    "Carol": 28,
}

// Map vide avec des types spécifiés
var scores map[string]int = map[string]int{}

// Syntaxe courte pour map vide
notes := map[string]float64{}
```

### Méthode 2 : Déclaration avec make()

```go
// Créer une map vide avec make()
prix := make(map[string]float64)

// Ajouter des éléments
prix["pain"] = 1.50
prix["lait"] = 0.95
prix["œufs"] = 2.30
```

### Méthode 3 : Déclaration sans initialisation

```go
var inventaire map[string]int // map = nil

// ATTENTION : il faut l'initialiser avant utilisation !
inventaire = make(map[string]int)
// ou
inventaire = map[string]int{}
```

## Types de clés et valeurs

```go
// Différents types de maps
ages := map[string]int{}           // clé: string, valeur: int
coordonnees := map[int]float64{}   // clé: int, valeur: float64
actif := map[string]bool{}         // clé: string, valeur: bool
listes := map[string][]int{}       // clé: string, valeur: slice d'int

// Les types de clés doivent être "comparables"
// ✅ Autorisés : string, int, float, bool, array
// ❌ Interdits : slice, map, function
```

## Opérations de base

### Ajouter et modifier des éléments

```go
fruits := make(map[string]int)

// Ajouter des éléments
fruits["pomme"] = 5
fruits["banane"] = 8
fruits["orange"] = 3

fmt.Println(fruits) // map[banane:8 orange:3 pomme:5]

// Modifier un élément existant
fruits["pomme"] = 7
fmt.Println(fruits["pomme"]) // 7
```

### Accéder aux éléments

```go
ages := map[string]int{
    "Alice": 25,
    "Bob":   30,
}

// Accès simple
fmt.Println(ages["Alice"]) // 25

// Accès avec vérification d'existence
age, existe := ages["Carol"]
if existe {
    fmt.Printf("L'âge de Carol est %d\n", age)
} else {
    fmt.Println("Carol n'est pas dans la map")
}

// Accès à une clé inexistante retourne la "zero value"
fmt.Println(ages["David"]) // 0 (zero value pour int)
```

### Supprimer des éléments

```go
scores := map[string]int{
    "Alice": 95,
    "Bob":   87,
    "Carol": 92,
}

fmt.Println("Avant:", scores) // map[Alice:95 Bob:87 Carol:92]

// Supprimer un élément
delete(scores, "Bob")
fmt.Println("Après:", scores) // map[Alice:95 Carol:92]

// Supprimer une clé inexistante ne cause pas d'erreur
delete(scores, "David") // Pas d'erreur
```

### Longueur d'une map

```go
utilisateurs := map[string]string{
    "admin": "password123",
    "user1": "secret456",
    "user2": "pass789",
}

fmt.Println("Nombre d'utilisateurs:", len(utilisateurs)) // 3
```

## Parcourir une Map

### Parcours avec range

```go
produits := map[string]float64{
    "café":     4.50,
    "thé":      3.20,
    "chocolat": 2.80,
}

// Parcourir clés et valeurs
for produit, prix := range produits {
    fmt.Printf("%s coûte %.2f€\n", produit, prix)
}

// Parcourir seulement les clés
for produit := range produits {
    fmt.Printf("Produit disponible: %s\n", produit)
}

// Parcourir seulement les valeurs
for _, prix := range produits {
    fmt.Printf("Prix: %.2f€\n", prix)
}
```

**Important :** L'ordre de parcours des maps est **aléatoire** ! Ne comptez jamais sur un ordre spécifique.

### Parcours ordonné

```go
import "sort"

scores := map[string]int{
    "Charlie": 85,
    "Alice":   95,
    "Bob":     87,
}

// Créer un slice des clés
var noms []string
for nom := range scores {
    noms = append(noms, nom)
}

// Trier les clés
sort.Strings(noms)

// Parcourir dans l'ordre alphabétique
fmt.Println("Scores par ordre alphabétique:")
for _, nom := range noms {
    fmt.Printf("%s: %d\n", nom, scores[nom])
}
```

## Vérifier l'existence d'une clé

```go
stock := map[string]int{
    "pommes":  15,
    "bananes": 8,
    "oranges": 0, // Attention : 0 est une valeur valide !
}

// Méthode correcte pour vérifier l'existence
quantite, existe := stock["pommes"]
if existe {
    fmt.Printf("Pommes en stock: %d\n", quantite)
}

// Vérifier plusieurs clés
produits := []string{"pommes", "poires", "oranges"}
for _, produit := range produits {
    if quantite, ok := stock[produit]; ok {
        fmt.Printf("%s: %d en stock\n", produit, quantite)
    } else {
        fmt.Printf("%s: produit non disponible\n", produit)
    }
}
```

## Maps de Maps (Maps imbriquées)

```go
// Représenter une base de données d'étudiants
etudiants := map[string]map[string]interface{}{
    "ET001": {
        "nom":    "Alice Dupont",
        "age":    20,
        "note":   15.5,
        "actif":  true,
    },
    "ET002": {
        "nom":    "Bob Martin",
        "age":    22,
        "note":   14.0,
        "actif":  false,
    },
}

// Accéder aux données imbriquées
if etudiant, existe := etudiants["ET001"]; existe {
    nom := etudiant["nom"].(string)  // Assertion de type
    age := etudiant["age"].(int)
    fmt.Printf("Étudiant: %s, âge: %d\n", nom, age)
}
```

## Maps avec des Slices comme valeurs

```go
// Grouper des éléments par catégorie
categories := map[string][]string{
    "fruits":   {"pomme", "banane", "orange"},
    "légumes":  {"carotte", "brocoli", "épinard"},
    "viandes":  {"bœuf", "porc", "agneau"},
}

// Ajouter un élément à une catégorie
categories["fruits"] = append(categories["fruits"], "kiwi")

// Afficher toutes les catégories
for categorie, items := range categories {
    fmt.Printf("%s: %v\n", categorie, items)
}
```

## Copier des Maps

### Attention : Maps sont des références

```go
original := map[string]int{"a": 1, "b": 2}

// Ceci ne fait PAS une copie !
mauvaiseCopie := original
mauvaiseCopie["a"] = 99

fmt.Println("Original:", original)       // map[a:99 b:2] - modifié !
fmt.Println("Copie:", mauvaiseCopie)    // map[a:99 b:2]
```

### Copie correcte

```go
original := map[string]int{"a": 1, "b": 2, "c": 3}

// Créer une vraie copie
copie := make(map[string]int)
for cle, valeur := range original {
    copie[cle] = valeur
}

// Modifier la copie n'affecte pas l'original
copie["a"] = 99
fmt.Println("Original:", original) // map[a:1 b:2 c:3] - intact
fmt.Println("Copie:", copie)      // map[a:99 b:2 c:3]
```

## Exemples pratiques

### Exemple 1 : Compteur de mots

```go
func compterMots(texte string) map[string]int {
    // Diviser le texte en mots (simplifié)
    mots := strings.Fields(strings.ToLower(texte))

    compteur := make(map[string]int)

    for _, mot := range mots {
        compteur[mot]++
    }

    return compteur
}

func main() {
    phrase := "Go est un langage. Go est rapide. Un langage rapide."
    resultat := compterMots(phrase)

    fmt.Println("Fréquence des mots:")
    for mot, count := range resultat {
        fmt.Printf("%s: %d\n", mot, count)
    }
}
```

### Exemple 2 : Annuaire téléphonique

```go
type Contact struct {
    Nom       string
    Telephone string
    Email     string
}

func main() {
    annuaire := map[string]Contact{
        "alice": {
            Nom:       "Alice Dupont",
            Telephone: "01-23-45-67-89",
            Email:     "alice@email.com",
        },
        "bob": {
            Nom:       "Bob Martin",
            Telephone: "01-98-76-54-32",
            Email:     "bob@email.com",
        },
    }

    // Rechercher un contact
    nom := "alice"
    if contact, trouve := annuaire[nom]; trouve {
        fmt.Printf("Contact trouvé:\n")
        fmt.Printf("Nom: %s\n", contact.Nom)
        fmt.Printf("Téléphone: %s\n", contact.Telephone)
        fmt.Printf("Email: %s\n", contact.Email)
    } else {
        fmt.Printf("Contact '%s' non trouvé\n", nom)
    }
}
```

### Exemple 3 : Cache simple

```go
type Cache struct {
    donnees map[string]string
}

func NouveauCache() *Cache {
    return &Cache{
        donnees: make(map[string]string),
    }
}

func (c *Cache) Get(cle string) (string, bool) {
    valeur, existe := c.donnees[cle]
    return valeur, existe
}

func (c *Cache) Set(cle, valeur string) {
    c.donnees[cle] = valeur
}

func (c *Cache) Delete(cle string) {
    delete(c.donnees, cle)
}

func (c *Cache) Taille() int {
    return len(c.donnees)
}

func main() {
    cache := NouveauCache()

    // Ajouter des données
    cache.Set("user:123", "Alice")
    cache.Set("user:456", "Bob")

    // Récupérer des données
    if nom, trouve := cache.Get("user:123"); trouve {
        fmt.Printf("Utilisateur trouvé: %s\n", nom)
    }

    fmt.Printf("Taille du cache: %d\n", cache.Taille())
}
```

## Performances et considérations

### Complexité temporelle

```go
// Toutes ces opérations sont en O(1) en moyenne :
m := make(map[string]int)

m["clé"] = 42        // Insertion: O(1)
valeur := m["clé"]   // Lecture: O(1)
delete(m, "clé")     // Suppression: O(1)
```

### Maps et concurrence

**⚠️ Attention :** Les maps ne sont **pas thread-safe** !

```go
// ❌ Dangereux en concurrence
var compteur = make(map[string]int)

func incrementer(cle string) {
    compteur[cle]++ // Race condition possible !
}

// ✅ Solution : utiliser sync.Map ou mutex
import "sync"

var (
    compteur = make(map[string]int)
    mutex    sync.RWMutex
)

func incrementerSecurise(cle string) {
    mutex.Lock()
    compteur[cle]++
    mutex.Unlock()
}

func lireSecurise(cle string) int {
    mutex.RLock()
    valeur := compteur[cle]
    mutex.RUnlock()
    return valeur
}
```

## Bonnes pratiques

### 1. Vérifiez toujours l'existence

```go
// ❌ Risqué
valeur := m["clé"] // Peut retourner zero value

// ✅ Sûr
if valeur, ok := m["clé"]; ok {
    // Utiliser valeur
}
```

### 2. Initialisez avant utilisation

```go
// ❌ Causera une panic
var m map[string]int
m["test"] = 42 // panic: assignment to entry in nil map

// ✅ Correct
var m = make(map[string]int)
// ou
m := map[string]int{}
```

### 3. Utilisez des clés cohérentes

```go
// ✅ Bon : clés cohérentes
utilisateurs := map[string]User{
    "user_123": user1,
    "user_456": user2,
}

// ❌ Évitez : mélange de formats
utilisateurs := map[string]User{
    "user_123": user1,
    "456":      user2, // Format incohérent
}
```

### 4. Considérez l'ordre si nécessaire

```go
// Si l'ordre est important, combinez map et slice
type OrdreDonnees struct {
    ordre  []string
    donnees map[string]interface{}
}

func (od *OrdreDonnees) Ajouter(cle string, valeur interface{}) {
    if _, existe := od.donnees[cle]; !existe {
        od.ordre = append(od.ordre, cle)
    }
    od.donnees[cle] = valeur
}
```

## Erreurs courantes

### 1. Modification pendant l'itération

```go
m := map[string]int{"a": 1, "b": 2, "c": 3}

// ❌ Dangereux : modifier pendant range
for cle := range m {
    if cle == "b" {
        delete(m, cle) // Peut causer des problèmes
    }
}

// ✅ Collectez d'abord, supprimez ensuite
var aSsupprimer []string
for cle := range m {
    if cle == "b" {
        aSupprimer = append(aSupprimer, cle)
    }
}
for _, cle := range aSupprimer {
    delete(m, cle)
}
```

### 2. Confusion entre zero value et absence

```go
scores := map[string]int{
    "Alice": 0, // Score de 0 (valide)
    "Bob":   5,
}

// ❌ Incorrect
if scores["Alice"] == 0 {
    fmt.Println("Alice n'a pas de score") // Faux !
}

// ✅ Correct
if score, existe := scores["Alice"]; existe {
    fmt.Printf("Score d'Alice: %d\n", score) // 0
} else {
    fmt.Println("Alice n'a pas de score")
}
```

## Exercices pratiques

### Exercice 1 : Basique

Créez un programme qui :

1. Crée une map pour stocker les notes d'étudiants (nom → note)
2. Ajoute 5 étudiants avec leurs notes
3. Calcule et affiche la moyenne des notes
4. Trouve et affiche l'étudiant avec la meilleure note

### Exercice 2 : Intermédiaire

Implémentez une fonction qui :

1. Prend un slice de strings en entrée
2. Retourne une map indiquant combien de fois chaque string apparaît
3. Affiche les résultats triés par fréquence (plus fréquent en premier)

### Exercice 3 : Avancé

Créez un système de gestion d'inventaire :

1. Map principale : `map[string]map[string]interface{}`
2. Chaque produit a : nom, prix, quantité, catégorie
3. Fonctions : ajouter produit, modifier stock, rechercher par catégorie
4. Fonction pour générer un rapport de stock faible (quantité < 5)

```go
package main

import (
 "fmt"
 "sort"
 "strings"
)

// ==========================================
// EXERCICE 1 : BASIQUE
// ==========================================

func exercice1() {
 fmt.Println("=== EXERCICE 1 : BASIQUE ===")

 // 1. Créer une map pour stocker les notes d'étudiants
 notes := make(map[string]float64)

 // 2. Ajouter 5 étudiants avec leurs notes
 notes["Alice"] = 16.5
 notes["Bob"] = 14.0
 notes["Charlie"] = 18.5
 notes["Diana"] = 12.5
 notes["Eve"] = 15.0

 fmt.Println("Notes des étudiants:")
 for nom, note := range notes {
  fmt.Printf("  %s: %.1f/20\n", nom, note)
 }

 // 3. Calculer et afficher la moyenne des notes
 var somme float64
 for _, note := range notes {
  somme += note
 }
 moyenne := somme / float64(len(notes))
 fmt.Printf("\nMoyenne de la classe: %.2f/20\n", moyenne)

 // 4. Trouver et afficher l'étudiant avec la meilleure note
 var meilleurEtudiant string
 var meilleureNote float64

 // Initialiser avec le premier étudiant
 premierTour := true
 for nom, note := range notes {
  if premierTour || note > meilleureNote {
   meilleurEtudiant = nom
   meilleureNote = note
   premierTour = false
  }
 }

 fmt.Printf("Meilleur étudiant: %s avec %.1f/20\n", meilleurEtudiant, meilleureNote)
 fmt.Println()
}

// ==========================================
// EXERCICE 2 : INTERMÉDIAIRE
// ==========================================

// Structure pour trier par fréquence
type FrequenceItem struct {
 Mot    string
 Compte int
}

// Fonction pour compter les occurrences
func compterOccurrences(mots []string) map[string]int {
 compteur := make(map[string]int)

 for _, mot := range mots {
  compteur[mot]++
 }

 return compteur
}

// Fonction pour trier et afficher par fréquence
func afficherParFrequence(compteur map[string]int) {
 // Convertir la map en slice pour pouvoir trier
 var items []FrequenceItem
 for mot, compte := range compteur {
  items = append(items, FrequenceItem{
   Mot:    mot,
   Compte: compte,
  })
 }

 // Trier par fréquence (ordre décroissant)
 sort.Slice(items, func(i, j int) bool {
  if items[i].Compte == items[j].Compte {
   // Si même fréquence, trier alphabétiquement
   return items[i].Mot < items[j].Mot
  }
  return items[i].Compte > items[j].Compte
 })

 // Afficher les résultats
 fmt.Println("Résultats triés par fréquence:")
 for i, item := range items {
  fmt.Printf("  %d. %s: %d occurrence(s)\n", i+1, item.Mot, item.Compte)
 }
}

func exercice2() {
 fmt.Println("=== EXERCICE 2 : INTERMÉDIAIRE ===")

 // Tests avec différents cas
 testCases := [][]string{
  {"pomme", "banane", "pomme", "orange", "banane", "pomme", "kiwi"},
  {"go", "est", "un", "langage", "go", "est", "rapide", "un", "langage", "rapide"},
  {"a", "b", "c", "a", "b", "a"},
  {"unique"},
  {}, // slice vide
 }

 for i, test := range testCases {
  fmt.Printf("\nTest %d: %v\n", i+1, test)

  if len(test) == 0 {
   fmt.Println("  Slice vide!")
   continue
  }

  // Compter les occurrences
  resultats := compterOccurrences(test)

  // Afficher les résultats bruts
  fmt.Println("  Map des occurrences:", resultats)

  // Afficher triés par fréquence
  afficherParFrequence(resultats)
 }

 // Test avec un texte réel
 fmt.Println("\n--- Test avec un texte réel ---")
 texte := "Go est un langage de programmation. Go est rapide et efficace. Un langage moderne."
 mots := strings.Fields(strings.ToLower(texte))

 // Nettoyer la ponctuation (version simplifiée)
 var motsNettoyes []string
 for _, mot := range mots {
  motPropre := strings.Trim(mot, ".,!?;:")
  if motPropre != "" {
   motsNettoyes = append(motsNettoyes, motPropre)
  }
 }

 fmt.Printf("Texte: %s\n", texte)
 fmt.Printf("Mots extraits: %v\n", motsNettoyes)

 resultats := compterOccurrences(motsNettoyes)
 afficherParFrequence(resultats)

 fmt.Println()
}

// ==========================================
// EXERCICE 3 : AVANCÉ
// ==========================================

// Système de gestion d'inventaire
type Inventaire struct {
 produits map[string]map[string]interface{}
}

// Créer un nouvel inventaire
func NouvelInventaire() *Inventaire {
 return &Inventaire{
  produits: make(map[string]map[string]interface{}),
 }
}

// Ajouter un produit
func (inv *Inventaire) AjouterProduit(id, nom string, prix float64, quantite int, categorie string) {
 inv.produits[id] = map[string]interface{}{
  "nom":       nom,
  "prix":      prix,
  "quantite":  quantite,
  "categorie": categorie,
 }
 fmt.Printf("Produit ajouté: %s (%s)\n", nom, id)
}

// Modifier le stock d'un produit
func (inv *Inventaire) ModifierStock(id string, nouvelleQuantite int) error {
 produit, existe := inv.produits[id]
 if !existe {
  return fmt.Errorf("produit %s non trouvé", id)
 }

 ancienneQuantite := produit["quantite"].(int)
 produit["quantite"] = nouvelleQuantite

 nom := produit["nom"].(string)
 fmt.Printf("Stock modifié pour %s (%s): %d → %d\n", nom, id, ancienneQuantite, nouvelleQuantite)
 return nil
}

// Rechercher par catégorie
func (inv *Inventaire) RechercherParCategorie(categorie string) []map[string]interface{} {
 var resultats []map[string]interface{}

 for id, produit := range inv.produits {
  if produit["categorie"].(string) == categorie {
   // Ajouter l'ID au produit pour le résultat
   produitAvecID := make(map[string]interface{})
   for k, v := range produit {
    produitAvecID[k] = v
   }
   produitAvecID["id"] = id
   resultats = append(resultats, produitAvecID)
  }
 }

 return resultats
}

// Générer un rapport de stock faible
func (inv *Inventaire) RapportStockFaible() []map[string]interface{} {
 var stockFaible []map[string]interface{}

 for id, produit := range inv.produits {
  quantite := produit["quantite"].(int)
  if quantite < 5 {
   produitAvecID := make(map[string]interface{})
   for k, v := range produit {
    produitAvecID[k] = v
   }
   produitAvecID["id"] = id
   stockFaible = append(stockFaible, produitAvecID)
  }
 }

 // Trier par quantité (plus faible en premier)
 sort.Slice(stockFaible, func(i, j int) bool {
  return stockFaible[i]["quantite"].(int) < stockFaible[j]["quantite"].(int)
 })

 return stockFaible
}

// Afficher tous les produits
func (inv *Inventaire) AfficherTous() {
 fmt.Println("\n=== INVENTAIRE COMPLET ===")
 if len(inv.produits) == 0 {
  fmt.Println("Aucun produit en inventaire.")
  return
 }

 for id, produit := range inv.produits {
  fmt.Printf("ID: %s\n", id)
  fmt.Printf("  Nom: %s\n", produit["nom"])
  fmt.Printf("  Prix: %.2f€\n", produit["prix"])
  fmt.Printf("  Quantité: %d\n", produit["quantite"])
  fmt.Printf("  Catégorie: %s\n", produit["categorie"])
  fmt.Println()
 }
}

// Obtenir des statistiques
func (inv *Inventaire) Statistiques() {
 if len(inv.produits) == 0 {
  fmt.Println("Aucun produit pour les statistiques.")
  return
 }

 // Compter par catégorie
 categories := make(map[string]int)
 var valeurTotale float64
 var quantiteTotale int

 for _, produit := range inv.produits {
  cat := produit["categorie"].(string)
  categories[cat]++

  prix := produit["prix"].(float64)
  quantite := produit["quantite"].(int)
  valeurTotale += prix * float64(quantite)
  quantiteTotale += quantite
 }

 fmt.Println("\n=== STATISTIQUES INVENTAIRE ===")
 fmt.Printf("Nombre total de produits: %d\n", len(inv.produits))
 fmt.Printf("Quantité totale en stock: %d\n", quantiteTotale)
 fmt.Printf("Valeur totale du stock: %.2f€\n", valeurTotale)

 fmt.Println("\nRépartition par catégorie:")
 for cat, count := range categories {
  fmt.Printf("  %s: %d produit(s)\n", cat, count)
 }
}

func exercice3() {
 fmt.Println("=== EXERCICE 3 : AVANCÉ ===")

 // Créer un inventaire
 inv := NouvelInventaire()

 // Ajouter des produits
 fmt.Println("\n--- Ajout de produits ---")
 inv.AjouterProduit("P001", "MacBook Pro", 2499.99, 3, "informatique")
 inv.AjouterProduit("P002", "iPhone 15", 1199.99, 12, "informatique")
 inv.AjouterProduit("P003", "AirPods", 179.99, 25, "informatique")
 inv.AjouterProduit("P004", "Chaise de bureau", 299.99, 2, "mobilier")
 inv.AjouterProduit("P005", "Bureau standing", 599.99, 1, "mobilier")
 inv.AjouterProduit("P006", "Lampe LED", 49.99, 8, "éclairage")
 inv.AjouterProduit("P007", "Câble USB-C", 19.99, 4, "accessoires")

 // Afficher l'inventaire complet
 inv.AfficherTous()

 // Modifier du stock
 fmt.Println("--- Modification de stock ---")
 inv.ModifierStock("P001", 1) // MacBook Pro: 3 → 1
 inv.ModifierStock("P007", 2) // Câble USB-C: 4 → 2

 if err := inv.ModifierStock("P999", 10); err != nil {
  fmt.Printf("Erreur: %v\n", err)
 }

 // Recherche par catégorie
 fmt.Println("\n--- Recherche par catégorie 'informatique' ---")
 produitsInfo := inv.RechercherParCategorie("informatique")
 for _, produit := range produitsInfo {
  fmt.Printf("- %s (%s): %d en stock à %.2f€\n",
   produit["nom"], produit["id"], produit["quantite"], produit["prix"])
 }

 fmt.Println("\n--- Recherche par catégorie 'mobilier' ---")
 produitsMobilier := inv.RechercherParCategorie("mobilier")
 for _, produit := range produitsMobilier {
  fmt.Printf("- %s (%s): %d en stock à %.2f€\n",
   produit["nom"], produit["id"], produit["quantite"], produit["prix"])
 }

 // Rapport de stock faible
 fmt.Println("\n--- Rapport de stock faible (< 5) ---")
 stockFaible := inv.RapportStockFaible()
 if len(stockFaible) == 0 {
  fmt.Println("Aucun produit en stock faible.")
 } else {
  fmt.Printf("⚠️  %d produit(s) en stock faible:\n", len(stockFaible))
  for i, produit := range stockFaible {
   fmt.Printf("%d. %s (%s): %d en stock\n",
    i+1, produit["nom"], produit["id"], produit["quantite"])
  }
 }

 // Afficher les statistiques
 inv.Statistiques()

 fmt.Println()
}

// ==========================================
// FONCTIONS UTILITAIRES BONUS
// ==========================================

// Fonction pour simuler un système de commande
func (inv *Inventaire) PasserCommande(id string, quantiteDemandee int) error {
 produit, existe := inv.produits[id]
 if !existe {
  return fmt.Errorf("produit %s non trouvé", id)
 }

 quantiteActuelle := produit["quantite"].(int)
 if quantiteActuelle < quantiteDemandee {
  return fmt.Errorf("stock insuffisant pour %s: %d demandé, %d disponible",
   produit["nom"], quantiteDemandee, quantiteActuelle)
 }

 nouvelleQuantite := quantiteActuelle - quantiteDemandee
 produit["quantite"] = nouvelleQuantite

 prix := produit["prix"].(float64)
 total := prix * float64(quantiteDemandee)

 fmt.Printf("✅ Commande validée:\n")
 fmt.Printf("   Produit: %s\n", produit["nom"])
 fmt.Printf("   Quantité: %d\n", quantiteDemandee)
 fmt.Printf("   Prix unitaire: %.2f€\n", prix)
 fmt.Printf("   Total: %.2f€\n", total)
 fmt.Printf("   Stock restant: %d\n", nouvelleQuantite)

 return nil
}

// Fonction pour sauvegarder l'inventaire (simulation)
func (inv *Inventaire) ExporterJSON() string {
 // En réalité, on utiliserait encoding/json
 // Ici on fait une version simplifiée pour la démonstration
 result := "{\n"
 result += "  \"produits\": {\n"

 i := 0
 for id, produit := range inv.produits {
  if i > 0 {
   result += ",\n"
  }
  result += fmt.Sprintf("    \"%s\": {\n", id)
  result += fmt.Sprintf("      \"nom\": \"%s\",\n", produit["nom"])
  result += fmt.Sprintf("      \"prix\": %.2f,\n", produit["prix"])
  result += fmt.Sprintf("      \"quantite\": %d,\n", produit["quantite"])
  result += fmt.Sprintf("      \"categorie\": \"%s\"\n", produit["categorie"])
  result += "    }"
  i++
 }

 result += "\n  }\n}"
 return result
}

func demonstrationAvancee() {
 fmt.Println("=== DÉMONSTRATION AVANCÉE ===")

 inv := NouvelInventaire()

 // Ajouter quelques produits
 inv.AjouterProduit("TEST001", "Produit Test", 99.99, 10, "test")

 // Simuler des commandes
 fmt.Println("\n--- Simulation de commandes ---")
 if err := inv.PasserCommande("TEST001", 3); err != nil {
  fmt.Printf("Erreur commande: %v\n", err)
 }

 if err := inv.PasserCommande("TEST001", 15); err != nil {
  fmt.Printf("Erreur commande: %v\n", err)
 }

 // Export JSON simulé
 fmt.Println("\n--- Export JSON (simulation) ---")
 jsonData := inv.ExporterJSON()
 fmt.Println(jsonData)
}

// ==========================================
// FONCTION PRINCIPALE
// ==========================================

func main() {
 fmt.Println("SOLUTIONS DES EXERCICES - MAPS")
 fmt.Println("===============================")

 exercice1()
 exercice2()
 exercice3()

 // Bonus: démonstration avancée
 demonstrationAvancee()
}
```

## Résumé

Les maps sont des structures de données essentielles en Go pour les associations clé-valeur :

**Points clés à retenir :**

- Les maps stockent des paires clé-valeur avec accès O(1)
- Elles doivent être initialisées avant utilisation
- L'ordre de parcours est aléatoire
- Toujours vérifier l'existence avec `valeur, ok := map[cle]`
- Elles ne sont pas thread-safe par défaut
- Utiliser `delete()` pour supprimer des éléments

**Cas d'usage typiques :**

- Caches et lookups rapides
- Compteurs et histogrammes
- Index et mappings
- Configuration et métadonnées

Les maps sont l'un des outils les plus puissants de Go pour organiser et accéder efficacement aux données.

⏭️
