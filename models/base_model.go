package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type BaseModel struct {
	DB *gorm.DB
}

type BaseSchema struct {
	ID        uint `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}
