package models

import (
	"github.com/jinzhu/gorm"
	"github.com/thoas/go-funk"
	"time"
)

type BaseModel struct {
	DB *gorm.DB
}

type BaseSchema struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

func MapPermit(m map[string]interface{}, permits ...string) {
	for key := range m {
		if !funk.ContainsString(permits, key) {
			delete(m, key)
		}
	}
}
