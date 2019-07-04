package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"ts-server/services"
)

type QuestionController struct {
	QuestionSvc *services.QuestionService
}

func NewQuestionController(questionSvc *services.QuestionService) *QuestionController {
	q := QuestionController{questionSvc}
	return &q
}

func (this *QuestionController) GetByIndex(ctx *gin.Context) {
	index := ctx.Params.ByName("index")
	indexInt, _ := strconv.ParseInt(index, 10, 64)
	q, e := this.QuestionSvc.GetByIndex(indexInt)
	if e != nil {
		ctx.String(http.StatusNotFound, "NotFound")
	}
	//obj, _ := json.Marshal(q)
	ctx.JSON(http.StatusOK, q)
}

func (this *QuestionController) GetAll(ctx *gin.Context) {
	questionList, error := this.QuestionSvc.GetAll()
	if error != nil {
		ctx.String(http.StatusNotFound, "NotFound")
	}
	ctx.JSON(http.StatusOK, questionList)
}
