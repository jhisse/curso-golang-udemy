package main

import (
	"fmt"
)

func notaParaConceito(n float64) string {
	var nota = int(n)
	switch nota {
	case 10:
		fallthrough // continua para o próximo case
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota Inválida"
	}
}

func main() {
	fmt.Println(notaParaConceito(10))
	fmt.Println(notaParaConceito(9.4))
	fmt.Println(notaParaConceito(4))
	fmt.Println(notaParaConceito(11))
	fmt.Println(notaParaConceito(2))
	fmt.Println(notaParaConceito(6.4))

}