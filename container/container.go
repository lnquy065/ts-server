package container

import (
	"github.com/google/wire"
	"ts-server/apis/v1"
	"ts-server/repositories"
	"ts-server/services"
)

func CreateQuestionController() *apis.QuestionController {
	panic(wire.Build(
		repositories.NewQuestionRepository,
		services.NewQuestionService,
		apis.NewQuestionController))
	return &apis.QuestionController{}
}
