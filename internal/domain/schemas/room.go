package schemas

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	Name     string  `gorm:"not null"`
	Floor    string  `gorm:"not null"`
	SchoolID uint    `gorm:"not null"`
	School   *School `gorm:"foreignKey:SchoolID"`
	Cameras  []Camera
}
