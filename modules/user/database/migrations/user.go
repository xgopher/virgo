package migrations

import(
	"app/database"
	"app/modules/user/models"
	"app/modules/user/database/seeds"
)

// User migration
type User struct {

}

// Create User Table and seed this table
func (u *User) Create(){
	if !database.DB.HasTable(&models.User{}) {
		database.DB.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&models.User{})
		user := seeds.User{}
		user.Seed()
	}
}