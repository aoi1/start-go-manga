package main

import "fmt"

func swap(x, y int) (int, int) {
	return y, x
}
func main() {
	result, _ := swap(1, 5)
	fmt.Println(result)
}
