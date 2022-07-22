package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/pusher/pusher-http-go"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Use(cors.New())

	pusherClient := pusher.Client{
		AppID:   os.Getenv("APPID"),
		Key:     os.Getenv("KEY"),
		Secret:  os.Getenv("SECRET"),
		Cluster: os.Getenv("CLUSTER"),
		Secure:  true,
	}

	app.Post("/api/messages", func(c *fiber.Ctx) error {

		var data map[string]string

		if err := c.BodyParser(&data); err != nil {
			return err
		}

		pusherClient.Trigger("chat", "message", data)

		return c.JSON([]string{})
	})

	app.Listen(":8000")
}
