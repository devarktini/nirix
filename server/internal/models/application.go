package models

import "gorm.io/gorm"

type Application struct {
	gorm.Model
	Name         string `gorm:"type:varchar(100);not null"`
	Description  string `gorm:"type:text"`
	Url          string `gorm:"type:varchar(255);not null;uniqueIndex"`
	AuthType     string `gorm:"type:varchar(50);not null"`
	AuthUrl      string `gorm:"type:varchar(255)"`
	ApiSecret    string `gorm:"type:varchar(255)"`
	CallbackUrl  string `gorm:"type:varchar(255)"`
	IconUrl      string `gorm:"type:varchar(255)"`
	Subscription string `gorm:"type:varchar(100)"`
	Pricing      string `gorm:"type:varchar(100)"`
}
