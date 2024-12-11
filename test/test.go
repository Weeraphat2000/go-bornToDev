package test

import (
	"fmt"
)

func PrintTest() { // ถ้าจะ export ออกไปให้ข้างนอกใช้งานได้ต้องขึ้นต้นด้วยตัวใหญ่
	fmt.Println("test")
	test2()
}

func test2() { // ขึ้นต้นด้วยตัวเล็ก ไม่สามารถ export ออกไปให้ข้างนอกใช้งานได้
	fmt.Println("test2")
}
