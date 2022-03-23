package info

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

	GetClasses(ctx *gin.Context) *GetClassResponse
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

type GetClassResponse struct {
	Classes []string
	Base    *util.Base
}

type GetInfoByIDRequest struct {
	UserID string `form:"user_id" binding:"required"`
}

type GetInfoByIDResponse struct {
	Base *util.Base
}