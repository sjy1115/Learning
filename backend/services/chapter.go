package services

import (
	"fmt"
	"learning/consts"
	"learning/dao"
	"learning/db/mysql"
	"learning/models"
	"learning/pkg/context"
	"learning/proto"
	"learning/utils"
	"strconv"
	"time"
)

func ChapterListHandler(c *context.Context, req *proto.ChapterListRequest) (resp *proto.ChapterListResponse, err error) {
	if c.UserToken.Role != consts.RoleTeacher {
		return nil, fmt.Errorf("permission denied")
	}

	resp = &proto.ChapterListResponse{}

	db := mysql.GetRds(c.Ctx).
		Model(&models.Chapter{}).
		Where("course_id = ?", req.CourseId)

	if len(req.Name) != 0 {
		db = db.Where("name like ?", utils.EscapeLIKE("%s"+req.Name+"%"))
	}

	err = db.Count(&resp.Total).Error
	if err != nil {
		return nil, err
	}

	var chapters []*models.Chapter
	err = db.Offset(utils.GetStartPage(req.Page, req.PageSize)).Limit(req.PageSize).Find(&chapters).Error
	if err != nil {
		return nil, err
	}

	for _, chapter := range chapters {
		// TODO learned and total

		resp.Items = append(resp.Items, &proto.ChapterListResponseItem{
			Id:           chapter.Id,
			Name:         chapter.Name,
			Introduction: chapter.Introduction,
			PdfUrl:       chapter.PdfUrl,
			CreateAt:     chapter.InsertTm.Unix(),
		})
	}

	return
}

func ChapterDetailHandler(c *context.Context, req *proto.ChapterDetailRequest) (resp *proto.ChapterDetailResponse, err error) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	chapter, err := dao.ChapterGetById(c.Ctx, id)
	if err != nil {
		return nil, err
	}

	fmt.Println(chapter)

	resp = &proto.ChapterDetailResponse{
		Id:           chapter.Id,
		Name:         chapter.Name,
		Introduction: chapter.Introduction,
		PdfUrl:       chapter.PdfUrl,
		CreateAt:     chapter.InsertTm.Unix(),
	}

	return resp, nil
}

func ChapterUpdateHandler(c *context.Context, req *proto.ChapterUpdateRequest) (resp *proto.ChapterUpdateResponse, err error) {
	chapter, err := dao.ChapterGetById(c.Ctx, req.Id)
	if err != nil {
		return nil, err
	}

	data := make(map[string]interface{})
	// TODO check unique
	if len(req.Name) != 0 {
		data["name"] = req.Name
	}
	if len(req.Introduction) != 0 {
		data["introduction"] = req.Introduction
	}

	err = dao.ChapterUpdateById(c.Ctx, req.Id, data)
	if err != nil {
		return nil, err
	}
	resp.Id = chapter.Id

	return resp, err
}

func ChapterCreateHandler(c *context.Context, req *proto.ChapterCreateRequest) (resp *proto.ChapterCreateResponse, err error) {
	if c.UserToken.Role != consts.RoleTeacher {
		return nil, fmt.Errorf("permission denied")
	}

	resp = &proto.ChapterCreateResponse{}

	chapter := models.Chapter{
		Name:         req.Name,
		CourseId:     req.CourseId,
		Introduction: req.Introduction,
		PdfUrl:       req.Pdf,
		VideoUrl:     req.Video,
		InsertTm:     time.Now(),
		UpdateTm:     time.Now(),
	}

	err = dao.Create(c.Ctx, &chapter)
	if err != nil {
		return nil, err
	}

	resp.Id = chapter.Id

	return
}

func ChapterDeleteHandler(c *context.Context, req *proto.ChapterDeleteRequest) (resp *proto.ChapterDeleteResponse, err error) {
	if c.UserToken.Role != consts.RoleTeacher {
		return nil, fmt.Errorf("permission denied")
	}

	resp = &proto.ChapterDeleteResponse{}

	err = dao.ChapterDeleteById(c.Ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return
}
