package models

import "time"

// User ...
type User struct {
	ID        int    `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Firstname string `gorm:"not null" form:"firstname" json:"firstname"`
	Lastname  string `gorm:"not null" form:"lastname" json:"lastname"`
	Username  string `gorm:"size:32;not null;unique_index" from:"username" json:"username"`
	Password  string `gorm:"not null" from:"password" json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// TableName 设置表名
func (model *User) TableName() string {
	return "users"
}
