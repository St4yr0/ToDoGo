package server

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func Server() {
	connStr := "postgresql://postgres:<password>@127.0.0.1/todos?sslmode=disable"
	// Connect to database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return IndexHandler(c, db)
	})

	app.Post("/", func(c *fiber.Ctx) error {
		return PostHandler(c, db)
	})

	app.Put("/update", func(c *fiber.Ctx) error {
		return PutHandler(c, db)
	})

	app.Delete("/delete", func(c *fiber.Ctx) error {
		return DeleteHandler(c, db)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}
