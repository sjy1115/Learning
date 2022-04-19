package chat

import (
	"github.com/gorilla/websocket"
	"learning/models"
)

type MessageType int

const (
	_ MessageType = iota
	OnlineMessage
	OfflineMessage
	CommonMessage
)

const (
	ClientTypeUser = "User"
	ClientTypeBot  = "Bot"
)

type Message struct {
	Type MessageType `json:"type"`
	Name string      `json:"name"`
}

type ClientId struct {
	Id     int    `json:"id"`
	Client Client `json:"client"`
}

type Client struct {
	Conn *websocket.Conn
	Type string
	User *models.User
}
