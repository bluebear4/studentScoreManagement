package model

import (
	"studentScoreManagement/consts"
	"studentScoreManagement/db"
)

type Role struct {
	RoleID   int    `json:"role_id,omitempty" gorm:"primary_key;comment:'0:管理员 1:教师 2:学生'"`
	RoleName string `json:"role_name,omitempty"`
	RoleCode string `json:"role_code,omitempty" gorm:"comment:'注册教师需要管理员的验证码'"`
}

func (Role) TableName() string {
	return "roles"
}

func (r *Role) isValid() bool {
	return r != nil && r.RoleID >= 0
}

func (r *Role) Create() error {
	if r.isValid() == false {
		return consts.GetError(consts.ErrCodeParameter)
	}
	return db.GetDatabase().Create(r).Error
}

func (r *Role) UpdateRoleCode() error {
	if r.isValid() == false {
		return consts.GetError(consts.ErrCodeParameter)
	}
	return db.GetDatabase().First(r).Where("role_id = ? and role_name = ?", r.RoleID, r.RoleName).
		Update("role_code", r.RoleCode).Error
}

func (r *Role) Find() (err error) {
	if r.isValid() == false {
		return consts.GetError(consts.ErrCodeParameter)
	}
	return db.GetDatabase().First(r, r.RoleID).Error
}
