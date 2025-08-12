package services

import (
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/pablohenriques/go-project/clients"
	"github.com/pablohenriques/go-project/dto"
	"github.com/pablohenriques/go-project/entity"
)

var listBooks = []entity.Book{}

func Home(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "active"})
}

func CreateBook(c *fiber.Ctx) error {
	var newBook dto.BookDTO

	if err := c.BodyParser(&newBook); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "JSON inválido"})
	}

	log.Printf("info id=%d book=%s", newBook.Id, newBook.Nome)

	modelBook := entity.Book{
		Id:       newBook.Id,
		Nome:     newBook.Nome,
		Detalhes: newBook.Detalhes,
		Data:     time.Now(),
	}

	listBooks = append(listBooks, modelBook)

	return c.Status(fiber.StatusOK).JSON(modelBook.ToResponseDTO())
}

func GetOneBook(c *fiber.Ctx) error {
	queryParam := c.Query("id")
	idBook, err := strconv.ParseInt(queryParam, 10, 64)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Parse Int não foi executado corretamente"})
	}

	for _, book := range listBooks {
		if book.Id == idBook {
			return c.Status(fiber.StatusOK).JSON(book.ToResponseDTO())
		}
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{"message": "Livro não encontrado"})
}

func GetAllBook(c *fiber.Ctx) error {
	if len(listBooks) < 1 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "não há livros cadastrados"})
	}

	return c.Status(fiber.StatusOK).JSON(listBooks)
}

func UpdateBook(c *fiber.Ctx) error {
	queryParam := c.Params("id")
	idBook, err := strconv.ParseInt(queryParam, 10, 64)
	var updateBook dto.BookUpdateDTO

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Parse Int não foi executado corretamente"})
	}

	if err := c.BodyParser(&updateBook); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "JSON inválido"})
	}

	for index, book := range listBooks {
		if book.Id == idBook {
			listBooks[index].Nome = updateBook.Nome
			listBooks[index].Detalhes = updateBook.Detalhes
			log.Print(listBooks[index])
			return c.Status(fiber.StatusOK).JSON(book.ToResponseDTO())
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "não há livros cadastrados"})
}

func DeleteBook(c *fiber.Ctx) error {
	queryParam := c.Params("id")
	idBook, err := strconv.ParseInt(queryParam, 10, 64)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Parse Int não foi executado corretamente"})
	}

	for index, book := range listBooks {
		if book.Id == idBook {
			listBooks = remove(listBooks, index)
			return c.Status(fiber.StatusOK).JSON(listBooks)
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Livro não encontrado"})

}

func remove(s []entity.Book, i int) []entity.Book {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func Erro500(c *fiber.Ctx) error {

	if time.Now().Unix()%2 == 0 {
		return errors.New("não foi possível conectar")
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "active"})
}

func CallCircuitBreaker(c *fiber.Ctx) error {
	clients.GetHttpBin()
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "active"})
}
