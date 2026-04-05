🔝 Retour au [Sommaire](/SOMMAIRE.md)

# 5-4 : Pointers

## Introduction

Un **pointeur** (pointer) est une variable qui stocke l'**adresse mémoire** d'une autre variable plutôt que sa valeur directe. C'est comme avoir l'adresse d'une maison au lieu de la maison elle-même. Les pointeurs permettent d'accéder et de modifier des données de manière indirecte et efficace.

## Analogie simple

Imaginez que vous avez un carnet d'adresses :

- **Valeur directe** : "Alice habite dans une maison bleue avec un jardin"
- **Pointeur** : "Alice habite au 123 Rue de la Paix"

Le pointeur (l'adresse) vous permet de retrouver Alice, de lui rendre visite, et même de modifier quelque chose chez elle. C'est exactement ce que font les pointeurs en programmation !

## Concepts de base

### Variables et adresses mémoire

```go
func main() {
    age := 25

    fmt.Printf("Valeur de age: %d\n", age)           // 25
    fmt.Printf("Adresse de age: %p\n", &age)         // 0xc0000b4008 (exemple)
    fmt.Printf("Type de &age: %T\n", &age)           // *int
}
```

**Explications :**

- `age` contient la valeur `25`
- `&age` donne l'adresse mémoire où `age` est stocké
- `*int` signifie "pointeur vers un int"

### Déclaration de pointeurs

```go
// Déclaration d'un pointeur vers un int
var p *int

// p est nil (pointeur nul) car il ne pointe vers rien
fmt.Printf("Pointeur p: %v\n", p)        // <nil>
fmt.Printf("Type de p: %T\n", p)         // *int
```

### Opérateurs de pointeurs

Go utilise deux opérateurs principaux :

**Opérateur `&` (address-of)** : Obtient l'adresse d'une variable

```go
nombre := 42
pointeur := &nombre  // pointeur contient l'adresse de nombre
```

**Opérateur `*` (dereference)** : Accède à la valeur pointée

```go
valeur := *pointeur  // valeur = 42 (le contenu de l'adresse)
```

## Exemple complet de base

```go
func main() {
    // Créer une variable
    x := 10
    fmt.Printf("x = %d, adresse: %p\n", x, &x)

    // Créer un pointeur vers x
    var ptr *int = &x
    fmt.Printf("ptr pointe vers: %p\n", ptr)
    fmt.Printf("valeur pointée: %d\n", *ptr)

    // Modifier la valeur via le pointeur
    *ptr = 20
    fmt.Printf("Après modification via pointeur:\n")
    fmt.Printf("x = %d\n", x)              // 20 !
    fmt.Printf("*ptr = %d\n", *ptr)        // 20

    // x et *ptr pointent vers la même mémoire
    fmt.Printf("x et *ptr sont égaux: %t\n", x == *ptr)  // true
}
```

## Pointeurs avec différents types

### Pointeurs vers strings

```go
func main() {
    nom := "Alice"
    ptrNom := &nom

    fmt.Printf("nom: %s\n", nom)           // Alice
    fmt.Printf("*ptrNom: %s\n", *ptrNom)   // Alice

    // Modifier via le pointeur
    *ptrNom = "Bob"
    fmt.Printf("nom après modification: %s\n", nom)  // Bob
}
```

### Pointeurs vers structs

```go
type Personne struct {
    Nom string
    Age int
}

func main() {
    p := Personne{Nom: "Alice", Age: 25}
    ptrPersonne := &p

    // Accès aux champs via le pointeur
    fmt.Printf("Nom: %s\n", (*ptrPersonne).Nom)  // Syntaxe explicite
    fmt.Printf("Age: %d\n", ptrPersonne.Age)      // Syntaxe simplifiée (Go fait la déréférence automatiquement)

    // Modification via le pointeur
    ptrPersonne.Age = 26
    fmt.Printf("Nouvel âge: %d\n", p.Age)  // 26
}
```

**Important :** Go permet l'accès direct aux champs d'une struct via un pointeur sans déréférencement explicite.

## Création de pointeurs avec `new()`

```go
func main() {
    // new() alloue de la mémoire et retourne un pointeur
    ptr := new(int)

    fmt.Printf("ptr: %p\n", ptr)      // Adresse allouée
    fmt.Printf("*ptr: %d\n", *ptr)    // 0 (zero value)

    // Assigner une valeur
    *ptr = 42
    fmt.Printf("*ptr après assignation: %d\n", *ptr)  // 42
}
```

### Comparaison : `new()` vs `&`

```go
func main() {
    // Méthode 1 : avec new()
    ptr1 := new(int)
    *ptr1 = 10

    // Méthode 2 : avec &
    x := 10
    ptr2 := &x

    fmt.Printf("*ptr1: %d\n", *ptr1)  // 10
    fmt.Printf("*ptr2: %d\n", *ptr2)  // 10

    // Les deux méthodes aboutissent au même résultat
}
```

## Pointeurs et fonctions

### Passage par valeur vs passage par pointeur

```go
// Fonction qui reçoit une copie (passage par valeur)
func incrementerParValeur(x int) {
    x++  // Modifie seulement la copie
    fmt.Printf("Dans la fonction: x = %d\n", x)
}

// Fonction qui reçoit un pointeur (passage par référence)
func incrementerParPointeur(x *int) {
    *x++  // Modifie la variable originale
    fmt.Printf("Dans la fonction: *x = %d\n", *x)
}

func main() {
    nombre := 5

    fmt.Printf("Avant: nombre = %d\n", nombre)  // 5

    incrementerParValeur(nombre)
    fmt.Printf("Après passage par valeur: nombre = %d\n", nombre)  // 5 (pas changé)

    incrementerParPointeur(&nombre)
    fmt.Printf("Après passage par pointeur: nombre = %d\n", nombre)  // 6 (modifié)
}
```

### Fonctions retournant des pointeurs

```go
func creerEntier(valeur int) *int {
    x := valeur  // Variable locale
    return &x    // Retourne l'adresse (Go gère la mémoire automatiquement)
}

func main() {
    ptr := creerEntier(42)
    fmt.Printf("Valeur: %d\n", *ptr)  // 42
}
```

**Note importante :** En Go, il est sûr de retourner l'adresse d'une variable locale. Le garbage collector s'occupe de la gestion mémoire.

## Pointeurs et structs

### Méthodes avec receivers par pointeur

```go
type CompteBancaire struct {
    Titulaire string
    Solde     float64
}

// Méthode avec receiver par valeur (pas de modification)
func (c CompteBancaire) AfficherSolde() {
    fmt.Printf("Solde de %s: %.2f€\n", c.Titulaire, c.Solde)
}

// Méthode avec receiver par pointeur (peut modifier)
func (c *CompteBancaire) Deposer(montant float64) {
    c.Solde += montant
    fmt.Printf("Dépôt de %.2f€ effectué\n", montant)
}

func (c *CompteBancaire) Retirer(montant float64) bool {
    if c.Solde >= montant {
        c.Solde -= montant
        fmt.Printf("Retrait de %.2f€ effectué\n", montant)
        return true
    }
    fmt.Printf("Solde insuffisant pour retirer %.2f€\n", montant)
    return false
}

func main() {
    compte := CompteBancaire{
        Titulaire: "Alice",
        Solde:     1000.0,
    }

    compte.AfficherSolde()  // 1000.00€

    // Go convertit automatiquement entre valeur et pointeur
    compte.Deposer(500.0)   // Équivalent à (&compte).Deposer(500.0)
    compte.AfficherSolde()  // 1500.00€

    compte.Retirer(200.0)
    compte.AfficherSolde()  // 1300.00€
}
```

### Pointeurs vers structs avec constructeurs

```go
type Voiture struct {
    Marque string
    Modele string
    Annee  int
}

func NouvelleVoiture(marque, modele string, annee int) *Voiture {
    return &Voiture{
        Marque: marque,
        Modele: modele,
        Annee:  annee,
    }
}

func (v *Voiture) Demarrer() {
    fmt.Printf("La %s %s de %d démarre ! 🚗\n", v.Marque, v.Modele, v.Annee)
}

func main() {
    voiture := NouvelleVoiture("Toyota", "Corolla", 2023)
    voiture.Demarrer()
}
```

## Pointeurs et slices/maps

### Slices et pointeurs

```go
func main() {
    // Les slices contiennent déjà une référence
    nombres := []int{1, 2, 3, 4, 5}

    fmt.Printf("Avant: %v\n", nombres)
    modifierSlice(nombres)
    fmt.Printf("Après: %v\n", nombres)  // Modifié !
}

func modifierSlice(s []int) {
    if len(s) > 0 {
        s[0] = 99  // Modifie le slice original
    }
}
```

### Pointeurs vers slices (pour redimensionner)

```go
func ajouterElement(s *[]int, element int) {
    *s = append(*s, element)  // Modifie le slice original
}

func main() {
    nombres := []int{1, 2, 3}
    fmt.Printf("Avant: %v\n", nombres)

    ajouterElement(&nombres, 4)
    ajouterElement(&nombres, 5)

    fmt.Printf("Après: %v\n", nombres)  // [1 2 3 4 5]
}
```

## Pointeurs nil et vérifications

### Gestion des pointeurs nil

```go
func traiterPointeur(p *int) {
    if p == nil {
        fmt.Println("Pointeur nil reçu")
        return
    }

    fmt.Printf("Valeur: %d\n", *p)
}

func main() {
    var p *int  // p est nil
    traiterPointeur(p)  // "Pointeur nil reçu"

    x := 42
    p = &x
    traiterPointeur(p)  // "Valeur: 42"
}
```

### Initialisation sécurisée

```go
func obtenirPointeurSecurise(valeur int) *int {
    resultat := new(int)
    *resultat = valeur
    return resultat
}

func main() {
    ptr := obtenirPointeurSecurise(42)
    // ptr n'est jamais nil ici
    fmt.Printf("Valeur sûre: %d\n", *ptr)
}
```

## Exemples pratiques

### Exemple 1 : Échange de valeurs

```go
func echanger(a, b *int) {
    temp := *a
    *a = *b
    *b = temp
}

func main() {
    x, y := 10, 20
    fmt.Printf("Avant échange: x=%d, y=%d\n", x, y)

    echanger(&x, &y)
    fmt.Printf("Après échange: x=%d, y=%d\n", x, y)
}
```

### Exemple 2 : Liste chaînée simple

```go
type Noeud struct {
    Valeur int
    Suivant *Noeud  // Pointeur vers le noeud suivant
}

type ListeChainee struct {
    Tete *Noeud
}

func (l *ListeChainee) Ajouter(valeur int) {
    nouveauNoeud := &Noeud{
        Valeur: valeur,
        Suivant: l.Tete,
    }
    l.Tete = nouveauNoeud
}

func (l *ListeChainee) Afficher() {
    courant := l.Tete
    fmt.Print("Liste: ")

    for courant != nil {
        fmt.Printf("%d ", courant.Valeur)
        courant = courant.Suivant
    }
    fmt.Println()
}

func (l *ListeChainee) Taille() int {
    count := 0
    courant := l.Tete

    for courant != nil {
        count++
        courant = courant.Suivant
    }

    return count
}

func main() {
    liste := &ListeChainee{}

    liste.Ajouter(3)
    liste.Ajouter(2)
    liste.Ajouter(1)

    liste.Afficher()  // Liste: 1 2 3
    fmt.Printf("Taille: %d\n", liste.Taille())
}
```

### Exemple 3 : Cache avec pointeurs

```go
type Cache struct {
    donnees map[string]*string  // Map de pointeurs vers strings
}

func NouveauCache() *Cache {
    return &Cache{
        donnees: make(map[string]*string),
    }
}

func (c *Cache) Set(cle, valeur string) {
    // Stocke un pointeur vers une copie de la valeur
    copieValeur := valeur
    c.donnees[cle] = &copieValeur
}

func (c *Cache) Get(cle string) (string, bool) {
    ptr, existe := c.donnees[cle]
    if !existe || ptr == nil {
        return "", false
    }
    return *ptr, true
}

func (c *Cache) Update(cle, nouvelleValeur string) bool {
    ptr, existe := c.donnees[cle]
    if !existe || ptr == nil {
        return false
    }

    *ptr = nouvelleValeur  // Modifie directement la valeur pointée
    return true
}

func (c *Cache) Delete(cle string) {
    delete(c.donnees, cle)
}

func main() {
    cache := NouveauCache()

    cache.Set("user:123", "Alice")
    cache.Set("user:456", "Bob")

    if valeur, ok := cache.Get("user:123"); ok {
        fmt.Printf("Trouvé: %s\n", valeur)
    }

    cache.Update("user:123", "Alice Dupont")

    if valeur, ok := cache.Get("user:123"); ok {
        fmt.Printf("Après mise à jour: %s\n", valeur)
    }
}
```

### Exemple 4 : Arbre binaire

```go
type NoeudArbre struct {
    Valeur int
    Gauche *NoeudArbre
    Droite *NoeudArbre
}

type ArbreBinaire struct {
    Racine *NoeudArbre
}

func (a *ArbreBinaire) Inserer(valeur int) {
    a.Racine = insererRecursif(a.Racine, valeur)
}

func insererRecursif(noeud *NoeudArbre, valeur int) *NoeudArbre {
    // Si le noeud est nil, créer un nouveau noeud
    if noeud == nil {
        return &NoeudArbre{Valeur: valeur}
    }

    // Insérer à gauche ou à droite selon la valeur
    if valeur < noeud.Valeur {
        noeud.Gauche = insererRecursif(noeud.Gauche, valeur)
    } else {
        noeud.Droite = insererRecursif(noeud.Droite, valeur)
    }

    return noeud
}

func (a *ArbreBinaire) ParcoursInfixe() {
    fmt.Print("Parcours infixe: ")
    parcoursInfixeRecursif(a.Racine)
    fmt.Println()
}

func parcoursInfixeRecursif(noeud *NoeudArbre) {
    if noeud != nil {
        parcoursInfixeRecursif(noeud.Gauche)
        fmt.Printf("%d ", noeud.Valeur)
        parcoursInfixeRecursif(noeud.Droite)
    }
}

func (a *ArbreBinaire) Rechercher(valeur int) bool {
    return rechercherRecursif(a.Racine, valeur)
}

func rechercherRecursif(noeud *NoeudArbre, valeur int) bool {
    if noeud == nil {
        return false
    }

    if valeur == noeud.Valeur {
        return true
    }

    if valeur < noeud.Valeur {
        return rechercherRecursif(noeud.Gauche, valeur)
    }

    return rechercherRecursif(noeud.Droite, valeur)
}

func main() {
    arbre := &ArbreBinaire{}

    // Insérer des valeurs
    valeurs := []int{5, 3, 7, 2, 4, 6, 8}
    for _, v := range valeurs {
        arbre.Inserer(v)
    }

    arbre.ParcoursInfixe()  // Parcours infixe: 2 3 4 5 6 7 8

    // Rechercher des valeurs
    fmt.Printf("Recherche 4: %t\n", arbre.Rechercher(4))   // true
    fmt.Printf("Recherche 9: %t\n", arbre.Rechercher(9))   // false
}
```

## Comparaison : Valeurs vs Pointeurs

### Quand utiliser des valeurs

```go
type Point struct {
    X, Y float64
}

// Utiliser des valeurs pour :
// - Petites structs
// - Données immutables
// - Calculs mathématiques

func (p Point) Distance() float64 {
    return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func (p Point) Ajouter(autre Point) Point {
    return Point{X: p.X + autre.X, Y: p.Y + autre.Y}
}
```

### Quand utiliser des pointeurs

```go
type BaseDeDonnees struct {
    Connexions map[string]*sql.DB
    Config     map[string]interface{}
}

// Utiliser des pointeurs pour :
// - Grosses structs
// - Données mutables
// - Partage de données
// - Optimisation mémoire

func (db *BaseDeDonnees) AjouterConnexion(nom string, conn *sql.DB) {
    db.Connexions[nom] = conn
}

func (db *BaseDeDonnees) ModifierConfig(cle string, valeur interface{}) {
    db.Config[cle] = valeur
}
```

## Performance et considérations

### Impact mémoire

```go
type PetiteStruct struct {
    A int
    B int
}

type GrosseStruct struct {
    Donnees [1000]int
    Nom     string
    Actif   bool
}

func traiterPetite(p PetiteStruct) {
    // Copie de 16 bytes - acceptable
}

func traiterPetitePtr(p *PetiteStruct) {
    // Copie de 8 bytes (pointeur) - légèrement plus d'indirection
}

func traiterGrosse(g GrosseStruct) {
    // Copie de ~8KB - inefficace !
}

func traiterGrossePtr(g *GrosseStruct) {
    // Copie de 8 bytes (pointeur) - efficace !
}
```

### Recommandations

```go
// ✅ Bon : petites structs par valeur
func calculerDistance(p1, p2 Point) float64 {
    // ...
}

// ✅ Bon : grosses structs par pointeur
func traiterDocument(doc *Document) error {
    // ...
}

// ✅ Bon : modification nécessaire = pointeur
func (c *CompteBancaire) Crediter(montant float64) {
    c.Solde += montant
}

// ✅ Bon : lecture seule de petite struct = valeur
func (p Point) String() string {
    return fmt.Sprintf("(%f, %f)", p.X, p.Y)
}
```

## Erreurs courantes

### 1. Déréférencement de pointeur nil

```go
// ❌ Dangereux
func mauvais() {
    var p *int
    fmt.Println(*p)  // Panic: runtime error
}

// ✅ Sûr
func bon() {
    var p *int
    if p != nil {
        fmt.Println(*p)
    } else {
        fmt.Println("Pointeur nil")
    }
}
```

### 2. Confusion entre valeur et pointeur

```go
type Personne struct {
    Nom string
}

// ❌ Ne modifie pas l'original
func changerNomMauvais(p Personne, nouveauNom string) {
    p.Nom = nouveauNom  // Modifie seulement la copie
}

// ✅ Modifie l'original
func changerNomBon(p *Personne, nouveauNom string) {
    p.Nom = nouveauNom  // Modifie l'original via le pointeur
}
```

### 3. Fuite de pointeurs locaux (attention)

```go
// ✅ Correct en Go (garbage collector gère)
func creerPersonne() *Personne {
    p := Personne{Nom: "Alice"}
    return &p  // Go rend ceci sûr
}

// ❌ Dans d'autres langages, ceci serait dangereux
// En Go, c'est parfaitement sûr grâce au GC
```

## Bonnes pratiques

### 1. Soyez cohérent avec les receivers

```go
type Utilisateur struct {
    ID   int
    Nom  string
    Age  int
}

// Si une méthode a besoin d'un pointeur receiver,
// utilisez des pointeurs pour toutes les méthodes

func (u *Utilisateur) ModifierAge(nouvelAge int) {
    u.Age = nouvelAge
}

func (u *Utilisateur) AfficherInfo() {  // Cohérent : utilise aussi un pointeur
    fmt.Printf("Utilisateur: %s, %d ans\n", u.Nom, u.Age)
}
```

### 2. Validez les pointeurs

```go
func traiterUtilisateur(u *Utilisateur) error {
    if u == nil {
        return fmt.Errorf("utilisateur ne peut pas être nil")
    }

    // Traiter l'utilisateur...
    return nil
}
```

### 3. Utilisez des constructeurs pour l'initialisation

```go
func NouvelUtilisateur(nom string, age int) *Utilisateur {
    return &Utilisateur{
        ID:  genererID(),
        Nom: nom,
        Age: age,
    }
}
```

## Exercices pratiques

### Exercice 1 : Basique

Créez des fonctions qui :

1. Échangent les valeurs de deux variables entières
2. Trouvent le minimum et maximum dans un slice et retournent des pointeurs vers ces valeurs
3. Incrémentent un compteur via un pointeur

### Exercice 2 : Intermédiaire

Implémentez une stack (pile) avec :

1. Structure utilisant des pointeurs pour les noeuds
2. Méthodes `Push()`, `Pop()`, `Peek()`, `IsEmpty()`
3. Gestion sécurisée des pointeurs nil

### Exercice 3 : Avancé

Créez un système de gestion de mémoire simple :

1. Pool d'objets réutilisables avec pointeurs
2. Fonction d'allocation et de libération
3. Compteur de références pour éviter les fuites

```go
package main

import (
 "fmt"
 "sync"
)

// ==========================================
// EXERCICE 1 : BASIQUE
// ==========================================

// 1. Fonction pour échanger deux variables entières
func echanger(a, b *int) {
 if a == nil || b == nil {
  fmt.Println("⚠️  Pointeurs nil détectés dans echanger()")
  return
 }

 temp := *a
 *a = *b
 *b = temp

 fmt.Printf("Échange effectué: a=%d, b=%d\n", *a, *b)
}

// 2. Fonction pour trouver min et max avec pointeurs vers les valeurs originales
func trouverMinMax(slice []int) (*int, *int) {
 if len(slice) == 0 {
  return nil, nil
 }

 // Initialiser avec le premier élément
 minPtr := &slice[0]
 maxPtr := &slice[0]

 // Parcourir le reste du slice
 for i := 1; i < len(slice); i++ {
  if slice[i] < *minPtr {
   minPtr = &slice[i]
  }
  if slice[i] > *maxPtr {
   maxPtr = &slice[i]
  }
 }

 return minPtr, maxPtr
}

// Alternative: retourner des pointeurs vers des copies
func trouverMinMaxCopie(slice []int) (*int, *int) {
 if len(slice) == 0 {
  return nil, nil
 }

 min := slice[0]
 max := slice[0]

 for _, valeur := range slice[1:] {
  if valeur < min {
   min = valeur
  }
  if valeur > max {
   max = valeur
  }
 }

 // Retourner des pointeurs vers des copies
 return &min, &max
}

// 3. Fonction pour incrémenter un compteur
func incrementerCompteur(compteur *int) {
 if compteur == nil {
  fmt.Println("⚠️  Compteur nil dans incrementerCompteur()")
  return
 }

 (*compteur)++
 fmt.Printf("Compteur incrémenté: %d\n", *compteur)
}

// Fonction bonus: incrémenter de n
func incrementerCompteurDe(compteur *int, increment int) {
 if compteur == nil {
  fmt.Println("⚠️  Compteur nil dans incrementerCompteurDe()")
  return
 }

 *compteur += increment
 fmt.Printf("Compteur incrémenté de %d: nouvelle valeur = %d\n", increment, *compteur)
}

func exercice1() {
 fmt.Println("=== EXERCICE 1 : BASIQUE ===")

 // Test 1: Échange de valeurs
 fmt.Println("\n--- Test 1: Échange de valeurs ---")
 x, y := 10, 20
 fmt.Printf("Avant échange: x=%d, y=%d\n", x, y)
 echanger(&x, &y)
 fmt.Printf("Après échange: x=%d, y=%d\n", x, y)

 // Test avec pointeurs nil
 var nilPtr *int
 echanger(&x, nilPtr) // Doit afficher un avertissement

 // Test 2: Trouver min/max
 fmt.Println("\n--- Test 2: Min/Max avec pointeurs ---")
 nombres := []int{5, 2, 8, 1, 9, 3, 7}
 fmt.Printf("Slice: %v\n", nombres)

 minPtr, maxPtr := trouverMinMax(nombres)
 if minPtr != nil && maxPtr != nil {
  fmt.Printf("Min: %d (adresse: %p)\n", *minPtr, minPtr)
  fmt.Printf("Max: %d (adresse: %p)\n", *maxPtr, maxPtr)

  // Démontrer que ce sont des pointeurs vers les éléments originaux
  fmt.Println("Modification via les pointeurs:")
  *minPtr = 0
  *maxPtr = 100
  fmt.Printf("Slice après modification: %v\n", nombres)
 }

 // Test avec slice vide
 fmt.Println("\nTest avec slice vide:")
 vide := []int{}
 minPtr, maxPtr = trouverMinMax(vide)
 if minPtr == nil && maxPtr == nil {
  fmt.Println("✅ Slice vide géré correctement (pointeurs nil)")
 }

 // Test avec version copie
 fmt.Println("\n--- Test avec version copie ---")
 nombres2 := []int{15, 3, 9, 1, 12}
 fmt.Printf("Slice: %v\n", nombres2)
 minCopie, maxCopie := trouverMinMaxCopie(nombres2)
 if minCopie != nil && maxCopie != nil {
  fmt.Printf("Min (copie): %d\n", *minCopie)
  fmt.Printf("Max (copie): %d\n", *maxCopie)

  // Modifier les copies n'affecte pas le slice original
  *minCopie = -1
  *maxCopie = 999
  fmt.Printf("Slice original inchangé: %v\n", nombres2)
 }

 // Test 3: Incrémenter compteur
 fmt.Println("\n--- Test 3: Incrément de compteur ---")
 compteur := 0
 fmt.Printf("Compteur initial: %d\n", compteur)

 incrementerCompteur(&compteur)
 incrementerCompteur(&compteur)
 incrementerCompteurDe(&compteur, 5)

 fmt.Printf("Compteur final: %d\n", compteur)

 // Test avec pointeur nil
 incrementerCompteur(nil) // Doit afficher un avertissement

 fmt.Println()
}

// ==========================================
// EXERCICE 2 : INTERMÉDIAIRE - STACK
// ==========================================

// Noeud de la stack
type NoeudStack struct {
 Valeur   interface{}
 Precedent *NoeudStack
}

// Structure Stack
type Stack struct {
 sommet *NoeudStack
 taille int
}

// Créer une nouvelle stack
func NouvelleStack() *Stack {
 return &Stack{
  sommet: nil,
  taille: 0,
 }
}

// Push: ajouter un élément au sommet
func (s *Stack) Push(valeur interface{}) {
 nouveauNoeud := &NoeudStack{
  Valeur:    valeur,
  Precedent: s.sommet,
 }
 s.sommet = nouveauNoeud
 s.taille++

 fmt.Printf("📥 Push: %v (taille: %d)\n", valeur, s.taille)
}

// Pop: retirer et retourner l'élément du sommet
func (s *Stack) Pop() (interface{}, bool) {
 if s.IsEmpty() {
  fmt.Println("⚠️  Tentative de Pop sur stack vide")
  return nil, false
 }

 valeur := s.sommet.Valeur
 s.sommet = s.sommet.Precedent
 s.taille--

 fmt.Printf("📤 Pop: %v (taille: %d)\n", valeur, s.taille)
 return valeur, true
}

// Peek: voir l'élément du sommet sans le retirer
func (s *Stack) Peek() (interface{}, bool) {
 if s.IsEmpty() {
  return nil, false
 }

 return s.sommet.Valeur, true
}

// IsEmpty: vérifier si la stack est vide
func (s *Stack) IsEmpty() bool {
 return s.sommet == nil
}

// Taille: retourner le nombre d'éléments
func (s *Stack) Taille() int {
 return s.taille
}

// Afficher la stack (pour debug)
func (s *Stack) Afficher() {
 if s.IsEmpty() {
  fmt.Println("Stack vide")
  return
 }

 fmt.Print("Stack (sommet → base): ")
 courant := s.sommet
 for courant != nil {
  fmt.Printf("%v ", courant.Valeur)
  courant = courant.Precedent
 }
 fmt.Printf("(taille: %d)\n", s.taille)
}

// Vider la stack
func (s *Stack) Vider() {
 for !s.IsEmpty() {
  s.Pop()
 }
 fmt.Println("🗑️  Stack vidée")
}

// Clone: créer une copie de la stack
func (s *Stack) Clone() *Stack {
 if s.IsEmpty() {
  return NouvelleStack()
 }

 // Utiliser une stack temporaire pour inverser l'ordre
 temp := NouvelleStack()
 courant := s.sommet

 // Copier dans l'ordre inverse
 for courant != nil {
  temp.Push(courant.Valeur)
  courant = courant.Precedent
 }

 // Créer la stack finale dans le bon ordre
 clone := NouvelleStack()
 for !temp.IsEmpty() {
  valeur, _ := temp.Pop()
  clone.Push(valeur)
 }

 return clone
}

func exercice2() {
 fmt.Println("=== EXERCICE 2 : INTERMÉDIAIRE - STACK ===")

 // Créer une nouvelle stack
 stack := NouvelleStack()

 // Test IsEmpty sur stack vide
 fmt.Printf("Stack vide? %t\n", stack.IsEmpty())

 // Test Peek sur stack vide
 if valeur, ok := stack.Peek(); ok {
  fmt.Printf("Sommet: %v\n", valeur)
 } else {
  fmt.Println("✅ Peek sur stack vide géré correctement")
 }

 // Test Push
 fmt.Println("\n--- Test Push ---")
 stack.Push(10)
 stack.Push("Hello")
 stack.Push(3.14)
 stack.Push(true)

 stack.Afficher()

 // Test Peek
 fmt.Println("\n--- Test Peek ---")
 if valeur, ok := stack.Peek(); ok {
  fmt.Printf("👀 Sommet (sans retirer): %v\n", valeur)
 }
 stack.Afficher() // Vérifier que rien n'a changé

 // Test Pop
 fmt.Println("\n--- Test Pop ---")
 for i := 0; i < 3; i++ {
  if valeur, ok := stack.Pop(); ok {
   fmt.Printf("Valeur récupérée: %v\n", valeur)
  }
  stack.Afficher()
 }

 // Ajouter plus d'éléments
 fmt.Println("\n--- Remplissage à nouveau ---")
 for i := 1; i <= 5; i++ {
  stack.Push(i * 10)
 }
 stack.Afficher()

 // Test Clone
 fmt.Println("\n--- Test Clone ---")
 clone := stack.Clone()
 fmt.Print("Original: ")
 stack.Afficher()
 fmt.Print("Clone: ")
 clone.Afficher()

 // Modifier l'original pour prouver l'indépendance
 stack.Push(999)
 fmt.Println("Après ajout de 999 à l'original:")
 fmt.Print("Original: ")
 stack.Afficher()
 fmt.Print("Clone (inchangé): ")
 clone.Afficher()

 // Vider la stack
 fmt.Println("\n--- Test Vider ---")
 stack.Vider()
 fmt.Printf("Stack vide après vidage? %t\n", stack.IsEmpty())

 // Test Pop sur stack vide
 fmt.Println("\n--- Test Pop sur stack vide ---")
 stack.Pop() // Doit afficher un avertissement

 fmt.Println()
}

// ==========================================
// EXERCICE 3 : AVANCÉ - POOL D'OBJETS
// ==========================================

// Objet géré par le pool
type ObjetPool struct {
 ID       int
 Donnees  []byte
 EnUse    bool
 RefCount int
}

// Interface pour les objets poolés
type Poolable interface {
 Reset()           // Remet l'objet à zéro
 GetID() int       // Retourne l'ID unique
 IsInUse() bool    // Vérifie si l'objet est utilisé
}

// Implémenter l'interface Poolable
func (obj *ObjetPool) Reset() {
 obj.Donnees = obj.Donnees[:0] // Vider le slice sans réallouer
 obj.EnUse = false
 obj.RefCount = 0
}

func (obj *ObjetPool) GetID() int {
 return obj.ID
}

func (obj *ObjetPool) IsInUse() bool {
 return obj.EnUse
}

// Pool de gestion de mémoire
type MemoryPool struct {
 objets        []*ObjetPool
 libres        []*ObjetPool
 utilises      map[int]*ObjetPool
 prochainID    int
 tailleMax     int
 mutex         sync.RWMutex
 statistiques  PoolStats
}

// Statistiques du pool
type PoolStats struct {
 TotalCrees    int
 TotalLiberes  int
 TotalAlloues  int
 PicUtilises   int
 ReutilisationRate float64
}

// Créer un nouveau pool
func NouveauMemoryPool(tailleMax int) *MemoryPool {
 return &MemoryPool{
  objets:       make([]*ObjetPool, 0, tailleMax),
  libres:       make([]*ObjetPool, 0, tailleMax),
  utilises:     make(map[int]*ObjetPool),
  prochainID:   1,
  tailleMax:    tailleMax,
  statistiques: PoolStats{},
 }
}

// Allouer un objet du pool
func (mp *MemoryPool) Allouer() *ObjetPool {
 mp.mutex.Lock()
 defer mp.mutex.Unlock()

 var obj *ObjetPool

 // Essayer de réutiliser un objet libre
 if len(mp.libres) > 0 {
  obj = mp.libres[len(mp.libres)-1]
  mp.libres = mp.libres[:len(mp.libres)-1]
  obj.Reset()
  fmt.Printf("♻️  Réutilisation objet ID %d\n", obj.ID)
 } else if len(mp.objets) < mp.tailleMax {
  // Créer un nouvel objet
  obj = &ObjetPool{
   ID:      mp.prochainID,
   Donnees: make([]byte, 0, 1024), // Capacité initiale
  }
  mp.prochainID++
  mp.objets = append(mp.objets, obj)
  mp.statistiques.TotalCrees++
  fmt.Printf("🆕 Création nouvel objet ID %d\n", obj.ID)
 } else {
  fmt.Printf("❌ Pool plein! Impossible d'allouer (max: %d)\n", mp.tailleMax)
  return nil
 }

 // Marquer comme utilisé
 obj.EnUse = true
 obj.RefCount = 1
 mp.utilises[obj.ID] = obj
 mp.statistiques.TotalAlloues++

 // Mettre à jour le pic d'utilisation
 if len(mp.utilises) > mp.statistiques.PicUtilises {
  mp.statistiques.PicUtilises = len(mp.utilises)
 }

 return obj
}

// Libérer un objet (décrémenter le compteur de références)
func (mp *MemoryPool) Liberer(obj *ObjetPool) bool {
 if obj == nil {
  fmt.Println("⚠️  Tentative de libération d'un objet nil")
  return false
 }

 mp.mutex.Lock()
 defer mp.mutex.Unlock()

 // Vérifier que l'objet est bien dans le pool
 if _, existe := mp.utilises[obj.ID]; !existe {
  fmt.Printf("⚠️  Objet ID %d non trouvé dans les objets utilisés\n", obj.ID)
  return false
 }

 // Décrémenter le compteur de références
 obj.RefCount--
 fmt.Printf("📉 Objet ID %d: RefCount = %d\n", obj.ID, obj.RefCount)

 // Si plus de références, libérer réellement
 if obj.RefCount <= 0 {
  delete(mp.utilises, obj.ID)
  obj.EnUse = false
  mp.libres = append(mp.libres, obj)
  mp.statistiques.TotalLiberes++
  fmt.Printf("🔓 Objet ID %d libéré et retourné au pool\n", obj.ID)
  return true
 }

 return false
}

// Incrémenter le compteur de références
func (mp *MemoryPool) AjouterReference(obj *ObjetPool) bool {
 if obj == nil {
  return false
 }

 mp.mutex.Lock()
 defer mp.mutex.Unlock()

 if _, existe := mp.utilises[obj.ID]; existe {
  obj.RefCount++
  fmt.Printf("📈 Objet ID %d: RefCount = %d\n", obj.ID, obj.RefCount)
  return true
 }

 return false
}

// Forcer la libération d'un objet (ignorer le RefCount)
func (mp *MemoryPool) ForcerLiberation(obj *ObjetPool) bool {
 if obj == nil {
  return false
 }

 mp.mutex.Lock()
 defer mp.mutex.Unlock()

 if _, existe := mp.utilises[obj.ID]; existe {
  delete(mp.utilises, obj.ID)
  obj.EnUse = false
  obj.RefCount = 0
  mp.libres = append(mp.libres, obj)
  mp.statistiques.TotalLiberes++
  fmt.Printf("🔨 Objet ID %d libéré de force\n", obj.ID)
  return true
 }

 return false
}

// Obtenir les statistiques du pool
func (mp *MemoryPool) Statistiques() PoolStats {
 mp.mutex.RLock()
 defer mp.mutex.RUnlock()

 stats := mp.statistiques
 stats.ReutilisationRate = 0
 if stats.TotalAlloues > 0 {
  stats.ReutilisationRate = float64(stats.TotalAlloues-stats.TotalCrees) / float64(stats.TotalAlloues) * 100
 }

 return stats
}

// Afficher l'état du pool
func (mp *MemoryPool) AfficherEtat() {
 mp.mutex.RLock()
 defer mp.mutex.RUnlock()

 fmt.Printf("\n=== ÉTAT DU POOL ===\n")
 fmt.Printf("Objets totaux créés: %d\n", len(mp.objets))
 fmt.Printf("Objets libres: %d\n", len(mp.libres))
 fmt.Printf("Objets utilisés: %d\n", len(mp.utilises))
 fmt.Printf("Capacité maximale: %d\n", mp.tailleMax)

 if len(mp.utilises) > 0 {
  fmt.Printf("Objets en cours d'utilisation:\n")
  for id, obj := range mp.utilises {
   fmt.Printf("  - ID %d (RefCount: %d)\n", id, obj.RefCount)
  }
 }

 stats := mp.Statistiques()
 fmt.Printf("\nStatistiques:\n")
 fmt.Printf("  Total créés: %d\n", stats.TotalCrees)
 fmt.Printf("  Total alloués: %d\n", stats.TotalAlloues)
 fmt.Printf("  Total libérés: %d\n", stats.TotalLiberes)
 fmt.Printf("  Pic d'utilisation: %d\n", stats.PicUtilises)
 fmt.Printf("  Taux de réutilisation: %.1f%%\n", stats.ReutilisationRate)
}

// Nettoyer le pool (libérer tous les objets)
func (mp *MemoryPool) Nettoyer() {
 mp.mutex.Lock()
 defer mp.mutex.Unlock()

 // Forcer la libération de tous les objets utilisés
 for id, obj := range mp.utilises {
  obj.Reset()
  mp.libres = append(mp.libres, obj)
  delete(mp.utilises, id)
 }

 fmt.Printf("🧹 Pool nettoyé: %d objets retournés\n", len(mp.libres))
}

func exercice3() {
 fmt.Println("=== EXERCICE 3 : AVANCÉ - POOL D'OBJETS ===")

 // Créer un pool avec une capacité maximale de 5 objets
 pool := NouveauMemoryPool(5)

 fmt.Println("\n--- Test d'allocation de base ---")

 // Allouer quelques objets
 obj1 := pool.Allouer()
 obj2 := pool.Allouer()
 obj3 := pool.Allouer()

 pool.AfficherEtat()

 // Test du compteur de références
 fmt.Println("\n--- Test compteur de références ---")
 if obj1 != nil {
  pool.AjouterReference(obj1) // RefCount = 2
  pool.AjouterReference(obj1) // RefCount = 3

  // Tentatives de libération
  pool.Liberer(obj1) // RefCount = 2
  pool.Liberer(obj1) // RefCount = 1
  pool.Liberer(obj1) // RefCount = 0, objet libéré
 }

 pool.AfficherEtat()

 // Test de réutilisation
 fmt.Println("\n--- Test de réutilisation ---")
 obj4 := pool.Allouer() // Doit réutiliser obj1
 obj5 := pool.Allouer() // Nouvel objet
 obj6 := pool.Allouer() // Nouvel objet

 pool.AfficherEtat()

 // Atteindre la limite du pool
 fmt.Println("\n--- Test limite du pool ---")
 obj7 := pool.Allouer() // Dernier objet possible
 obj8 := pool.Allouer() // Doit échouer (pool plein)

 pool.AfficherEtat()

 // Libérer quelques objets pour faire de la place
 fmt.Println("\n--- Libération d'objets ---")
 if obj2 != nil {
  pool.Liberer(obj2)
 }
 if obj3 != nil {
  pool.Liberer(obj3)
 }

 // Maintenant on peut allouer à nouveau
 obj9 := pool.Allouer() // Doit réussir (réutilise obj2 ou obj3)

 pool.AfficherEtat()

 // Test de libération forcée
 fmt.Println("\n--- Test libération forcée ---")
 if obj4 != nil {
  pool.AjouterReference(obj4) // RefCount = 2
  fmt.Printf("Avant libération forcée - RefCount obj4: %d\n", obj4.RefCount)
  pool.ForcerLiberation(obj4) // Force la libération malgré RefCount > 1
 }

 pool.AfficherEtat()

 // Statistiques finales
 fmt.Println("\n--- Statistiques finales ---")
 stats := pool.Statistiques()
 fmt.Printf("📊 Statistiques du pool:\n")
 fmt.Printf("   Objets créés: %d\n", stats.TotalCrees)
 fmt.Printf("   Allocations totales: %d\n", stats.TotalAlloues)
 fmt.Printf("   Libérations totales: %d\n", stats.TotalLiberes)
 fmt.Printf("   Pic d'utilisation: %d objets\n", stats.PicUtilises)
 fmt.Printf("   Taux de réutilisation: %.1f%%\n", stats.ReutilisationRate)

 // Nettoyer le pool
 fmt.Println("\n--- Nettoyage final ---")
 pool.Nettoyer()
 pool.AfficherEtat()

 fmt.Println()
}

// ==========================================
// FONCTIONS BONUS ET DÉMONSTRATIONS
// ==========================================

// Démonstration des différences entre pointeurs et valeurs
func demonstrationPerformance() {
 fmt.Println("=== DÉMONSTRATION PERFORMANCE ===")

 // Structure pour tester
 type GrosseStructure struct {
  Donnees [1000]int
  Nom     string
 }

 // Création d'une grosse structure
 gs := GrosseStructure{Nom: "Test"}
 for i := 0; i < 1000; i++ {
  gs.Donnees[i] = i
 }

 // Fonction avec passage par valeur (copie)
 passerParValeur := func(gs GrosseStructure) {
  // Toute la structure est copiée (8KB environ)
  gs.Nom = "Modifié en copie"
 }

 // Fonction avec passage par pointeur (référence)
 passerParPointeur := func(gs *GrosseStructure) {
  // Seul le pointeur est passé (8 bytes)
  gs.Nom = "Modifié par pointeur"
 }

 fmt.Printf("Nom avant: %s\n", gs.Nom)

 passerParValeur(gs)
 fmt.Printf("Après passage par valeur: %s\n", gs.Nom) // Inchangé

 passerParPointeur(&gs)
 fmt.Printf("Après passage par pointeur: %s\n", gs.Nom) // Modifié

 fmt.Println("💡 Pour les grosses structures, les pointeurs sont plus efficaces!")
}

// Test de gestion mémoire avancée
func testGestionMemoire() {
 fmt.Println("\n=== TEST GESTION MÉMOIRE AVANCÉE ===")

 // Simuler des fuites de mémoire et leur prévention
 type Ressource struct {
  ID       int
  Donnees  *[]byte
  Callback func()
 }

 // Fonction pour créer une ressource avec cleanup
 creerRessource := func(id int, taille int) *Ressource {
  donnees := make([]byte, taille)

  ressource := &Ressource{
   ID:      id,
   Donnees: &donnees,
   Callback: func() {
    fmt.Printf("🧹 Nettoyage ressource %d\n", id)
   },
  }

  return ressource
 }

 // Créer et nettoyer des ressources
 ressources := make([]*Ressource, 0, 3)

 for i := 1; i <= 3; i++ {
  r := creerRessource(i, 1024*i) // Tailles variables
  ressources = append(ressources, r)
  fmt.Printf("📦 Ressource %d créée (%d bytes)\n", r.ID, 1024*i)
 }

 // Nettoyer les ressources
 fmt.Println("\nNettoyage des ressources:")
 for _, r := range ressources {
  if r.Callback != nil {
   r.Callback()
  }
 }

 // Les slices sont automatiquement nettoyés par le GC de Go
 ressources = nil
 fmt.Println("✅ Toutes les ressources libérées")
}

// ==========================================
// FONCTION PRINCIPALE
// ==========================================

func main() {
 fmt.Println("SOLUTIONS DES EXERCICES - POINTERS")
 fmt.Println("===================================")

 exercice1()
 exercice2()
 exercice3()

 // Bonus: démonstrations avancées
 demonstrationPerformance()
 testGestionMemoire()

 fmt.Println("\n🎉 Tous les exercices terminés!")
 fmt.Println("💡 Les pointeurs permettent une gestion mémoire efficace et un partage de données sécurisé en Go.")
}
```

## Résumé

Les pointeurs sont un outil puissant en Go pour :

**Avantages :**

- Efficacité mémoire (pas de copie)
- Modification de données partagées
- Structures de données complexes (listes, arbres)
- Polymorphisme avec interfaces

**Points clés à retenir :**

- `&` pour obtenir l'adresse, `*` pour déréférencer
- Vérifiez toujours les pointeurs nil
- Utilisez des pointeurs pour les grosses structs et les modifications
- Go gère la mémoire automatiquement (garbage collector)
- Les méthodes peuvent avoir des receivers par valeur ou par pointeur

**Règle générale :**

- Petites données immutables → valeurs
- Grosses données ou modifications → pointeurs

Les pointeurs permettent d'écrire du code Go efficace et expressif, tout en restant sûr grâce au garbage collector.

⏭️
