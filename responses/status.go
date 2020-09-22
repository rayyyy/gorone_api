package responses

// Status リクエストのステータス
type Status struct {
	Status     string `json:"status"`
	RequestID  string `json:"request_id"`
	TotalCount int    `json:"total_count"`
	Completed  int    `json:"completed"`
}
