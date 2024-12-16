package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// gin
	// echo
	// fiber == express.js
	fmt.Println("Hello, Fiber!")
	//
	app := fiber.New() // Create a new Fiber instance

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,PATCH,DELETE",
		// AllowHeaders: "*",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println("Hello, Fiber!")
		return c.SendString("Hello, Fiber!")
	})
	// print(books) // print คือ แสดง memory address ของตัวแปร
	app.Get("/books", middleware, getBooks)
	app.Get("/books/:id", getBook)
	app.Post("/books", createBook)
	app.Patch("/books/:id", updateBook)
	app.Delete("/books/:id", deleteBook)

	app.Post("/upload", uploadFile)

	app.Listen(":3000") // Listen on port 3000
	// ถ้ายิ่งไปที่ที่ยังไม่มี port รองรับ จะได้ Method Not Allowed 405 (go fiber เป็นคนจัดการให้)
}

func uploadFile(c *fiber.Ctx) error {

	// a, error := c.FormFile("documents") // รับ file ที่ส่งมาจาก form ที่มี enctype="multipart/form-data"
	// a จะได้ address ของ file ที่ส่งมา แค่ 1 file
	// fmt.Println("aaa", a.Filename, "error", error) // a คือ ข้อมูลของ file ที่ส่งมาจาก form ที่มี enctype="multipart/form-data"

	// Parse the multipart form file คือ การรับ file ที่ส่งมา จาก form ที่มี enctype="multipart/form-data"
	form, err := c.MultipartForm()
	if err != nil {
		fmt.Println("err", err)
		// request Content-Type has bad boundary or is not multipart/form-data คือ ไม่มี multipart/form-data เข้ามา

		// request Content-Type has bad boundary or is not multipart/form-data คือ ไม่มี multipart/form-data เข้ามา
		return err
	}
	// fmt.Println("form", form) // form คือ ข้อมูลที่ส่งมาจาก form ที่มี enctype="multipart/form-data"

	// Get all files from "documents" key
	files := form.File["documents"] // จะได้ address ของ file ที่ส่งมา (files คือ slice ของ file)
	// *files[0] คือ การดึงค่าออกมาจาก address นั้น
	// fmt.Println("files", files, *files[0])
	// Iterate over files
	for _, file := range files {
		// Save the files to disk
		fmt.Println("name", file.Filename)
		c.SaveFile(file, fmt.Sprintf("./uploads/%s", file.Filename))
	}
	// Return success message
	return c.JSON(fiber.Map{
		"message": "Files uploaded successfully",
	})
}
