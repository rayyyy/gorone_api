package responses

// Request リクエスト時のレスポンス
type Request struct {
	Message   string `json:"message"`
	RequestID string `json:"request_id"`
}
