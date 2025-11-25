# 🎮 Power 4 - Jeu en Ligne

Une implémentation complète du jeu Power 4 jouable sur **localhost** avec une interface web moderne et interactive.

## 🚀 Démarrage rapide

### Avec Go directement
```bash
go build -o server/server ./server
./server/server
```

### Via terminal
```bash
cd /home/thoma/git/Power4
go run ./server/main.go
```

## 🌐 Accès au jeu

Une fois le serveur démarré, accédez au jeu à l'adresse :
```
http://localhost:9090
```

## 🎯 Comment jouer

1. **Deux joueurs** : Rouge (R) et Jaune (J)
2. **Placement** : Cliquez sur les boutons en haut pour placer vos pions dans une colonne
3. **Victoire** : Alignez 4 pions horizontalement, verticalement ou en diagonale
4. **Égalité** : Si la grille se remplit sans gagnant
5. **Nouvelle partie** : Cliquez sur "🔄 NOUVELLE PARTIE"

## 📁 Structure du projet

```
Power4/
├── go.mod                  # Configuration du module Go
├── request.go             # Structures (compatibilité)
├── START.sh              # Script de lancement
├── server/
│   └── main.go           # Serveur HTTP et API
├── Power4-classic/
│   ├── fichier.go        # Logique de la grille
│   ├── fonction.go       # Fonctions exportées
│   └── main/
│       ├── page.html     # Interface du jeu
│       ├── page.css      # Styles (utilise les images)
│       └── images/
│           ├── bouton.png          # Boutons de colonnes
│           ├── grille_puissance4.png # Grille de jeu
│           ├── imgpower4.png       # Logo
│           ├── jeton_r.png        # Pion rouge
│           └── jeton_j.png        # Pion jaune
```

## 🎨 Interface

- **Grille interactive** : Affichée avec la vraie image de grille
- **Pions animés** : Chute fluide avec animation
- **Boutons** : Images authentiques pour chaque colonne
- **Statut** : Affiche le joueur actuel et l'état du jeu
- **Responsive** : S'adapte à tous les écrans

## 🔌 API REST

### État du jeu
```
GET http://localhost:9090/api/state
```

### Placer un pion
```
POST http://localhost:9090/api/move
Content-Type: application/json

{"colonne": 0}  // 0 à 6
```

### Réinitialiser
```
POST http://localhost:9090/api/reset
```

## 🔧 Technologies

- **Backend** : Go 1.22.2
- **Frontend** : HTML5 + JavaScript vanilla + CSS3
- **Communication** : JSON REST API

## 📝 Fichiers clés

- `server/main.go` : Gère les routes, l'état du jeu et les réponses JSON
- `Power4-classic/fichier.go` : Logique du jeu (grille, victoire, placement)
- `Power4-classic/main/page.html` : Rendu du jeu côté client
- `Power4-classic/main/page.css` : Stylisation avec images intégrées

## ✅ Fonctionnalités

✅ Jeu complet 2 joueurs
✅ Interface web moderne
✅ Images pour tous les éléments visuels
✅ Détection de victoire (4 directions)
✅ Gestion des limites de grille
✅ Animations fluides
✅ API RESTful
✅ Support multi-joueur en temps réel

Bon jeu ! 🎉
