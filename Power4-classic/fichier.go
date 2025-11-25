package fichier

import "sync"

// Initialisation de la grille
var grille [6][7]string
var mu sync.Mutex

// InitialiserGrille initialise une nouvelle grille vide
func InitialiserGrille() {
	mu.Lock()
	defer mu.Unlock()
	for i := 0; i < 6; i++ {
		for j := 0; j < 7; j++ {
			grille[i][j] = "."
		}
	}
}

// ObtenirGrille retourne une copie de la grille
func ObtenirGrille() [6][7]string {
	mu.Lock()
	defer mu.Unlock()
	return grille
}

// Fonctions
func PlacerPion(symbole string, colonne int) bool {
	mu.Lock()
	defer mu.Unlock()
	if colonne < 0 || colonne >= 7 {
		return false
	}
	for i := 5; i >= 0; i-- { // de bas en haut
		if grille[i][colonne] == "." {
			grille[i][colonne] = symbole
			return true
		}
	}
	return false
}

// Victoire vérifie si le joueur a gagné
func Victoire(symbole string) bool {
	mu.Lock()
	defer mu.Unlock()

	// Vérifie lignes horizontales
	for i := 0; i < 6; i++ {
		for j := 0; j <= 3; j++ {
			if grille[i][j] == symbole &&
				grille[i][j+1] == symbole &&
				grille[i][j+2] == symbole &&
				grille[i][j+3] == symbole {
				return true
			}
		}
	}

	// Vérifie lignes verticales
	for i := 0; i <= 2; i++ {
		for j := 0; j < 7; j++ {
			if grille[i][j] == symbole &&
				grille[i+1][j] == symbole &&
				grille[i+2][j] == symbole &&
				grille[i+3][j] == symbole {
				return true
			}
		}
	}

	// Vérifie diagonales descendantes (/)
	for i := 3; i < 6; i++ {
		for j := 0; j <= 3; j++ {
			if grille[i][j] == symbole &&
				grille[i-1][j+1] == symbole &&
				grille[i-2][j+2] == symbole &&
				grille[i-3][j+3] == symbole {
				return true
			}
		}
	}

	// Vérifie diagonales montantes (\)
	for i := 0; i <= 2; i++ {
		for j := 0; j <= 3; j++ {
			if grille[i][j] == symbole &&
				grille[i+1][j+1] == symbole &&
				grille[i+2][j+2] == symbole &&
				grille[i+3][j+3] == symbole {
				return true
			}
		}
	}
	return false
}

// GrillePleine vérifie si la grille est complètement remplie
func GrillePleine() bool {
	mu.Lock()
	defer mu.Unlock()
	for i := 0; i < 6; i++ {
		for j := 0; j < 7; j++ {
			if grille[i][j] == "." {
				return false
			}
		}
	}
	return true
}
