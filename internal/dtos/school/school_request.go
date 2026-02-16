package school

import "github.com/MatheusMikio/eduGuard_api/internal/domain/models"

type Request struct {
	Name   string        `json:"name" binding:"required"`
	CNPJ   string        `json:"cnpj" binding:"required,len=18"`
	Adress models.Adress `json:"adress" binding:"required"`
}
