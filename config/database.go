package config

import (
	"github.com/WilliamsStudentsOnline/wso-backend/models"
	"github.com/jinzhu/gorm"
	"log"
)

import _ "github.com/jinzhu/gorm/dialects/sqlite"

func LoadDatabase() *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatalln("failed to connect database")
	}

	db.LogMode(true)

	// Would use a better migration library in real version
	err = db.AutoMigrate(
		&models.User{},
		&models.Department{}).Error
	if err != nil {
		log.Fatalln(err)
	}

	return db
}

func CloseDatabase(db *gorm.DB) {
	if err := db.Close(); err != nil {
		log.Fatalln(err)
	}
}