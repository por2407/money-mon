package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/moneymon/internal/core/domain"
	"github.com/moneymon/internal/infrastructure"
)

func main() {
	// เชื่อมต่อ Database
	db, err := infrastructure.NewPostgresDB()
	if err != nil {
		log.Fatal("Could not connect to database:", err)
	}

	db.AutoMigrate(&domain.User{}, &domain.Character{})

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	fmt.Println("Server running on port 8080")
	if err := app.Listen(":8080"); err != nil {
		fmt.Println("server error: ", err)
	}
}
