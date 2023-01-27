package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Récuperation d'une matrice dans un fichier texte
func TxtToMatrix(fichier string) [][]int {
	data, err := ioutil.ReadFile(fichier)
	lines := string(data)
	if err != nil {
		fmt.Println(err)
		matrix := make([][]int, 1)
		strings.Split(lines, "\n")

	}

}

// Multiplication d'un matrice a de taille m*n avec une matrice b de taille p*q
func MultiplicationMatrix(a [][]int, b [][]int) [][]int {
	m := len(a)
	n := len(a[0])
	p := len(b)
	q := len(b[0])
	//Verification que les matrices a et b aient des tailles compatibles
	if n != p {
		fmt.Println("Erreur de taille")
		os.Exit(1)
	}
	//Création de la mtrice a*b
	ab := make([][]int, m)
	for ligne := range ab {
		ab[ligne] = make([]int, q)
	}
	//Multiplication de a et b
	for i := 0; i < m; i++ {
		for k := 0; k < n; k++ {
			for j := 0; j < q; j++ {
				ab[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return ab
}

func main() {
	a := [][]int{[]int{1, 2, 3}, []int{3, 2, 1}, []int{2, 1, 1}}
	b := [][]int{[]int{1, 1}, []int{1, 2}}
	fmt.Println(MultiplicationMatrix(a, b))
}
