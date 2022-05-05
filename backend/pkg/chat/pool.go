package chat

import (
	"context"
	"github.com/gorilla/websocket"
	"sync"
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

				OnlineChan:  make(chan *Client, 0),
				OfflineChan: make(chan *Client, 0),
				MessageChan: make(chan *Message, 0),
			}
		},
	}
}

func Process(ctx context.Context, courseId int, userId int, conn *websocket.Conn) {
	room, ok := Rooms[courseId]
	if !ok {
		room = Pool.Get().(*Room)
		room.CourseId = courseId
		Rooms[courseId] = room

		go room.Broadcast(context.TODO())
	}
	room.Process(ctx, conn, userId)
}
