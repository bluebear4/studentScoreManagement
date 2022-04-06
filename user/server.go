package user

import (
	"github.com/gin-gonic/gin"
	"studentScoreManagement/util"
)

type Service interface {
	CreateUser(ctx *gin.Context, req *CreateUserRequest) *CreateUserResponse
	ValidatePassword(ctx *gin.Context, req *ValidatePasswordRequest) *ValidatePasswordResponse
	ChangePassword(ctx *gin.Context, req *ChangePasswordRequest) *ChangePasswordResponse
}

type CreateUserRequest struct {
	ID         string `form:"id"  binding:"required" json:"id,omitempty"`
	PassWord   string `form:"pass_word" binding:"required" json:"pass_word,omitempty"`
	RoleID     int    `form:"role_id" binding:"required" json:"role_id,omitempty"`
	VerifyCode string `form:"verify_code" json:"verify_code,omitempty"`
}

type CreateUserResponse struct {
	UserID string
	Base   *util.Base
}

type ValidatePasswordRequest struct {
	ID       string `form:"id"  binding:"required" json:"id,omitempty"`
	PassWord string `form:"pass_word" binding:"required" json:"pass_word,omitempty"`
}

type ValidatePasswordResponse struct {
	Base *util.Base
}

type ChangePasswordRequest struct {
	ID          *string `form:"id"  json:"id,omitempty"`
	OldPassWord string  `form:"old_pass_word" binding:"required" json:"old_pass_word,omitempty"`
	NewPassWord string  `form:"new_pass_word" binding:"required" json:"new_pass_word,omitempty"`
}

type ChangePasswordResponse struct {
	Base *util.Base
}
