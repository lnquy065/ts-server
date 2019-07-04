package models

import "github.com/jinzhu/gorm"

type Exam struct {
	gorm.Model
	Index       int
	LicenseType string
	Question    []Question `gorm:"many2many:examination.exam_question;"`
}

func (this *Exam) TableName() string {
	return "examination.ts_exam"
}
