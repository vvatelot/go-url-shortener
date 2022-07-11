package repositories

import (
	"errors"

	"github.com/vvatelot/url-shortener/api/entities"
	"github.com/vvatelot/url-shortener/config"
)

func ListLinks(size, page int) ([]entities.Link, error) {
	var links []entities.Link

	result := config.Database.Preload("Clicks").Limit(size).Offset((page - 1) * size).Find(&links)
	if result.RowsAffected == 0 {
		return nil, errors.New("not found")
	}

	return links, nil
}

func CountAllLinks() int64 {
	var count int64

	config.Database.Model(&entities.Link{}).Count(&count)

	return count
}

func GetLinkByID(id int) (entities.Link, error) {
	var link entities.Link

	result := config.Database.Preload("Clicks").Find(&link, id)

	if result.RowsAffected == 0 {
		return link, errors.New("not found")
	}

	return link, nil
}
