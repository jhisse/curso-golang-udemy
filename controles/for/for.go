package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for i < 5 { // não há while em go
		fmt.Println(i)
		i++
	}

	for j := 0; j <= 20; j += 2 {
		fmt.Println(j)
	}

	for { //laço infinito
		fmt.Println("Loop infinito")
		time.Sleep(time.Second)
	}
}
