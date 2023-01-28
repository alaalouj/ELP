package main

import (
	"os"
    "bufio"
    "fmt"
    "net"
    "strconv"
    "strings"
)

func main() {
    // Connect to the server
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer conn.Close()

    // Prompt the user to enter matrix A
    fmt.Println("Enter matrix A (one row per line, values separated by spaces):")
    a := [][]int{}
    scanner := bufio.NewScanner(os.Stdin)
    for i := 0; scanner.Scan(); i++ {
        row := scanner.Text()
		if row == "end" {
			t := []int{999}
			a = append(a, t)
			break
		}
		row = strings.TrimSpace(row)
        values := strings.Split(row, " ")
        var intValues []int
        for _, value := range values {
            intValue, err := strconv.Atoi(value)
            if err != nil {
				// il ne faut pas entrer d'espace après le dernier nombre d'une ligne de matrice
                fmt.Println("error:", err)
                return
            }
            intValues = append(intValues, intValue)
        }
        a = append(a, intValues)
    }
	

    // Send matrix A to the server
    writeMatrix(conn, a)
	a = append(a[:len(a)-1])
	

    // Prompt the user to enter matrix B
    fmt.Println("Enter matrix B (one row per line, values separated by spaces):")
    b := [][]int{}
    scanner = bufio.NewScanner(os.Stdin)
    for i := 0; scanner.Scan(); i++ {
        row := scanner.Text()
		
		if row == "end" {
			t := []int{999}
			b = append(b, t)
			break
		}
		row = strings.TrimSpace(row)
        values := strings.Split(row, " ")
        var intValues []int
        for _, value := range values {
            intValue, err := strconv.Atoi(value)
            if err != nil {
                fmt.Println("Error:", err)
                return
            }
            intValues = append(intValues, intValue)
        }
        b = append(b, intValues)
    }
	

    // Send matrix B to the server
    writeMatrix(conn, b)
	b = append(b[:len(b)-1])

    // Read the result from the server
    c, err := readMatrix(conn)

    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Print the result
    fmt.Println("Result:")
    for _, row := range c {
        for _, value := range row {
            fmt.Print(value, " ")
        }
        fmt.Println()
    }
}

// readMatrix reads a matrix from a connection
func readMatrix(conn net.Conn) ([][]int, error) {
    matrix := [][]int{}
    scanner := bufio.NewScanner(conn)
    for scanner.Scan() {
        row := scanner.Text()
		row = strings.TrimSpace(row)
        values := strings.Split(row, " ")
        var intValues []int
        for _, value := range values {
            intValue, err := strconv.Atoi(value)
            if err != nil {
                return nil, err
            }
            intValues = append(intValues, intValue)
        }
        matrix = append(matrix, intValues)
    }
    if err := scanner.Err(); err != nil {
        return nil, err
    }
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
