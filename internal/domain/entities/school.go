package entities

import "gorm.io/gorm"

type School struct {
	gorm.Model
	Name   string `gorm:"not null"`
	CNPJ   string `gorm:"uniqueIndex;size:18;not null"`
	Adress Adress
	Rooms  []Room
	Users  []User
}
