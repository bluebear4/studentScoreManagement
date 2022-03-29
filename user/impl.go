package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strings"
	"studentScoreManagement/consts"
	"studentScoreManagement/db"
	"studentScoreManagement/model"
	"studentScoreManagement/util"
)

var (
	server Service
)

func init() {
	server = &ServiceImpl{}
}

type ServiceImpl struct {
}

func (s *ServiceImpl) ChangeValidateCode(_ *gin.Context, req *ChangeValidateCodeRequest) *ChangeValidateCodeResponse {
	role := &model.Role{
		RoleID:   req.RoleID,
		RoleCode: req.ValidateCode,
	}
	if err := role.UpdateRoleCode(); err != nil {
		return &ChangeValidateCodeResponse{
			Base: util.NewBase(consts.ErrCodeErrorUserOrPassword, err),
		}
	}

	return &ChangeValidateCodeResponse{
		Base: util.NewBase(consts.ErrCodeSuccess),
	}
}

func (s *ServiceImpl) CreateUser(_ *gin.Context, req *CreateUserRequest) *CreateUserResponse {

	role := model.Role{RoleID: req.RoleID}

	if req.RoleID == consts.RoleIDAdmin || role.Find() != nil {
		//admin不可注册
		//不存在角色不可注册
		return &CreateUserResponse{Base: util.NewBase(consts.ErrCodeParameter)}
	}

	if role.RoleCode != nil && *role.RoleCode != req.VerifyCode {
		//验证码不对不可注册
		//由管理员控制修改
		return &CreateUserResponse{Base: util.NewBase(consts.ErrCodeValidate)}
	}

	//事务
	err := db.GetDatabase().Transaction(func(tx *gorm.DB) error {
		//用户注册
		user := &model.User{
			ID:       req.ID,
			PassWord: req.PassWord,
		}
		if err := user.Create(); err != nil {
			if strings.Contains(err.Error(), "Duplicate") {
				return consts.GetError(consts.ErrCodeUserDuplicate)
			}
			return err
		}
		//角色绑定
		uid2role := &model.UserID2RoleID{
			UserID: user.ID,
			RoleID: req.RoleID,
		}
		return uid2role.Create()
	})
	if err != nil {
		return &CreateUserResponse{
			Base: util.NewBase(consts.ErrCodeFail, err),
		}
	}

	return &CreateUserResponse{
		UserID: req.ID,
		Base:   util.NewBase(consts.ErrCodeSuccess),
	}
}

func (s *ServiceImpl) ValidatePassword(_ *gin.Context, req *ValidatePasswordRequest) *ValidatePasswordResponse {
	user := &model.User{
		ID:       req.ID,
		PassWord: req.PassWord,
	}
	if err := user.Find(); err != nil {
		return &ValidatePasswordResponse{
			Base: util.NewBase(consts.ErrCodeErrorUserOrPassword),
		}
	}

	return &ValidatePasswordResponse{
		Base: util.NewBase(consts.ErrCodeSuccess),
	}
}

func (s *ServiceImpl) ChangePassword(_ *gin.Context, req *ChangePasswordRequest) *ChangePasswordResponse {
	user := &model.User{
		ID:       req.ID,
		PassWord: req.OldPassWord,
	}
	if err := user.UpdatePassword(req.NewPassWord); err != nil {
		return &ChangePasswordResponse{
			Base: util.NewBase(consts.ErrCodeErrorUserOrPassword),
		}
	}

	return &ChangePasswordResponse{
		Base: util.NewBase(consts.ErrCodeSuccess),
	}
}
