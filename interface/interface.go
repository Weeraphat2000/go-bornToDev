package main

import (
	"fmt"
	"math"
	"strconv"
)

type Geometry interface {
	area()
	perim()
}

type rectangular struct {
	width, height int
}
type circle struct {
	radius int
}

func area(r rectangular, unit string) (int, string) {
	return r.width * r.height, unit
}

func (r rectangular) area() int {
	return r.width * r.height
}

func (r rectangular) perim() int {
	return 2*r.width + 2*r.height
}

func (c circle) area() int {
	return int(math.Pi * float64(c.radius) * float64(c.radius))
}
func (c circle) perim() int {
	return int(2 * math.Pi * float64(c.radius))
}

func main() {
	r := rectangular{width: 10, height: 5}
	c := circle{radius: 5}
	fmt.Println("r", r)
	fmt.Println("Rectangle"+" "+"Area:", r.area())

	areaValue, unit := area(r, "cm")
	fmt.Println("Area:", areaValue, unit)

	fmt.Println("Rectangle"+" "+"Perimeter:", r.perim())

	fmt.Println("Circle"+" "+"Area:", c.area())
	fmt.Println("Circle"+" "+"Perimeter:", c.perim())

	PrintValue("Hello")
	PrintValue(true)
	PrintValue(123)
	PrintValue(123.45)

	a, b := any(123123).(float64)
	fmt.Println(a, b) // ถ้า b เป็น false แสดงว่าไม่สามารถแปลงเป็น float64 ได้และ a จะเป็น 0 ถ้า b เป็น true แสดงว่าสามารถแปลงเป็น float64 ได้ และ a จะเป็นค่าที่แปลงเป็น float64

	hun := bicycle{brand: "HUN", price: 1000, amount: 10}
	fmt.Println("Brand:", hun.Brand())
	fmt.Println("Amount:", hun.Bmount())
	fmt.Println("Total Price:", hun.totalPrice("THB"))

	//
	//
	//
	//
	//
	// mikelopster
	dog := Dog{Name: "Dog"}
	person := Person{Name: "Person"}
	makeSound(dog) // makeSound(dog) จะเรียกใช้ method Speak ของ struct Dog, makeSound จะเช็คว่า dog มี method Speak หรือไม่ ถ้ามีจะเรียกใช้ method Speak ของ struct Dog
	// fmt.Println("Dog:", dog.Speak())

	fmt.Println("Person:", person.Speak())
	makeSound(person) // makeSound(person) จะเรียกใช้ method Speak ของ struct Person, makeSound จะเช็คว่า person มี method Speak หรือไม่ ถ้ามีจะเรียกใช้ method Speak ของ struct Person

	walk(dog)
	fmt.Println("Dog:", dog.Walk())

}

type TestInterface interface {
	~string | ~bool | ~int | ~float64
}

// ฟังก์ชัน Generic ที่รองรับ TestInterface
func PrintValue[T TestInterface](value T) {
	fmt.Println(value)
	// ถ้าเข้ามาเป็น string จะทำงานในส่วนนี้
	switch v := any(value).(type) {
	case string:
		fmt.Println("String:", v)
	// ถ้าเข้ามาเป็น bool จะทำงานในส่วนนี้
	case bool:
		fmt.Println("Boolean:", v)
	// ถ้าเข้ามาเป็น int จะทำงานในส่วนนี้
	case int:
		fmt.Println("Int:", v)
	// ถ้าเข้ามาเป็น float64 จะทำงานในส่วนนี้
	case float64:
		fmt.Println("Float:", v)
	// ถ้าเข้ามาเป็น type อื่นๆ จะทำงานในส่วนนี้
	default:
		fmt.Println(v)
	}
}

type bicycle struct {
	brand  string
	price  int
	amount int
}

// การเขียน method ใน struct เหมือนกับการเขียน function แต่ต้องขึ้นต้นด้วยชื่อ struct ที่ต้องการเขียน method ตามด้วยชื่อ method และ parameter ที่ต้องการ และ return type ที่ต้องการ
func (b bicycle) totalPrice(unit string) string {
	fmt.Println("Unit:", unit)
	return strconv.Itoa(b.price*b.amount) + " " + unit
}

func (b bicycle) Bmount() int {
	return b.amount
}

func (b bicycle) Brand() string {
	return b.brand
}

// mikelopster
// speaker interface
// ถ้า struct ต้องการ implement interface จะต้องมี method ที่กำหนดใน interface ทุกตัว
type Speaker interface {
	Speak() string // method ที่ต้องมีใน struct ที่จะ implement interface นี้ได้
	Walk() string
}

// dog struct
type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return d.Name + " says Woof!"
}

func (d Dog) Walk() string {
	return d.Name + " is walking"
}

type Person struct {
	Name string
}

func (p Person) Speak() string {
	return p.Name + " says Hello!"
}
func (p Person) Walk() string {
	return p.Name + " is walking"
}

func makeSound(s Speaker) {
	fmt.Println(s.Speak())
}
func walk(s Speaker) {
	fmt.Println(s.Walk())
}
