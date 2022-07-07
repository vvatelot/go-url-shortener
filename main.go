package main

import (
	"os"

	"github.com/gofiber/template/html"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
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

	app.Get("/r/:key", handlers.Redirect)
	app.Static("/", "./public")

	app.Use(basicauth.New(config.BasicAuthConfig()))
	app.Get("/", front.HandleHomePage)
	app.Get("/new", front.HandleNewPage)

	api := app.Group("/api")
	links := api.Group("/links")
	links.Get("/:id", handlers.GetLink)
	links.Get("/", handlers.GetLinks)
	links.Post("/", handlers.AddLink)
	links.Patch("/:id", handlers.UpdateLink)
	links.Delete("/:id", handlers.DeleteLink)
	links.Put("/:id/switch-active", handlers.ActivateLink)

	clicks := api.Group("/clicks")
	clicks.Get("/:id", handlers.GetClick)
	clicks.Get("/", handlers.GetClicks)

	app.Listen(":" + os.Getenv("APP_PORT"))
}
