package main

// go install github.com/swaggo/swag/cmd/swag@latest // ติดตั้ง swag
// go get github.com/gofiber/swagger // ติดตั้ง swagger โดยใช้ go get
// swag init // สร้าง docs โดยใช้ swag

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/golang-jwt/jwt/v4"

	"github.com/joho/godotenv"

	jwtware "github.com/gofiber/jwt/v2"

	"github.com/gofiber/swagger"
	_ "github.com/weeraphat2000/docs" // docs คือ ชื่อ folder ที่เก็บ swagger ที่สร้างขึ้นมา
)

type User struct {
	email    string
	password string
}

var user = User{
	email:    "user@gmail.com",
	password: "1234",
}

func login(c *fiber.Ctx) error {
	type Login struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var loginVals Login
	err := c.BodyParser(&loginVals)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if loginVals.Email != user.email || loginVals.Password != user.password {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Email or password is incorrect",
		})
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	// token.Claims เป็นอินเตอร์เฟซที่เก็บข้อมูลของ claims (เช่น payload ของ JWT)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = loginVals.Email
	claims["role"] = "admin"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	fmt.Println("token", token)
	fmt.Println("claims", claims) // map[exp:1631530007 email: "..."]

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secretkey")) // สร้าง token และ encode ด้วย secretkey
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	fmt.Println("t", t) // token ที่ถูก encode แล้ว

	return c.JSON(fiber.Map{"token": t})
}

func checkMidleware(c *fiber.Ctx) error {

	fmt.Println("locals(user)", c.Locals("user"))

	user := c.Locals("user").(*jwt.Token)
	// c.Locals("user") คือ การดึงค่า user ออกมาจาก context ของ c
	// *jwt.Token คือ การแปลง user ให้เป็น pointer ของ jwt.Token
	claims := user.Claims.(jwt.MapClaims)
	// user.Claims.(jwt.MapClaims) คือ การดึงค่า claims ออกมาจาก user
	fmt.Println("claims", claims) // claims คือ map เช่น map[exp:1631530007 email: "..."]
	email := claims["email"]
	role := claims["role"]
	fmt.Println("email", email)
	if role != "admin" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	// คือ เพิ่มค่าใน context ของ c โดยใช้ key ชื่อ email และ value คือ email
	c.Locals("email", email)

	// เพื่ม key ชื่อ Hun และ value คือ Hun ใน context ของ c
	c.Locals("Hun", "Hun")

	return c.Next()
}

// เหมือนจะต้องเอาไปเขียนที่ main.go แต่เราเอามาเขียนที่นี่เพื่อให้เข้าใจง่ายขึ้น
// @title Book API
// @description This is a sample server for a book API.
// @version 1.0
// @host localhost:8080
// @BasePath /
// @schemes http
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// gin
	// echo
	// fiber == express.js
	fmt.Println("Hello, Fiber!")
	//
	app := fiber.New() // Create a new Fiber instance

	app.Get("/swagger/*", swagger.HandlerDefault) // สร้าง swagger UI โดยใช้ swagger.HandlerDefault

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,PATCH,DELETE",
		// AllowHeaders: "*",
	}))
	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println("Hello, Fiber!")
		return c.SendString("Hello, Fiber!")
	})

	app.Post("/login", login)

	// app.User คือ การใช้ middleware ทุก route ที่ตามหลัง
	// app.Use(jwtware.New(jwtware.Config{
	// 	SigningKey: []byte("secretkey"),
	// })) // ต้องมี token ถึงจะเข้าได้ ถ้าไม่มีจะเข้าไม่ได้ และต้องเป็น token ที่ถูก encode ด้วย secretkey และถูกต้องด้วย
	// เช็คให้อัตโนมัติว่าส่งเข้ามาใน header authorization หรือไม่ ถ้าไม่ส่งเข้ามาจะไม่สามารถเข้าได้

	app.Get("/testJWTMiddleware",

		// ต้องส่ง authorization bearer token ที่ถูก encode ด้วย secretkey และถูกต้องด้วย
		jwtware.New(jwtware.Config{
			SigningKey: []byte("secretkey"),
		}),

		checkMidleware,

		func(c *fiber.Ctx) error {

			email := c.Locals("email")
			fmt.Println("email in testJWTMiddleware", email)

			// ถ้าไม่มี email ก็จะได้ nil(nil คือ ค่าว่าง (null))
			if email == nil {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "ทำไมไม่มี email",
				})
			}

			return c.JSON(fiber.Map{
				"email": c.Locals("email"),
				"Hun":   c.Locals("Hun"),
			})
		})

	app.Use(func(c *fiber.Ctx) error {
		fmt.Printf("URL: %s, Method: %s, Time: %s\n", c.OriginalURL(), c.Method(), time.Now())
		return c.Next()
	})

	// print(books) // print คือ แสดง memory address ของตัวแปร
	app.Get("/books",
		middleware,
		// jwtware.New(jwtware.Config{
		// 	SigningKey: []byte("secretkey"),
		// }),
		// checkMidleware,
		getBooks)
	app.Get("/books/:id", getBook)
	app.Post("/books", createBook)
	app.Patch("/books/:id", updateBook)
	app.Delete("/books/:id", deleteBook)

	app.Post("/upload", uploadFile)

	app.Get("/getENV", getENV)

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

func getENV(c *fiber.Ctx) error {
	err := godotenv.Load() // โหลด .env
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	fmt.Println("err", err)

	env := os.Getenv("DATABASE_URL") // ดึงค่าจาก .env จะต้องมี godotenv.Load() ก่อน
	fmt.Println("env", env)
	return c.JSON(fiber.Map{
		"databaseUrl": os.Getenv("DATABASE_URL"),
	})
}
