package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"ts-server/services"
)

type ExamController struct {
	ExamService *services.ExamService
}

func NewExamController(examSvc *services.ExamService) *ExamController {
	return &ExamController{examSvc}
}

func (this *ExamController) GetAll(ctx *gin.Context) {
	examList, err := this.ExamService.GetAll()
	if err != nil {
		ctx.String(http.StatusNotFound, "NotFound")
	}
	ctx.JSON(http.StatusOK, examList)
}

func (this *ExamController) GetById(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	idInt, _ := strconv.ParseInt(id, 10, 64)
	exam, err := this.ExamService.GetById(idInt)
	if err != nil {
		ctx.String(http.StatusNotFound, "NotFound")
	}
	ctx.JSON(http.StatusOK, exam)
}

func (this *ExamController) GetByIdWithQuestions(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	idInt, _ := strconv.ParseInt(id, 10, 64)
	exam, err := this.ExamService.GetByIdWithQuestions(idInt)
	if err != nil {
		ctx.String(http.StatusNotFound, "NotFound")
	}
	ctx.JSON(http.StatusOK, exam)
}

func (this *ExamController) GetAllWithQuestions(ctx *gin.Context) {
	examList, err := this.ExamService.GetAllWithQuestions()
	if err != nil {
		ctx.String(http.StatusNotFound, "NotFound")
	}
	ctx.JSON(http.StatusOK, examList)
}

func (this *ExamController) GetRandom(ctx *gin.Context) {
	licenseType := ctx.Params.ByName("licenseType")
	var pack1Type, pack2Type string
	var pack1QuantityInt, pack2QuantityInt int64

	switch licenseType {
	case "B2":
		{
			pack1Type = "LT"
			pack1QuantityInt = 15
			pack2Type = "TH"
			pack2QuantityInt = 15
		}
	}

	exam, err := this.ExamService.GetRandom(pack1Type, pack1QuantityInt, pack2Type, pack2QuantityInt, licenseType)
	if err != nil {
		ctx.String(http.StatusNotFound, "NotFound")
	}
	ctx.JSON(http.StatusOK, exam)
}
