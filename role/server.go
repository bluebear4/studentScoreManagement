package role

import (
	"github.com/gin-gonic/gin"
	"studentScoreManagement/util"
)

type Service interface {
	ChangeValidateCode(ctx *gin.Context, req *ChangeValidateCodeRequest) *ChangeValidateCodeResponse
	GetValidateCode(ctx *gin.Context, req *GetValidateCodeRequest) *GetValidateCodeResponse
	GetRoleID(ctx *gin.Context) *GetRoleIDResponse
}

type ChangeValidateCodeRequest struct {
	RoleID       int    `form:"role_id"  binding:"required" json:"role_id,omitempty"`
	ValidateCode string `form:"validate_code" json:"validate_code,omitempty"`
}

type ChangeValidateCodeResponse struct {
	Base *util.Base
}

type GetValidateCodeRequest struct {
	RoleID int `form:"role_id"  binding:"required" json:"role_id,omitempty"`
}

type GetValidateCodeResponse struct {
	ValidateCode string
	Base         *util.Base
}

type GetRoleIDResponse struct {
	RoleID int
	Base   *util.Base
}
