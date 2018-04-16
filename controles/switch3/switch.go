package main

import (
	"fmt"
	"time"
)

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "função"
	default:
		return "desconhecido"
	}
}

func main() {
	fmt.Println(tipo(2.3))
	fmt.Println(tipo(time.Now()))
	fmt.Println(tipo(2))
	fmt.Println(tipo("Olá"))
	fmt.Println(tipo(4))
}
