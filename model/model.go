package model

import "time"

type Area struct {
	Id         int    `gorm:"primark_key" json:"id"`
	Code       int    `gorm:"primary_key, not null;column:code" form:"code" json:"code"`
	Name       string `gorm:"primary_key"`
	ParentCode int    `gorm:"column:parent_code;primary_key" form:"parent_code" json:"parent_code"`
	Level      string
	XAxis      float64
	YAxis      float64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (b *Area) TableName() string {
	return "cities"
}
