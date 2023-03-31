package main

import (
	"fmt"
	"time"
)

var result bool

func counter_num(n int, sum_count chan int) {
	x := <-sum_count
	// fmt.Println("start at ", x)
	for i := 0; i < 200; i++ {
		x = x + 1
	}

	go func() {
		sum_count <- x
	}()

	fmt.Println("From", n, ":", x)
}

func main() {
	sum_count := make(chan int)

	go func() {
		sum_count <- 0
	}()

	for i := 0; i < 5; i++ {
		go counter_num(i, sum_count)
	}
	//(noteเอาไว้ถามอาจารย์อีกที)ที่ใส่time sleep ไว้ตรงนี้เพราะตอนที่ไม่ใส่ตัวที่นับสุดท้ายไม่แสดงผลเลยแล้วก็ทำให้sumนับถึงแค่800
	//คิดว่าthreadรองสุดท้ายกับthreadน่าจะเสร็จในเวลาเดียวกันพร้อมกัน แต่ตัวcommandไม่สามารถแสดงผลลัพธ์พร้อมกันได้ เลยหน่วงให้มันแสดงผล แต่ถ้าเป็นแบบนั้นทำไมตอนไม่ใส่sumถึงได้800ทั้งที่ควรเป็น1000
	//เพิ่มเติม:
	//เคยลองใส่time sleep ใน f-counter_num อยู่ -->แต่เจอdead lockเป็นบางครั้ง
	//เคยลองใส่ในloopบนนี้เหมือนกัน แต่ออกมาหน้าตามันเหมือนsequentialเลย ก็เลยย้ายออกมาไว้ตรงนี้
	time.Sleep(1 * time.Second)
	fmt.Println("Summary: ", <-sum_count)
}
