package main

import (
	"digicert_book_api/controllers"
	"digicert_book_api/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.InitDB()
	app := fiber.New()

	app.Get("/", controllers.GetAllBooks)
	app.Get("/book/:id", controllers.GetBookById)
	app.Post("/book", controllers.CreateBook)
	app.Put("/book/:id", controllers.UpdateBookById)
	app.Delete("/book/:id", controllers.DeleteBookById)

	app.Listen(":8080")
}
