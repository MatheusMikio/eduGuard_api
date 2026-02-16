package school

import "github.com/MatheusMikio/eduGuard_api/internal/domain/models"

type Update struct {
	Name   *string        `json:"name,omitempty"`
	CNPJ   *string        `json:"cnpj,omitempty" binding:"omitempty,len=18"`
	Adress *models.Adress `json:"adress,omitempty"`
}
