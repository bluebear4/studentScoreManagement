package util

import (
	"github.com/gin-gonic/gin"
	"studentScoreManagement/consts"
)

type Base struct {
	Code    int
	Message string
}

func (b *Base) ChangeToGinJson(extras ...map[string]interface{}) (httpCode int, H gin.H) {
	extra := map[string]interface{}{
		"Code":    b.Code,
		"Message": b.Message,
	}
	if len(extras) > 0 {
		extra["Data"] = extras[0]
	}
	return consts.GetHttpCode(b.Code), extra

}

func NewBase(code int, errors ...error) *Base {
	base := &Base{
		Code: code,
	}
	if len(errors) > 0 {
		base.Message = errors[0].Error()
	} else {
		base.Message = consts.GetErrorMsg(code)
	}
	return base
}
