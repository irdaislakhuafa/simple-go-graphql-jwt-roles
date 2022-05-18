package config

import (
	"fmt"
	"log"
	"os"

	"github.com/irdaislakhuafa/simple-go-graphql-jwt-roles/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// to get DB from other package
func GetDB() *gorm.DB {
	return db
}

// for initialize database connection
func InitDB() {
	log.Println("entering method to initialize database connection")

	url := getStringUrlFromEnv()
	var err error
	db, err = gorm.Open(mysql.Open(*url), &gorm.Config{})
	if err != nil {
		panic("error while connect to database: " + err.Error())
	}

	log.Printf(`success database connected to "%s"`, os.Getenv("DB_DATABASE"))
	enableMigrations()
}

func getStringUrlFromEnv() *string {
	log.Println("entering method to generate strung url from .env")
	url := `%s:%s@tcp(%s:%s)/%s?multiStatements=true&parseTime=true`
	url = fmt.Sprintf(
		url,
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)

	log.Println("success generate string database url")
	return &url
}

func enableMigrations() {
	log.Println("entering method to migrate table to database")
	GetDB().AutoMigrate(&entities.User{}, &entities.Role{})
	log.Println("success migrating table")
}
