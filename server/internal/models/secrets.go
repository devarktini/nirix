package models

import "gorm.io/gorm"

type Secret struct {
	gorm.Model
	Name        string `gorm:"type:varchar(100);not null;uniqueIndex"`
	Value       string `gorm:"type:text;not null"`
	Environment string `gorm:"type:varchar(50);not null"`
}
