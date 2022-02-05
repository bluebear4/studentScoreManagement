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

//ChangePassword(ctx context.Context, req *ChangePasswordRequest) *ChangePasswordResponse

type CreateUserRequest struct {
	UserName   string `form:"user_name"`
	PassWord   string `form:"pass_word"`
	RoleID     int    `form:"role_id"`
	VerifyCode string `form:"verify_code"`
}

type CreateUserResponse struct {
	UserID string
	Base   *util.Base
}

type ValidatePasswordRequest struct {
	UserName string `form:"user_name"`
	PassWord string `form:"pass_word"`
}

type ValidatePasswordResponse struct {
	Base *util.Base
}

type ChangePasswordRequest struct {
	UserName    string `form:"user_name"`
	OldPassWord string `form:"old_pass_word"`
	NewPassWord string `form:"new_pass_word"`
}

type ChangePasswordResponse struct {
	Base *util.Base
}
