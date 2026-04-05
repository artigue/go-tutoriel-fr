🔝 Retour au [Sommaire](/SOMMAIRE.md)

# 5. Structures de données

## Introduction

Les structures de données sont les fondations de tout programme informatique. En Go, elles permettent d'organiser, de stocker et de manipuler efficacement vos données. Cette section couvre les principales structures de données natives du langage Go, chacune ayant ses propres caractéristiques et cas d'usage.

## Vue d'ensemble des structures de données en Go

Go propose plusieurs types de structures de données intégrées, chacune optimisée pour des scénarios spécifiques :

### Arrays (Tableaux)

Les tableaux sont des collections d'éléments de même type, avec une taille fixe définie à la compilation. Ils offrent un accès rapide aux éléments par index mais manquent de flexibilité en raison de leur taille immutable.

### Slices (Tranches)

Les slices sont des abstractions dynamiques construites au-dessus des tableaux. Elles constituent l'une des structures de données les plus utilisées en Go, offrant flexibilité et performance pour la manipulation de collections d'éléments.

### Maps (Cartes/Dictionnaires)

Les maps sont des structures de données clé-valeur qui permettent de stocker et récupérer des données basées sur une clé unique. Elles sont équivalentes aux dictionnaires dans d'autres langages.

### Structs (Structures)

Les structs permettent de regrouper des données de différents types sous une même entité. Elles constituent la base de la programmation orientée données en Go et remplacent les classes d'autres langages.

### Pointers (Pointeurs)

Les pointeurs stockent l'adresse mémoire d'une variable plutôt que sa valeur directe. Ils permettent un contrôle fin de la mémoire et facilitent le partage efficace de données entre fonctions.

## Pourquoi ces structures sont-elles importantes ?

### Performance

Go privilégie la performance et la simplicité. Ces structures de données sont conçues pour être efficaces en mémoire et rapides à l'exécution, avec des optimisations intégrées au niveau du compilateur.

### Sécurité de type

Toutes ces structures bénéficient du système de types fort de Go, ce qui permet de détecter de nombreuses erreurs à la compilation plutôt qu'à l'exécution.

### Simplicité et lisibilité

La syntaxe de Go pour manipuler ces structures est volontairement simple et intuitive, favorisant la lisibilité du code et réduisant les erreurs.

## Concepts clés à retenir

### Valeurs vs Références

- **Arrays et Structs** : Passés par valeur par défaut (copie complète)
- **Slices et Maps** : Contiennent des références vers les données sous-jacentes
- **Pointers** : Références explicites vers l'adresse mémoire

### Zéro Values (Valeurs nulles)

Chaque type en Go a une "zero value" - une valeur par défaut sûre :

- Array : tous les éléments à leur zero value
- Slice : `nil`
- Map : `nil`
- Struct : tous les champs à leur zero value
- Pointer : `nil`

### Allocation mémoire

Go gère automatiquement l'allocation mémoire avec son garbage collector, mais comprendre comment ces structures utilisent la mémoire vous aidera à écrire du code plus efficace.

## Comment aborder cette section

Cette section est organisée de manière progressive :

1. **Arrays et Slices** : Les fondamentaux des collections indexées
2. **Maps** : Les structures clé-valeur pour l'accès associatif
3. **Structs** : L'organisation de données complexes
4. **Pointers** : La gestion fine de la mémoire et des références

Chaque sous-section inclura :

- Les concepts théoriques
- Des exemples pratiques
- Les bonnes pratiques
- Les pièges courants à éviter
- Des exercices pour consolider vos connaissances

## Prérequis

Avant de commencer cette section, assurez-vous de maîtriser :

- La syntaxe de base de Go
- Les types de données primitifs
- Les fonctions et leur utilisation
- Les structures de contrôle (if, for, etc.)

## Objectifs d'apprentissage

À la fin de cette section, vous devriez être capable de :

- Choisir la structure de données appropriée pour vos besoins
- Manipuler efficacement arrays, slices, maps, structs et pointers
- Comprendre les implications de performance de chaque structure
- Éviter les erreurs courantes liées à la gestion mémoire
- Écrire du code Go idiomatique utilisant ces structures

Commençons maintenant par explorer en détail chacune de ces structures de données essentielles !

⏭️
