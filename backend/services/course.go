package services

import (
	lcontext "context"
	"fmt"
	"learning/consts"
	"learning/dao"
	"learning/db/mysql"
	"learning/models"
	"learning/pkg/chat"
	"learning/pkg/context"
	"learning/pkg/jwt"
	"learning/pkg/oss"
	"learning/proto"
	"learning/utils"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func CourseListHandler(c *context.Context, req *proto.CourseListRequest) (resp *proto.CourseListResponse, err error) {

	if c.UserToken.Role != consts.ROLE_STUDENT {
		return nil, fmt.Errorf("permission denied")
	}

	resp = &proto.CourseListResponse{}

	db := mysql.GetRds(c.Ctx).Model(&models.Course{})

	if len(req.Name) != 0 {
		db = db.Where("name = ?", req.Name)
	}

	if len(req.Semester) != 0 {
		db = db.Where("semester = ?", req.Semester)
	}

	err = db.Count(&resp.Total).Error
	if err != nil {
		return nil, err
	}

	var courses []*models.Course
	err = db.Offset(utils.GetStartPage(req.Page, req.PageSize)).Limit(req.PageSize).Find(&courses).Error
	if err != nil {
		return nil, err
	}

	for _, course := range courses {
		teacher, err := dao.TeacherGetByCourseId(c.Ctx, course.Id)
		if err != nil {
			return nil, err
		}

		studentNum, err := dao.StudentNumGetByCourseId(c.Ctx, course.Id)
		if err != nil {
			return nil, err
		}

		// TODO invite code

		item := &proto.CourseListResponseItem{
			ID:         course.Id,
			Name:       course.Name,
			Semester:   course.Semester,
			Teacher:    teacher.Name,
			StudentNum: studentNum,
			CreateTm:   course.InsertTm.Unix(),
		}

		resp.Items = append(resp.Items, item)
	}

	return
}

func CourseDetailHandler(c *context.Context, req *proto.CourseDetailRequest) (resp *proto.CourseDetailResponse, err error) {

	return
}

func CourseCreateHandler(c *context.Context, req *proto.CourseCreateRequest) (resp *proto.CourseCreateResponse, err error) {
	if c.UserToken.Role != consts.ROLE_TEACHER {
		return nil, fmt.Errorf("permission denied")
	}

	resp = &proto.CourseCreateResponse{}

	course := &models.Course{
		Name:         req.Name,
		Semester:     req.Semester,
		Introduction: req.Introduction,
		InsertTm:     time.Now(),
		UpdateTm:     time.Now(),
	}

	tx := mysql.GetRds(c.Ctx).Begin()

	err = dao.Create(tx, course)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	fh, err := c.FormFile("avatar")
	if err != nil {
		return nil, err
	}

	file, err := fh.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	filePath := fmt.Sprintf("%s/%d/%s", consts.AvatarPrefix, course.Id, fh.Filename)

	err = oss.Bucket.PutReader(filePath, file, fh.Size)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = dao.UpdateColumn(tx, course.Id, map[string]interface{}{
		"avatar":    filePath,
		"update_tm": time.Now(),
	})
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	resp.Id = course.Id

	return
}

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
		utils.Error(c, utils.ErrorCode, err)
		return err
	}

	token := req.Token

	userToken, err := jwt.ParseToken(token)
	if err != nil {
		utils.Error(c, utils.ErrorCode, err)
		return err
	}

	go chat.RoomInstance.Process(lcontext.TODO(), conn, userToken.UserId)

	return nil
}
