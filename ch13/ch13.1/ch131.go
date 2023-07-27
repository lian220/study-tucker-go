package main

import "fmt"

type House struct {
	Address  string
	Size     int
	Price    float64
	Category string
}

func main() {
	var house House
	house.Address = "seoul"
	house.Size = 25
	house.Price = 10
	house.Category = "apartment"

	fmt.Printf("주소:%s 사이즈:%d평 가격:%v억원 종류:%s\n", house.Address, house.Size, house.Price, house.Category)
}
