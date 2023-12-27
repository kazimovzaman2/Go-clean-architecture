package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/kazimovzaman2/clean-architecture/api/routes"
	_ "github.com/kazimovzaman2/clean-architecture/docs"
	"github.com/kazimovzaman2/clean-architecture/pkg/book"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// @title Clean Architecture example
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name Zaman Kazimov
// @contact.email kazimovzaman2@gmail.com
// @license.name MIT
// @host localhost:8080
// @BasePath /
func main() {
	db, cancel, err := databaseConnection()
	if err != nil {
		log.Fatal("Database connection Error $s", err)
	}
	fmt.Println("Database connection success!")
	bookCollection := db.Collection("books")
	bookRepo := book.NewRepo(bookCollection)
	bookService := book.NewService(bookRepo)

	app := fiber.New()
	app.Use(cors.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Welcome to the page."))
	})

	app.Get("/swagger/*", swagger.HandlerDefault)

	api := app.Group("/api")
	routes.BookRouter(api, bookService)
	defer cancel()

	log.Fatal(app.Listen(":8080"))
}

func databaseConnection() (*mongo.Database, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb://localhost:27017/fiber").SetServerSelectionTimeout(
		5*time.Second))

	if err != nil {
		cancel()
		return nil, nil, err
	}
	db := client.Database("books")
	return db, cancel, nil
}
