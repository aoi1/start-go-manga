package main

import (
	"fmt"
)

type Speaker interface {
	Speak() string
}
type Dog struct {
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
}

func (c Cat) Speak() string {
	return "Meow!"
}
func main() {
	speakers := []Speaker{Dog{}, Cat{}}
	for _, speaker := range speakers {
		fmt.Println(speaker.Speak())
	}
}
