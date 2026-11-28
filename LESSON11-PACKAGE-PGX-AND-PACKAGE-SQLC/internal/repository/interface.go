package repository

import (
	"context"
	"hoc-golang/internal/db/sqlc"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(ctx context.Context, input sqlc.CreateUserParams) (sqlc.User, error)
	FindByUuid(ctx context.Context, uuid uuid.UUID) (sqlc.User, error)
}
