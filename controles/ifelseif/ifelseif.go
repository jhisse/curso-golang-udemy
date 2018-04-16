package main

import (
	"fmt"
)

func notaParaConceiro(nota float64) string {
	if nota >= 9 && nota <= 10 {
		return "A"
	} else if nota >=8 && nota < 9 {
		return "B"
	} else if nota >= 5 && nota < 8 {
		return "C"
	} else if nota >= 3 && nota < 5 {
		return "D"
	}
	return "E"
}

func main() {
	fmt.Println(notaParaConceiro(3.5))
	fmt.Println(notaParaConceiro(5.6))
	fmt.Println(notaParaConceiro(9.2))
}