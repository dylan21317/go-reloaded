package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . entree.txt sortie.txt")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Erreur de lecture du fichier %s : %v\n", inputFile, err)
		return
	}

	processedText := Process(string(content))

	err = os.WriteFile(outputFile, []byte(processedText), 0644)
	if err != nil {
		fmt.Printf("Erreur d'écriture du fichier %s : %v\n", outputFile, err)
		return
	}

	fmt.Println("text modifier")
}

//test
