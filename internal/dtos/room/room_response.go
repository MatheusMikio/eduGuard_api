package room

import "github.com/MatheusMikio/eduGuard_api/internal/domain/schemas"

type Response struct {
	ID       uint             `json:"id"`
	Name     string           `json:"name"`
	Floor    string           `json:"floor"`
	SchoolID uint             `json:"schoolId"`
	Cameras  []CamerasSummary `json:"cameras,omitempty"`
}

type CamerasSummary struct {
	ID       uint                 `json:"id"`
	Label    string               `json:"label"`
	Location string               `json:"location"`
	Status   schemas.CameraStatus `json:"status"`
	RoomID   *uint                `json:"roomId,omitempty"`
}
