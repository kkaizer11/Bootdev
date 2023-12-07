package main

import "fmt"

func ft_reverse(s string) string {
	runed := []rune(s)
	for i := 0; i < len(s)/2; i++ {
		runed[i], runed[len(s)-1-i] = runed[len(s)-1-i], runed[i]
	}
	return string(runed)
}

func main() {
	x := ft_reverse("-121")
	fmt.Println(x)
}
