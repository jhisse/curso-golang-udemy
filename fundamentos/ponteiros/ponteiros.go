package main

import (
	"fmt"
)

func main() {
	i := 1
	
	var p *int = nil

	p = &i
	*p++
	i++

	fmt.Println(i)
	fmt.Println(*p)
	fmt.Println(&i)
	fmt.Println(p)

}