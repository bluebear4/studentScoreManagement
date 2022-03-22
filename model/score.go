package model

import (
	"studentScoreManagement/consts"
	"studentScoreManagement/db"
)

type Scores []Score

type Score struct {
	//学号 工号
	ID      string  `json:"id,omitempty" gorm:"primary_key;index:id_subject"`
	Subject string  `json:"subject" gorm:"index:id_subject"`
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

func (s *Score) Delete() error {
	if s.isValid() == false {
		return consts.GetError(consts.ErrCodeParameter)
	}
	return db.GetDatabase().Delete(s).Error
}

func (s *Scores) Find(IDs []string, Subject ...string) error {
	if len(Subject) > 0 {
		return db.GetDatabase().Find(s).Where("id IN ? AND subject >= ?", IDs, Subject[0]).Error
	}
	return db.GetDatabase().Find(s).Where("id IN ?", IDs).Error
}
