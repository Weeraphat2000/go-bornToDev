package main

import "fmt"

func main() {

	text1 := "Hello"
	text2 := "World"
	func(text1 string, text2 string) string {
		fmt.Println(text1, text2)
		return text1 + text2
	}(text1, text2)
}
