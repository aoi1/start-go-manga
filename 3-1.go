package main

import "fmt"

type People struct {
	Name string
	Age  int
}

func main() {
	var people People = People{"mayu", 28}
	fmt.Println(people.Name)
	fmt.Println(people.Age)
}
