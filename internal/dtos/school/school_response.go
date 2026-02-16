package school

import (
	"time"

	"github.com/MatheusMikio/eduGuard_api/internal/domain/models"
)

type Response struct {
	ID        uint          `json:"id"`
	Name      string        `json:"name"`
	CNPJ      string        `json:"cnpj"`
	Adress    models.Adress `json:"adress"`
	RoomIDs   []uint        `json:"room_ids,omitempty"`
	UserIDs   []uint        `json:"user_ids,omitempty"`
	CameraIDs []uint        `json:"camera_ids,omitempty"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}
