package main

import (
	"fmt"
)

func callerA(c chan string) {
	c <- "Hello World!"
	close(c)
}

func callerB(c chan string) {
	c <- "Hola Mundo!"
	close(c)
}

func main() {
	a, b := make(chan string), make(chan string)

	go callerA(a)
	go callerB(b)

	var msg string
	ok1, ok2 := true, true
	for ok1 || ok2 {

		select {
		case msg, ok1 = <-a:
			if ok1 {
				fmt.Printf("%s from A, ok1 is %v now\n", msg, ok1)
			}
			fmt.Println("ok1", ok1)
		case msg, ok2 = <-b:
			if ok2 {
				fmt.Printf("%s from B, ok2 is %v now\n", msg, ok2)
			}
			fmt.Println("ok2", ok2)
			// default:
			// 	fmt.Println("Default")
		}
	}
}
