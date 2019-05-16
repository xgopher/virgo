package controllers

import (
	"strconv"

	"app/database"
	"app/services/pagination"
	"app/modules/user/services"
	"app/modules/user/models"

	"github.com/gin-gonic/gin"
)

// UserController ...
type UserController struct {
}

type UserInfo struct {
	models.User
}

// LoginResult 登录结果结构
type loginResult struct {
    Token string `json:"token"`
    UserInfo `json:"user"`
}

func NewUserController() *UserController {
	return &UserController{}
}

// login
func(i *UserController) Login(c *gin.Context) {
	db := database.DB

	username := c.Params.ByName("username")
	password := c.Params.ByName("password")

	var user models.User
	// SELECT * FROM users WHERE username='' and password='';
	if db.Where(&models.User{Username: username, Password: password}).First(&user).Error != nil{
		c.JSON(401, gin.H{"error": "username or password error"})
	}
	
	token, err := services.GetTokenFromUser(user)
	if err != nil {
		c.JSON(401, gin.H{"error": err})
	}

	data := loginResult{
		Token: token,
		UserInfo: UserInfo {
			User: user,
		},
    }
	c.JSON(200, gin.H{
        "status": 0,
        "msg":    "登录成功！",
        "data":   data,
    })
}

// Index 用户列表（分页）
func (i *UserController) Index(c *gin.Context) {

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	PerPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "0"))
	
	// Connection to the database
	db := database.DB

	var users []models.User
	// SELECT * FROM users

	paginator, err := pagination.Pagging(&pagination.Param{
        DB:      db,
        Page:    page,
        PerPage:   PerPage,
        OrderBy: []string{"id desc"},
	}, &users)
	
	// db.Find(&users)

	if err != nil {
		c.JSON(200, gin.H{
			"status": 1,
        	"msg": err,
		})
	}else{
		// Display JSON result
		c.JSON(200, gin.H{
			"status": 0,
			"msg": "success!",
			"data": users,
			"meta": paginator,
		})
	}
}

// Store 新增用户
func (i *UserController) Store(c *gin.Context) {
	db := database.DB

	var user models.User
	c.Bind(&user)

	if user.Firstname != "" && user.Lastname != "" {
		// INSERT INTO "users" (name) VALUES (user.Name);
		db.Create(&user)
		// Display error
		c.JSON(201, gin.H{"success": user})
	} else {
		// Display error
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}
}

// GetUsers 获取所有用户
func (i *UserController) GetUsers(c *gin.Context) {
	// Connection to the database
	db := database.DB

	var users []models.User
	// SELECT * FROM users
	db.Find(&users)

	// Display JSON result
	c.JSON(200, users)
}

// Show 获取单个用户信息
func (i *UserController) Show(c *gin.Context) {
	// Connection to the database
	db := database.DB

	id := c.Params.ByName("id")
	var user models.User
	// SELECT * FROM users WHERE id = 1;
	db.First(&user, id)

	if user.ID != 0 {
		// Display JSON result
		c.JSON(200, user)
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "User not found"})
	}
}

// Update 更新用户信息
func (i *UserController) Update(c *gin.Context) {
	// Connection to the database
	db := database.DB

	// Get id user
	id := c.Params.ByName("id")
	var user models.User
	// SELECT * FROM users WHERE id = 1;
	db.First(&user, id)

	if user.Firstname != "" && user.Lastname != "" {

		if user.ID != 0 {
			var newUser models.User
			c.Bind(&newUser)

			result := models.User{
				ID:        user.ID,
				Firstname: newUser.Firstname,
				Lastname:  newUser.Lastname,
			}

			// UPDATE users SET firstname='newUser.Firstname', lastname='newUser.Lastname' WHERE id = user.ID;
			db.Save(&result)
			// Display modified data in JSON message "success"
			c.JSON(200, gin.H{"success": result})
		} else {
			// Display JSON error
			c.JSON(404, gin.H{"error": "User not found"})
		}

	} else {
		// Display JSON error
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}
}

// Destroy 删除用户
func (i *UserController) Destroy(c *gin.Context) {
	// Connection to the database
	db := database.DB

	// Get id user
	id := c.Params.ByName("id")
	var user models.User
	// SELECT * FROM users WHERE id = 1;
	db.First(&user, id)

	if user.ID != 0 {
		// DELETE FROM users WHERE id = user.ID
		db.Delete(&user)
		// Display JSON result
		c.JSON(200, gin.H{"success": "User #" + id + " deleted"})
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "User not found"})
	}
}
