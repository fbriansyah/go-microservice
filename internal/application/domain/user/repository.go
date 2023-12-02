package user

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	GetAll(ctx context.Context) ([]*User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*User, error)
	Update(ctx context.Context, user *User) error
	Deleted(ctx context.Context, id uuid.UUID) error
	Save(ctx context.Context, user *User) error
	GetByEmail(ctx context.Context, email string) (*User, error)
}
