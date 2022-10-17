package main

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
	"github.com/vvatelot/url-shortener/api/handlers"
	"github.com/vvatelot/url-shortener/config"
	"github.com/vvatelot/url-shortener/front"
)

//go:embed i18n/*
var embedDirI18n embed.FS

//go:embed static/*
var embedDirStatic embed.FS

//go:embed views/*
var embedDirViews embed.FS

var ENV *config.Environment = config.GetEnvironment()

func main() {
	config.ENV = ENV
	config.EmbedDirI18n = embedDirI18n
	viewsFileSystem, err := fs.Sub(embedDirViews, "views")
	if err != nil {
		panic(err)
	}
	app := fiber.New(fiber.Config{
		Views: html.NewFileSystem(http.FS(viewsFileSystem), ".html"),
	})

	if ENV.Env == "dev" {
		app.Use(logger.New())
	}
	config.Connect()
	config.GetTranslation()

	app.Get("/r/:key", handlers.Redirect)
	app.Use("/static", filesystem.New(filesystem.Config{
		Root:       http.FS(embedDirStatic),
		PathPrefix: "static",
	}))

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

	app.Listen(":" + ENV.AppPort)
}
