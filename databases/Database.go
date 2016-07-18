package databases

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm_practice/models"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost user=gchasepatron dbname=gormtest sslmode=disable password=psqlpassword")

	if err != nil {
		panic(err)
		return nil
	}

	db.AutoMigrate(&models.User{}, &models.School{})
	//db.Model(&models.User{}).Related(&models.School{})
	return db
}
