package entities

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Floor    string
	SchoolId uint `gorm:"not null"`
	School   *School
	Cameras  []Camera
}
