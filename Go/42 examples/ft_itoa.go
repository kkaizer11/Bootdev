package main

import "fmt"

func itoa(num int) string {
	return fmt.Sprint(num)
}

func main() {
	s := itoa(-121)
	fmt.Printf("Printando o valor de %s que é o tipo: %T", s, s)
}
