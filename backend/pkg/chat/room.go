package chat

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"learning/dao"
	"learning/models"
	"sync"
	"time"
)

type Room struct {
	CourseId int

	onlineNum int

	mux     sync.Mutex
	Clients []*ClientId

	// chan
	OnlineChan  chan *Client
	OfflineChan chan *Client
	MessageChan chan *Message

	stopChan chan struct{}
}

func NewChatRoom(courseId int) *Room {
	return &Room{
		CourseId: courseId,

		Clients:     make([]*ClientId, 0),
		OnlineChan:  make(chan *Client, 10),
		OfflineChan: make(chan *Client, 10),
		MessageChan: make(chan *Message, 10),
	}
}

func (r *Room) OnlineNum() int {
	r.mux.Lock()
	defer r.mux.Unlock()

	return r.onlineNum
}

func (r *Room) Process(ctx context.Context, conn *websocket.Conn, userId int) {
	user, err := dao.UserGetById(ctx, userId)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"id":  userId,
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
		Id:     userId,
		Client: client,
	})
	r.onlineNum++
	r.mux.Unlock()

	r.OnlineChan <- client
	logrus.WithFields(logrus.Fields{
		"id":        userId,
		"user name": user.Name,
		"online":    r.OnlineNum(),
		"room":      r.CourseId,
	}).Info("user online")

	defer func() {
		r.OfflineChan <- client

		r.mux.Lock()
		for i, client := range r.Clients {
			if client.Id == userId {
				r.Clients = append(r.Clients[:i], r.Clients[i+1:]...)
				break
			}
		}
		r.onlineNum--

		r.mux.Unlock()

		if r.OnlineNum() == 0 {
			release(r)
			logrus.WithFields(logrus.Fields{
				"id": userId,
			}).Info("chat room closed")
		}
		r.stopChan <- struct{}{}
		conn.Close()
	}()

	for {
		var msg SendMessage
		err = conn.ReadJSON(&msg)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"user id": userId,
				"err":     err.Error(),
			}).Error("read message failed")
			break
		}

		r.MessageChan <- &Message{
			Type:     MsgCommon,
			Msg:      msg.Msg,
			Sender:   userId,
			Avatar:   user.Avatar,
			Name:     client.User.Name,
			SendTime: time.Now().Unix(),
		}
	}
}

func (r *Room) Broadcast() {
	for {
		select {
		case client := <-r.OnlineChan:
			msg := &Message{
				Name:     client.User.Name,
				Msg:      fmt.Sprintf("%s加入聊天室", client.User.Name),
				Type:     MsgOnline,
				SendTime: time.Now().Unix(),
			}

			r.mux.Lock()
			msg.OnlineNum = r.onlineNum
			r.mux.Unlock()

			fmt.Println("welcome", client.User.Name)
			r.MessageChan <- msg
		case client := <-r.OfflineChan:
			msg := &Message{
				Name:     client.User.Name,
				Msg:      fmt.Sprintf("%s离开聊天室", client.User.Name),
				Type:     MsgOffline,
				SendTime: time.Now().Unix(),
			}

			r.mux.Lock()
			msg.OnlineNum = r.onlineNum
			r.mux.Unlock()

			fmt.Println("bye", client.User.Name)
			r.MessageChan <- msg
		case msg := <-r.MessageChan:
			fmt.Println("broadcast", msg.Name)
			r.SendMessage(msg)
		case <-time.After(time.Second * 60 * 2):
			logrus.WithFields(logrus.Fields{
				"room": r.CourseId,
				"num":  r.OnlineNum(),
			}).Info("online number")
		case <-r.stopChan:
			return
		}
	}
}

func (r *Room) SendMessage(msg *Message) {
	switch msg.Type {
	case MsgOnline, MsgOffline:
		for _, clientId := range r.Clients {
			//if clientId.Client.User.Name != msg.Name {
			err := clientId.Client.Conn.WriteJSON(msg)
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"err": err.Error(),
				}).Error("write message failed")
			}
		}
		//}
	case MsgRobot:

	case MsgCommon:
		var msgs []*models.ChatHistory
		for _, clientId := range r.Clients {
			err := clientId.Client.Conn.WriteJSON(msg)
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"err": err.Error(),
				}).Error("write message failed")
			}
			msgs = append(msgs, &models.ChatHistory{
				RoomId:   clientId.Id,
				CourseId: r.CourseId,
				From:     0,
				To:       clientId.Client.User.Id,
				SentTm:   time.Unix(msg.SendTime, 0),
				InsertTm: time.Now(),
			})
		}
		err := dao.Create(context.TODO(), &msgs)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"msg": msg,
				"err": err.Error(),
			})
		}
	}
}
