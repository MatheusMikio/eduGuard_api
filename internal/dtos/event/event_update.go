package event

import (
	"time"

	"github.com/MatheusMikio/eduGuard_api/internal/domain/schemas"
)

type Update struct {
	ID           uint                `json:"id" binding:"required"`
	Status       schemas.EventStatus `json:"status"`
	ReviewedByID *uint               `json:"reviewedById"`
	ReviewedAt   *time.Time          `json:"reviewedAt"`
	Notes        string              `json:"notes"`
}
