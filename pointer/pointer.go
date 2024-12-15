package main

import "fmt"

func main() {
	fmt.Println("pointer")
	x := 5
	x += 5
	fmt.Println(x) // 10
	// &x คือ การเข้าถึงตำแหน่งของตัวแปร x
	changValue(&x) // ส่งค่า x ไปที่ function changValue โดยใช้ & นำหน้าตัวแปร คือ การส่งค่าตำแหน่งของตัวแปรไปที่ function และเปลี่ยนค่าใน function จะเปลี่ยนค่าของตัวแปรที่ส่งมาด้วย
	fmt.Println(x) // 100

	// การสร้าง pointer
	var y *int = &x // ประกาศตัวแปร y ที่เป็น pointer ของ int
	// กำหนดค่าให้ y โดยใช้ & นำหน้าตัวแปร คือ การเข้าถึงตำแหน่งของตัวแปร
	fmt.Println(y)  // 0xc0000b6010 คือ ตำแหน่งของตัวแปร x
	fmt.Println(x)  // 100
	fmt.Println(*y) // 100 คือ ค่าที่อยู่ในตำแหน่งของ pointer นั้น
	*y = 1
	fmt.Println(*y) // 1
	fmt.Println(x)  // 1
	// ช้างบนเป็นการเปลี่ยนค่าของตัวแปร x โดยใช้ pointer ที่ชี้ไปที่ตัวแปร x ถ้าเปลี่ยนค่าของของ y จะเปลี่ยนค่าของ x ด้วย เพราะว่ามันเป็นการ copy by reference (pass by value) (มันเก็บที่อยู่ของตัวแปร x และ y ไว้ในตำแหน่งเดียวกัน)

	xx := 5
	yy := xx
	yy = 10
	fmt.Println(xx) // 5
	fmt.Println(yy) // 10
	// ช้างบนเป็นการเปลี่ยนค่าของตัวแปร yy โดยไม่ใช้ pointer ถ้าเปลี่ยนค่าของตัวแปร yy จะไม่เปลี่ยนค่าของตัวแปร xx เพราะว่ามันเป็นการ copy by value

	//
	//
	emp := Employee{Name: "John", Salary: 1000}
	emp.Name = "Hunter"
	fmt.Println(emp.Name) // Hunter
	emp.ChangeName("Doe") // การเปลี่ยนค่าของ struct โดยใช้ pointer จะเปลี่ยนค่าของ struct นั้นๆ
	fmt.Println(emp.Name) // Doe

	per := Person{Name: "John", Salary: 1000}
	ChangeSalary(&per, 2000) // การเปลี่ยนค่าของ struct โดยใช้ pointer จะเปลี่ยนค่าของ struct นั้นๆ
	fmt.Println(per.Salary)  // 2000

	//
	//
	test := t{name: "John"}
	res, i := test.changeName("Doe")
	fmt.Println(res, i) // name 123123
}

// pass by reference
func changValue(prt *int) { // *int คือ การรับค่าที่เป็น pointer ของ int คือ การรับค่าที่อยู่ในตำแหน่งของ pointer นั้น
	*prt = 100
	// *prt คือ การเข้าถึงค่าที่อยู่ในตำแหน่งของ pointer นั้น
}

type Employee struct {
	Name   string
	Salary int
}

func (e *Employee) ChangeName(newName string) {
	e.Name = newName
}

type Person struct {
	Name   string
	Salary int
}

func ChangeSalary(e *Person, salary int) {
	e.Salary = salary
}

type t struct {
	name string
}

func (t t) changeName(name string) (string, int) {
	fmt.Println("changeName", name)
	return "name", 123123
}
