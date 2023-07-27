package main

import "fmt"

func main() {
	var a int = 500
	var p *int

	p = &a

	fmt.Printf("p의 값: %p\n", p)
	fmt.Printf("p의 메모리의 값 : %d\n", *p)

	*p = 100
	fmt.Printf("a의 값: %d\n", a)
	fmt.Printf("p의 메모리의 값 : %d\n", *p)
}
