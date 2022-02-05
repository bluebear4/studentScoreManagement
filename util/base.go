package util

import (
	"github.com/gin-gonic/gin"
	"studentScoreManagement/consts"
)

type Base struct {
	Code    int
	Message string
}

func (b *Base) ChangeToGinJson() (httpCode int, H gin.H) {
	return consts.GetHttpCode(b.Code), gin.H{
		"Code":    b.Code,
		"Message": b.Message,
	}
}

func NewBase(code int, errors ...error) *Base {
	base := &Base{
		Code: code,
	}
	if len(errors) > 0 {
		base.Message = errors[0].Error()
	} else {
		base.Message = consts.GetError(code).Error()
	}
	return base
}
