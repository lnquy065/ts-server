package services

import (
	"ts-server/models"
	"ts-server/repositories"
)

type ExamService struct {
	ExamRepository *repositories.ExamRepository
}

func NewExamService(examRepo *repositories.ExamRepository) *ExamService {
	return &ExamService{examRepo}
}

func (this *ExamService) GetAll() ([]*models.Exam, error) {
	return this.ExamRepository.GetAll()
}

func (this *ExamService) GetById(id int64) (*models.Exam, error) {
	return this.ExamRepository.GetById(id)
}

func (this *ExamService) GetByIdWithQuestions(id int64) (*models.Exam, error) {
	return this.ExamRepository.GetByIdWithQuestion(id)
}

func (this *ExamService) GetAllWithQuestions() ([]*models.Exam, error) {
	return this.ExamRepository.GetAllWithQuestions()
}
