package models

import (
	"github.com/jinzhu/gorm"
)

type School struct {
	gorm.Model
	Name  string
	State string
}
