package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	for i := 0; i < 5; i++ {
		go sayHello(i, c)
	}
	for i := 0; i < 5; i++ {
		id := <-c
		fmt.Println("said hello to ", id)
	}
}

func sayHello(id int, c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hi ", id)
	c <- id
}
