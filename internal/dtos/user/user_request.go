package user

import "github.com/MatheusMikio/eduGuard_api/internal/domain/schemas"

type User struct {
	Name     string       `gorm:"not null"`
	Email    string       `gorm:"uniqueIndex;not null"`
	Password string       `gorm:"not null"`
	Role     schemas.Role `gorm:"not null;default:'VIEWER'"`
	SchoolID uint         `gorm:"not null"`
}
