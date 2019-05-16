package models

import "app/models"

// UserProfile ...
type UserProfile struct {
	ID        int    `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	UserID    int    `form:"user_id" json:"user_id"`
	Avatar    string `gorm:"not null" form:"avatar" json:"avatar"`

	models.Date
}

// TableName 设置表名
func (model *UserProfile) TableName() string {
	return "user_profiles"
}
