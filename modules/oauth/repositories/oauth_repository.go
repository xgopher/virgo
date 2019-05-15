package repositories

import "app/database"
import "app/modules/oauth/models"

type OauthRepository struct {
}

// FindByUnionID ...
func (i *OauthRepository) FindByUnionID(unionID string) {
	var count int

	db := database.DB
	db.Model(&models.OauthWechat{}).Where("unionid = ?", unionID).Count(&count)

	if count > 0 {
		// ....
	} else {
		// ....
	}

	return
}

// FindByOpenID ...
func (i *OauthRepository) FindByOpenID(openID string) {

	return
}

// FindByUnionIDOrCreate
func (i *OauthRepository) FindByUnionIDOrCreate(unionID string) {
	return
}

// FindByOpenIDOrCreate ...
func (i *OauthRepository) FindByOpenIDOrCreate(openID string) {
	return
}