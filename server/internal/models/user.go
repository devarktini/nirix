package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100);not null"`
	Username string `gorm:"type:varchar(50);uniqueIndex;not null"`
	Email    string `gorm:"type:varchar(100);uniqueIndex;not null"`
	PolicyID uint
	Policy   Policy `gorm:"foreignKey:PolicyID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
