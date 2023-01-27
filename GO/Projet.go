package main

import (
	"fmt"
	"os"
	"sync"
)

// Récuperation d'une matrice dans un fichier texte
/*
func TxtToMatrix(fichier string) [][]int {
	data, err := ioutil.ReadFile(fichier)
	lines := string(data)
	if err != nil {
		fmt.Println(err)
		matrix := make([][]int, 1)
		strings.Split(lines, "\n")

	}

}*/

// Fonction goroutine :
func multiply(a, b, ab [][]int, row, col int, wg *sync.WaitGroup) {
	//Multiplication de a et b
	defer wg.Done()
	sum := 0
	for i := 0; i < len(a[0]); i++ {
		sum += a[row][i] * b[i][col]
	}
	ab[row][col] = sum
}

// Multiplication d'un matrice a de taille m*n avec une matrice b de taille p*q
func MultiplicationMatrix(a, b [][]int) {
	m := len(a) //nombre de lignes __
	n := len(a[0])
	p := len(b)
	q := len(b[0]) // ||
	//Verification que les matrices a et b aient des tailles compatibles
	if n != p {
		fmt.Println("Erreur de taille")
		os.Exit(1)
	}
	//Création de la matrice a*b
	ab := make([][]int, m)
	for ligne := range ab {
		ab[ligne] = make([]int, q)
	}

	// en utilisant les goroutines :

	var wg sync.WaitGroup
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b[0]); j++ {
			wg.Add(1)
			go multiply(a, b, ab, i, j, &wg)
		}
	}
	wg.Wait()
	fmt.Println("Result: ", ab)

}

func main() {
	a := [][]int{{1, 2, 3}, {3, 2, 1}, {5, 3, 2}}
	b := [][]int{{1, 1, 1}, {1, 2, 5}, {1, 2, 5}}
	MultiplicationMatrix(a, b)
}
