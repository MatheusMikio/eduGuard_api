package schemas

import (
	"time"

	"gorm.io/gorm"
)

type EventType string

const (
	EventAgglomeration EventType = "AGGLOMERATION"
	EventAbruptMotion  EventType = "ABRUPT_MOTION"
	EventSurrounding   EventType = "SURROUNDING"
)

type EventStatus string

const (
	EventStatusPending  EventStatus = "PENDING"
	EventStatusReviewed EventStatus = "REVIEWED"
	EventStatusIgnored  EventStatus = "IGNORED"
)

type Event struct {
	gorm.Model
	CameraID     uint        `gorm:"not null"`
	Camera       *Camera     `gorm:"foreignKey:CameraID"`
	Type         EventType   `gorm:"not null"`
	RiskScore    float64     `gorm:"not null"`
	Status       EventStatus `gorm:"not null;default:'PENDING'"`
	StartedAt    time.Time   `gorm:"not null"`
	EndedAt      time.Time   `gorm:"not null"`
	VideoClipURL string      `gorm:"type:text"`
	ReviewedByID *uint
	ReviewedBy   *User `gorm:"foreignKey:ReviewedByID"`
	ReviewedAt   *time.Time
	Notes        string
}
