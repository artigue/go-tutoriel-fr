🔝 Retour au [Sommaire](/SOMMAIRE.md)

# 6. Méthodes et interfaces

## Introduction

Après avoir maîtrisé les fondamentaux de Go (variables, fonctions, structures de données), nous entrons dans un domaine plus avancé et puissant : les méthodes et les interfaces. Ces concepts sont au cœur de la programmation orientée objet en Go et constituent l'un des aspects les plus élégants et distinctifs du langage.

## Pourquoi les méthodes et interfaces sont-elles importantes ?

### 1. **Abstraction et encapsulation**

Les méthodes permettent d'associer des comportements à des types de données, créant ainsi une abstraction naturelle. Au lieu d'avoir des fonctions éparpillées qui manipulent des structures, nous pouvons grouper logiquement les fonctionnalités avec les données qu'elles manipulent.

### 2. **Polymorphisme à la Go**

Go n'a pas de classes au sens traditionnel, mais les interfaces offrent un mécanisme de polymorphisme encore plus puissant et flexible. Une interface en Go définit un contrat de comportement plutôt qu'une hiérarchie de types.

### 3. **Composition over inheritance**

Go favorise la composition plutôt que l'héritage. Les interfaces permettent de créer des systèmes modulaires et flexibles où les types peuvent implémenter plusieurs interfaces sans contraintes hiérarchiques.

## Concepts clés à retenir

### **Méthodes**

- Une méthode est une fonction associée à un type spécifique
- Elle possède un "receiver" qui détermine sur quel type elle opère
- Les méthodes peuvent être définies sur n'importe quel type (pas seulement les structs)

### **Interfaces**

- Une interface définit un ensemble de signatures de méthodes
- Un type implémente une interface simplement en définissant toutes ses méthodes
- L'implémentation est implicite (pas de mot-clé "implements")
- Les interfaces permettent le polymorphisme et la testabilité

### **Philosophie Go**
>
> "Don't design with interfaces, discover them" - Rob Pike

Cette philosophie encourage à créer des types concrets d'abord, puis à découvrir les interfaces naturelles qui émergent de l'utilisation.

## Ce que vous allez apprendre

Dans ce chapitre, nous explorerons :

1. **Comment définir et utiliser des méthodes** - Associer des comportements à vos types personnalisés
2. **Les interfaces et le polymorphisme** - Créer des abstractions flexibles et réutilisables
3. **L'interface vide et les assertions de type** - Gérer les types dynamiques en toute sécurité
4. **Composition vs héritage** - Comprendre l'approche Go pour structurer le code

## Prérequis

Avant de commencer ce chapitre, assurez-vous de maîtriser :

- Les types de base et les structs
- Les fonctions et leurs signatures
- Les pointeurs et leur utilisation
- Les packages et la visibilité (public/private)

## Exemple motivant

Prenons un exemple simple pour illustrer la puissance des méthodes et interfaces :

```go
// Sans méthodes/interfaces (approche procédurale)
type Rectangle struct {
    Width, Height float64
}

func CalculateArea(r Rectangle) float64 {
    return r.Width * r.Height
}

func CalculatePerimeter(r Rectangle) float64 {
    return 2 * (r.Width + r.Height)
}

// Avec méthodes/interfaces (approche orientée objet)
type Shape interface {
    Area() float64
    Perimeter() float64
}

type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

// Maintenant, toute fonction peut accepter n'importe quelle forme
func PrintShapeInfo(s Shape) {
    fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}
```

La deuxième approche est plus flexible, extensible et expressive. Elle nous permet de traiter différents types de formes de manière uniforme.

## Prêt à commencer ?

Les méthodes et interfaces transformeront votre façon de concevoir et structurer vos programmes Go. Elles ouvrent la porte à des patterns de conception élégants et à une architecture logicielle plus maintenable.

Commençons par la définition des méthodes...

⏭️
