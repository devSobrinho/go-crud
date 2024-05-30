package requestSSE

type SendMessageRequest struct {
	ID      string `json:"id" binding:"omitempty"`
	Message string `json:"message" binding:"required"`
}
