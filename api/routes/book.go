package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kazimovzaman2/clean-architecture/api/handlers"
	"github.com/kazimovzaman2/clean-architecture/pkg/book"
)

func BookRouter(app fiber.Router, service book.Service) {
	app.Get("/books", handlers.GetBooks(service))
	app.Post("/books", handlers.AddBook(service))
	app.Put("/books", handlers.UpdateBook(service))
	app.Delete("/books", handlers.RemoveBook(service))
}
