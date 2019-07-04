package services

import (
	"ts-server/models"
	"ts-server/repositories"
)

type ExamService struct {
	ExamRepository     *repositories.ExamRepository
	QuestionRepository *repositories.QuestionRepository
}

func NewExamService(examRepo *repositories.ExamRepository, questRepo *repositories.QuestionRepository) *ExamService {
	return &ExamService{examRepo, questRepo}
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

func (this *ExamService) GetRandom(
	pack1Type string, pack1Quantity int64,
	pack2Type string, pack2Quantity int64,
	licenseType string) (*models.Exam, error) {

	var questionList []*models.Question

	// Get pack1
	q1List, error := this.QuestionRepository.GetRandom(pack1Quantity, pack1Type, licenseType)
	if error != nil {
		return nil, error
	}
	// Get pack2
	q2List, error := this.QuestionRepository.GetRandom(pack2Quantity, pack2Type, licenseType)
	if error != nil {
		return nil, error
	}

	questionList = append(q1List, q2List...)

	exam := &models.Exam{
		Index:       -1,
		LicenseType: licenseType,
		Questions:   questionList,
	}

	return exam, nil
}
