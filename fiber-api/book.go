package main

// go install github.com/swaggo/swag/cmd/swag@latest // ติดตั้ง swag
// swag init // สร้าง docs โดยใช้ swag

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Book struct {
	// ID คือ field ที่ต้องการรับเข้ามา
	// int คือ type ของ field นั้น
	// json:"id" คือ tag ที่ใช้เพื่อบอกว่า field นี้จะถูกแปลงเป็น json โดยใช้ชื่อ id แล้วจะแปลงให้อยู่ใน field ที่มีชื่อ ID ของ struct นี้ (map กับ struct แล้วได้ข้อมูลออกมา (เป็น json ต้องมี key ที่เป็น string และ value ที่เป็น string, number, boolean, array, object))
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	// รับเข้ามาหรือส่งกลับไป field จะอยู่ในรูปแบบ json ด้านขวา (metadata json)
}

var books = []*Book{
	// ทำให้เป็น pointer เพราะการจัดการกับ memory ให้เป็นประสิทธิภาพมากขึ้น

	// 	ถ้าไม่ใช้ pointer:
	// ทุกครั้งที่เพิ่ม struct ลงใน slice จะมีการคัดลอกค่าของ struct ทั้งหมด
	// ถ้า struct มีขนาดใหญ่ (ฟิลด์เยอะ) จะทำให้สิ้นเปลืองหน่วยความจำ
	{
		ID:     1,
		Title:  "The Art of Computer Programming",
		Author: "Donald Knuth",
	}, {
		ID:     2,
		Title:  "Structure and Interpretation of Computer Programs",
		Author: "Harold Abelson",
	},
}

// Handler functions
// getBooks godoc
// @Summary Get all books
// @Description Get details of all books
// @Tags books
// @Accept  json
// @Produce  json
// @Success 200 {array} Book
// @Router /books [get]
// @Security ApiKeyAuth
func getBooks(c *fiber.Ctx) error {
	return c.JSON(books) // ส่งค่า books ออกไปในรูปแบบ json
}

type message struct {
	Message string `json:"message"`
}

func getBook(c *fiber.Ctx) error {
	//
	// sort := c.Query("sort") // รับค่า sort จาก query parameter
	// fmt.Println("query", sort)
	//

	//
	// authen := c.Get("Authorization") // รับค่า Authorization จาก header
	// fmt.Println("authen", authen)
	//

	// fmt.Println(fiber.ErrBadRequest.Code) // 400
	// fmt.Println(fiber.ErrBadRequest)      // Bad Request

	id, err := strconv.Atoi(c.Params("id")) // รับค่า id จาก url parameter
	// strconv.Atoi คือ แปลง string เป็น int
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(err.Error()) // strconv.Atoi: parsing "...": invalid syntax
	}

	for _, book := range books {
		if book.ID == id {
			return c.JSON(book) // ส่งค่า book ออกไปในรูปแบบ json
		}
	}
	return c.Status(fiber.StatusNotFound).SendString("Book not found") // ส่งค่าออกไปในรูปแบบ string

	// return c.Status(fiber.StatusNotFound).JSON(
	// 	message{
	// 		Message: "Book not found",
	// 	},
	// ) // ส่งค่าออกไปในรูปแบบ json

	// return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
	// 	"message": "Book not found",
	// }) // ส่งค่าออกไปในรูปแบบ json
}

func createBook(c *fiber.Ctx) error {
	type requestBody struct { // สร้าง struct ที่มี field ที่ต้องการรับเข้ามาใน body
		Title  string `json:"title"`
		Author string `json:"author"`
	}
	var body requestBody // สร้างตัวแปร body ที่มี type เป็น request
	// body := new(requestBody)   // สร้างตัวแปร body ที่มี type เป็น request แล้วก็จอง address ให้กับตัวแปรนั้น
	err := c.BodyParser(&body) // รับค่าจาก body แล้วเก็บไว้ในตัวแปร body ถ้ามี error จะ return ค่าออกไป ถ้าไม่มี error จะเก็บค่าไว้ในตัวแปร body
	// c.BodyParser คือ method ที่ใช้ในการรับค่าจาก body ให้ตัวแปร body
	// &body คือการส่งตัวแปร body ไปให้ method BodyParser แล้ว method นั้นจะเก็บค่าที่ได้ไว้ในตัวแปร body
	// & นำหน้า body ใน c.BodyParser(&body) ในภาษา go ถ้าไม่มี & จะเป็นการส่งค่าไปแบบ pass by value แล้วจะไม่สามารถเปลี่ยนค่าในตัวแปร body ได้
	// c.BodyParser คือ แปลง JSON, XML หรือ Form Data จาก Body ของคำขอ (Request Body) ให้เป็น Go struct หรือ map
	fmt.Println("body", body)
	fmt.Println("err", err)
	if err != nil {
		return c.Status(400).SendString("Bad Request")
	}
	book := &Book{ // & สร้าง pointer ของ Book เพื่อจัดการกับ memory ให้เป็นประสิทธิภาพมากขึ้น
		ID:     len(books) + 1,
		Title:  body.Title,
		Author: body.Author,
	}
	books = append(books, book) // เพิ่ม book ลงใน books
	return c.JSON(book)
}

func updateBook(c *fiber.Ctx) error {
	type requestBody struct {
		Title  string `json:"title"`
		Author string `json:"author"`
	}
	var body requestBody
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(400).SendString("Bad Request na")
	}
	id, err := strconv.Atoi(c.Params("id")) // strconv.Atoi คือ แปลง string เป็น int ถ้าไม่สำเร็จจะ return error
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	for _, book := range books {
		if book.ID == id {
			// fmt.Sprintf("%d", book.ID) คือการแปลงค่า book.ID จาก int เป็น string
			book.Title = body.Title
			book.Author = body.Author
			return c.JSON(book)
		}
	}
	return c.Status(404).SendString("Book not found")
}

func deleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	for i, book := range books {
		if fmt.Sprintf("%d", book.ID) == id {
			books = append(books[:i], books[i+1:]...)
			return c.SendString("Book is deleted")
		}
	}
	return c.Status(404).SendString("Book not found")
}

func middleware(c *fiber.Ctx) error {
	fmt.Println("middleware")
	return c.Next()
}
