package main

import (
	"bufio" // ใช้สำหรับรับค่าจากคีย์บอร์ด
	"fmt"
	"os"      // ใช้สำหรับรับค่าจากคีย์บอร์ด
	"strconv" // ใช้สำหรับแปลง type
	"strings" // ใช้สำหรับการจัดการ string, slice
)

var reader = bufio.NewReader(os.Stdin) // สร้าง reader สำหรับรับค่าจากคีย์บอร์ด

func main() {
	value1 := getInput("Enter first number: ")
	value2 := getInput("Enter second number: ")
	fmt.Println(value1, value2)

	// การใช้งาน package strconv
	fmt.Println(" ")
	fmt.Println("strconv")
	// แปลง string เป็น int
	// strconv.Atoi	string	int	ใช้ได้เฉพาะ base 10
	str := "123"
	num, err := strconv.Atoi(str) // แปลง string เป็น int ถ้าเป็น string ที่ไม่ใช่ตัวเลขจะเกิด error
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Number:", num) // Output: Number: 123

	// แปลง int เป็น string
	num2 := 123
	str2 := strconv.Itoa(num2)
	fmt.Println("String:", str2) // Output: String: 123

	// แปลง string เป็น float64
	str3 := "123.45"
	floatNum3, err3 := strconv.ParseFloat(str3, 64) // 64 คือ bitSize ของ float ถ้าเป็น string ที่ไม่ใช่ตัวเลขจะเกิด error
	if err3 != nil {
		fmt.Println("Error:", err3)
		return
	}
	fmt.Println("Float:", floatNum3)                // Output: Float: 123.45
	fmt.Println(strconv.ParseInt("123456", 10, 64)) // Output: 123456 <nil> 10 คือ base ของเลข 64 คือ bitSize ของ int

	// แปลง float64 เป็น string
	num5 := 123.456789
	str5 := strconv.FormatFloat(num5, 'f', 2, 64)
	fmt.Println("String:", str5) // Output: String: 123.46 ควบคุมทศนิยมและรูปแบบ 2 คือจำนวนทศนิยม 64 คือ bitSize ของ float f คือรูปแบบทศนิยม

	// แปลง string เป็น bool
	str6 := "true"
	boolean6, err6 := strconv.ParseBool(str6)
	if err6 != nil {
		fmt.Println("Error:", err6)
		return
	}
	fmt.Println("Boolean:", boolean6) // Output: Boolean: true
	str66 := "asdfasdfasdf"
	boolean66, err66 := strconv.ParseBool(str66)
	if err66 != nil {
		fmt.Println("Error:", err66, "asdfasdfasdf")
		return
	}
	fmt.Println("Boolean:", boolean66) // Output: Boolean: true

	// แปลง bool เป็น string
	boolean7 := true
	str7 := strconv.FormatBool(boolean7)
	fmt.Println("String:", str7) // Output: String: true
	// strconv.Itoa	int	string
	// strconv.ParseInt	string	int64	รองรับ base (2, 10, 16, ฯลฯ)
	// strconv.FormatInt	int64	string	ระบุ base
	// strconv.ParseFloat	string	float64
	// strconv.FormatFloat	float64	string	ควบคุมทศนิยมและรูปแบบ
	// strconv.ParseBool	string	bool	รับค่า "true" หรือ "false"
	// strconv.FormatBool	bool	string

	// การใช้งาน package strings
	fmt.Println(" ")
	fmt.Println("strings")
	fmt.Println(strings.Contains("hello world", "world"))             // Output: true
	fmt.Println(strings.Contains("hello world", "golang"))            // Output: false
	fmt.Println(strings.Index("hello world", "world"))                // Output: 6
	fmt.Println(strings.Index("hello world", "golang"))               // Output: -1
	fmt.Println(strings.HasPrefix("golang", "go"))                    // Output: true
	fmt.Println(strings.HasSuffix("golang", "lang"))                  // Output: true
	fmt.Println(strings.ToUpper("hello"))                             // Output: HELLO
	fmt.Println(strings.ToLower("WORLD"))                             // Output: world
	fmt.Println(strings.Title("hello world"))                         // Output: Hello World
	fmt.Println(strings.Trim("  hello world  ", " "))                 // Output: hello world
	fmt.Println(strings.Split("a,b,c", ","))                          // Output: [a b c]
	fmt.Println(strings.SplitN("a,b,c", ",", 2))                      // Output: [a b,c]
	fmt.Println(strings.Fields("hello world golang"))                 // Output: [hello world golang]
	fmt.Println(strings.Join([]string{"hello", "world"}, " "))        // Output: hello world
	fmt.Println(strings.Replace("hello world", "world", "golang", 1)) // Output: hello golang
	fmt.Println(strings.ReplaceAll("foo bar foo", "foo", "baz"))      // Output: baz bar baz
	fmt.Println(strings.Repeat("go", 3))                              // Output: gogogo

}

func getInput(prompt string) float64 {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	fmt.Println(err, "err")
	if err != nil { // nil คือค่าว่าง

		// ถ้าเกิด error ให้แสดง error ออกมาและ return ค่า 0
		fmt.Println(err)
		return 0
	}
	return value
}
