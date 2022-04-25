package models

import "time"

type Chapter struct {
	Id           int       `json:"id" gorm:"primary_key;column:id"`
	CourseId     int       `json:"course_id" column:"course_id"`
	Name         string    `json:"name" column:"name"`
	Introduction string    `json:"introduction" column:"introduction"`
	PdfUrl       string    `json:"pdf_url" gorm:"column:pdf_url"`
	VideoUrl     string    `json:"video_url" gorm:"column:video_url"`
	InsertTm     time.Time `json:"insert_tm" gorm:"column:insert_tm"`
	UpdateTm     time.Time `json:"update_tm" gorm:"column:update_tm"`
}

func (C *Chapter) TableName() string {
	return "chapter"
}
