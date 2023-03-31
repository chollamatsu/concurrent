// main say "hello" goroutine say "world"
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func pong(wg *sync.WaitGroup) {
	fmt.Println("goroutine: world")
	wg.Done()
}
func main() {
	fmt.Println("main: hello")
	wg.Add(1)
	go pong(&wg)
	wg.Wait()
}
