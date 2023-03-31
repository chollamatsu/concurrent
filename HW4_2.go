/* The Barbershop (Originated by Dijkstra)

ร้านตัดผมมีเก้าอี้ n ที่นั่ง ใช้นั่งตัดผม 1 ที่นั่ง

-กรณีไม่มีลูกค้า ช่างตัดผมจะนอนหลับ

-กรณีที่ลูกค้าเข้าร้านและเก้าอี้นั่งรอเต็ม ลูกค้าจะออกจากร้านไป

-กรณีช่างตัดผมกำลังทำงาน แต่เก้าอี้นั่งรอว่าง ลูกค้าจะนั่งรอที่เก้าอี้ที่ว่าง

-กรณีที่ลูกค้าเข้าร้านขณะที่ช่างตัดผมหลับ ลูกค้าจะปลุกช่างตัดผม

-กรณีลูกค้าตัดผมเสร็จแล้ว และที่นั่งตัดผมว่าง ช่างตัดผมจะเรียกลูกค้าที่นั่งรอมาตัดผมต่อไปโดยไม่ไปนอนหลับ */
/* Hint:

/ Customer threads should invoke a function named getHairCut.
- If a customer thread arrives when the shop is full, it can invoke balk(หยุดชะงัก,หยุดนิ่ง), which does not return. ->if the shop is full, does not return
/ The barber thread should invoke cutHair.
- When the barber invokes cutHair there should be exactly one thread invoking getHairCut --> แสดงว่า cuthair & get haircut เป็น go func
/ customers counts the number of customers in the shop; it is protected by mutex. --> ให้invoke func coustomer couting
- The barber waits on customer until a customer enters the shop, then the customer waits on barber until the barber signals him to take a seat.-->barberตัดผมเสร็จ ส่งsignal customer คนต่อไป
- After the haircut, the customer signals customerDone and waits on barberDone. -->sinals: customerDone&barberDone
*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var mutex sync.Mutex
var wg sync.WaitGroup

//customerDone

func getHairCut(count chan int, barberDone chan bool, customerDone chan bool) {
	// fmt.Println("customer invoke get hair cut")
	if len(count) < cap(count) {
		// fmt.Println("customer wait list:", len(count))
		customerCounting(count)
	}
}

func cutHair(count chan int, barberDone chan bool, customerDone chan bool) {
	// fmt.Println("barber invoke cut hair")

	for i := 0; i < 5; i++ { //ทำจนลูกค้าหมด
		time.Sleep(time.Duration(rand.Int63n(1e9)))
		if len(count) > 0 {
			fmt.Println("baber wait list:", len(count))
			// mutex.Lock()
			<-count
			// mutex.Unlock()
			fmt.Println("ตัดผมเสร็จ")
		}
	}
	wg.Done()
}

func customerCounting(count chan int) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Duration(rand.Int63n(1e9)))
		if len(count) < cap(count) {
			count <- 1 //เพิ่มลูกค้าเข้าbuffer
			fmt.Println("ลูกค้าเข้าร้าน")
		}
	}
}

func main() {
	wg.Add(1)
	count := make(chan int, 5)
	barberDone := make(chan bool)
	customerDone := make(chan bool)

	// for i := 0; i < 5; i++ {
	go getHairCut(count, barberDone, customerDone)
	go cutHair(count, barberDone, customerDone)
	// }

	wg.Wait()

	fmt.Println("after ของที่อยู่ใน buffer =", len(count))
}
