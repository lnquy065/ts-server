package models

import "github.com/jinzhu/gorm"

type Answer struct {
	gorm.Model
	QuestionId int
	Answer     string
	IsCorrect  bool
}

func (this *Answer) TableName() string {
	return "examination.ts_answer"
}
