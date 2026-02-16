package camera

type Response struct {
	ID        uint   `json:"id"`
	Label     string `json:"label"`
	Location  string `json:"location"`
	StreamURL string `json:"streamUrl"`
	Status    string `json:"status"`
	SchoolID  uint   `json:"schoolId"`
	RoomID    *uint  `json:"roomId,omitempty"`
}
