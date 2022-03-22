package teacher

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	"studentScoreManagement/util"
)

type Service interface {
	UploadInfo(ctx *gin.Context, req *UploadInfoRequest) *UploadInfoResponse
	AddInfo(ctx *gin.Context, req *AddInfoRequest) *AddInfoResponse
	UpdateInfo(ctx *gin.Context, req *UpdateInfoRequest) *UpdateInfoResponse
	DeleteInfo(ctx *gin.Context, req *DeleteInfoRequest) *DeleteInfoResponse

	UploadScore(ctx *gin.Context, req *UploadScoreRequest) *UploadScoreResponse
	AddScore(ctx *gin.Context, req *AddScoreRequest) *AddScoreResponse
	UpdateScore(ctx *gin.Context, req *UpdateScoreRequest) *UpdateScoreResponse
	DeleteScore(ctx *gin.Context, req *DeleteScoreRequest) *DeleteScoreResponse
}

type UploadInfoRequest struct {
	File *excelize.File
}

type UploadInfoResponse struct {
	SuccessCount int
	FailCount    int
	Base         *util.Base
}

type AddInfoRequest struct {
	UserID string  `form:"user_id" binding:"required"`
	Name   string  `form:"name" binding:"required"`
	Class  *string `form:"class,omitempty"`
}

type AddInfoResponse struct {
	Base *util.Base
}

type UpdateInfoRequest struct {
	UserID string  `form:"user_id" binding:"required"`
	Name   string  `form:"name" binding:"required"`
	Class  *string `form:"class,omitempty"`
}

type UpdateInfoResponse struct {
	Base *util.Base
}

type DeleteInfoRequest struct {
	UserID string `form:"user_id" binding:"required"`
}

type DeleteInfoResponse struct {
	Base *util.Base
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
	UserID  string  `form:"user_id" binding:"required"`
	Subject string  `form:"subject" binding:"required"`
	Score   float64 `form:"score" binding:"required"`
}

type AddScoreResponse struct {
	Base *util.Base
}

type UpdateScoreRequest struct {
	UserID  string  `form:"user_id" binding:"required"`
	Subject string  `form:"subject" binding:"required"`
	Score   float64 `form:"score" binding:"required"`
}

type UpdateScoreResponse struct {
	Base *util.Base
}

type DeleteScoreRequest struct {
	UserID  string `form:"user_id" binding:"required"`
	Subject string `form:"subject" binding:"required"`
}

type DeleteScoreResponse struct {
	Base *util.Base
}
