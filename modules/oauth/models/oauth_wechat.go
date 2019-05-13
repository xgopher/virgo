package models

import "time"

// OauthWechat 微信用户
type OauthWechat struct {
	ID        uint   `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Openid    string `gorm:"not null" form:"openid" json:"openid"`
	Unionid   string `gorm:"not null" form:"unionid" json:"unionid"`
	Nickname  string `gorm:"not null" form:"nickname" json:"nickname"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// TableName 设置表名
func (model *OauthWechat) TableName() string {
	return "oauth_wechat"
}
