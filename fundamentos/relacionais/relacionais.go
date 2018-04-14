package main

import (
	"time"
	"fmt"
)

func main() {
	fmt.Println("Strings", "Maça" == "Maça")
	fmt.Println("!=", 3 != 2)
	fmt.Println("<", 3 < 2)
	fmt.Println(">", 3 > 2)
	fmt.Println("<=", 3 <= 2)
	fmt.Println(">=", 3 >= 2)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println("Datas:", d1 == d2) // Não compara a referência de memória
	fmt.Println("Datas:", d1.Equal(d2))

	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{"João"}
	p2 := Pessoa{"João"}
	fmt.Println("Pessoas:", p1 == p2) // Não leva em consideração o endereço de memória
	fmt.Println("Pessoas:", &p1 == &p2) // Compara se é o mesmo elemento na memória

}