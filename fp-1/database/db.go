package database

import (
	"fmt"
	"log"

	"project-1/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "postgres"
	dbPort   = "5432"
	dbname   = "hacktiv8-project-1"
	db       *gorm.DB
	err      error
)

func StartDB() {
	// using ssl
	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=Asia/Shanghai", host, user, password, dbname, dbPort)
	// without ssl
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=Asia/Shanghai", host, user, password, dbname, dbPort)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database :", err)
	}

	fmt.Println("Database connection success.")
	_ = db.AutoMigrate(models.Todo{})

}

func GetDB() *gorm.DB {
	return db
}
