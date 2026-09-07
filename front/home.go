package front

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/vvatelot/url-shortener/api/repositories"
	"github.com/vvatelot/url-shortener/config"
	"github.com/vvatelot/url-shortener/utils"
)

var translation config.Translation

func HandleHomePage(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}
	size, err := strconv.Atoi(c.Query("size"))
	if err != nil {
		size = 10
	}

	flash := c.Query("flash")
	message := c.Query("message")

	links, _ := repositories.ListLinks(size, page)

	CountAllLinks := repositories.CountAllLinks()

	return c.Render("index", fiber.Map{
		"e":          config.T,
		"Links":      links,
		"Pagination": utils.GetPagination(CountAllLinks, page, size),
		"Flash":      flash,
		"Message":    message,
	})
}

func HandleNewPage(c *fiber.Ctx) error {
	return c.Render("editLink", fiber.Map{
		"e":      config.T,
		"IsEdit": false,
	})
}

func HandleEditPage(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid id")
	}

	link, err := repositories.GetLinkByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Not found")
	}

	return c.Render("editLink", fiber.Map{
		"e":      config.T,
		"IsEdit": true,
		"Link":   link,
	})
}

func HandleLinkDetailPage(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid id")
	}

	link, err := repositories.GetLinkByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Not found")
	}

	totalClicks, err := repositories.CountAllClicksByLinkID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error while loading stats")
	}

	return c.Render("linkDetail", fiber.Map{
		"e":           config.T,
		"Link":        link,
		"TotalClicks": totalClicks,
	})
}
