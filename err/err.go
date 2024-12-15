package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := divide(10, 5)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("result:", result)

	err = login("admin1", "password")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("login success")
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero") // สร้าง error โดยใช้ errors.New โดยสร้าง error จาก string
	}
	return a / b, nil // nil คือ null ในภาษา Go
}

type LoginError struct {
	Username string
	Password string
}

// สร้าง method ชื่อ Error ที่รับ pointer ของ LoginError และ return string // ต้องชื่อ Error เพื่อให้เป็น error interface
func (e *LoginError) Error() string {
	return fmt.Sprintf("error: %s %s", e.Username, e.Password) // ใช้ fmt.Sprintf ในการสร้าง string จาก format ของ string และค่าที่ต้องการแทนที่
}

func login(username, password string) error {
	if username != "admin" {
		return &LoginError{Username: username, Password: password} // สร้าง pointer ของ LoginError โดยใช้ & นำหน้า struct คือการสร้าง pointer ของ struct ถ้า username ไม่ใช่ admin จะ return error ของ LoginError โดยใช้ pointer ของ LoginError
	}
	return nil
}
