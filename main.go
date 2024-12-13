// go run main.go เอาไว้ run ไฟล์ main.go
// go mod init สร้างไฟล์ go.mod ในโปรเจค (go.mod จะเก็บข้อมูลเกี่ยวกับโปรเจค)
// go build สร้างไฟล์ exe ในโปรเจค เอา file exe ไป run ได้

package main

import (
	"fmt"  // เอาไว้ print ข้อความ
	"math" // เอาไว้ใช้งานพวกคำนวณเลข (pow, sqrt, max, min และอื่นๆ)
	"strconv"
	"strings" // เอาไว้ใช้งานพวก string (uppercase, lowercase, contains, replace, split, join, length, index, substring, slice, map, if else, switch case และอื่นๆ

	"github.com/weeraphat2000/test"
)

// fmt is a package that provides I/O functions

func main() {

	// การใช้งาน any
	fmt.Println(" ")
	fmt.Println("any")
	// any คือ จะเอาไวว้เช็คว่า value, type ได้
	// any = interface{} คือ เป็น type อะไรก็ได้
	// any(ค่าที่ต้องการเช็ค).(type ที่ต้องการเช็ค) ถ้าเป็น type ที่ต้องการจะ return 2 ค่า คือ value และ true ถ้าไม่ใช่ type ที่ต้องการจะ return 2 ค่า คือ 0(type int), 0.0(type float64), ""(type string) และ false
	Value, Check := any(12312).(string)
	fmt.Println(Value, Check) // การใช้งาน type assertion

	fmt.Println("Hello, World!") // fp key ลัดเพื่อใช้ fmt.Println
	num := 10                    // ถ้าประกาศตัวแปรแบบนี้จะไม่สามารถกำหนด type ได้
	var num2 int = 20            // ถ้าประกาศตัวแปรแบบนี้จะสามารถกำหนด type ได้ หรือไม่ก็ไม่กำหนด type ก็ได้
	fmt.Println(num + num2)
	var num4, num5 int = 40, 50
	fmt.Println(num4 + num5)
	numssss := 10e3
	fmt.Println(numssss, "10e3")
	fmt.Println(float64(20) + 20.2222)
	num6, num7 := 30, 40
	fmt.Println(num6 % num7)
	fmt.Println(
		math.Pow(2, 3),
		math.Sqrt(16),
		math.Max(10, 20), // ได้ค่าที่มากที่สุด ได้แค่ 2 ค่า
		math.Min(10, 20), // ได้ค่าที่น้อยที่สุด ได้แค่ 2 ค่า
		math.Floor(20.8), // ปัดเศษลง
		math.Ceil(20.1),  // ปัดเศษขึ้น
		math.Abs(-10),    // ค่าสัมบูรณ์
		math.Round(20.5), // ปัดเศษ 0.5 ขึ้น
		math.Log10(100),  // ลอการิทึม ฐาน 10
		math.Log2(8),     // ลอการิทึม ฐาน 2
		math.Log(100),    // ลอการิทึม ฐาน e
		math.Exp(3),      // ค่า e ยกกำลัง 3

	)
	const pi float64 = 3.14159265359 // const คือค่าคงที่
	fmt.Println(pi)
	var name = "John"
	fmt.Println(name)
	fmt.Println("hello" + " " + "world")
	fmt.Println("age:", 30)
	fmt.Println("age: " + fmt.Sprint(30))
	fmt.Println("age: " + string(30))
	fmt.Printf("age: %d\n", 30)
	var isTrue bool = true
	fmt.Println(isTrue)

	// การใช้งาน string
	fmt.Println(" ")
	fmt.Println("string")
	var str string = "Hello, World!"
	fmt.Println(str)
	fmt.Println(strings.ToUpper(str), "string uppercase")
	fmt.Println(strings.ToLower(str), "string lowercase")
	fmt.Println(strings.Contains(str, "World"), "string contains")
	fmt.Println(strings.ReplaceAll(str, "World", "Golang"), "string replace")
	fmt.Println(strings.Split(str, ","), "string split")
	fmt.Println(strings.Join([]string{"Hello", "World"}, ", "), "string join")
	fmt.Println(len(str), "length")
	fmt.Println(str[0], "index")
	fmt.Println(str[2:5], "substring")
	fmt.Println(str[7:])
	fmt.Println(str[:5])
	fmt.Println(str[:])

	// การใช้งาน boolean
	fmt.Println(" ")
	fmt.Println("boolean")
	var isTrue2 bool = true
	var isFalse bool = false
	fmt.Println(isTrue2 && isFalse) // การใช้งาน and
	fmt.Println(isTrue2 || isFalse) // การใช้งาน or
	fmt.Println(!isTrue2)           // การใช้งาน not

	// operator
	fmt.Println(" ")
	fmt.Println("operator")
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 <= 2)

	// การใช้งาน array
	fmt.Println(" ")
	fmt.Println("array")
	var arr [3]int // การประกาศ array แบบไม่กำหนดค่า
	arr[0] = 10
	arr[1] = 20
	fmt.Println(arr, "array") // [10 20 0] 0 คือค่า default ของ int, float และ string คือ 0, 0.0 และ ""
	fmt.Println(arr[1])
	fmt.Println(len(arr))
	arr2 := [3]int{10, 20, 30} // การประกาศ array แบบกำหนดค่า
	fmt.Println(arr2)

	// การใช้งาน slice
	fmt.Println(" ")
	fmt.Println("slice")
	var slice []int = []int{1, 2, 3, 4, 5}
	fmt.Println(slice, "slice")
	fmt.Println(slice[1:3], "slice")
	fmt.Println(slice[2:])
	fmt.Println(slice[:3])             // [1 2 3]
	slice = append(slice, 6, 7, 8)     // การเพิ่มค่าใน slice [1 2 3 4 5 6 7 8]
	fmt.Println(slice, "slice append") // [1 2 3 4 5 6 7 8]
	newSlice := slice[1:3]
	fmt.Println(newSlice, "new slice")    // [2 3]
	newSlice[0] = 100                     // แก้ไขค่าใน slice จะส่งผลกระทบกับ slice ต้นฉบับด้วย เพราะ slice จะเก็บค่า reference ของ array ต้นฉบับ
	fmt.Println(slice)                    // [1 100 3 4 5 6 7 8]
	fmt.Println(newSlice)                 // [100 3]
	newSlice2 := make([]int, 3)           // การสร้าง slice โดยกำหนดความยาวของ slice
	copy(newSlice2, slice)                // การ copy ค่าจาก slice ต้นฉบ
	fmt.Println(newSlice2, "new slice 2") // [1 100 3]
	newSlice2[0] = 200                    // แก้ไขค่าใน newSlice2 จะไม่ส่งผลกระทบกับ slice ต้นฉบับ เพราะ newSlice2 จะเก็บค่า reference ของ array ใหม่
	fmt.Println(slice, "slice")           // [1 100 3 4 5 6 7 8]
	fmt.Println(newSlice, "new slice")    // [100 3]
	fmt.Println(newSlice2, "new slice 2") // [200 100 3]
	slice2 := []int{1, 2, 3, 4, 5, 6, 7}
	slice2 = append(slice2[:3], slice2[5:]...) // การลบค่าใน slice
	fmt.Println(slice2, "slice 2")
	list := []int{}
	fmt.Println(list, "list") // output: [] ถ้าไม่มีค่าให้แสดงค่าเป็น []
	var list2 []int = nil
	fmt.Println(list2, "list2") // output: [] ถ้าไม่มีค่าให้แสดงค่าเป็น []

	// การใช้งาน map
	fmt.Println(" ")
	fmt.Println("map")
	// map[keyType]valueType
	var m map[string]int = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	} // initial value
	fmt.Println(m, "map")
	fmt.Println(m["one"], "map value")
	m["four"] = 4 // เพิ่มค่าใน map โดยใช้ key กับ value
	fmt.Println(m)
	delete(m, "two") // delete ใช้ได้เฉพาะ map
	m["one"] = 10    // การแก้ไขค่าใน map
	fmt.Println(m, "map delete")
	fmt.Println(m["one"])
	value, isExist := m["two"] // การเช็คว่า key นั้นมีอยู่ใน map หรือไม่ ถ้ามีจะ return value กับ true ถ้าไม่มีจะ return 0 กับ false
	fmt.Println(value, isExist, "map check key")
	data := map[string]interface{}{
		"name":    "John",
		"age":     30,
		"isAdmin": true,
		"height":  1.75,
	}
	fmt.Println(data, "map interface")
	age, ok := data["age"].(int) // การเช็คว่า value นั้นมี type ตามที่ต้องการหรือไม่ และเช็คว่า value นั้นเป็น int หรือไม่
	if ok {
		fmt.Println("Age is an integer:", age)
	} else {
		fmt.Println("Age is not an integer")
	}
	NAME, isString := data["name"].(string)
	if isString {
		fmt.Println(NAME)
	} else {
		fmt.Println(0)
	}

	// การใช้งาน struct
	fmt.Println(" ")
	fmt.Println("struct")
	type person struct {
		name string
		age  int
	}
	type superPerson struct {
		power int
		person
		person2 person
	}
	p1 := person{name: "John", age: 30}
	fmt.Println(p1)
	fmt.Println(p1.name)
	// delete(p1, "name") // ไม่สามารถใช้ delete กับ struct ได้
	p1.name = "Doe" // การแก้ไขค่าใน struct
	fmt.Println(p1.name)
	sp1 := superPerson{
		power:   100,
		person:  person{name: "John", age: 30},
		person2: person{name: "Doe", age: 40},
	}
	fmt.Println(sp1, "superPerson")
	fmt.Println(sp1.name)
	fmt.Println(sp1.person2.name)
	sp1.person2.name = "hun" // การแก้ไขค่าใน struct ซ้อน struct
	fmt.Println(sp1.person2.name)
	type employee struct {
		name   string
		weight int
		enable bool
	}
	employeeList := []employee{
		{
			name:   "12",
			weight: 1234,
			enable: true,
		},
		{
			name: "23",
		},
	}
	employeeList[0].weight = 123123123213
	fmt.Println(employeeList, "employeeList")
	fmt.Println(employeeList[0], employeeList[0].name)
	fmt.Println(employeeList[1].weight, "weight") // ถ้าไม่มี key นั้นจะได้ 0 ออกมา int
	fmt.Println(employeeList[1].name, "name")     // ถ้าไม่มี key นั้นจะได้ "" ออกมา string
	fmt.Println(employeeList[1].enable, "enable") // ถ้าไม่มี key นั้นจะได้ false ออกมา bool
	// fmt.Println(employeeList[12])                 // ถ้าไม่มี index นั้นจะได้ error
	data12 := []int{1, 2}
	index12 := 12
	if index12 >= 0 && index12 < len(data) {
		fmt.Println("Accessing index:", data12[index12])
	} else {
		fmt.Println("Index out of range:", index12)
	}
	ex := test.Export{
		Export1: "Export1",
		Export2: "Export2",
	}
	fmt.Println(ex, "export")

	// การใช้งาน pointer
	fmt.Println(" ")
	fmt.Println("pointer")
	var num8 int = 10
	var ptr *int = &num8 // *int คือ pointer ที่เก็บค่า reference ของตัวแปร num8
	// &num8 คือการเข้าถึง address ของตัวแปร num8
	fmt.Println(num8, "num8")
	fmt.Println(ptr, "pointer")
	fmt.Println(*ptr, "pointer value")
	*ptr = 20 // การแก้ไขค่าใน pointer
	fmt.Println(num8)

	// การใช้งาน function
	fmt.Println(" ")
	fmt.Println("function")
	fmt.Println(add(10, 20))                // การเรียกใช้งาน function return ค่ากลับมา 1 ค่า
	fmt.Println(add2(10, 20))               // การเรียกใช้งาน function return ค่ากลับมา 3 ค่า
	value3, isTrue3, float3 := add2(10, 20) // การเรียกใช้งาน function return ค่ากลับมา 3 ค่า และเก็บค่าไว้ในตัวแปร
	fmt.Println(value3, isTrue3, float3)
	value4 := add(10, 20) // จะ return ค่ากลับมา 1 ค่า และเก็บค่าไว้ในตัวแปร เรียงตามลำดับที่ return กลับมา
	fmt.Println(value4, value4 > 20)
	value5, array5 := add3(20, 20) // จะ return ค่ากลับมา 2 ค่า และเก็บค่าไว้ในตัวแปร เรียงตามลำดับที่ return กลับมา
	fmt.Println(value5, array5[1:])
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	numss := 123
	fmt.Println(numss)
	changeNum(numss)
	fmt.Println(numss, "changeNum") // การส่งค่าไปใน function แล้วแก้ไขค่าใน function จะไม่ส่งผลกระทบกับค่าตัวแปรที่ส่งไป
	fmt.Println(rectangularArea(rectType{width: 10, height: 5}, "cm^2"), "rectangularArea")
	fmt.Println(triangularArea(10, 5), "triangular")
	fmt.Println(triangularArea(10, 5).area, "triangularArea")

	// การใช้งาน interface
	// interface คือ สามารถเก็บค่าได้หลายชนิด และเก็บค่าได้หลายชนิดในตัวเดียวกัน
	fmt.Println(" ")
	fmt.Println("interface")
	fmt.Println(circle{radius: 10}.area())

	// การใช้งาน if else
	fmt.Println(" ")
	fmt.Println("if else, condition")
	nums := 10
	if nums > 10 {
		fmt.Println("more than 10")
	} else if nums == 10 {
		fmt.Println("equal 10")
	} else {
		fmt.Println("less than 10")
	}

	// การใช้งาน switch case
	fmt.Println(" ")
	fmt.Println("switch case")
	nums2 := 5
	switch nums2 {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("default") // ถ้าไม่ตรงกับ case ใดๆเลยจะทำงาน case default
	}

	// การใช้งาน loop
	fmt.Println(" ")
	fmt.Println("for loop")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	nums3 := []int{1, 2, 3, 4, 5}
	for index, value := range nums3 {
		if index == 2 {
			continue
		}
		fmt.Println(index, value)
	}
	for _, value := range nums3 {
		fmt.Println(value)
	}
	for index := range nums3 {
		fmt.Println(index)
	}
	for index := 0; index < len(nums3); index++ {
		fmt.Println(index, nums3[index])
	}
	type structOfArray struct {
		name string
		age  int
		nums []int
	}
	structs := []structOfArray{
		{name: "John", age: 30, nums: []int{1, 2, 3}},
		{name: "Doe", age: 40},
	}
	for index, value := range structs {
		fmt.Println(index, value, "struct of array")
		for index2, value2 := range value.nums { // การ loop ซ้อน loop ถ้ามีก็มี ถ้าไม่มีก็ไม่ทำงาน ไม่ error
			fmt.Println(index2, value2)
		}
	}

	// การใช้งาน while loop
	fmt.Println(" ")
	fmt.Println("while loop")
	nums4 := 0
	for nums4 < 5 {
		fmt.Println(nums4)
		nums4++
	}
	for {
		fmt.Println("true")
		break
	}

	// การใช้งาน do while loop
	fmt.Println(" ")
	fmt.Println("do while loop")
	nums5 := 0
	for {
		fmt.Println(nums5)
		nums5++
		if nums5 >= 5 {
			break
		}
	}

	test.PrintTest() // package

	//
	//
	PrintValue("John") // Output: John
	PrintValue(123)    // Output: 123
	PrintValue(3.14)   // Output: 3.14
	PrintValue(true)   // Output: true
	//
	//
	var number23 int
	fmt.Scanf("%d", &number23)

}

func add(num1 int, num2 int) int {
	return num1 + num2
}

func add2(num1 int, num2 int) (int, bool, float64) {
	var boo bool
	if num1 > num2 {
		boo = true
	} else {
		boo = false
	}
	return num1 + num2, boo, 40.4
}

func add3(num1 int, num2 int) (int, []int) {
	return num1 + num2, []int{1, 2, 3}
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// กำหนด Generic Constraint ที่รองรับหลายชนิด
type TestInterface interface {
	~string | ~bool | ~int | ~float64
}

// ฟังก์ชัน Generic ที่รองรับ TestInterface
func PrintValue[T TestInterface](value T) {
	fmt.Println(value)
}

func sum(num1 ...int) int { // รับค่าแบบไม่จำกัดจำนวนแต่ต้องเป็น int
	total := 0
	for _, value := range num1 {
		total += value
	}
	return total
}

func changeNum(numss int) {
	numss = 20
}

type rectType struct {
	width, height int
}

func rectangularArea(r rectType, unit string) string {
	return strconv.Itoa(r.width*r.height) + unit // strconv.Itoa คือ แปลง int เป็น string
}

type triangularReturnType struct {
	area, perimeter int
}

func triangularArea(base, height int) triangularReturnType {
	area := base * height / 2
	perimeter := base + height + int(math.Sqrt(float64(base*base+height*height)))
	return triangularReturnType{area: area, perimeter: perimeter}
}

//
// strconv.Atoi	string เป็น int ใช้ได้เฉพาะ base 10
// strconv.ParseInt	string เป็น int64 รองรับ base (2, 10, 16, ฯลฯ)
// strconv.FormatInt	int64 เป็น string ระบุ base
// strconv.ParseFloat	string เป็น float64
// strconv.FormatFloat	float64 เป็น string ควบคุมทศนิยมและรูปแบบ
// strconv.ParseBool	string เป็น bool รับค่า "true" หรือ "false"
// strconv.FormatBool	bool เป็น string
// strconv.Itoa	int เป็น string

//
// strings.Contains ค้นหาคำนั้นมีคำที่ต้องการหรือไม่ เช่น strings.Contains("hello world", "world") // Output: true
// strings.ToUpper แปลง string เป็น uppercase เช่น strings.ToUpper("hello world") // Output: HELLO WORLD
// strings.ToLower แปลง string เป็น lowercase เช่น strings.ToLower("HELLO WORLD") // Output: hello world
// strings.Title แปลง string เป็น title เช่น strings.Title("hello world") // Output: Hello World
// strings.Trim ตัดช่องว่างหน้าและหลัง string เช่น strings.Trim(" hello world ", " ") // Output: hello world
// strings.Split แยก string เป็น array เช่น strings.Split("a,b,c", ",") // Output: [a b c]
// strings.Join รวม array เป็น string เช่น strings.Join([]string{"hello", "world"}, " ") // Output: hello world
// strings.Replace แทนที่ string เช่น strings.Replace("hello world", "world", "golang", 1) เลข 1 คือ จำนวนที่จะแทนที่ // Output: hello golang
// strings.ReplaceAll แทนที่ string ทั้งหมด เช่น strings.ReplaceAll("foo bar foo", "foo", "baz") // Output: baz bar baz
// strings.Repeat ทำซ้ำ string เช่น strings.Repeat("go", 3) // Output: gogogo
// strings.Index หา index ของ string เช่น strings.Index("hello world", "world") // Output: 6
// strings.HasPrefix ตรวจสอบว่า string นั้นขึ้นต้นด้วยคำที่ต้องการหรือไม่ เช่น strings.HasPrefix("golang", "go") // Output: true
// strings.HasSuffix ตรวจสอบว่า string นั้นลงท้ายด้วยคำที่ต้องการหรือไม่ เช่น strings.HasSuffix("golang", "lang") // Output: true
// strings.Fields แยก string ด้วยช่องว่าง เช่น strings.Fields("hello world golang") // Output: [hello world golang]
// strings.SplitN แยก string ด้วยคำที่ต้องการและระบุจำนวนที่แยก เช่น strings.SplitN("a,b,c", ",", 2) // Output: [a b,c]
