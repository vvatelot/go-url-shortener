package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/vvatelot/url-shortener/api/entities"
	"github.com/vvatelot/url-shortener/api/repositories"
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

func GetLinkClickStats(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid id")
	}

	if err := repositories.EnsureLinkExists(id); err != nil {
		return c.Status(http.StatusNotFound).SendString("Not found")
	}

	window := c.Query("window", "7d")
	var days int
	switch window {
	case "7d":
		days = 7
	case "30d":
		days = 30
	default:
		return c.Status(http.StatusBadRequest).SendString("Invalid window")
	}

	since := time.Now().UTC().Truncate(24*time.Hour).AddDate(0, 0, -(days - 1))
	buckets, total, err := repositories.CountClicksByDay(id, since)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Error while loading stats")
	}

	return c.Status(http.StatusOK).JSON(entities.ClickStatsResponse{
		LinkID:  id,
		Window:  window,
		Total:   total,
		Buckets: buckets,
	})
}
