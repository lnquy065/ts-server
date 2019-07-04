package repositories

import (
	"ts-server/db"
	"ts-server/models"
)

type QuestionRepository struct {
}

func NewQuestionRepository() *QuestionRepository {
	return &QuestionRepository{}
}

func (this *QuestionRepository) FindOneByIndex(index int64) (*models.Question, error) {
	var q models.Question

	if err := db.Pgdb.Preload("Answers").
		Where("index = ?", index).
		Find(&q).Error; err != nil {
		return nil, err
	}

	return &q, nil
}

func (this *QuestionRepository) FindAll() ([]*models.Question, error) {
	var questionList []*models.Question

	err := db.Pgdb.Preload("Answers").Find(&questionList).Error
	if err != nil {
		return nil, err
	}

	return questionList, nil
}
