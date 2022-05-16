package chat

import (
	"context"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

var (
	Pool  sync.Pool
	Rooms = make(map[int]*Room)
)

func init() {
	Pool = sync.Pool{
		New: func() interface{} {
			return &Room{
				Clients: make([]*ClientId, 0),

				OnlineChan:  make(chan *Client, 10),
				OfflineChan: make(chan *Client, 10),
				MessageChan: make(chan *Message, 10),
			}
		},
	}

	go func() {
		for {
			select {
			case <-time.After(time.Second * 60 * 5):
				logrus.WithFields(logrus.Fields{
					"rooms": Rooms,
				}).Info("chat rooms")
			}
		}
	}()
}

func Process(ctx context.Context, courseId int, userId int, conn *websocket.Conn) {
	room, ok := Rooms[courseId]
	if !ok {
		room = Pool.Get().(*Room)
		room.CourseId = courseId
		Rooms[courseId] = room

		logrus.WithFields(logrus.Fields{
			"courseId": courseId,
		}).Info("New room created")

		go room.Broadcast()
	}
	room.Process(ctx, conn, userId)
}

func release(room *Room) {
	Pool.Put(room)
	delete(Rooms, room.CourseId)
}
