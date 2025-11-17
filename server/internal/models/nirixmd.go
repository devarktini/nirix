package models

import "gorm.io/gorm"

/*
NIRIX Metadata Model
*/
type NirixData struct {
	gorm.Model
	CustomerId string `gorm:"type:varchar(100);not null"`
	LicenseKey string `gorm:"type:varchar(255);not null;uniqueIndex"`
	ExpiryDate string `gorm:"type:varchar(50);not null"`
	Status     string `gorm:"type:varchar(50);not null"`
}
