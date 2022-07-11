package entities

type APIListResponse[T any] struct {
	Data       []T
	Pagination APIListResponsePagination
}

type APIListResponsePagination struct {
	TotalCount   int64
	PageLast     int
	PageSize     int
	PageCurrent  int
	PagePrevious int
	PageNext     int
}
