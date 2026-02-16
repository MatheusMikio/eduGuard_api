package camera

import "github.com/MatheusMikio/eduGuard_api/internal/domain/schemas"

type Update struct {
	ID        uint                 `json:"id" binding:"required"`
	Label     *string              `json:"label,omitempty"`
	Location  *string              `json:"location,omitempty"`
	StreamURL *string              `json:"streamUrl,omitempty"`
	RoomID    *uint                `json:"roomId,omitempty"`
	Status    schemas.CameraStatus `json:"status"`
}
