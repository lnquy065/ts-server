package models

import (
	"github.com/jinzhu/gorm"
)

type Question struct {
	gorm.Model
	Question     string
	QuestionImg  string
	QuestionType string
	LicenseType  string
	Index        int
	Answers      []Answer `gorm:"foreignkey:QuestionId"`
}

func (this *Question) TableName() string {
	return "examination.ts_question"
}
