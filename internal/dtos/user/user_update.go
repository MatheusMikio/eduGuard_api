package user

import "github.com/MatheusMikio/eduGuard_api/internal/domain/schemas"

type Update struct {
	ID       uint          `json:"id"`
	Name     *string       `json:"name,omitempty"`
	Email    *string       `json:"email,omitempty" binding:"omitempty,email"`
	Password *string       `json:"password,omitempty" binding:"omitempty,min=6"`
	Role     *schemas.Role `json:"role,omitempty"`
	SchoolID *uint         `json:"schoolId,omitempty"`
}
