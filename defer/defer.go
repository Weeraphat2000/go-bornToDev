package main

import "fmt"

func main() {
	// defer คือการทำงานของ function นั้นจะถูกเรียกใช้งานตอนที่ก่อนจะจบ function ที่เรียกใช้งาน defer นั้น
	fmt.Println("start")
	defer fmt.Println(add(10, 20))
	fmt.Println("end")

	f()

	//	start
	// end
	// 1
	// 3
	// 2
	// 4
	// 5
	// 30
}

func add(a, b int) int {
	return a + b
}

func f() {
	fmt.Println("1")
	defer s()
	defer fmt.Println("2")
	fmt.Println("3")
}

func s() {
	fmt.Println("4")
	fmt.Println("5")
}
