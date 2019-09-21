package main

import "fmt"

type People struct {
	Name string
	Age  int
}

func Countup(p *People) {
	p.Age = p.Age + 1
}

func main() {
	var people *People = &People{"mayu", 28}
	Countup(people)
	fmt.Println(people.Age)
}
