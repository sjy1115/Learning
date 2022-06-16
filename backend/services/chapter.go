package services

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"learning/consts"
	"learning/dao"
	"learning/db/cache"
	"learning/db/mysql"
	"learning/models"
	"learning/pkg/context"
	"learning/proto"
	"learning/utils"
	"strconv"
	"time"
)

func ChapterListHandler(c *context.Context, req *proto.ChapterListRequest) (resp *proto.ChapterListResponse, err error) {

	resp = &proto.ChapterListResponse{}

	isTeacher := c.UserToken.Role == consts.RoleTeacher

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

		item := proto.ChapterListResponseItem{
			Id:           chapter.Id,
			Name:         chapter.Name,
			Introduction: chapter.Introduction,
			PdfUrl:       chapter.PdfUrl,
			VideoUrl:     chapter.VideoUrl,
			CreateAt:     chapter.InsertTm.Unix(),
		}

		if isTeacher {
			learnNum, err := dao.StudentLearnNumBetByChapterId(c.Ctx, chapter.Id)
			if err != nil {
				return nil, err
			}

			signNum, err := dao.StudentSignInNumHandler(c.Ctx, chapter.Id)
			if err != nil {
				return nil, err
			}

			total, err := dao.StudentNumGetByCourseId(c.Ctx, chapter.CourseId)
			if err != nil {
				return nil, err
			}

			item.LearnNum = learnNum
			item.Total = total
			item.SignNum = signNum
		} else {
			cus, err := dao.StudentIsLearnedByStudentId(c.Ctx, c.UserToken.UserId, chapter.Id)
			if err != nil {
				return nil, err
			}

			if len(cus) > 0 {
				item.Learned = cus[0].SignIn == 1 && cus[0].Status == 1
				item.PostSignIn = cus[0].SignIn
				item.SignIn = cus[0].Status
				item.Score = int64(cus[0].Score)
			}
		}

		resp.Items = append(resp.Items, &item)
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
		VideoUrl:     chapter.VideoUrl,
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

	exercise := models.Exercises{
		ChapterId: chapter.Id,
		InsertTm:  time.Now(),
		UpdateTm:  time.Now(),
	}
	err = dao.Create(c.Ctx, &exercise)
	if err != nil {
		return nil, err
	}

	return
}

func ChapterLearnHandler(c *context.Context, req *proto.ChapterLearnRequest) (resp *proto.ChapterLearnResponse, err error) {
	if c.UserToken.Role != consts.RoleStudent {
		return nil, fmt.Errorf("permission denied")
	}

	cu, err := dao.ChapterUserGetByChapterIdAndUserId(c.Ctx, c.UserToken.UserId, req.ChapterId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			cu := models.ChapterUser{
				UserID:    c.UserToken.UserId,
				ChapterID: req.ChapterId,
				InsertTm:  time.Now(),
				UpdateTm:  time.Now(),
			}
			err = dao.Create(c.Ctx, &cu)
			if err != nil {
				return nil, err
			}
		}
	} else {
		err = dao.ChapterUserUpdateByChapterIdAndUserId(c.Ctx, cu.UserID, cu.ChapterID, map[string]interface{}{})
		if err != nil {

		}
	}

	return
}

func ChapterDeleteHandler(c *context.Context, req *proto.ChapterDeleteRequest) (resp *proto.ChapterDeleteResponse, err error) {
	chapterIdStr := c.Param("id")
	chapterId, _ := strconv.Atoi(chapterIdStr)

	if c.UserToken.Role != consts.RoleTeacher {
		return nil, fmt.Errorf("permission denied")
	}

	resp = &proto.ChapterDeleteResponse{}

	err = dao.ChapterDeleteById(c.Ctx, chapterId)
	if err != nil {
		return nil, err
	}

	return
}

func PostSignInHandler(c *context.Context, req *proto.PostSignInRequest) (resp *proto.PostSignInResponse, err error) {
	if c.UserToken.Role != consts.RoleTeacher {
		return nil, fmt.Errorf("permission denied")
	}

	code := utils.RandomString(4)

	err = cache.SetEx(c.Ctx, code, cache.SignInKey(req.ChapterId), 5*60)
	if err != nil {
		return
	}

	cu, err := dao.ChapterUserGetByChapterIdAndUserId(c.Ctx, c.UserToken.UserId, req.ChapterId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			cu = models.ChapterUser{
				UserID:    c.UserToken.UserId,
				ChapterID: req.ChapterId,
				SignIn:    1,
				InsertTm:  time.Now(),
				UpdateTm:  time.Now(),
			}
			err = mysql.GetRds(c.Ctx).Create(&cu).Error
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	} else {
		err = dao.ChapterUserUpdateByChapterIdAndUserId(c.Ctx, cu.UserID, cu.ChapterID, map[string]interface{}{
			"sign_in":   1,
			"update_tm": time.Now(),
		})
	}

	resp = &proto.PostSignInResponse{
		Code: code,
	}

	return
}

func SignInHandler(c *context.Context, req *proto.SignInRequest) (resp *proto.SignInResponse, err error) {
	if c.UserToken.Role != consts.RoleStudent {
		return nil, fmt.Errorf("permission denied")
	}

	code, err := cache.Get(c.Ctx, cache.SignInKey(req.ChapterId))
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, fmt.Errorf("签到已过期")
		}
		return nil, err
	}

	if code != req.Code {
		return nil, fmt.Errorf("验证码不正确")
	}

	cu, err := dao.ChapterUserGetByChapterIdAndUserId(c.Ctx, c.UserToken.UserId, req.ChapterId)
	if err != nil {
		return nil, err
	}

	err = dao.ChapterUserUpdateByChapterIdAndUserId(c.Ctx, cu.UserID, cu.ChapterID, map[string]interface{}{
		"status":    1,
		"update_tm": time.Now(),
	})
	if err != nil {
		return nil, err
	}

	resp = &proto.SignInResponse{
		Status: "OK",
	}

	return
}
