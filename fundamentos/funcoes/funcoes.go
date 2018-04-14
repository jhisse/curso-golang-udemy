package main

import (
	"fmt"
)

func somar(x int, y int) int { 
	return x + y
}

func main() {
	x := 10
	y:= 32

	fmt.Println("Somando", x, "+", y, "=", somar(x, y))
}