package event

import (
	"time"

	"github.com/MatheusMikio/eduGuard_api/internal/domain/schemas"
)

type Request struct {
	CameraID     uint                `json:"cameraId"`
	Type         schemas.EventType   `json:"type"`
	RiskScore    float64             `json:"riskScore"`
	Status       schemas.EventStatus `json:"status"`
	StartedAt    time.Time           `json:"startedAt"`
	EndedAt      time.Time           `json:"endetAt"`
	VideoClipURL string              `json:"videoCLipURL"`
	ReviewedByID *uint               `json:"reviewedById"`
	ReviewedAt   *time.Time
	Notes        string
}
