package responses

// Request リクエスト時のレスポンス
type Request struct {
	Message   string `json:"message"`
	RequestID uint   `json:"request_id"`
}
