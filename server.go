package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	_, err = gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"),
		),
		// PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		panic("Application could not start because of database connection")
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello Fiber")
	})

	err = app.Listen(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
	if err != nil {
		panic("Application could not start because app port was not set")
	}
}
