package score

import (
	"github.com/gin-gonic/gin"
	"strconv"
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

func (s ServiceImpl) GetSubjects(_ *gin.Context) *GetSubjectResponse {
	var scores model.Scores
	response := &GetSubjectResponse{}

	if err := scores.GetSubjects(); err != nil {
		response.Base = util.NewBase(consts.ErrCodeFail, err)
		return response
	}

	response.Base = util.NewBase(consts.ErrCodeSuccess)
	for _, score := range scores {
		response.Subjects = append(response.Subjects, score.Subject)
	}
	return response
}

func (s *ServiceImpl) AddScore(_ *gin.Context, req *AddScoreRequest) *AddScoreResponse {
	score := &model.Score{
		Subject: req.Subject,
		UserID:  req.UserID,
		Name:    req.Name,
		Score:   req.Score,
	}
	if err := score.Create(); err != nil {
		return &AddScoreResponse{
			Base: util.NewBase(consts.ErrCodeFail, err),
		}
	}

	return &AddScoreResponse{
		Base: util.NewBase(consts.ErrCodeSuccess),
	}
}

func (s *ServiceImpl) UploadScore(_ *gin.Context, req *UploadScoreRequest) (response *UploadScoreResponse) {
	response = &UploadScoreResponse{Base: util.NewBase(consts.ErrCodeSuccess)}
	xlsx := req.File
	// 获取excel中具体的列的值
	rows := xlsx.GetRows("Sheet1")
	// 循环刚刚获取到的表中的值
	for key, row := range rows {
		// 去掉标题行
		if key > 0 {
			if len(row) != 4 {
				response.FailCount++
				continue
			}
			num, err := strconv.ParseFloat(row[3], 64)
			if err != nil {
				response.FailCount++
				continue
			}
			score := &model.Score{
				UserID:  row[0],
				Name:    row[1],
				Subject: row[2],
				Score:   num,
			}
			if score.Create() != nil {
				response.FailCount++
			} else {
				response.SuccessCount++
			}
		}
	}
	return
}

func (s *ServiceImpl) UpdateScore(_ *gin.Context, req *UpdateScoreRequest) *UpdateScoreResponse {
	score := &model.Score{
		UserID:  req.UserID,
		Subject: req.Subject,
		Score:   req.Score,
	}
	if req.Name != nil {
		score.Name = *req.Name
	}
	if err := score.Update(); err != nil {
		return &UpdateScoreResponse{
			Base: util.NewBase(consts.ErrCodeFail, err),
		}
	}

	return &UpdateScoreResponse{
		Base: util.NewBase(consts.ErrCodeSuccess),
	}
}

func (s *ServiceImpl) DeleteScore(_ *gin.Context, req *DeleteScoreRequest) *DeleteScoreResponse {
	score := &model.Score{
		UserID:  req.UserID,
		Subject: req.Subject,
	}
	if err := score.Delete(); err != nil {
		return &DeleteScoreResponse{
			Base: util.NewBase(consts.ErrCodeFail, err),
		}
	}

	return &DeleteScoreResponse{
		Base: util.NewBase(consts.ErrCodeSuccess),
	}
}

func (s ServiceImpl) GetScoresByID(_ *gin.Context, req *GetScoresByIDRequest) *GetScoresByIDResponse {
	var scores model.Scores
	var err error
	if req.Subject != nil {
		err = scores.Find([]string{req.ID}, *req.Subject)
	} else {
		err = scores.Find([]string{req.ID})
	}
	if err != nil {
		return &GetScoresByIDResponse{
			Scores: nil,
			Base:   util.NewBase(consts.ErrCodeFail, err),
		}
	}
	return &GetScoresByIDResponse{
		Scores: scores,
		Base:   util.NewBase(consts.ErrCodeSuccess),
	}
}

func (s ServiceImpl) GetScoresByClass(_ *gin.Context, req *GetScoresByClassRequest) *GetScoresByClassResponse {
	var infos model.UserInfos

	response := &GetScoresByClassResponse{
		Scores: nil,
		Base:   nil,
	}
	var scores model.Scores
	var err error
	if req.Class != nil {
		if err := infos.GetIDByClass(*req.Class); err != nil {
			response.Base = util.NewBase(consts.ErrCodeFail, err)
			return response
		}
		var ids []string
		for _, info := range infos {
			ids = append(ids, info.UserID)
		}

		if req.Subject != nil {
			err = scores.Find(ids, *req.Subject)
		} else {
			err = scores.Find(ids)
		}
	} else {
		err = scores.FindAll()
	}
	if err != nil {
		response.Base = util.NewBase(consts.ErrCodeFail, err)
		return response
	}
	return &GetScoresByClassResponse{
		Scores: scores,
		Base:   util.NewBase(consts.ErrCodeSuccess),
	}
}
