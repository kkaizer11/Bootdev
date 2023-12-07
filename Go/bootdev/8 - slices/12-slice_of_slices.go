package main

import "fmt"

func createMatrix(rows, cols int) [][]int {
	matriz := [][]int{} // muita atenção com as chaves
	for i := 0; i < rows; i++ {
		linha := []int{}
		for j := 0; j < cols; j++ {
			linha = append(linha, i*j)
		}
		matriz = append(matriz, linha)
	}
	return matriz
}

// dont edit below this line

func test(rows, cols int) {
	fmt.Printf("Creating %v x %v matrix...\n", rows, cols)
	matrix := createMatrix(rows, cols)
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
	fmt.Println("===== END REPORT =====")
}

func main() {
	test(3, 3)
	test(5, 5)
	test(10, 10)
	test(15, 15)
}
