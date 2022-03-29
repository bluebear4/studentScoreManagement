package score

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	"studentScoreManagement/model"
	"studentScoreManagement/util"
)

type Service interface {
	GetSubjects(ctx *gin.Context) *GetSubjectResponse

	AddScore(ctx *gin.Context, req *AddScoreRequest) *AddScoreResponse
	UploadScore(ctx *gin.Context, req *UploadScoreRequest) *UploadScoreResponse
	UpdateScore(ctx *gin.Context, req *UpdateScoreRequest) *UpdateScoreResponse
	DeleteScore(ctx *gin.Context, req *DeleteScoreRequest) *DeleteScoreResponse

	GetScoresByID(ctx *gin.Context, req *GetScoresByIDRequest) *GetScoresByIDResponse
	GetScoresByClass(ctx *gin.Context, req *GetScoresByClassRequest) *GetScoresByClassResponse
}

type GetScoresByClassRequest struct {
	Class   string  `form:"class"  binding:"required" json:"class,omitempty"`
	Subject *string `form:"subject,omitempty" json:"subject,omitempty"`
}

type GetScoresByClassResponse struct {
	Scores []model.Score
	Base   *util.Base
}

type GetScoresByIDRequest struct {
	ID      string  `form:"id"  binding:"required" json:"id,omitempty"`
	Subject *string `form:"subject,omitempty" json:"subject,omitempty"`
}

type GetScoresByIDResponse struct {
	Scores []model.Score
	Base   *util.Base
}

type GetSubjectResponse struct {
	Subjects []string
	Base     *util.Base
}

type UploadScoreRequest struct {
	File *excelize.File
}

type UploadScoreResponse struct {
	SuccessCount int
	FailCount    int
	Base         *util.Base
}

type AddScoreRequest struct {
	UserID  string  `form:"user_id"  binding:"required" json:"user_id,omitempty"`
	Name    string  `form:"name" binding:"required" json:"name,omitempty"`
	Subject string  `form:"subject" binding:"required" json:"subject,omitempty"`
	Score   float64 `form:"score" binding:"required" json:"score,omitempty"`
}

type AddScoreResponse struct {
	Base *util.Base
}

type UpdateScoreRequest struct {
	UserID  string  `form:"user_id"  binding:"required" json:"user_id,omitempty"`
	Subject string  `form:"subject" binding:"required" json:"subject,omitempty"`
	Name    *string `form:"name" json:"name,omitempty"`
	Score   float64 `form:"score" binding:"required" json:"score,omitempty"`
}

type UpdateScoreResponse struct {
	Base *util.Base
}

type DeleteScoreRequest struct {
	UserID  string `form:"user_id"  binding:"required" json:"user_id,omitempty"`
	Subject string `form:"subject" binding:"required" json:"subject,omitempty"`
}

type DeleteScoreResponse struct {
	Base *util.Base
}
