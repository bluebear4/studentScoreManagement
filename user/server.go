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
	ID         string `form:"id" binding:"required"`
	PassWord   string `form:"pass_word" binding:"required"`
	RoleID     int    `form:"role_id" binding:"required"`
	VerifyCode string `form:"verify_code"`
}

type CreateUserResponse struct {
	UserID string
	Base   *util.Base
}

type ValidatePasswordRequest struct {
	ID       string `form:"id" binding:"required"`
	PassWord string `form:"pass_word" binding:"required"`
}

type ValidatePasswordResponse struct {
	Base *util.Base
}

type ChangePasswordRequest struct {
	ID          string `form:"id" binding:"required"`
	OldPassWord string `form:"old_pass_word" binding:"required"`
	NewPassWord string `form:"new_pass_word" binding:"required"`
}

type ChangePasswordResponse struct {
	Base *util.Base
}
