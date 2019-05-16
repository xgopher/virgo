package pagination

import (
	"fmt"
	"testing"

	"app/database"
	"app/modules/user/models"
)

func TestPagging(t *testing.T) {

	// 初始化DB链接池
	database.InitDb()
	// 服务停止时清理数据库链接
	defer database.DB.Close()

	db := database.DB
	param := Param {
        DB:      db,
        Page:    1,
        PerPage: 10,
        OrderBy: []string{"id desc"},
	}
	var users []models.User

	paginator, err := Pagging(&param, &users)

	fmt.Printf("%v %v\n", paginator, err)
}
