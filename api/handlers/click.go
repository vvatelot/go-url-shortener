package handlers

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/vvatelot/url-shortener/api/entities"
	"github.com/vvatelot/url-shortener/config"
)

func GetClick(c *fiber.Ctx) error {
	id := c.Params("id")
	var click entities.Click
	result := config.Database.Find(&click, id)
	if result.RowsAffected == 0 {
		return c.Status(http.StatusNotFound).SendString("Not found")
	}
	return c.Status(http.StatusOK).JSON(&click)
}

func GetClicks(c *fiber.Ctx) error {
	var clicks []entities.Click
	size, err := strconv.Atoi(c.Query("size"))
	if err != nil {
		size = 10
	}
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}
	result := config.Database.Limit(size).Offset((page - 1) * size).Find(&clicks)
	if result.RowsAffected == 0 {
		return c.Status(http.StatusNotFound).SendString("Not found")
	}
	return c.Status(http.StatusOK).JSON(clicks)
}
