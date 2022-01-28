package controllers

import (
	"digicert_book_api/database"
	"encoding/json"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllBooks(ctx *fiber.Ctx) error {
	books, err := database.GetAllBooks()

	if err != nil {
		return err
	}

	return ctx.JSON(books)
}

func GetBookById(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		return err
	}

	book, err := database.GetBookById(id)

	if err != nil {
		return err
	}

	return ctx.JSON(book)
}

func UpdateBookById(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		return err
	}

	var book database.Book

	json.Unmarshal(ctx.Body(), &book)

	err = database.UpdateBookById(id, book)

	if err != nil {
		return err
	}

	return nil
}

func DeleteBookById(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		return err
	}

	err = database.DeleteBookById(id)

	if err != nil {
		return err
	}

	return nil
}

func CreateBook(ctx *fiber.Ctx) error {
	var book database.Book

	json.Unmarshal(ctx.Body(), &book)

	err := database.CreateBook(book)

	if err != nil {
		return err
	}

	return nil
}
