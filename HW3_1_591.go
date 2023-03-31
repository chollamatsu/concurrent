// HW3_1: ออกแบบ Bounded Buffer Problem โดยใช้ Channel
// แทนการใช้ Mutex

package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var m sync.Mutex

var buffer_chan = make(chan []byte, 10)
var all sync.WaitGroup

func main() {
	rand.Seed(10)
	all.Add(1)
	go writer('a')
	go writer('b')
	go consumer()
	all.Wait()
}

func writer(c byte) {
	var buffer = make([]byte, 0, 10)

	for i := 0; i < 5; i++ { //ถ้าไม่มีforนี้มันจะไม่ออกลูปconsumer(ถ้าaกับbหมดแล้วนะ) i ของ consumerจะกลายเป็น 1 ตลอดแล้วก็ไม่ออกลูป
		time.Sleep(time.Duration(rand.Int63n(1e9)))
		// m.Lock()

		lb := len(buffer)
		if len(buffer_chan) == 0 && len(buffer) == 0 {
			buffer = buffer[:1]
			buffer[0] = c
			buffer_chan <- buffer
		} else {
			buffer = <-buffer_chan
			lb = len(buffer) //re-assign new current index -->essential!
		}
		if lb < cap(buffer) { //len buffer less than buffer หาว่าbuffer เต็มยัง
			buffer = buffer[:lb+1] //increase len buffer to contain the new character.
			buffer[lb] = c         //push c to buffer -->pop first a & first b -->จะเหลือmiddle
			buffer_chan <- buffer
			fmt.Printf("'%c' written to buffer.     buffer contents: %s\n", c, string(buffer))
		}
		// m.Unlock()
	}
}

func consumer() { //consumer จะทำ ไม่เกินแค่ 5 ครั้ง เท่านั้น ครบ 5 แล้วมันจะออกloop -> all.doneจบการทำงานเลย
	buffer := make([]byte, 0, 10)

	a := []byte{'a'}
	b := []byte{'b'}
	for i := 0; i < 5; {
		time.Sleep(time.Duration(rand.Int63n(1e9)))
		// m.Lock()
		buffer = <-buffer_chan
		ai := bytes.Index(buffer, a) //get the first index of 'a'
		bi := bytes.Index(buffer, b) //get the first index of 'b'
		if ai >= 0 && bi >= 0 {
			if ai > bi {
				ai, bi = bi, ai
			}
			copy(buffer[bi:], buffer[bi+1:])
			copy(buffer[ai:], buffer[ai+1:])
			buffer = buffer[:len(buffer)-2]
			buffer_chan <- buffer
			fmt.Printf("pair removed from buffer.  buffer contents: %s\n", string(buffer))
			i++
		}
		// m.Unlock()
	}
	all.Done()
}
