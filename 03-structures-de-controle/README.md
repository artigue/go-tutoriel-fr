🔝 Retour au [Sommaire](/SOMMAIRE.md)

# 3. Structures de contrôle

## Introduction

Les structures de contrôle sont les éléments fondamentaux qui permettent de diriger l'exécution d'un programme. En Go, ces structures déterminent l'ordre dans lequel les instructions sont exécutées et permettent de créer des programmes dynamiques et intelligents.

## Qu'est-ce qu'une structure de contrôle ?

Une structure de contrôle est une construction syntaxique qui permet de contrôler le flux d'exécution d'un programme. Elle détermine quelles instructions seront exécutées, dans quel ordre, et combien de fois.

## Les types de structures de contrôle en Go

Go propose plusieurs types de structures de contrôle, chacune ayant un rôle spécifique :

### 1. Structures conditionnelles

Ces structures permettent d'exécuter différentes portions de code selon des conditions spécifiques :

- **if/else** : Exécute du code si une condition est vraie ou fausse
- **switch** : Permet de tester une valeur contre plusieurs cas possibles

### 2. Structures itératives

Ces structures permettent de répéter l'exécution d'un bloc de code :

- **for** : La seule boucle disponible en Go, mais très polyvalente
- **range** : Utilisé avec `for` pour itérer sur des collections

### 3. Structures de gestion d'erreurs

Go a une approche particulière pour gérer les erreurs :

- **Vérification explicite des erreurs** : Convention idiomatique de Go
- **defer** : Exécute du code à la fin d'une fonction
- **panic/recover** : Mécanisme d'exception (à utiliser avec parcimonie)

## Philosophie de Go concernant les structures de contrôle

Go adopte une approche minimaliste et claire :

### Simplicité avant tout

- **Pas de while ou do-while** : Seulement `for` pour toutes les boucles
- **Pas d'opérateur ternaire** : Utilisation explicite de `if/else`
- **Syntaxe claire** : Pas de parenthèses obligatoires autour des conditions

### Exemple de cette philosophie

```go
// Go privilégie la clarté
if x > 0 {
    return "positif"
} else {
    return "négatif ou nul"
}

// Plutôt qu'un opérateur ternaire comme dans d'autres langages
// return x > 0 ? "positif" : "négatif ou nul"  // N'existe pas en Go
```

## Caractéristiques importantes

### 1. Accolades obligatoires

En Go, les accolades `{}` sont toujours obligatoires, même pour une seule instruction :

```go
// Correct
if condition {
    doSomething()
}

// Incorrect - ne compile pas
if condition
    doSomething()
```

### 2. Pas de parenthèses autour des conditions

Contrairement à d'autres langages, Go n'exige pas de parenthèses autour des conditions :

```go
// Go style - recommandé
if x > 10 {
    // ...
}

// Aussi valide mais moins idiomatique
if (x > 10) {
    // ...
}
```

### 3. Déclaration et initialisation dans les conditions

Go permet de déclarer et initialiser des variables directement dans les structures de contrôle :

```go
if err := someFunction(); err != nil {
    // err n'est accessible que dans ce bloc
    return err
}
```

## Bonnes pratiques générales

### 1. Gestion des erreurs

En Go, la gestion d'erreurs est explicite et fait partie intégrante du flux de contrôle :

```go
result, err := riskyOperation()
if err != nil {
    // Toujours gérer l'erreur
    return fmt.Errorf("échec de l'opération : %w", err)
}
// Continuer avec result
```

### 2. Return early pattern

Go encourage le pattern "return early" pour réduire l'imbrication :

```go
func processData(data []int) error {
    if len(data) == 0 {
        return errors.New("données vides")
    }

    if !isValid(data) {
        return errors.New("données invalides")
    }

    // Traitement principal
    return nil
}
```

### 3. Éviter l'imbrication excessive

```go
// Préférer ceci
if !condition {
    return
}
// code principal

// Plutôt que ceci
if condition {
    // beaucoup de code imbriqué
}
```

## Ce que nous allons voir

Dans les sections suivantes, nous explorerons en détail :

- Les conditions avec `if/else` et `switch`
- Les différentes formes de boucles `for`
- L'utilisation de `range` pour itérer sur des collections
- Les techniques de gestion d'erreurs spécifiques à Go

Chaque structure sera accompagnée d'exemples pratiques et de cas d'usage courants pour vous permettre de maîtriser parfaitement le contrôle de flux en Go.

⏭️
