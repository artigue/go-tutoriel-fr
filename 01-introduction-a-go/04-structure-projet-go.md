🔝 Retour au [Sommaire](/SOMMAIRE.md)

# 1-4 : Structure d'un projet Go

## Introduction

Maintenant que vous savez créer un programme Go basique, il est temps d'apprendre à organiser votre code de manière professionnelle. Une bonne structure de projet vous aidera à :

- Maintenir votre code facilement
- Collaborer avec d'autres développeurs
- Faire évoluer votre application
- Suivre les conventions de la communauté Go

## Évolution de l'organisation Go

### Avant Go 1.11 : GOPATH

**Ancienne méthode (à connaître mais plus utilisée) :**

```
$GOPATH/
├── bin/        # Exécutables compilés
├── pkg/        # Packages compilés
└── src/        # Code source
    └── github.com/
        └── utilisateur/
            └── projet/
```

### Depuis Go 1.11 : Go Modules (Moderne)

**Nouvelle approche (recommandée) :**

- Les projets peuvent être n'importe où sur votre système
- Chaque projet a son propre fichier `go.mod`
- Gestion automatique des dépendances
- Plus flexible et plus simple

## Anatomie d'un module Go

### Le fichier go.mod

**Créer un nouveau module :**

```bash
mkdir mon-projet
cd mon-projet
go mod init github.com/votre-nom/mon-projet
```

**Contenu du fichier `go.mod` :**

```go
module github.com/votre-nom/mon-projet

go 1.21

require (
    // Les dépendances apparaîtront ici automatiquement
)
```

**Explication des parties :**

- `module` : Nom unique de votre module
- `go 1.21` : Version minimale de Go requise
- `require` : Liste des dépendances externes

### Le fichier go.sum

**Apparaît automatiquement quand vous ajoutez des dépendances :**

```

github.com/gin-gonic/gin v1.9.1 h1:4idEAncQnU5cB7BeOkPtxjfCSye0AAm1R0RVIqJ+Jmg=
github.com/gin-gonic/gin v1.9.1/go.mod h1:hPrL7YrpYKXt5YId3A/Tnip5kqbEAP+KLuI3SUcPTeU=
```

**Rôle :**

- Vérifie l'intégrité des dépendances
- Assure la reproductibilité des builds
- Ne pas modifier manuellement

## Structure d'un projet simple

### Projet basique (un seul fichier)

````go

mon-petit-projet/
├── go.mod
├── main.go
└── README.md
```go
**Exemple de `main.go` :**

```go
package main

import "fmt"

func main() {
    fmt.Println("Mon petit projet Go")
}
````

### Projet avec plusieurs fichiers

```go
calculatrice/
├── go.mod
├── main.go
├── operations.go
└── README.md
```

**main.go :**

```go
package main

import "fmt"

func main() {
    resultat := addition(5, 3)
    fmt.Printf("5 + 3 = %d\n", resultat)
}
```

**operations.go :**

```go
package main

func addition(a, b int) int {
    return a + b
}

func soustraction(a, b int) int {
    return a - b
}
```

**Points importants :**

- Tous les fichiers du même répertoire doivent avoir le même `package`
- Les fonctions commençant par une majuscule sont publiques
- Les fonctions commençant par une minuscule sont privées

## Structure d'un projet moyen

### Organisation recommandée

```go
mon-api/
├── go.mod
├── go.sum
├── main.go
├── README.md
├── handlers/
│   ├── user.go
│   └── product.go
├── models/
│   ├── user.go
│   └── product.go
├── database/
│   └── connection.go
└── config/
    └── config.go
```

### Exemple concret

**main.go (point d'entrée) :**

```go
package main

import (
    "fmt"
    "mon-api/handlers"
    "mon-api/config"
)

func main() {
    cfg := config.Load()
    fmt.Printf("Serveur démarré sur le port %s\n", cfg.Port)

    handlers.SetupRoutes()
}
```

**config/config.go :**

```go
package config

type Config struct {
    Port     string
    Database string
}

func Load() *Config {
    return &Config{
        Port:     "8080",
        Database: "localhost:5432",
    }
}
```

**handlers/user.go :**

```go
package handlers

import "fmt"

func GetUser() {
    fmt.Println("Récupération d'un utilisateur")
}

func CreateUser() {
    fmt.Println("Création d'un utilisateur")
}
```

**models/user.go :**

```go
package models

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

func (u *User) Validate() bool {
    return u.Name != "" && u.Email != ""
}
```

## Structure d'un grand projet

### Layout Standard (Recommandation communautaire)

```text
grand-projet/
├── go.mod
├── go.sum
├── README.md
├── Makefile
├── cmd/                    # Applications principales
│   ├── api/
│   │   └── main.go
│   └── worker/
│       └── main.go
├── internal/               # Code privé de l'application
│   ├── app/
│   ├── pkg/
│   └── platform/
├── pkg/                    # Code réutilisable par d'autres
│   ├── logger/
│   └── database/
├── api/                    # Spécifications API
│   └── openapi.yaml
├── web/                    # Assets web
│   ├── static/
│   └── templates/
├── configs/                # Fichiers de configuration
│   └── config.yaml
├── scripts/                # Scripts d'automatisation
│   └── deploy.sh
├── test/                   # Tests d'intégration
│   └── integration/
├── docs/                   # Documentation
│   └── architecture.md
└── .github/                # GitHub Actions
    └── workflows/
```

### Explication des répertoires

#### `/cmd`

**Rôle :** Applications principales de votre projet

**Exemple :**

```text
cmd/
├── api/
│   └── main.go      # Point d'entrée de l'API
├── worker/
│   └── main.go      # Point d'entrée du worker
└── migration/
    └── main.go      # Outil de migration
```

#### `/internal`

**Rôle :** Code privé qui ne peut pas être importé par d'autres projets

**Avantage :** Go empêche automatiquement l'import de code dans `/internal`

```go
// Ceci NE FONCTIONNE PAS depuis un autre projet :
import "github.com/autre/projet/internal/database"
```

#### `/pkg`

**Rôle :** Code réutilisable par d'autres projets

**Exemple :**

```go
// Ceci fonctionne depuis d'autres projets :
import "github.com/votre/projet/pkg/logger"
```

#### `/api`

**Rôle :** Spécifications et documentation des APIs

**Contenu typique :**

- Fichiers OpenAPI/Swagger
- Schémas JSON
- Documentation des protocoles

## Conventions de nommage

### Packages

**Bonnes pratiques :**

```go
// ✅ Bon : court, descriptif, minuscules
package user
package http
package json

// ❌ Éviter : trop long, majuscules, underscores
package user_management
package HTTP
package JsonParser
```

### Fichiers

**Convention :**

```txt
user.go           # ✅ minuscules
user_service.go   # ✅ underscore pour séparer
UserService.go    # ❌ éviter les majuscules
user-service.go   # ❌ éviter les tirets
```

### Répertoires

**Structure recommandée :**

```txt
models/           # ✅ pluriel
handlers/         # ✅ pluriel
config/           # ✅ singulier (un seul concept)
database/         # ✅ singulier (une seule base)
```

## Gestion des dépendances

### Ajouter une dépendance

```bash
# Ajouter une dépendance
go get github.com/gin-gonic/gin

# Ajouter une version spécifique
go get github.com/gin-gonic/gin@v1.9.1
```

**Le fichier `go.mod` se met à jour automatiquement :**

```go
module mon-projet

go 1.21

require (
    github.com/gin-gonic/gin v1.9.1
)
```

### Utiliser la dépendance

```go
package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello World!",
        })
    })
    r.Run(":8080")
}
```

### Nettoyer les dépendances

```bash
# Supprimer les dépendances inutilisées
go mod tidy

# Télécharger les dépendances
go mod download
```

## Exemple pratique : Projet Todo List

### Structure complète

```txt
todo-app/
├── go.mod
├── main.go
├── README.md
├── models/
│   └── todo.go
├── handlers/
│   └── todo.go
├── storage/
│   └── memory.go
└── static/
    └── index.html
```

### Implémentation

**go.mod :**

```go
module todo-app

go 1.21

require github.com/gin-gonic/gin v1.9.1
```

**models/todo.go :**

```go
package models

import "time"

type Todo struct {
    ID        int       `json:"id"`
    Title     string    `json:"title"`
    Completed bool      `json:"completed"`
    CreatedAt time.Time `json:"created_at"`
}

func NewTodo(title string) *Todo {
    return &Todo{
        Title:     title,
        Completed: false,
        CreatedAt: time.Now(),
    }
}
```

**storage/memory.go :**

```go
package storage

import (
    "todo-app/models"
    "sync"
)

type MemoryStorage struct {
    todos  []models.Todo
    nextID int
    mu     sync.RWMutex
}

func NewMemoryStorage() *MemoryStorage {
    return &MemoryStorage{
        todos:  make([]models.Todo, 0),
        nextID: 1,
    }
}

func (s *MemoryStorage) Create(todo *models.Todo) {
    s.mu.Lock()
    defer s.mu.Unlock()

    todo.ID = s.nextID
    s.nextID++
    s.todos = append(s.todos, *todo)
}

func (s *MemoryStorage) GetAll() []models.Todo {
    s.mu.RLock()
    defer s.mu.RUnlock()

    return append([]models.Todo(nil), s.todos...)
}
```

**handlers/todo.go :**

```go
package handlers

import (
    "net/http"
    "todo-app/models"
    "todo-app/storage"

    "github.com/gin-gonic/gin"
)

type TodoHandler struct {
    storage *storage.MemoryStorage
}

func NewTodoHandler(s *storage.MemoryStorage) *TodoHandler {
    return &TodoHandler{storage: s}
}

func (h *TodoHandler) CreateTodo(c *gin.Context) {
    var req struct {
        Title string `json:"title" binding:"required"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    todo := models.NewTodo(req.Title)
    h.storage.Create(todo)

    c.JSON(http.StatusCreated, todo)
}

func (h *TodoHandler) GetTodos(c *gin.Context) {
    todos := h.storage.GetAll()
    c.JSON(http.StatusOK, todos)
}
```

**main.go :**

```go
package main

import (
    "todo-app/handlers"
    "todo-app/storage"

    "github.com/gin-gonic/gin"
)

func main() {
    // Initialiser le stockage
    store := storage.NewMemoryStorage()

    // Initialiser les handlers
    todoHandler := handlers.NewTodoHandler(store)

    // Configuration du routeur
    r := gin.Default()

    // Routes API
    api := r.Group("/api/v1")
    {
        api.GET("/todos", todoHandler.GetTodos)
        api.POST("/todos", todoHandler.CreateTodo)
    }

    // Servir les fichiers statiques
    r.Static("/static", "./static")
    r.LoadHTMLGlob("static/*")
    r.GET("/", func(c *gin.Context) {
        c.HTML(200, "index.html", nil)
    })

    // Démarrer le serveur
    r.Run(":8080")
}
```

### Lancer le projet

```bash
# Initialiser les dépendances
go mod tidy

# Lancer l'application
go run main.go

# Ou compiler puis exécuter
go build -o todo-app
./todo-app
```

## Bonnes pratiques

### 1. Organisation logique

- Grouper le code par fonctionnalité, pas par type
- Séparer les préoccupations (handlers, models, storage)
- Utiliser des noms descriptifs

### 2. Visibilité des packages

```go
// ✅ Fonction publique (majuscule)
func CreateUser() {}

// ✅ Fonction privée (minuscule)
func validateEmail() {}
```

### 3. Documentation

```go
// Package user fournit les fonctionnalités de gestion des utilisateurs.
package user

// User représente un utilisateur de l'application.
type User struct {
    ID   int
    Name string
}

// CreateUser crée un nouvel utilisateur avec le nom donné.
func CreateUser(name string) *User {
    return &User{Name: name}
}
```

### 4. Tests

```
projet/
├── handlers/
│   ├── user.go
│   └── user_test.go      # Tests dans le même package
├── models/
│   ├── user.go
│   └── user_test.go
```

## Outils utiles

### Commandes de gestion de projet

```bash
# Initialiser un module
go mod init nom-du-projet

# Nettoyer les dépendances
go mod tidy

# Vérifier les dépendances
go mod verify

# Voir l'arbre des dépendances
go mod graph

# Créer un vendor (optionnel)
go mod vendor
```

### Outils de développement

```bash
# Formater tout le projet
go fmt ./...

# Vérifier tout le projet
go vet ./...

# Lancer tous les tests
go test ./...

# Build pour différentes plateformes
GOOS=linux GOARCH=amd64 go build
GOOS=windows GOARCH=amd64 go build
```

## Récapitulatif

**Ce que vous avez appris :**

- ✅ Évolution de GOPATH vers Go Modules
- ✅ Structure d'un projet Go moderne
- ✅ Organisation du code par packages
- ✅ Conventions de nommage
- ✅ Gestion des dépendances
- ✅ Exemple pratique complet

**Points clés :**

1. **Go Modules** : La méthode moderne pour organiser les projets
2. **Séparation des responsabilités** : Chaque package a un rôle précis
3. **Visibilité** : Majuscule = public, minuscule = privé
4. **Structure standard** : `/cmd`, `/internal`, `/pkg` pour les gros projets
5. **Dépendances** : `go mod` gère tout automatiquement

**Commandes essentielles :**

```bash
go mod init projet        # Créer un nouveau module
go get package           # Ajouter une dépendance
go mod tidy              # Nettoyer les dépendances
go fmt ./...             # Formater tout le projet
go test ./...            # Tester tout le projet
```

---

_Félicitations ! Vous maîtrisez maintenant l'organisation d'un projet Go. Dans la prochaine section, nous plongerons dans la syntaxe de base avec les variables et les types de données !_

⏭️
