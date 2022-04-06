package role

import (
	"github.com/gin-gonic/gin"
	"studentScoreManagement/consts"
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

func (s *ServiceImpl) GetRoleID(ctx *gin.Context) *GetRoleIDResponse {
	userID, err := util.GetUserID(ctx)
	if err != nil {
		return &GetRoleIDResponse{
			Base: util.NewBase(consts.ErrCodeFail, err),
		}
	}
	role := &model.UserID2RoleID{
		UserID: userID,
	}
	if err := role.Find(); err != nil {
		return &GetRoleIDResponse{
			Base: util.NewBase(consts.ErrCodeFail, err),
		}
	}

	return &GetRoleIDResponse{
		RoleID: role.RoleID,
		Base:   util.NewBase(consts.ErrCodeSuccess),
	}
}

func (s *ServiceImpl) GetValidateCode(_ *gin.Context, req *GetValidateCodeRequest) *GetValidateCodeResponse {
	role := &model.Role{
		RoleID: req.RoleID,
	}
	if err := role.Find(); err != nil {
		return &GetValidateCodeResponse{
			Base: util.NewBase(consts.ErrCodeFail, err),
		}
	}

	return &GetValidateCodeResponse{
		ValidateCode: role.RoleCode,
		Base:         util.NewBase(consts.ErrCodeSuccess),
	}
}

func (s *ServiceImpl) ChangeValidateCode(_ *gin.Context, req *ChangeValidateCodeRequest) *ChangeValidateCodeResponse {
	role := &model.Role{
		RoleID:   req.RoleID,
		RoleCode: req.ValidateCode,
	}
	if err := role.UpdateRoleCode(); err != nil {
		return &ChangeValidateCodeResponse{
			Base: util.NewBase(consts.ErrCodeFail, err),
		}
	}

	return &ChangeValidateCodeResponse{
		Base: util.NewBase(consts.ErrCodeSuccess),
	}
}
