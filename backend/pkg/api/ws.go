package api

type MessageResponse struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

type ErrorMessage struct {
	Message string `json:"message"`
}

type UpdateNotification struct {
	Message string `json:"message"`
	Count  int    `json:"count"`
}