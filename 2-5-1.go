package main

import "fmt"

func add(x int, y int) int {
	return x + y
}
func main() {
	result := add(1, 5)
	fmt.Println(result)
}
