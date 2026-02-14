package entities

import (
	"time"

	"gorm.io/gorm"
)

type Role string

const (
	RoleAdmin       Role = "ADMIN"
	RoleCoordinator Role = "COORDINATOR"
	RoleViewer      Role = "VIEWER"
)

type User struct {
	gorm.Model
	Name          string `gorm:"not null"`
	Email         string `gorm:"uniqueIndex;not null"`
	PasswordHash  string `gorm:"not null"`
	Role          Role   `gorm:"not null;default:'VIEWER'"`
	SchoolId      uint   `gorm:"not null"`
	School        *School
	RefreshToken  *string `gorm:"type:text"`
	TokenIssuedAt *time.Time
}
