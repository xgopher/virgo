package controllers

import (
	"app/database"
	"app/modules/user/models"
	"github.com/gin-gonic/gin"
)

// UserController ...
type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

// ...
func (i *UserController) Index(c *gin.Context) {
	// Connection to the database
	db := database.DB

	var users []models.User
	// SELECT * FROM users
	db.Find(&users)

	// Display JSON result
	c.JSON(200, users)
}

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

func (i *UserController) GetUsers(c *gin.Context) {
	// Connection to the database
	db := database.DB

	var users []models.User
	// SELECT * FROM users
	db.Find(&users)

	// Display JSON result
	c.JSON(200, users)
}

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
