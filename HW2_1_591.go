package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var sum = 0
var mu sync.Mutex

func counter(n int, wg *sync.WaitGroup) {
	for i := 0; i < 10000; i++ {
		//note: ที่เลือกlockตรงนี้เพราะคิดว่าน่าจะต้องlocklส่วนที่เป็น atomic แล้วคิดว่า sum+=1เป็นส่วน atomic
		mu.Lock()
		sum = sum + 1
		mu.Unlock()
	}
	fmt.Println("From", n, ":", sum)
	wg.Done()

}

func main() {

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go counter(i, &wg)
	}

	wg.Wait()
	fmt.Println("Summary: ", sum)
}
