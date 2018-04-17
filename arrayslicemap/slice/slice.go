package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} // array
	s1 := []int{1, 2, 3}  // slice
	fmt.Println(a1, s1)
	fmt.Println("A1:", reflect.TypeOf(a1))
	fmt.Println("S1:", reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5} // array
	s2 := a2[1:3]
	fmt.Println(a2, s2)

	s3 := a2[:2]
	fmt.Println(a2, s3)
	fmt.Println(&s3[1], &a2[1])

	s4 := s3[:2]
	fmt.Println(&s4[1], &s3[1], &a2[1])
}
