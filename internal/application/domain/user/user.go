package user

import (
	"github.com/fbriansyah/go-microservice/internal/application/domain"
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID
	Name     string
	Email    string
	Password string
	Status   domain.Status
}

func NewUser(name, email, password string) *User {
	status := domain.Active
	return &User{
		Name:     name,
		Email:    email,
		Password: password,
		Status:   status,
	}
}

func (u *User) Active() *User {
	u.Status = domain.Active
	return u
}

func (u *User) Inactive() *User {
	u.Status = domain.Inactive
	return u
}

func (u *User) Deleted() *User {
	u.Status = domain.Deleted
	return u
}
