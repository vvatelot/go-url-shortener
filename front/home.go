package front

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/vvatelot/url-shortener/api/entities"
	"github.com/vvatelot/url-shortener/config"
)

func HandleHomePage(c *fiber.Ctx) error {
	var links []entities.Link

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}
	size, err := strconv.Atoi(c.Query("size"))
	if err != nil {
		size = 10
	}

	config.Database.Model(&entities.Link{}).Preload("Clicks").Limit(size).Offset((page - 1) * size).Find(&links)

	return c.Render("index", fiber.Map{
		"PageTitle": "Liste des liens",
		"Links":     links,
	})
}

func HandleNewPage(c *fiber.Ctx) error {
	return c.Render("editLink", fiber.Map{
		"PageTitle": "Ajouter un lien",
	})
}
