package utils

import (
	"math"

	"github.com/vvatelot/url-shortener/api/entities"
)

func GetLastPage(countAll int64, size int) int {
	return int(math.Ceil(float64(countAll) / float64(size)))
}

func GetPreviousPage(currentPage int) int {
	if currentPage == 1 {
		return 1
	}

	return currentPage - 1
}

func GetNextPage(currentPage int, lastPage int) int {
	if currentPage == lastPage {
		return lastPage
	}

	return currentPage + 1
}

func GetPagination(total int64, page int, size int) entities.APIListResponsePagination {
	lastPage := GetLastPage(total, size)
	pagination := entities.APIListResponsePagination{
		TotalCount:   total,
		PageLast:     lastPage,
		PageSize:     size,
		PageCurrent:  page,
		PagePrevious: GetPreviousPage(page),
		PageNext:     GetNextPage(page, lastPage),
	}

	return pagination
}
