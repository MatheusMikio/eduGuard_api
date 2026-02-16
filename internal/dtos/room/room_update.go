package room

type Update struct {
	ID    uint   `json:"id" binding:"required"`
	Name  string `json:"name"`
	Floor string `json:"floor"`
}
