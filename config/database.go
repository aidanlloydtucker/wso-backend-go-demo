package config

import (
	"log"

	"github.com/aidanlloydtucker/wso-backend-go-demo/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func LoadDatabase(cfg *Config) *gorm.DB {
	db, err := gorm.Open(cfg.DatabaseType, cfg.DatabaseArgs)
	if err != nil {
		log.Fatalln("failed to connect database")
	}

	if cfg.IsDevelopment() || cfg.IsTest() {
		db.LogMode(true)
	}

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
