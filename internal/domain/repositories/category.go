package repositories

import (
	"context"

	"github.com/dmarins/my-finances/internal/domain/entities"
)

type ICategoryRepository interface {
	GetByName(ctx context.Context, userID, name string) (*entities.Category, error)
	CreateCategory(ctx context.Context, category *entities.Category) error
}
