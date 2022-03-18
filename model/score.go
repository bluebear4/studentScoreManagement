package model

import (
	"studentScoreManagement/consts"
	"studentScoreManagement/db"
)

type Score struct {
	//学号 工号
	ID      string  `json:"id,omitempty" gorm:"primary_key"`
	Subject string  `json:"subject"`
	Score   float64 `json:"score"`
}

func (Score) TableName() string {
	return "score"
}

func (s *Score) isValid() bool {
	return s != nil && s.ID != ""
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
	return db.GetDatabase().Save(s).Error
}
