package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 10)
	go count(5, c)
	for e := range c {
		fmt.Println(e)
	}
	fmt.Println("main終了")
}

func count(n int, c chan int) {
	for i := 0; i < n; i++ {
		c <- i
	}
	time.Sleep(1 * time.Second)
	fmt.Println("ループ終了")
	time.Sleep(5 * time.Second)
	close(c)
}
