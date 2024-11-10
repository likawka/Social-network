package api

type NotificationInfo struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
	IDRef int `json:"idRef"`
}
