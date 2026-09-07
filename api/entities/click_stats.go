package entities

type ClickStatPoint struct {
	Date  string `json:"date"`
	Count int64  `json:"count"`
}

type ClickStatsResponse struct {
	LinkID  int              `json:"link_id"`
	Window  string           `json:"window"`
	Total   int64            `json:"total"`
	Buckets []ClickStatPoint `json:"buckets"`
}
