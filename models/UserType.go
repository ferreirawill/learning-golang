package models

import "gorm.io/gorm"

type UserType struct {
	gorm.Model
	Type string `gorm:"foreignKey:ID"`
}