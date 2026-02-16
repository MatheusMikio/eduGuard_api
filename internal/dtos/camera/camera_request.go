package camera

type Request struct {
	Label     string `json:"label" binding:"required"`
	Location  string `json:"location" binding:"required"`
	StreamURL string `json:"streamUrl" binding:"required"`
	SchoolID  uint   `json:"schoolId" binding:"required"`
	RoomID    *uint  `json:"roomId,omitempty"`
}
