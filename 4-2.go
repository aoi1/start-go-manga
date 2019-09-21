package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		go sayHello(i)
	}
	time.Sleep(4 * time.Second)
}
func sayHello(id int) {
	time.Sleep(3 * time.Second)
	fmt.Println("hello", id)
}
