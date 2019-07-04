package repositories

import (
	"ts-server/db"
	"ts-server/models"
)

type ExamRepository struct {
}

func NewExamRepository() *ExamRepository {
	return &ExamRepository{}
}

func (this *ExamRepository) GetAll() ([]*models.Exam, error) {
	var examList []*models.Exam

	err := db.Pgdb.Find(&examList).Error
	if err != nil {
		return nil, err
	}
	return examList, nil
}

func (this *ExamRepository) GetById(id int64) (*models.Exam, error) {
	var exam models.Exam

	err := db.Pgdb.Where("id = ?", id).Find(&exam).Error
	if err != nil {
		return nil, err
	}
	return &exam, nil
}

func (this *ExamRepository) GetAllWithQuestions() ([]*models.Exam, error) {
	var examList []*models.Exam

	err := db.Pgdb.Preload("Questions").Preload("Questions.Answers").Find(&examList).Error
	if err != nil {
		return nil, err
	}
	return examList, nil
}

func (this *ExamRepository) GetByIdWithQuestion(id int64) (*models.Exam, error) {
	var exam models.Exam

	err := db.Pgdb.Preload("Questions").Preload("Questions.Answers").Where("id = ?", id).Find(&exam).Error
	if err != nil {
		return nil, err
	}
	return &exam, err
}
