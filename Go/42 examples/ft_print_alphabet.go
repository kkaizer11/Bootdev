package main

import "fmt"

func ft_print_alphabet() {
	var letter = 'a'

	for letter <= 'z' {
		fmt.Printf("%c", letter)
		letter += 1
	}
}
