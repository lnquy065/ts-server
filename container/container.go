package container

import (
	"ts-server/apis/v1"
	"ts-server/repositories"
	"ts-server/services"
)

func CreateQuestionController() *apis.QuestionController {
	questionRepository := repositories.NewQuestionRepository()
	questionService := services.NewQuestionService(questionRepository)
	questionController := apis.NewQuestionController(questionService)
	return questionController
}

func CreateExamController() *apis.ExamController {
	examRepository := repositories.NewExamRepository()
	examService := services.NewExamService(examRepository)
	examController := apis.NewExamController(examService)
	return examController
}
