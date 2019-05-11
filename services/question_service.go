package services

import (
	"ts-server/models"
	"ts-server/repositories"
)

type QuestionService struct {
	QuestionRepo *repositories.QuestionRepository
}

func NewQuestionService(questionRepo *repositories.QuestionRepository) *QuestionService {
	var qs = QuestionService{questionRepo}
	return &qs
}

func (this *QuestionService) GetByIndex(index int64) (*models.Question, error) {
	return this.QuestionRepo.FindOneByIndex(index)
}
