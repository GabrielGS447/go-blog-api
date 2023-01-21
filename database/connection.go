package database

import (
	"fmt"

	"github.com/gabrielgs447/go-blog-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect(url string, reset bool) error {
	var err error

	db, err = gorm.Open(mysql.Open(url), &gorm.Config{}) // Schema needs to be created manually
	if err != nil {
		return err
	}

	if reset {
		resetDB()
	} else {
		db.AutoMigrate(&models.User{}, &models.Post{})
	}

	fmt.Println("Connection Opened to Database")
	return nil
}

func Disconnect() {
	sqlDB, _ := db.DB()
	sqlDB.Close()
}

func resetDB() {
	db.Migrator().DropTable(&models.User{}, &models.Post{})
	db.AutoMigrate(&models.User{}, &models.Post{})
	seedUsers()
	seedPosts()
}
