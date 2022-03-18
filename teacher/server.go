package teacher

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	"studentScoreManagement/util"
)

type Service interface {
	UploadInfo(ctx *gin.Context, req *UploadInfoRequest) *UploadInfoResponse
	UpdateInfo(ctx *gin.Context, req *UpdateInfoRequest) *UpdateInfoResponse
	AddInfo(ctx *gin.Context, req *AddInfoRequest) *AddInfoResponse
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
	UserID string `json:"user_id,omitempty"`
	Name   string `json:"name"`
	Class  string `json:"class,omitempty"`
}

type AddInfoResponse struct {
	Base *util.Base
}

type UpdateInfoRequest struct {
	UserID string `json:"user_id,omitempty"`
	Name   string `json:"name"`
	Class  string `json:"class,omitempty"`
}

type UpdateInfoResponse struct {
	Base *util.Base
}
