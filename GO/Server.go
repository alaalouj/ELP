package main

import (
    "bufio"
    "fmt"
    "net"
    "strconv"
    "strings"
    "os"
    "sync"
)

func main() {
    // Listen on port 8080
    listener, err := net.Listen("tcp", "localhost:8080")
    if err != nil {
        //fmt.Println("Error1:", err)
        return
    }
    defer listener.Close()

    // Accept connections
    for {
        conn, err := listener.Accept()
        if err != nil {
            //fmt.Println("Error2:", err)
            continue
        }
        go handleConnection(conn)
    }
}

// handleConnection handles a connection
func handleConnection(conn net.Conn) {
    defer conn.Close()

    // Read matrix A from the client
    fmt.Println("First Matrix:")
    a, err := readMatrix(conn)
    //fmt.Println("a")
    if err != nil {
        fmt.Println("Error3:", err)
        return
    }

    // Read matrix B from the client
    fmt.Println("Second Matrix:")
    b, err := readMatrix(conn)
    //fmt.Println("b")
    if err != nil {
        fmt.Println("Error4:", err)
        return
    }
    
    // Multiply the matrices
    c := multiplyMatrices(a, b)

    // Send the result to the client
    writeMatrix(conn, c)

}

// readMatrix reads a matrix from a connection
func readMatrix(conn net.Conn) ([][]int, error) {
    matrix := [][]int{}
    scanner := bufio.NewScanner(conn)
    for scanner.Scan() {
        row := scanner.Text()
        //fmt.Println(row)
        row = strings.TrimSpace(row)
        if row == "999" {
			break
		}
        
        values := strings.Split(row, " ")
        var intValues []int
        for _, value := range values {
            intValue, err := strconv.Atoi(value)
            //fmt.Println(intValue)
            if err != nil {
                //fmt.Println("a7")
                return nil, err
            }
            intValues = append(intValues, intValue)
            //fmt.Println(intValues)

        }
        matrix = append(matrix, intValues)
        
    }
    if err := scanner.Err(); err != nil {
        fmt.Println("2")
        return nil, err
    }
    fmt.Println(matrix)
    return matrix, nil
}

// writeMatrix writes a matrix to a connection
func writeMatrix(conn net.Conn, matrix [][]int) {
    for _, row := range matrix {
        for _, value := range row {
            fmt.Fprint(conn, value, " ")
        }
        fmt.Fprintln(conn)
    }
}

// multiplyMatrices multiplies two matrices
func multiplyMatrices(a, b [][]int) [][]int {
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
    return ab
}

func multiply(a, b, ab [][]int, row, col int, wg *sync.WaitGroup) {
	//Multiplication de a et b
	defer wg.Done()
	sum := 0
	for i := 0; i < len(a[0]); i++ {
		sum += a[row][i] * b[i][col]
	}
	ab[row][col] = sum
}
