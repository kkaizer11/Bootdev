package main

// import "fmt"

func ft_swap(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

// func main(){
//   var(
//     x = 10
//     y = 42
//   )
//   fmt.Printf("O valor de x:%d e o de y:%d\n", x, y)
//   ft_swap(&x, &y)
//   fmt.Printf("O novo valor de x:%d e o de y:%d\n", x, y)

// }
