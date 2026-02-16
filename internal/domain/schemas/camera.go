package schemas

import "gorm.io/gorm"

type CameraStatus string

const (
	CameraStatusOnline  CameraStatus = "ONLINE"
	CameraStatusOffline CameraStatus = "OFFLINE"
)

type Camera struct {
	gorm.Model
	Label     string       `gorm:"not null"`
	Location  string       `gorm:"not null"`
	StreamURL string       `gorm:"not null"`
	Status    CameraStatus `gorm:"not null;default:'OFFLINE'"`
	SchoolID  uint         `gorm:"not null"`
	School    *School      `gorm:"foreignKey:SchoolID"`
	RoomID    *uint        `gorm:"default:null"`
	Room      *Room        `gorm:"foreignKey:RoomID"`
	Events    []Event
}
