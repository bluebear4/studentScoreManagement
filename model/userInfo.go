package model

import (
	"studentScoreManagement/consts"
	"studentScoreManagement/db"
)

type UserInfo struct {
	//学号 工号
	UserID string `json:"user_id,omitempty" gorm:"primary_key"`
	Name   string `json:"name,omitempty"`
	Class  string `json:"class" gorm:"index"`
}

func (UserInfo) TableName() string {
	return "user_info"
}

func (s *UserInfo) isValid() bool {
	return s != nil && s.UserID != ""
}

func (s *UserInfo) Create() error {
	if s.isValid() == false {
		return consts.GetError(consts.ErrCodeParameter)
	}
	return db.GetDatabase().Create(s).Error
}

func (s *UserInfo) Find() error {
	if s.isValid() == false {
		return consts.GetError(consts.ErrCodeParameter)
	}
	return db.GetDatabase().First(s).Error
}

func (s *UserInfo) Update() error {
	if s.isValid() == false {
		return consts.GetError(consts.ErrCodeParameter)
	}
	return db.GetDatabase().Save(s).Error
}
