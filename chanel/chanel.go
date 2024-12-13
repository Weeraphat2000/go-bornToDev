package main

import (
	"fmt"
	"time"
)

func main() {

	// chanel คือการสร้างช่องสื่อสารระหว่าง goroutine กับ goroutine อื่นๆ
	// การสร้าง chanel ใช้เครื่องหมาย <-

	go f("direct") // การสร้าง goroutine ใช้ go นำหน้า function ที่ต้องการสร้าง goroutine คือ การทำงานของ function นั้นจะไม่รอให้จบก่อน ฟังก์ชัน f จะทำงานแบบ asynchronous โดยไม่รอให้ฟังก์ชันนั้นเสร็จสิ้น
	go f("goroutine")
	time.Sleep(5 * time.Second)

	go func(msg string) {
		fmt.Println(msg)
	}("anonymous goroutine")

	ch := make(chan string, 1) // สร้าง chanel โดยใช้ make โดยกำหนดชนิดของข้อมูลที่จะส่งผ่าน chanel คือ string และกำหนดขนาดของ chanel คือ 1
	process(ch, "Hello")       // process คือ function ที่สร้างขึ้นเพื่อส่งข้อมูลผ่าน chanel

	fmt.Println("chanel", <-ch) // รับข้อมูลจาก chanel โดยใช้เครื่องหมาย <- โดยเมื่อรับข้อมูลจะรอจนกว่าจะมีข้อมูลถูกส่งผ่าน chanel มา
	time.Sleep(5 * time.Second)

}

func f(from string) {
	for i := 0; i < 10; i++ {
		fmt.Println(from, ":", i)
	}
}

func process(c chan string, data string) {
	c <- data // ส่งข้อมูลผ่าน chanel โดยใช้เครื่องหมาย <- โดยส่งค่า value ผ่าน chanel c
}
