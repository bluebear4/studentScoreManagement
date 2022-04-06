package model

import (
	"studentScoreManagement/consts"
	"studentScoreManagement/db"
)

type Scores []Score

type Score struct {
	//学号 工号
	Subject string  `json:"subject" gorm:"primary_key"`
	UserID  string  `json:"user_id,omitempty" gorm:"primary_key"`
	Name    string  `json:"name" gorm:"not null"`
	Score   float64 `json:"score" gorm:"not null"`
}

func (Score) TableName() string {
	return "score"
}

func (s *Score) isValid() bool {
	return s != nil && s.UserID != ""
}

func (s *Score) Create() error {
	if s.isValid() == false {
		return consts.GetError(consts.ErrCodeParameter)
	}

	return db.GetDatabase().Create(s).Error
}

func (s *Score) Find() error {
	if s.isValid() == false {
		return consts.GetError(consts.ErrCodeParameter)
	}
	return db.GetDatabase().First(s).Error
}

func (s *Score) Update() error {
	if s.isValid() == false {
		return consts.GetError(consts.ErrCodeParameter)
	}
	return db.GetDatabase().Model(s).Updates(s).Error
}

func (s *Score) Delete() error {
	if s.isValid() == false {
		return consts.GetError(consts.ErrCodeParameter)
	}
	return db.GetDatabase().Delete(s).Error
}

func (s *Scores) Find(IDs []string, Subject ...string) error {
	if len(Subject) > 0 {
		return db.GetDatabase().Where("subject = ? AND user_id IN ?", Subject[0], IDs).Find(s).Error
	}
	return db.GetDatabase().Where("user_id IN ?", IDs).Find(s).Error
}

func (s *Scores) GetSubjects() error {
	return db.GetDatabase().Distinct("subject").Find(s).Error
}

func (s *Scores) FindAll() error {
	return db.GetDatabase().Find(s).Error
}
