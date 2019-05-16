package models

import "app/models"

// User ...
type User struct {
	ID        int    `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Firstname string `gorm:"not null" form:"firstname" json:"firstname"`
	Lastname  string `gorm:"not null" form:"lastname" json:"lastname"`
	Username  string `gorm:"size:32;not null;unique_index" from:"username" json:"username"`
	Password  string `gorm:"not null" from:"password" json:"-"`

	models.Date

	Profile   UserProfile `gorm:"ForeignKey:ID;AssociationForeignKey:UserID"`
}

// TableName 设置表名
func (model *User) TableName() string {
	return "users"
}
