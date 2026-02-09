package ports

import (
	"context"

	"github.com/moneymon/internal/core/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) error
	GetByUsername(ctx context.Context, username string) (*domain.User, error)
}
