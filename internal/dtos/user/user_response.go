package user

import (
	"time"

	"github.com/MatheusMikio/eduGuard_api/internal/domain/schemas"
)

type Response struct {
	ID        uint          `json:"id"`
	Name      string        `json:"name"`
	Email     string        `json:"email"`
	Role      schemas.Role  `json:"role"`
	SchoolID  uint          `json:"school_id"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}
