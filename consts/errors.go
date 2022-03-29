package consts

import (
	"errors"
	"net/http"
)

const (
	//StatusOK
	ErrCodeSuccess = iota

	// StatusBadRequest
	ErrCodeFail
	ErrCodeParameter = 10001 + iota
	ErrCodeNotLogin
	ErrCodeValidate
	ErrCodeUserDuplicate
	ErrCodeErrorUserOrPassword
	ErrCodeErrorAuthFail

	//StatusInternalServerError
	ErrCodeRetry = 20001 + iota
)

func GetError(code int) error {
	if code == 0 {
		return nil
	}
	return errors.New(GetErrorMsg(code))
}

func GetErrorMsg(code int) string {
	switch code {
	case ErrCodeSuccess:
		return "操作成功"
	case ErrCodeFail:
		return "操作失败"
	case ErrCodeParameter:
		return "参数错误"
	case ErrCodeNotLogin:
		return "未登录或登录已超时"
	case ErrCodeValidate:
		return "验证码错误"
	case ErrCodeUserDuplicate:
		return "用户已存在"
	case ErrCodeRetry:
		return "请稍后重试"
	case ErrCodeErrorUserOrPassword:
		return "用户名或密码错误"
	case ErrCodeErrorAuthFail:
		return "无此操作权限"
	default:
		return "未知错误"
	}
}

func GetHttpCode(code int) int {
	if code == 0 {
		return http.StatusOK
	} else if code < 20001 {
		return http.StatusOK
		//return http.StatusBadRequest
	} else {
		return http.StatusInternalServerError
	}
}
