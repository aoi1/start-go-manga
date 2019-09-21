package main

import (
	"fmt"
)

type Basket struct {
	X, Y int
}

func (b Basket) Sum() int {
	return b.X + b.Y
}

func main() {
	var b Basket = Basket{1, 2}
	fmt.Println(b.Sum())
}
