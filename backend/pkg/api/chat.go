package api

import "time"

type Chat struct {
	Hash          string             `json:"hash"`
	ChatType      string             `json:"chatType"`
	ChatName      string             `json:"chatName"`
	Users         []UserResponseInfo `json:"users"`
	CreatedAt     time.Time          `json:"createdAt"`
	LastMasege    string             `json:"lastMasege"`
	LastSender    UserResponseInfo   `json:"lastSnder"`
	LastMasegTime time.Time          `json:"lastMasegTime"`
}

type CrateChat struct {
	Member UserResponseInfo `json:"member"`
}

type CreateChatResponse struct {
	Hash string `json:"hash"`
}

type SendMessageRequest struct {
	ChatHash string `json:"chatHash"`
	Pyload   Message `json:"pyload"`
}

type Message struct {
	Hash      string           `json:"hash"`
	Content   string           `json:"content"`
	Sender    UserResponseInfo `json:"sender"`
	CreatedAt time.Time        `json:"createdAt"`
}

type ChatList struct {
	IndividualChats []Chat `json:"individualChats"`
	GroupChats      []Chat `json:"groupChats"`
}

type OpenChat struct {
	Hash     string    `json:"hash"`
	Messages []Message `json:"messages"`
}
