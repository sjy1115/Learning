package services

import (
	"errors"
	"fmt"
	"learning/pkg/chat"
	"learning/pkg/jwt"
	"net/http"
	"strconv"
	"time"

	"learning/consts"
	"learning/dao"
	"learning/db/mysql"
	"learning/models"
	"learning/pkg/context"
	"learning/pkg/oss"
	"learning/proto"
	"learning/utils"

	"github.com/gorilla/websocket"
)

func CourseListHandler(c *context.Context, req *proto.CourseListRequest) (resp *proto.CourseListResponse, err error) {

	resp = &proto.CourseListResponse{}

	isTeacher := c.UserToken.Role == consts.RoleTeacher

	db := mysql.GetRds(c.Ctx).
		Model(&models.Course{}).
		Joins("JOIN course_user cu ON cu.course_id = course.id AND cu.user_id = ?", c.UserToken.UserId)

	if len(req.Name) != 0 {
		name := "%" + req.Name + "%"
		db = db.Where("name LIKE ?", name)
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
		item := &proto.CourseListResponseItem{
			ID:       course.Id,
			Name:     course.Name,
			Semester: course.Semester,
			CreateTm: course.InsertTm.Unix(),
		}

		teacher, err := dao.TeacherGetByCourseId(c.Ctx, course.Id)
		if err != nil {
			return nil, err
		}
		item.Teacher = teacher.Name
		item.TeacherAvatar = teacher.Avatar

		if isTeacher {
			studentNum, err := dao.StudentNumGetByCourseId(c.Ctx, course.Id)
			if err != nil {
				return nil, err
			}
			item.StudentNum = studentNum
			item.InviteCode = course.InviteCode
		} else {
			item.Avatar = course.Avatar
			item.Introduction = course.Introduction
		}

		resp.Items = append(resp.Items, item)
	}

	return
}

func CourseDetailHandler(c *context.Context, req *proto.CourseDetailRequest) (resp *proto.CourseDetailResponse, err error) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	course, err := dao.CourseGetById(c.Ctx, id)
	if err != nil {
		return nil, err
	}

	teacher, err := dao.TeacherGetByCourseId(c.Ctx, course.Id)
	if err != nil {
		return nil, err
	}

	resp = &proto.CourseDetailResponse{
		ID:           course.Id,
		Name:         course.Name,
		Semester:     course.Semester,
		Teacher:      teacher.Name,
		Avatar:       course.Avatar,
		Introduction: course.Introduction,
		//StudentNum: course.StudentNum,
		CreateTm: course.InsertTm.Unix(),
	}

	return
}

func CourseUpdateHandler(c *context.Context, req *proto.CourseUpdateRequest) (resp *proto.CourseUpdateResponse, err error) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	course, err := dao.CourseGetById(c.Ctx, id)
	if err != nil {
		return nil, err
	}

	data := make(map[string]interface{})
	if len(req.Name) != 0 {
		data["name"] = req.Name
	}
	if len(req.Semester) != 0 {
		data["semester"] = req.Semester
	}
	if len(req.Introduction) != 0 {
		data["introduction"] = req.Introduction
	}

	err = dao.CourseUpdateById(c.Ctx, id, data)
	if err != nil {
		return nil, err
	}

	resp = &proto.CourseUpdateResponse{
		Id: course.Id,
	}

	return
}

func CourseCreateHandler(c *context.Context, req *proto.CourseCreateRequest) (resp *proto.CourseCreateResponse, err error) {
	if c.UserToken.Role != consts.RoleTeacher {
		return nil, fmt.Errorf("permission denied")
	}

	tx := mysql.GetRds(c.Ctx).Begin()

	resp = &proto.CourseCreateResponse{}

	course := &models.Course{
		Name:         req.Name,
		Semester:     req.Semester,
		Introduction: req.Introduction,
		Avatar:       req.Avatar,
		InviteCode:   utils.RandomString(32),
		InsertTm:     time.Now(),
		UpdateTm:     time.Now(),
	}

	err = dao.Create(c.Ctx, course, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	courseUser := models.CourseUser{
		CourseID: course.Id,
		UserID:   c.UserToken.UserId,
		Status:   consts.OK,
		InsertTm: time.Now(),
		UpdateTm: time.Now(),
	}

	err = dao.Create(c.Ctx, &courseUser, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	resp.Id = course.Id

	return
}

func CourseDeleteHandler(c *context.Context, req *proto.CourseDeleteRequest) (resp *proto.CourseDeleteResponse, err error) {
	courseIdStr := c.Param("id")
	courseId, _ := strconv.Atoi(courseIdStr)

	if c.UserToken.Role != consts.RoleTeacher {
		return nil, fmt.Errorf("permission denied")
	}

	resp = &proto.CourseDeleteResponse{}

	err = dao.CourseDeleteById(c.Ctx, courseId)
	if err != nil {
		return nil, err
	}

	return
}

func ListChapterStudentHandler(c *context.Context, req *proto.ListChapterStudentRequest) (resp *proto.ListChapterStudentResponse, err error) {
	if c.UserToken.Role != consts.RoleTeacher {
		return nil, errors.New("permission denied")
	}

	resp = new(proto.ListChapterStudentResponse)

	var studentScores []struct {
		Id     int64  `json:"id" gorm:"column:id"`
		Name   string `json:"name" gorm:"column:name"`
		Avatar string `json:"avatar" gorm:"column:avatar"`
	}

	err = mysql.GetRds(c.Ctx).
		Model(&models.CourseUser{}).
		Joins("LEFT JOIN user u ON u.id = course_user.user_id AND u.role = ?", consts.RoleStudent).
		Where("course_id = ?", req.CourseId).
		Count(&resp.Total).
		Offset(utils.GetStartPage(req.Page, req.PageSize)).
		Limit(req.PageSize).
		Select("u.id AS id, u.name AS name, u.avatar AS avatar").
		Scan(&studentScores).
		Error
	if err != nil {
		return nil, err
	}

	for _, studentScore := range studentScores {
		var learned bool
		score, exist, err := dao.ScoreGetByChapterIdAndUserId(c.Ctx, int(studentScore.Id), req.ChapterId)
		if err != nil {
			return nil, err
		}
		if exist {
			learned = true
		}

		item := proto.ListCourseStudentItem{
			Id:      studentScore.Id,
			Name:    studentScore.Name,
			Learned: learned,
			Score:   int64(score),
			Avatar:  studentScore.Avatar,
		}
		resp.Items = append(resp.Items, &item)
	}

	return
}

func StudentSignInHandler(ctx *context.Context, req *proto.StudentSignInRequest) (resp *proto.StudentSignInResponse, err error) {
	if ctx.UserToken.Role != consts.RoleStudent {
		return nil, errors.New("permission denied")
	}

	return
}

func JoinCourseHandler(c *context.Context, req *proto.JoinCourseRequest) (resp *proto.JoinCourseResponse, err error) {
	if c.UserToken.Role != consts.RoleStudent {
		return nil, fmt.Errorf("permission denied")
	}

	resp = &proto.JoinCourseResponse{}

	course, err := dao.CourseGetByInviteCode(c, req.InviteCode)
	if err != nil {
		return nil, err
	}

	student, err := dao.UserGetById(c.Ctx, c.UserToken.UserId)
	if err != nil {
		return nil, err
	}

	userCourse := models.CourseUser{
		CourseID: course.Id,
		UserID:   student.Id,
		Status:   consts.OK,
		InsertTm: time.Now(),
		UpdateTm: time.Now(),
	}

	err = dao.Create(c.Ctx, &userCourse)
	if err != nil {
		return nil, err
	}

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

	go chat.Process(c.Ctx, req.CourseId, userToken.UserId, conn)

	return nil
}
