package scoreView

import (
	"github.com/gin-gonic/gin"
	"studentScoreManagement/model"
	"studentScoreManagement/util"
)

type Service interface {
	GetScoreByID(ctx *gin.Context, req *GetScoreByIDRequest) *GetScoreByIDResponse
	GetScoreByClass(ctx *gin.Context, req *GetScoreByClassRequest) *GetScoreByClassResponse
	GetClass(ctx *gin.Context) *GetClassResponse
}

type GetScoreByClassRequest struct {
	Class   string  `form:"class" binding:"required"`
	Subject *string `form:"subject,omitempty"`
}

type GetScoreByClassResponse struct {
	Scores []model.Score
	Base   *util.Base
}

type GetScoreByIDRequest struct {
	ID      string  `form:"id" binding:"required"`
	Subject *string `form:"subject,omitempty"`
}

type GetScoreByIDResponse struct {
	Scores []model.Score
	Base   *util.Base
}

type GetClassResponse struct {
	Classes []string
	Base    *util.Base
}
