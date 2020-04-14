package main

import (
	"fmt"
	"time"
)

func throw(c chan int) {
	for i := 0; i < 5; i++ {
		c <- i
		fmt.Println("threw >>", i)
	}
}

func catch(c chan int) {
	for i := 0; i < 5; i++ {
		num := <-c
		fmt.Println("caught <<", num)
	}
}

func main() {
	c := make(chan int, 3)
	go throw(c)
	go catch(c)

	time.Sleep(10 * time.Millisecond)
}
