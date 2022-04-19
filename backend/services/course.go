package services

import (
	lcontext "context"
	"github.com/gorilla/websocket"
	"learning/pkg/chat"
	"learning/pkg/context"
	"learning/pkg/jwt"
	"learning/pkg/oss"
	"learning/proto"
	"learning/utils"
	"net/http"
)

func UploadCourseHandler(c *context.Context, req *proto.UploadCourseRequest) (resp interface{}, err error) {
	//fileName := req.Filename

	fh, err := c.FormFile("file")
	if err != nil {
		return nil, err
	}

	file, err := fh.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	err = oss.Bucket.PutReader(fh.Filename, file, fh.Size)
	if err != nil {
		return nil, err
	}

	return
}

func StartChatHandler(c *context.Context, req *proto.StartChatRequest) error {
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}).Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		utils.Error(c, err)
		return err
	}

	token := req.Token

	userToken, err := jwt.ParseToken(token)
	if err != nil {
		utils.Error(c, err)
		return err
	}

	go chat.RoomInstance.Process(lcontext.TODO(), conn, userToken.UserId)

	return nil
}
