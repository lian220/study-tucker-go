package main

import "fmt"

func main() {
	i := 0
	for i < 10 {
		fmt.Print(i, ", ")
		i++
	}
	fmt.Print(i)
}
