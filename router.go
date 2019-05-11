package main

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	// Controllers
	questionCtrl := CreateQuestionController()

	apiV1 := router.Group("v1")
	{
		questions := apiV1.Group("questions")
		{
			questions.GET("/:index", questionCtrl.GetByIndex)
		}
	}
	return router
}
