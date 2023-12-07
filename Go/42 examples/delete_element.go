// package main

// import "fmt"

func delete_element(numbers []int, index int) []int {
	numbers = append(numbers[:index], numbers[index+1:]...)
	return numbers
}

// func main() {
// 	nums := []int{0, 1, 2, 3, 4, 5, 6}
// 	fmt.Println("Antes da função", nums)
// 	delete_element(nums, 2)
// 	fmt.Println("Após a função", nums)
// }
