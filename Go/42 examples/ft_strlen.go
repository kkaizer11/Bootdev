package main

import "fmt"

func ft_strlen(str string) {
	i := 0
	for range str {
		i += 1
	}
	fmt.Println(i)
}

// func main(){
//   g := "Banana"
//   ft_strlen(g)
// }
