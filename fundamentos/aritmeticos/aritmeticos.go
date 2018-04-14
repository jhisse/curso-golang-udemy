package main

import (
	"math"
	"fmt"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Soma =>", a + b)
	fmt.Println("Subtração =>", a - b)
	fmt.Println("Divisão =>", a / b) // resultado em inteiro
	fmt.Println("Multiplicação =>", a * b)
	fmt.Println("Módulo =>", a % b)

	// bitwise
	fmt.Println("AND =>", a & b) // 11 & 10 = 10
	fmt.Println("OR =>", a | b) // 11 | 10 = 11
	fmt.Println("XOR =>", a ^ b) // 11 ^ 10 = 10

	c:= 3.0
	d := 2.0

	fmt.Println("Maior =>", math.Max(float64(a), float64(b)))
	fmt.Println("Maior =>", math.Min(c, d))
	fmt.Println("Exponenciação =>", math.Pow(c, d))

}