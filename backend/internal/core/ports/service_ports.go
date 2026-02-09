package ports

import (
	"context"

	"github.com/moneymon/internal/core/domain"
)

type UserService interface {
	Register(ctx context.Context, req *domain.RegisterRequest) (*domain.User, error)
	Login(ctx context.Context, req *domain.LoginRequest) (string, error)
}
