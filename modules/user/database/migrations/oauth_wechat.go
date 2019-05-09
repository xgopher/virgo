package migrations

import(
	"app/database"
	"app/modules/user/models"
	"app/modules/user/database/seeds"
)

// OauthWechat migration
type OauthWechat struct {

}

// Create OauthWechat Table and seed this table
func (u *OauthWechat) Create(){
	if !database.DB.HasTable(&models.OauthWechat{}) {
		database.DB.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&models.OauthWechat{})
		oauth_wechat := seeds.OauthWechat{}
		oauth_wechat.Seed()
	}
}