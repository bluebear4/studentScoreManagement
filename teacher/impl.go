package teacher

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

func (s *ServiceImpl) UploadScore(_ *gin.Context, req *UploadScoreRequest) (response *UploadScoreResponse) {
	response = &UploadScoreResponse{Base: util.NewBase(consts.ErrCodeSuccess)}
	xlsx := req.File
	// 获取excel中具体的列的值
	rows := xlsx.GetRows("Sheet1")
	// 循环刚刚获取到的表中的值
	for key, row := range rows {
		// 去掉标题行
		if key > 0 {
			if len(row) != 3 {
				response.FailCount++
				continue
			}
			num, err := strconv.ParseFloat(row[2], 64)
			if err != nil {
				response.FailCount++
				continue
			}
			score := &model.Score{
				ID:      row[0],
				Subject: row[1],
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

func (s *ServiceImpl) AddScore(_ *gin.Context, req *AddScoreRequest) *AddScoreResponse {
	score := &model.Score{
		ID:      req.UserID,
		Subject: req.Subject,
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

func (s *ServiceImpl) UpdateScore(_ *gin.Context, req *UpdateScoreRequest) *UpdateScoreResponse {
	score := &model.Score{
		ID:      req.UserID,
		Subject: req.Subject,
		Score:   req.Score,
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
		ID:      req.UserID,
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

func (s *ServiceImpl) DeleteInfo(_ *gin.Context, req *DeleteInfoRequest) *DeleteInfoResponse {
	info := &model.UserInfo{
		UserID: req.UserID,
	}
	if err := info.Delete(); err != nil {
		return &DeleteInfoResponse{
			Base: util.NewBase(consts.ErrCodeFail, err),
		}
	}

	return &DeleteInfoResponse{
		Base: util.NewBase(consts.ErrCodeSuccess),
	}
}

func (s *ServiceImpl) AddInfo(_ *gin.Context, req *AddInfoRequest) *AddInfoResponse {
	info := &model.UserInfo{
		UserID: req.UserID,
		Name:   req.Name,
	}
	if req.Class != nil {
		info.Class = *req.Class
	}
	if err := info.Create(); err != nil {
		return &AddInfoResponse{
			Base: util.NewBase(consts.ErrCodeFail, err),
		}
	}

	return &AddInfoResponse{
		Base: util.NewBase(consts.ErrCodeSuccess),
	}
}

func (s *ServiceImpl) UpdateInfo(_ *gin.Context, req *UpdateInfoRequest) *UpdateInfoResponse {
	info := &model.UserInfo{
		UserID: req.UserID,
		Name:   req.Name,
	}
	if req.Class != nil {
		info.Class = *req.Class
	}
	if err := info.Update(); err != nil {
		return &UpdateInfoResponse{
			Base: util.NewBase(consts.ErrCodeFail, err),
		}
	}

	return &UpdateInfoResponse{
		Base: util.NewBase(consts.ErrCodeSuccess),
	}
}

func (s *ServiceImpl) UploadInfo(_ *gin.Context, req *UploadInfoRequest) (response *UploadInfoResponse) {
	response = &UploadInfoResponse{Base: util.NewBase(consts.ErrCodeSuccess)}
	xlsx := req.File
	// 获取excel中具体的列的值
	rows := xlsx.GetRows("Sheet1")
	// 循环刚刚获取到的表中的值
	for key, row := range rows {
		// 去掉标题行
		if key > 0 {
			if len(row) != 3 && len(row) != 2 {
				response.FailCount++
				continue
			}
			info := &model.UserInfo{
				UserID: row[0],
				Name:   row[1],
			}
			if len(row) == 3 {
				info.Class = row[2]
			}
			if info.Create() != nil {
				response.FailCount++
			} else {
				response.SuccessCount++
			}
		}
	}
	return
}
