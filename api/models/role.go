package models

import "github.com/jinzhu/gorm"

type Role struct {
	gorm.Model
	Role        string `gorm:"size:100;not null;unique" json:"role"`
	AccessLevel string `gorm:"size:1;not null" json:"access_level"`
}
