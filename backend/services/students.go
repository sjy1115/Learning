package services

import (
	"learning/dao"
	"learning/models"
	"learning/pkg/context"
	"learning/proto"
	"learning/utils"
	"path/filepath"
)

func StudentsListHandler(c *context.Context, req *proto.StudentListReq) (*proto.StudentListResp, error) {
	students, err := dao.GetPageUser(c.Request.Context(), req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}

	resp := proto.StudentListResp{
		Results: make([]proto.StudentListItem, 0),
	}

	for _, student := range students {
		item := proto.StudentListItem{
			Id:    student.Id,
			Name:  student.Name,
			Sex:   student.Sex,
			Age:   student.Age,
			Phone: student.Phone,
			Photo: student.Photo,
		}
		resp.Results = append(resp.Results, item)
	}

	return &resp, nil
}

func StudentCreateHandler(c *context.Context, req *proto.StudentCreateReq) (*proto.StudentCreateResp, error) {
	fh, err := c.FormFile("file")
	if err != nil {
		return nil, err
	}

	// TODO OSS support
	// dest := oss.put()

	// 先随机生成字符串
	dest := filepath.Join("", utils.RandomString(32)+".png")

	err = c.SaveUploadedFile(fh, dest)
	if err != nil {
		return nil, err
	}

	student := models.User{
		Name:  req.Name,
		Age:   req.Age,
		Sex:   req.Sex,
		Phone: req.Phone,
		Photo: dest,
	}

	err = dao.CreateUser(c.Request.Context(), &student)
	if err != nil {
		return nil, err
	}

	return &proto.StudentCreateResp{
		Id: student.Id,
	}, nil
}

func StudentDetailHandler(c *context.Context, req *proto.StudentDetailReq) (*proto.StudentDetailResp, error) {
	student, err := dao.GetUserById(c.Request.Context(), req.Id)
	if err != nil {
		return nil, err
	}

	return &proto.StudentDetailResp{
		Name:  student.Name,
		Age:   student.Age,
		Sex:   student.Sex,
		Phone: student.Phone,
		Photo: student.Photo,
	}, nil
}

func StudentUpdateHandler(c *context.Context, req *proto.StudentUpdateReq) (*proto.StudentUpdateResp, error) {

	return nil, nil
}

func StudentDeleteHander(c *context.Context, req *proto.StudentDeleteReq) error {
	err := dao.DeleteUserById(c.Request.Context(), req.Id)
	return err
}
