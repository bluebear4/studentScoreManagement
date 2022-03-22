package scoreView

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

func (s ServiceImpl) GetClass(_ *gin.Context) *GetClassResponse {

	var infos model.UserInfos
	response := &GetClassResponse{}

	if err := infos.GetClass(); err != nil {
		response.Base = util.NewBase(consts.ErrCodeFail, err)
		return response
	}

	response.Base = util.NewBase(consts.ErrCodeSuccess)
	for _, info := range infos {
		response.Classes = append(response.Classes, info.Class)
	}
	return response
}

func (s ServiceImpl) GetScoreByID(_ *gin.Context, req *GetScoreByIDRequest) *GetScoreByIDResponse {
	var scores model.Scores
	var err error
	if req.Subject != nil {
		err = scores.Find([]string{req.ID}, *req.Subject)
	} else {
		err = scores.Find([]string{req.ID})
	}
	if err != nil {
		return &GetScoreByIDResponse{
			Scores: nil,
			Base:   util.NewBase(consts.ErrCodeFail, err),
		}
	}
	return &GetScoreByIDResponse{
		Scores: scores,
		Base:   util.NewBase(consts.ErrCodeSuccess),
	}
}

func (s ServiceImpl) GetScoreByClass(_ *gin.Context, req *GetScoreByClassRequest) *GetScoreByClassResponse {
	var infos model.UserInfos

	response := &GetScoreByClassResponse{
		Scores: nil,
		Base:   nil,
	}

	if err := infos.GetIDByClass(req.Class); err != nil {
		response.Base = util.NewBase(consts.ErrCodeFail, err)
		return response
	}
	var ids []string
	for _, info := range infos {
		ids = append(ids, info.UserID)
	}

	var scores model.Scores
	var err error
	if req.Subject != nil {
		err = scores.Find(ids, *req.Subject)
	} else {
		err = scores.Find(ids)
	}

	if err != nil {
		response.Base = util.NewBase(consts.ErrCodeFail, err)
		return response
	}

	return &GetScoreByClassResponse{
		Scores: scores,
		Base:   util.NewBase(consts.ErrCodeSuccess),
	}
}
