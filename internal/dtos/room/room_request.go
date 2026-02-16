package room

type Request struct {
	Name     string `gorm:"not null"`
	Floor    string `gorm:"not null"`
	SchoolID uint   `gorm:"not null"`
}
