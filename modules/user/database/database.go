package database

import(
	"app/modules/user/database/migrations"	
)

func Migrate() {
	user := migrations.User{}
	user.Create()
}