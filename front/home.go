package front

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/vvatelot/url-shortener/api/repositories"
	"github.com/vvatelot/url-shortener/utils"
)

func HandleHomePage(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}
	size, err := strconv.Atoi(c.Query("size"))
	if err != nil {
		size = 10
	}

	links, _ := repositories.ListLinks(size, page)

	CountAllLinks := repositories.CountAllLinks()

	return c.Render("index", fiber.Map{
		"PageTitle":  "Liste des liens",
		"Links":      links,
		"Pagination": utils.GetPagination(CountAllLinks, page, size),
	})
}

func HandleNewPage(c *fiber.Ctx) error {
	return c.Render("editLink", fiber.Map{
		"PageTitle": "Ajouter un lien",
	})
}
