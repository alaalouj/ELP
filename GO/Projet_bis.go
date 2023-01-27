package main

/*
import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Verifie si erreur lors de la lecture d'un fichier
func check(e error) {
	if e != nil {
		panic(e)
	}
}
func strToInt(txt string) int {
	val, err := strconv.Atoi(txt)
	check(err)
	return val
}

// Lire une matrice a partir d'un fichier
func ReadMatrix(fichier string) [][]int {
	data, err := os.ReadFile(fichier)
	lignes := string(data)
	check(err)
	test := strings.Split(lignes, "\r")
	nbLignes := len(test)
	stringA := make([][]string, nbLignes)
	for i := 0; i < nbLignes; i++ {
		stringA[i] = strings.Split(test[i], " ")
	}
	nbColonnes := len(stringA[0])
	a := make([][]int, nbLignes)
	for j := 0; j < nbLignes; j++ {
		a[j] = make([]int, nbColonnes)
		for k := 0; k < nbColonnes; k++ {
			a[j][k] = strToInt(stringA[j][k])
		}
	}
	return a
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
	//a := [][]int{[]int{1, 2, 3}, []int{3, 2, 1}, []int{2, 1, 1}}
	//b := [][]int{[]int{1, 1, 5}, []int{1, 2, 4}, []int{1, 2, 4}}
	fmt.Print(ReadMatrix("MatriceA.txt"))
}
*/
