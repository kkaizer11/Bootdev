package main

import "fmt"

func ft_print_numbers() {
	var number = 0

	for number <= 9 {
		fmt.Printf("%d", number)
		number += 1
	}
}
