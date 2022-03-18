package teacher

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

func (s *ServiceImpl) AddInfo(_ *gin.Context, req *AddInfoRequest) *AddInfoResponse {
	info := &model.UserInfo{
		UserID: req.UserID,
		Name:   req.Name,
		Class:  req.Class,
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
		Class:  req.Class,
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
			err := info.Create()
			if err != nil {
				response.FailCount++
			} else {
				response.SuccessCount++
			}
		}
	}
	return
}
