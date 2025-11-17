package models

import "gorm.io/gorm"

type Activities struct {
	gorm.Model
	UserID      uint
	User        User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Action      string `gorm:"type:varchar(100);not null"`
	Description string `gorm:"type:text"`
}
