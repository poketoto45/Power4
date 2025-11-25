package main

// GameState représente l'état actuel du jeu
type GameState struct {
	Grille    [6][7]string `json:"grille"`
	JoueurAct string       `json:"joueur_actuel"`
	Statut    string       `json:"statut"` // "en_cours", "gagne_X", "gagne_O", "egalite"
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
