package model

import (
	"crypto/md5"
	"fmt"
	"studentScoreManagement/consts"
	"studentScoreManagement/db"
)

type User struct {
	//学号 工号
	ID       string `json:"id,omitempty" gorm:"primary_key"`
	PassWord string `json:"pass_word,omitempty" gorm:"not null"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) isValid() bool {
	return u != nil && u.ID != "" && u.PassWord != ""
}

func (u User) Create() error {
	if u.isValid() == false {
		return consts.GetError(consts.ErrCodeParameter)
	}

	//不明文存储密码
	u.PassWord = fmt.Sprintf("%x", md5.Sum([]byte(u.PassWord)))

	return db.GetDatabase().Create(u).Error
}

func (u *User) Find() error {
	if u.isValid() == false {
		return consts.GetError(consts.ErrCodeParameter)
	}

	passWord := fmt.Sprintf("%x", md5.Sum([]byte(u.PassWord)))

	return db.GetDatabase().First(u, "id = ? and pass_word = ?", u.ID, passWord).Error
}

func (u *User) UpdatePassword(newPassword string) error {
	if u.isValid() == false {
		return consts.GetError(consts.ErrCodeParameter)
	}

	newPassword = fmt.Sprintf("%x", md5.Sum([]byte(newPassword)))

	return db.GetDatabase().First(u).Update("pass_word", newPassword).Error
}
