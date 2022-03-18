package model

import (
	"studentScoreManagement/consts"
	"studentScoreManagement/db"
)

type UserID2RoleID struct {
	UserID string `json:"user_id,omitempty" gorm:"primary_key"`
	RoleID int    `json:"role_id,omitempty"`
}

func (UserID2RoleID) TableName() string {
	return "uid_to_rid"
}

func (r *UserID2RoleID) isValid() bool {
	return r != nil && r.UserID != "" && r.RoleID >= 0
}

func (r *UserID2RoleID) Create() error {
	if r.isValid() == false {
		return consts.GetError(consts.ErrCodeParameter)
	}
	return db.GetDatabase().Create(r).Error
}

func (r *UserID2RoleID) Find() (err error) {
	if r.isValid() == false {
		return consts.GetError(consts.ErrCodeParameter)
	}
	return db.GetDatabase().First(r).Error
}
