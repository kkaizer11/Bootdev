package main

// import "fmt"

func ft_swap(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

func ft_bubble_sort(tab []int) []int {
	var (
		i    = 0
		j    = 0
		size = len(tab)
	)
	for i < size {
		i++ // pode ser assim
		for j < size-1-i {
			if tab[j] > tab[j+1] {
				ft_swap(&tab[j], &tab[j+1])
			}
			j += 1 // ou assim
		}
	}
	return tab
}

// TODO: Adicionar de fato a funçao random
/*
 func main() {
 	rand_num := []int{90, -5, 13, 58, 27}
 	ft_bubble_sort(rand_num)
 	fmt.Println(rand_num)
 }
*/
