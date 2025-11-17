package models

import "gorm.io/gorm"

type Roles string

const (
	AdminRole        Roles = "admin"
	AllAccessRole    Roles = "user"
	NoAccessRole     Roles = "no-access"
	CustomAccessRole Roles = "custom"
)

type Policy struct {
	gorm.Model
	Role Roles         `gorm:"type:varchar(50);not null"`
	Apps []Application `gorm:"many2many:policy_apps;"`
}
