package chat

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"learning/dao"
	"sync"
)

var RoomInstance *Room

func init() {
	RoomInstance = &Room{
		Clients: make([]*ClientId, 0),

		OnlineChan:  make(chan *Client, 10),
		OfflineChan: make(chan *Client, 10),
		MessageChan: make(chan *Message, 10),
	}

	go RoomInstance.Broadcast(context.TODO())
}

type Room struct {
	CourseId int

	onlineNum int

	mux     sync.Mutex
	Clients []*ClientId

	// chan
	OnlineChan  chan *Client
	OfflineChan chan *Client
	MessageChan chan *Message
}

func NewChatRoom(courseId int) *Room {
	return &Room{
		CourseId: courseId,

		OnlineChan:  make(chan *Client, 10),
		OfflineChan: make(chan *Client, 10),
		MessageChan: make(chan *Message, 10),
	}
}

func (r *Room) OnlineNum() int {
	return r.onlineNum
}

func (r *Room) Process(ctx context.Context, conn *websocket.Conn, id int) {
	user, err := dao.GetUserById(ctx, id)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"id":  id,
			"err": err.Error(),
		}).Error("get user failed")
		return
	}

	client := &Client{
		Type: ClientTypeUser,
		User: &user,
		Conn: conn,
	}

	r.mux.Lock()
	r.Clients = append(r.Clients, &ClientId{
		Id:     id,
		Client: client,
	})
	r.onlineNum++
	r.mux.Unlock()

	r.OnlineChan <- client
	fmt.Println("online:", client.User.Name)

	defer func() {
		r.OfflineChan <- client

		r.mux.Lock()
		for i, client := range r.Clients {
			if client.Id == id {
				r.Clients = append(r.Clients[:i], r.Clients[i+1:]...)
				break
			}
		}
		r.onlineNum--
		r.mux.Unlock()

		conn.Close()
	}()

	for {
		_, data, err := conn.ReadMessage()
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"err": err.Error(),
			}).Error("read message failed")
			break
		}

		r.MessageChan <- &Message{
			Type: MsgCommon,
			Msg:  string(data),
			Name: client.User.Name,
		}
	}
}

func (r *Room) Broadcast(ctx context.Context) {
	for {
		select {
		case client := <-r.OnlineChan:
			msg := &Message{
				Name: client.User.Name,
				Msg:  fmt.Sprintf("%s加入聊天室", client.User.Name),
				Type: MsgOnline,
			}

			r.mux.Lock()
			msg.OnlineNum = r.onlineNum
			r.mux.Unlock()

			fmt.Println("welcome", client.User.Name)
			r.MessageChan <- msg
		case client := <-r.OfflineChan:
			msg := &Message{
				Name: client.User.Name,
				Msg:  fmt.Sprintf("%s离开聊天室", client.User.Name),
				Type: MsgOffline,
			}

			r.mux.Lock()
			msg.OnlineNum = r.onlineNum
			r.mux.Unlock()

			fmt.Println("bye", client.User.Name)
			r.MessageChan <- msg
		case msg := <-r.MessageChan:
			fmt.Println("broadcast", msg.Name)
			r.SendMessage(msg)
		}
	}
}

func (r *Room) SendMessage(msg *Message) {
	switch msg.Type {
	case MsgOnline, MsgOffline:
		for _, clientId := range r.Clients {
			if clientId.Client.User.Name != msg.Name {
				err := clientId.Client.Conn.WriteJSON(msg)
				if err != nil {
					logrus.WithFields(logrus.Fields{
						"err": err.Error(),
					}).Error("write message failed")
				}
			}
		}
	case MsgRobot:

	case MsgCommon:
		for _, clientId := range r.Clients {
			err := clientId.Client.Conn.WriteJSON(msg)
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"err": err.Error(),
				}).Error("write message failed")
			}
		}
	}
}
