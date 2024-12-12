package main

import (
	"fmt"
	"math"
)

type Name interface {
	area()
	perim()
}

type rectangular struct {
	width, height int
}
type circle struct {
	radius int
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

	fmt.Println("Rectangle"+" "+"Area:", r.area())
	fmt.Println("Rectangle"+" "+"Perimeter:", r.perim())

	fmt.Println("Circle"+" "+"Area:", c.area())
	fmt.Println("Circle"+" "+"Perimeter:", c.perim())

	PrintValue("Hello")
	PrintValue(true)
	PrintValue(123)
	PrintValue(123.45)

	a, b := any(123123).(float64)
	fmt.Println(a, b)

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
