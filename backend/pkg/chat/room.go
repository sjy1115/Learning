package chat

import (
	"github.com/gorilla/websocket"
)

type Room struct {
	CourseId int

	onlineNum int
	Clients   []*ClientId

	// chan
	OnlineChan  chan *Client
	OfflineChan chan *Client
	MessageChan chan *Message
}

func (r *Room) OnlineNum() int {
	return r.onlineNum
}

func (r *Room) Process(conn *websocket.Conn, id int) {

}

func (r *Room) Broadcast() {
	for {
		select {
		case client := <-r.OnlineChan:
			msg := &Message{
				Name: client.User.Name,
			}
			r.MessageChan <- msg
		case client := <-r.OfflineChan:
			msg := &Message{
				Name: client.User.Name,
			}
			r.MessageChan <- msg
		case msg := <-r.MessageChan:
			r.SendMessage(msg)
		}
	}
}

func (r *Room) SendMessage(msg *Message) {
	switch msg.Type {
	case OnlineMessage:
	case OfflineMessage:
	case CommonMessage:
	}
}
