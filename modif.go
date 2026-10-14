package main

import (
	"strconv"
	"strings"
)

// Process prend le texte brut en entrée et applique toutes les modifications demandées
func Process(text string) string {
	// Découpe le texte en une liste de mots en se basant sur les espaces
	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {
		// Gestion de la conversion hexadécimale (hex)
		if words[i] == "(hex)" && i > 0 {
			// Convertit le mot précédent de l'hexadécimal (base 16) vers le décimal (base 10)
			val, err := strconv.ParseInt(words[i-1], 16, 64)
			if err == nil {
				words[i-1] = strconv.FormatInt(val, 10)
				// Supprime le tag "(hex)" du tableau de mots
				words = append(words[:i], words[i+1:]...)
				i-- // Recule l'index pour s'adapter à la suppression
			}
			// Gestion de la conversion binaire (bin)
		} else if words[i] == "(bin)" && i > 0 {
			// Convertit le mot précédent du binaire (base 2) vers le décimal (base 10)
			val, err := strconv.ParseInt(words[i-1], 2, 64)
			if err == nil {
				words[i-1] = strconv.FormatInt(val, 10)
				// Supprime le tag "(bin)" du tableau de mots
				words = append(words[:i], words[i+1:]...)
				i--
			}
		}
	}

	// Reassemble tous les mots de la liste en une seule chaîne de caractères séparée par des espaces
	return strings.Join(words, " ")
}
