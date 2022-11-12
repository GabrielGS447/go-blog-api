package db

import (
	"fmt"
	"os"

	"github.com/gabrielgaspar447/go-blog-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(reset bool) {
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPass := os.Getenv("MYSQL_PASS")
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlPort := os.Getenv("MYSQL_PORT")
	mysqlDb := os.Getenv("MYSQL_DB")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysqlUser, mysqlPass, mysqlHost, mysqlPort, mysqlDb)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) // Schema needs to be created manually
	if err != nil {
		panic("failed to connect database")
	}

	if reset {
		resetDB(db)
	} else {
		db.AutoMigrate(&models.User{}, &models.Post{})
	}

	fmt.Println("Connection Opened to Database")
	DB = db
}

func resetDB(db *gorm.DB) {
	db.Migrator().DropTable(&models.User{}, &models.Post{})
	db.AutoMigrate(&models.User{}, &models.Post{})
	seed(db)
}
