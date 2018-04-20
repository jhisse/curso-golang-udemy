package main

import (
	"fmt"
)

func troca(p1, p2 int) (segundo int, primeiro int) {
	segundo = p2
	primeiro = p1
	return
}

func main() {
	x, y := troca(1, 2)
	fmt.Println(x)
	fmt.Println(y)

	primeiro, segundo := troca(5, 7)
	fmt.Println(primeiro)
	fmt.Println(segundo)
}
