package entities

import "gorm.io/gorm"

type CameraStatus string

const (
	CameraStatusOnline  CameraStatus = "ONLINE"
	CameraStatusOffline CameraStatus = "OFFLINE"
)

type Camera struct {
	gorm.Model
	Label     string       `gorm:"not null"`
	StreamURL string       `gorm:"not null"`
	Status    CameraStatus `gorm:"not null;default:'OFFLINE'"`
	RoomID    uint         `gorm:"not null"`
	Room      *Room
	Events    []Event
}
