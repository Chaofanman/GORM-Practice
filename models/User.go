package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Firstname string
	Lastname  string
	Age       int64
	School    School
	SchoolID  int64
}

//type UserModel struct{}

// var db = databases.InitDB()

// func (user UserModel) AllUsers() (users []User) {
// 	db.Find(&users)

// 	return users
// }
