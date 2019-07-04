package main

import (
	"github.com/gin-gonic/gin"
	"ts-server/container"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	// Controllers
	questionCtrl := container.CreateQuestionController()
	examCtrl := container.CreateExamController()

	apiV1 := router.Group("v1")
	{
		questions := apiV1.Group("questions")
		{
			questions.GET("/:index", questionCtrl.GetByIndex)
			questions.GET("/", questionCtrl.GetAll)
		}

		exams := apiV1.Group("exams-non-questions")
		{
			exams.GET("/:id", examCtrl.GetById)
			exams.GET("/", examCtrl.GetAll)
		}

		examsWithQ := apiV1.Group("exams-inc-questions")
		{
			examsWithQ.GET("/:id", examCtrl.GetByIdWithQuestions)
			examsWithQ.GET("/", examCtrl.GetAllWithQuestions)
		}

		examsWithQRandom := apiV1.Group("exams/random")
		{
			examsWithQRandom.GET("/licenseType/:licenseType", examCtrl.GetRandom)
		}
	}
	return router
}
