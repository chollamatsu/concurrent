package main

import (
	"fmt"
	"time"
)

func player(name string, c chan string, status chan bool) {

	for i := 0; 1 < 5; i++ {
		if i == 4 {
			status <- true
			break
		}
		action := <-c
		fmt.Println(name, action)
		time.Sleep(1 * time.Second)

		switch action {
		case "ping":
			c <- "pong"
		default:
			c <- "ping"
		}
	}

}

func main() {
	c := make(chan string)
	status := make(chan bool)
	// c <- "ping" //why this 'c' is deadlock!
	go player("Bob: ", c, status)
	go player("Alice: ", c, status)
	c <- "ping"
	for {
		if <-status {
			// fmt.Println(<-status)
			break
		}
	}
}
