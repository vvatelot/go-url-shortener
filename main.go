package main

import (
	"os"

	"github.com/gofiber/template/html"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/joho/godotenv/autoload"
	"github.com/vvatelot/url-shortener/api/handlers"
	"github.com/vvatelot/url-shortener/config"
	"github.com/vvatelot/url-shortener/front"
)

func main() {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	if os.Getenv("ENV") == "dev" {
		app.Use(logger.New())
	}
	config.Connect()

	app.Get("/", front.HandleHomePage)

	app.Get("/links/:id", handlers.GetLink)
	app.Get("/links", handlers.GetLinks)
	app.Post("/links", handlers.AddLink)
	app.Patch("/links/:id", handlers.UpdateLink)
	app.Delete("/links/:id", handlers.DeleteLink)

	app.Get("/clicks/:id", handlers.GetClick)
	app.Get("/clicks/", handlers.GetClicks)

	app.Get("/r/:key", handlers.Redirect)

	app.Listen(":" + os.Getenv("APP_PORT"))
}