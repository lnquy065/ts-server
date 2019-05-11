package models

import "github.com/jinzhu/gorm"

type Question struct {
	gorm.Model
	Question    string
	QuestionImg string
	LicenseType string
	Index       int
	Answers     []Answer `gorm:"foreignkey:QuestionId"`
}

type QuestionDTO struct {
}

func (this *Question) TableName() string {
	return "examination.ts_question"
}
