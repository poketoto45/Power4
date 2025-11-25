package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	fichier "main.go/Power4-classic"
)

// GameState représente l'état actuel du jeu
type GameState struct {
	Grille    [6][7]string `json:"grille"`
	JoueurAct string       `json:"joueur_actuel"`
	Statut    string       `json:"statut"` // "en_cours", "gagne_R", "gagne_J", "egalite"
}

// MoveRequest représente une tentative de placement d'un pion
type MoveRequest struct {
	Colonne int `json:"colonne"`
}

// MoveResponse représente la réponse après un coup
type MoveResponse struct {
	Succes  bool         `json:"succes"`
	Grille  [6][7]string `json:"grille"`
	Statut  string       `json:"statut"`
	Message string       `json:"message"`
}

// ResetResponse représente la réponse après un reset
type ResetResponse struct {
	Message string       `json:"message"`
	Grille  [6][7]string `json:"grille"`
}

var gameState = &GameState{
	Grille:    [6][7]string{},
	JoueurAct: "R",
	Statut:    "en_cours",
}

func init() {
	fichier.InitialiserGrille()
	gameState.Grille = fichier.ObtenirGrille()
}

// handleJeu retourne la page HTML du jeu
func handleJeu(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.ServeFile(w, r, filepath.Join("Power4-classic/main", "page.html"))
}

// handleStatic sert les fichiers statiques
func handleStatic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, filepath.Join("Power4-classic/main", r.URL.Path[8:]))
}

// handleState retourne l'état actuel du jeu
func handleState(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	gameState.Grille = fichier.ObtenirGrille()
	json.NewEncoder(w).Encode(gameState)
}

// handleMove traite un coup du joueur
func handleMove(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req MoveRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	response := &MoveResponse{
		Succes: false,
		Grille: fichier.ObtenirGrille(),
		Statut: gameState.Statut,
	}

	// Valider la colonne
	if req.Colonne < 0 || req.Colonne >= 7 {
		response.Message = "Colonne invalide"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	// Placer le pion
	if !fichier.PlacerPion(gameState.JoueurAct, req.Colonne) {
		response.Message = "Colonne pleine"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Succes = true
	response.Grille = fichier.ObtenirGrille()
	response.Message = "Pion placé avec succès"

	// Vérifier si le joueur a gagné
	if fichier.Victoire(gameState.JoueurAct) {
		gameState.Statut = "gagne_" + gameState.JoueurAct
		response.Statut = gameState.Statut
		response.Message = "Joueur " + gameState.JoueurAct + " a gagné !"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	// Vérifier l'égalité
	if fichier.GrillePleine() {
		gameState.Statut = "egalite"
		response.Statut = gameState.Statut
		response.Message = "Égalité !"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	// Changer de joueur
	if gameState.JoueurAct == "R" {
		gameState.JoueurAct = "J"
	} else {
		gameState.JoueurAct = "R"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// handleReset réinitialise le jeu
func handleReset(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fichier.InitialiserGrille()
	gameState.Grille = fichier.ObtenirGrille()
	gameState.JoueurAct = "R"
	gameState.Statut = "en_cours"

	response := &ResetResponse{
		Message: "Jeu réinitialisé",
		Grille:  gameState.Grille,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Routes
	http.HandleFunc("/", handleJeu)
	http.HandleFunc("/static/", handleStatic)
	http.HandleFunc("/api/state", handleState)
	http.HandleFunc("/api/move", handleMove)
	http.HandleFunc("/api/reset", handleReset)

	fmt.Println("🎮 Serveur Power 4 démarré sur http://localhost:9091")
	log.Fatal(http.ListenAndServe(":9091", nil))
}
