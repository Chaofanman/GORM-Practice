package main

import (
	_ "fmt"
	"github.com/gin-gonic/gin"
	"gorm_practice/controllers"
	_ "gorm_practice/models"
)

func main() {

	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		user := new(controllers.UserController)

		v1.GET("/users", user.GetUsers)
		v1.GET("/users/:id", user.GetUser)
		v1.POST("/users", user.PostUser)
		v1.PUT("/users/:id", user.UpdateUser)
		v1.DELETE("/users/:id", user.DeleteUser)
	}

	router.Run(":9090")
}
