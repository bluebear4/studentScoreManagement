package info

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	"studentScoreManagement/model"
	"studentScoreManagement/util"
)

type Service interface {
	UploadInfo(ctx *gin.Context, req *UploadInfoRequest) *UploadInfoResponse
	AddInfo(ctx *gin.Context, req *AddInfoRequest) *AddInfoResponse
	UpdateInfo(ctx *gin.Context, req *UpdateInfoRequest) *UpdateInfoResponse
	DeleteInfo(ctx *gin.Context, req *DeleteInfoRequest) *DeleteInfoResponse
	GetInfos(ctx *gin.Context, req *GetInfosRequest) *GetInfosResponse

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
	UserID string  `form:"user_id"  binding:"required" json:"user_id,omitempty"`
	Name   string  `form:"name" binding:"required" json:"name,omitempty"`
	Class  *string `form:"class,omitempty" json:"class,omitempty"`
}

type AddInfoResponse struct {
	Base *util.Base
}

type UpdateInfoRequest struct {
	UserID string  `form:"user_id"  binding:"required" json:"user_id,omitempty"`
	Name   string  `form:"name" binding:"required" json:"name,omitempty"`
	Class  *string `form:"class,omitempty" json:"class,omitempty"`
}

type UpdateInfoResponse struct {
	Base *util.Base
}

type DeleteInfoRequest struct {
	UserID string `form:"user_id"  binding:"required" json:"user_id,omitempty"`
}

type DeleteInfoResponse struct {
	Base *util.Base
}

type GetClassResponse struct {
	Classes []string
	Base    *util.Base
}

type GetInfosRequest struct {
}

type GetInfosResponse struct {
	Infos model.UserInfos
	Base  *util.Base
}
