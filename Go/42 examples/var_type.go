package main

import "fmt"

func main(){

  var(
    i = 21
    j = 'A'
    g = "00"
  )
  fmt.Printf("i do tipo %T com valor %d\n", i, i)
  fmt.Printf("j do tipo %T com valor %c\n", j, j)
  fmt.Printf("g do tipo %T com valor %s\n", g, g)
}
