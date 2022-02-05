package model

import (
	"crypto/md5"
	"fmt"
	"github.com/google/uuid"
	"studentScoreManagement/consts"
	"studentScoreManagement/db"
)

type User struct {
	ID       string `json:"id,omitempty" gorm:"primary_key"`
	Name     string `json:"name,omitempty" gorm:"unique"`
	PassWord string `json:"pass_word,omitempty"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) isValid() bool {
	return u != nil && u.Name != "" && u.PassWord != ""
}

func (u *User) Create() error {
	if u.isValid() == false {
		return consts.GetError(consts.ErrCodeParameter)
	}

	//不明文存储密码
	u.PassWord = fmt.Sprintf("%x", md5.Sum([]byte(u.PassWord)))
	u.ID = uuid.New().String()

	return db.GetDatabase().Create(u).Error
}

func (u *User) Find() error {
	if u.isValid() == false {
		return consts.GetError(consts.ErrCodeParameter)
	}

	u.PassWord = fmt.Sprintf("%x", md5.Sum([]byte(u.PassWord)))

	return db.GetDatabase().First(u).
		Where("name = ? and pass_word = ?", u.Name, u.PassWord).Error
}

func (u *User) UpdatePassword(newPassword string) error {
	if u.isValid() == false {
		return consts.GetError(consts.ErrCodeParameter)
	}

	newPassword = fmt.Sprintf("%x", md5.Sum([]byte(newPassword)))

	return db.GetDatabase().First(u).Where("pass_word = ?", u.PassWord).
		Update("pass_word", newPassword).Error
}
