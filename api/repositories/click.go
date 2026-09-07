package repositories

import (
	"errors"
	"time"

	"github.com/vvatelot/url-shortener/api/entities"
	"github.com/vvatelot/url-shortener/config"
)

type dayCountRow struct {
	Day   string
	Count int64
}

func CountClicksByDay(linkID int, since time.Time) ([]entities.ClickStatPoint, int64, error) {
	var rows []dayCountRow

	dayExpr := "DATE(created_at)"
	if config.Database.Dialector.Name() == "postgres" {
		dayExpr = "TO_CHAR(created_at AT TIME ZONE 'UTC', 'YYYY-MM-DD')"
	}

	result := config.Database.Model(&entities.Click{}).
		Select(dayExpr+" AS day, COUNT(*) AS count").
		Where("link_id = ? AND created_at >= ?", linkID, since).
		Group(dayExpr).
		Order("day ASC").
		Scan(&rows)

	if result.Error != nil {
		return nil, 0, result.Error
	}

	byDay := make(map[string]int64, len(rows))
	var total int64
	for _, row := range rows {
		byDay[row.Day] = row.Count
		total += row.Count
	}

	now := time.Now().UTC().Truncate(24 * time.Hour)
	start := since.UTC().Truncate(24 * time.Hour)
	buckets := make([]entities.ClickStatPoint, 0)
	for day := start; !day.After(now); day = day.AddDate(0, 0, 1) {
		key := day.Format("2006-01-02")
		buckets = append(buckets, entities.ClickStatPoint{
			Date:  key,
			Count: byDay[key],
		})
	}

	return buckets, total, nil
}

func CountAllClicksByLinkID(linkID int) (int64, error) {
	var count int64
	result := config.Database.Model(&entities.Click{}).Where("link_id = ?", linkID).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}

func EnsureLinkExists(linkID int) error {
	var count int64
	result := config.Database.Model(&entities.Link{}).Where("id = ?", linkID).Count(&count)
	if result.Error != nil {
		return result.Error
	}
	if count == 0 {
		return errors.New("not found")
	}
	return nil
}
