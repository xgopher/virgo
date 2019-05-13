package database

import (
	"app/modules/oauth/database/migrations"
)

func Migrate() {
	oauth_wechat := migrations.OauthWechat{}
	oauth_wechat.Create()
}
