package schemas

import (
	"github.com/MatheusMikio/eduGuard_api/internal/domain/models"
	"gorm.io/gorm"
)

type School struct {
	gorm.Model
	Name    string        `gorm:"not null"`
	CNPJ    string        `gorm:"uniqueIndex;size:18;not null"`
	Adress  models.Adress `gorm:"embedded"`
	Rooms   []Room
	Users   []User
	Cameras []Camera
}
