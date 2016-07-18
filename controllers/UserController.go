package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm_practice/databases"
	"gorm_practice/models"
	"strconv"
)

type UserController struct{}

//var userModel = new(models.UserModel)

var db = databases.InitDB()

func (ctrl UserController) GetUsers(c *gin.Context) {
	var users []models.User

	//raw sql in gorm
	//.scan saves it into a var
	err := db.Raw("SELECT * FROM users WHERE deleted_at IS  NULL").Scan(&users).Error

	//relation
	for i, _ := range users {
		db.Model(users[i]).Related(&users[i].School)
	}

	if err != nil {
		c.JSON(404, gin.H{"error": "Query error in GetUsers"})
		panic(err)
	}

	if users == nil {
		c.JSON(404, gin.H{"error": "No rows in user table"})
	}

	c.JSON(200, users)
}

func (ctrl UserController) GetUser(c *gin.Context) {
	var user models.User

	id := c.Params.ByName("id")
	user_id, _ := strconv.ParseInt(id, 0, 64)
	fmt.Println(id)

	//raw sql
	err := db.Raw("SELECT * FROM users WHERE ID=$1 AND deleted_at IS NULL", user_id).Scan(&user).Error

	//relation
	db.Model(user).Related(&user.School)

	if err != nil {
		c.JSON(404, gin.H{"error": "Query error in GetUser"})
		panic(err)
	}

	c.JSON(200, user)
}

func (ctrl *UserController) PostUser(c *gin.Context) {
	var ParsedUser models.User

	//bind the post to the model
	c.Bind(&ParsedUser)

	content := &models.User{
		Firstname: ParsedUser.Firstname,
		Lastname:  ParsedUser.Lastname,
		Age:       ParsedUser.Age,
		SchoolID:  ParsedUser.SchoolID,
	}

	db.Create(&content)

	c.JSON(200, content)
}

func (ctrl *UserController) UpdateUser(c *gin.Context) {
	var ParsedUser models.User
	var user models.User

	id := c.Params.ByName("id")
	user_id, _ := strconv.ParseInt(id, 0, 64)

	c.Bind(&ParsedUser)
	// fmt.Println("Parsed User: ", ParsedUser.Firstname, ParsedUser.Lastname, ParsedUser.Age, ParsedUser.SchoolID)

	//Get user to be updated
	err := db.Raw("SELECT * FROM users WHERE ID=$1 AND deleted_at IS  NULL", user_id).Scan(&user).Error

	if err != nil {
		c.JSON(404, gin.H{"error": "Error getting user to update"})
		panic(err)
	}

	err = db.Model(&user).Updates(map[string]interface{}{
		"Firstname": ParsedUser.Firstname,
		"Lastname":  ParsedUser.Lastname,
		"Age":       ParsedUser.Age,
		"SchoolID":  ParsedUser.SchoolID,
	}).Error

	if err != nil {
		c.JSON(404, gin.H{"error": "Query error in UpdateUser"})
		panic(err)
	}

	c.JSON(200, user)
}

func (ctrl *UserController) DeleteUser(c *gin.Context) {
	var user models.User

	id := c.Params.ByName("id")
	user_id, _ := strconv.ParseInt(id, 0, 64)

	err := db.Raw("SELECT * FROM users WHERE ID=$1 AND deleted_at IS  NULL", user_id).Scan(&user).Error

	if err != nil {
		c.JSON(404, gin.H{"error": "Error getting user to update"})
		panic(err)
	}

	err = db.Delete(&user).Error

	if err != nil {
		c.JSON(404, gin.H{"error": "Error deleting user"})
		panic(err)
	}

	c.JSON(200, gin.H{"Success": "Deleted user"})

}
