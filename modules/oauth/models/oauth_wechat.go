package models

import "time"

// OauthWechat 微信用户
type OauthWechat struct {
	ID       uint   `gorm:"AUTO_INCREMENT" form:"id" json:"id"`                           // 自增 id
	UID      uint64 `gorm:"index;column:uid;not null;default:0" form:"uid" json:"uid"`    // 分布式user_id, 用雪花算法生成
	OpenID   string `gorm:"index;column:openid;not null" form:"openid" json:"openid"`     // 用户的唯一标识
	UnionID  string `gorm:"column:unionid;unique;not null" form:"unionid" json:"unionid"` // 只有在用户将公众号绑定到微信开放平台帐号后，才会出现该字段。
	Nickname string `gorm:"not null" form:"nickname" json:"nickname"`                     // 昵称
	// 用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），
	// 用户没有头像时该项为空。若用户更换头像，原有头像URL将失效。
	HeadImgURL string `gorm:"column:headimgurl;not null" form:"headimgurl" json:"headimgurl"`
	Sex        int    `gorm:"type:tinyint(3);" form:"sex" json:"sex"`                               // 性别, 值为1时是男性, 值为2时是女性, 值为0时是未知
	Province   string `gorm:"type:varchar(32);not null;default:''" form:"province" json:"province"` // 省
	City       string `gorm:"type:varchar(32);not null;default:''" form:"city" json:"city"`         // 市
	Country    string `gorm:"type:varchar(32);not null;default:''" form:"country" json:"country"`   // 国家, 如中国为CN
	Privilege  string `gorm:"not null;default:''" form:"privilege" json:"privilege"`                // 用户特权信息，json 数组，如微信沃卡用户为（chinaunicom）
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// TableName 设置表名
func (model *OauthWechat) TableName() string {
	return "oauth_wechat"
}
