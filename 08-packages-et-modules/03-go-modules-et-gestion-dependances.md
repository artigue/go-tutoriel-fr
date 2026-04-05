🔝 Retour au [Sommaire](/SOMMAIRE.md)

# 8-3 : Go modules et gestion des dépendances

## Qu'est-ce qu'un module Go ?

Un **module** en Go est comme un "conteneur" pour votre projet qui :

- Définit le nom de votre projet
- Liste toutes les dépendances externes (autres packages/bibliothèques)
- Spécifie les versions exactes de ces dépendances
- Assure que votre projet fonctionne de manière identique sur tous les ordinateurs

Imaginez un module comme la "recette" complète de votre projet, avec tous les ingrédients et leurs quantités exactes !

## Pourquoi les modules sont-ils importants ?

### Avant les modules (le problème)

```bash
# Problèmes de l'ancien système GOPATH :
- Tous les projets dans un seul dossier
- Impossible d'avoir plusieurs versions d'une même bibliothèque
- Partage de code compliqué
- Builds non reproductibles
```

### Avec les modules (la solution)

```bash
# Avantages des modules :
✅ Chaque projet est indépendant
✅ Gestion fine des versions
✅ Builds reproductibles
✅ Partage facile de code
✅ Sécurité et intégrité des dépendances
```

## Créer votre premier module

### 1. Initialisation d'un module

```bash
# Créer un nouveau dossier pour votre projet
mkdir mon-projet
cd mon-projet

# Initialiser le module
go mod init mon-projet
```

Cela crée un fichier `go.mod` :

```go
module mon-projet

go 1.21
```

### 2. Structure recommandée

```text
mon-projet/
├── go.mod          # Définition du module
├── go.sum          # Checksums de sécurité (généré automatiquement)
├── main.go         # Point d'entrée
├── utils/
│   └── helpers.go
└── models/
    └── user.go
```

## Le fichier go.mod expliqué

### Exemple complet

```go
module github.com/monnom/mon-super-projet

go 1.21

require (
    github.com/gin-gonic/gin v1.9.1
    github.com/gorilla/mux v1.8.0
    golang.org/x/crypto v0.10.0
)

exclude (
    github.com/ancien/package v1.0.0
)

replace (
    github.com/fork/package => github.com/monfork/package v1.2.3
)
```

### Décortiquons chaque partie

**1. Module name**

```go
module github.com/monnom/mon-super-projet
```

- Nom unique de votre module
- Souvent l'URL GitHub/GitLab de votre projet
- Utilisé par d'autres pour importer votre code

**2. Version Go**

```go
go 1.21
```

- Version minimale de Go requise
- Assure la compatibilité

**3. Dépendances requises**

```go
require (
    github.com/gin-gonic/gin v1.9.1    # Version exacte
    github.com/gorilla/mux v1.8.0      # Version exacte
)
```

**4. Exclusions (optionnel)**

```go
exclude (
    github.com/buggy/package v1.0.0    # Exclut une version problématique
)
```

**5. Remplacements (optionnel)**

```go
replace (
    github.com/original => github.com/myfork v1.2.3
)
```

## Exemple pratique : Application web simple

Créons une application web qui utilise des dépendances externes.

### 1. Initialisation du projet

```bash
mkdir webapp-example
cd webapp-example
go mod init webapp-example
```

### 2. Création du code principal

```go
// main.go
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    // Créer un routeur Gin
    r := gin.Default()

    // Route simple
    r.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello World!",
            "status":  "success",
        })
    })

    // Route avec paramètre
    r.GET("/user/:name", func(c *gin.Context) {
        name := c.Param("name")
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello " + name,
            "user":    name,
        })
    })

    // Démarrer le serveur
    r.Run(":8080")
}
```

### 3. Ajout des dépendances

```bash
# Ajouter la dépendance Gin
go get github.com/gin-gonic/gin

# Votre go.mod est automatiquement mis à jour !
```

Le fichier `go.mod` devient :

```go
module webapp-example

go 1.21

require github.com/gin-gonic/gin v1.9.1

require (
    github.com/bytedance/sonic v1.9.1 // indirect
    github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311 // indirect
    // ... autres dépendances indirectes
)
```

### 4. Tester l'application

```bash
go run main.go
# Ouvrir http://localhost:8080 dans votre navigateur
```

## Commandes essentielles

### 1. go mod init

```bash
# Créer un nouveau module
go mod init nom-du-module

# Exemples :
go mod init monapp
go mod init github.com/monnom/monapp
go mod init example.com/moncompany/monapp
```

### 2. go get

```bash
# Ajouter une dépendance
go get github.com/gorilla/mux

# Ajouter une version spécifique
go get github.com/gorilla/mux@v1.8.0

# Mettre à jour vers la dernière version
go get -u github.com/gorilla/mux

# Mettre à jour toutes les dépendances
go get -u ./...
```

### 3. go mod tidy

```bash
# Nettoyer les dépendances
go mod tidy

# Supprime les dépendances inutilisées
# Ajoute les dépendances manquantes
# Met à jour go.sum
```

### 4. go mod download

```bash
# Télécharger les dépendances sans les compiler
go mod download
```

### 5. go list

```bash
# Lister toutes les dépendances
go list -m all

# Voir les versions disponibles d'un package
go list -m -versions github.com/gin-gonic/gin
```

## Gestion des versions

### Semantic Versioning (SemVer)

Go utilise le versioning sémantique : `v1.2.3`

```text
v1.2.3
│ │ │
│ │ └─ PATCH : corrections de bugs
│ └─── MINOR : nouvelles fonctionnalités (compatibles)
└───── MAJOR : changements incompatibles
```

### Exemples de contraintes de version

```bash
# Version exacte
go get github.com/gin-gonic/gin@v1.9.1

# Dernière version mineure de v1
go get github.com/gin-gonic/gin@v1

# Dernière version patch de v1.9
go get github.com/gin-gonic/gin@v1.9

# Commit spécifique
go get github.com/gin-gonic/gin@abc1234

# Branche spécifique
go get github.com/gin-gonic/gin@master
```

## Exemple avancé : API avec base de données

Créons une API plus complexe avec plusieurs dépendances.

### 1. Initialisation

```bash
mkdir api-example
cd api-example
go mod init api-example
```

### 2. Ajout des dépendances

```bash
go get github.com/gin-gonic/gin       # Framework web
go get github.com/jinzhu/gorm         # ORM pour base de données
go get github.com/lib/pq              # Driver PostgreSQL
go get github.com/joho/godotenv       # Variables d'environnement
```

### 3. Code de l'application

```go
// main.go
package main

import (
    "log"
    "net/http"
    "os"

    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    _ "github.com/lib/pq"
    "github.com/joho/godotenv"
)

type User struct {
    ID    uint   `json:"id" gorm:"primary_key"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

var db *gorm.DB
var err error

func main() {
    // Charger les variables d'environnement
    err := godotenv.Load()
    if err != nil {
        log.Println("Aucun fichier .env trouvé")
    }

    // Connexion à la base de données
    dbURL := os.Getenv("DATABASE_URL")
    if dbURL == "" {
        dbURL = "user=postgres password=password dbname=testdb sslmode=disable"
    }

    db, err = gorm.Open("postgres", dbURL)
    if err != nil {
        log.Fatal("Erreur de connexion à la base de données")
    }
    defer db.Close()

    // Migration automatique
    db.AutoMigrate(&User{})

    // Configuration Gin
    r := gin.Default()

    // Routes
    r.GET("/users", getUsers)
    r.POST("/users", createUser)
    r.GET("/users/:id", getUser)

    // Démarrer le serveur
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    r.Run(":" + port)
}

func getUsers(c *gin.Context) {
    var users []User
    db.Find(&users)
    c.JSON(http.StatusOK, users)
}

func createUser(c *gin.Context) {
    var user User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    db.Create(&user)
    c.JSON(http.StatusCreated, user)
}

func getUser(c *gin.Context) {
    id := c.Param("id")
    var user User

    if err := db.First(&user, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Utilisateur non trouvé"})
        return
    }

    c.JSON(http.StatusOK, user)
}
```

### 4. Fichier .env

```bash
# .env
DATABASE_URL=user=postgres password=password dbname=testdb sslmode=disable
PORT=8080
```

Le fichier `go.mod` final :

```go
module api-example

go 1.21

require (
    github.com/gin-gonic/gin v1.9.1
    github.com/jinzhu/gorm v1.9.16
    github.com/joho/godotenv v1.4.0
    github.com/lib/pq v1.10.9
)

require (
    github.com/bytedance/sonic v1.9.1 // indirect
    github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311 // indirect
    // ... autres dépendances indirectes
)
```

## Publier votre module

### 1. Préparer le code

```bash
# Structure recommandée pour un module public
mon-package/
├── go.mod
├── README.md
├── LICENSE
├── .gitignore
├── main.go
└── examples/
    └── basic/
        └── main.go
```

### 2. Tagging des versions

```bash
# Créer un tag Git pour une version
git tag v1.0.0
git push origin v1.0.0

# Go utilise les tags Git pour les versions !
```

### 3. Documentation

```go
// Package mathutils fournit des utilitaires mathématiques.
//
// Ce package contient des fonctions pour des opérations
// mathématiques courantes non disponibles dans le package math standard.
package mathutils

// Add additionne deux nombres entiers.
//
// Exemple :
//   result := mathutils.Add(5, 3) // result = 8
func Add(a, b int) int {
    return a + b
}
```

## Résolution de problèmes courants

### 1. Dépendances en conflit

```bash
# Voir les détails des conflits
go mod graph

# Forcer une version spécifique
go mod edit -require=github.com/problematic/package@v1.2.3
go mod tidy
```

### 2. Modules locaux

```go
// go.mod
module monapp

require (
    monpackagelocal v0.0.0
)

replace monpackagelocal => ./local-package
```

### 3. Vérifier l'intégrité

```bash
# Vérifier que go.sum est à jour
go mod verify

# Régénérer go.sum
rm go.sum
go mod tidy
```

## Bonnes pratiques

### 1. Nommage des modules

```bash
# ✅ Bon : utilise un domaine
go mod init github.com/monnom/monprojet
go mod init example.com/company/project

# ❌ Évitez : noms trop génériques
go mod init utils
go mod init helpers
```

### 2. Gestion des versions

```bash
# ✅ Bon : versions spécifiques en production
go get github.com/gin-gonic/gin@v1.9.1

# ⚠️ Attention : versions flottantes en développement
go get github.com/gin-gonic/gin@latest
```

### 3. Fichier .gitignore

```gitignore
# Binaires
*.exe
*.exe~
*.dll
*.so
*.dylib

# Fichiers de test
*.test

# Couverture de code
*.out

# Dépendances vendor (si utilisées)
vendor/
```

## Exercice pratique

Créez un module "task-manager" qui :

1. **Initialise un module** avec `go mod init`
2. **Utilise les dépendances** :
   - `github.com/gin-gonic/gin` pour le serveur web
   - `github.com/google/uuid` pour générer des IDs
3. **Implémente** :
   - Structure `Task` avec ID, Title, Done
   - Routes GET, POST pour les tâches
   - Stockage en mémoire

**Test :**

```bash
curl http://localhost:8080/tasks
curl -X POST -H "Content-Type: application/json" \
  -d '{"title":"Apprendre Go"}' \
  http://localhost:8080/tasks
```

## Résumé

Les modules Go permettent :

**📦 Organisation :**

- Projets indépendants
- Gestion fine des versions
- Distribution facile

**🔒 Sécurité :**

- Checksums d'intégrité (go.sum)
- Versions verrouillées
- Builds reproductibles

**⚡ Productivité :**

- Dépendances automatiques
- Mises à jour simples
- Outils intégrés

**📋 Commandes clés :**

- `go mod init` : créer un module
- `go get` : ajouter/mettre à jour des dépendances
- `go mod tidy` : nettoyer les dépendances
- `go mod verify` : vérifier l'intégrité

Dans la section suivante, nous verrons comment documenter efficacement vos packages avec godoc !

⏭️
