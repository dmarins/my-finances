package repositories_test

import (
	"fmt"
	"testing"

	"github.com/dmarins/my-finances/internal/adapters/repositories"
	"github.com/dmarins/my-finances/internal/domain/entities"
	"github.com/stretchr/testify/assert"
)

func TestGetByName(t *testing.T) {
	userID := "123e4567-e89b-12d3-a456-426614174000"
	sut := repositories.NewCategoryRepository(db)

	tests := []struct {
		name         string
		categoryName string
		wantName     string
		wantNil      bool
	}{
		{"existent category", "receitas", "Receitas", false},
		{"non-existent category", "sbrubles", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			category, err := sut.GetByName(ctx, userID, tt.categoryName)
			assert.NoError(t, err)
			if tt.wantNil {
				assert.Nil(t, category)
			} else {
				assert.Equal(t, tt.wantName, category.Name)
			}
		})
	}
}

func TestGetByName_WhenDbFails(t *testing.T) {
	userID := "123e4567-e89b-12d3-a456-426614174000"
	sut := repositories.NewCategoryRepository(failedDb)

	category, err := sut.GetByName(ctx, userID, "receitas")

	assert.Error(t, err)
	assert.Nil(t, category)
}

func TestCreateCategory(t *testing.T) {
	userID := "123e4567-e89b-12d3-a456-426614174000"
	sut := repositories.NewCategoryRepository(db)

	tests := []struct {
		name     string
		category *entities.Category
		wantErr  bool
	}{
		{
			name: "new category",
			category: &entities.Category{
				PK:   fmt.Sprintf("USER#%s", userID),
				SK:   "CATEGORY#new-category",
				Name: "New Category",
			},
			wantErr: false,
		},
		{
			name: "existing category",
			category: &entities.Category{
				PK:   fmt.Sprintf("USER#%s", userID),
				SK:   "CATEGORY#receitas",
				Name: "Receitas",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := sut.CreateCategory(ctx, tt.category)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, "categoria já existe", err.Error())
			} else {
				assert.NoError(t, err)

				created, err := sut.GetByName(ctx, userID, "new-category")

				assert.NoError(t, err)
				assert.NotNil(t, created)
				assert.Equal(t, tt.category.Name, created.Name)
			}
		})
	}
}

func TestCreateCategory_WhenDbFails(t *testing.T) {
	userID := "123e4567-e89b-12d3-a456-426614174000"
	sut := repositories.NewCategoryRepository(failedDb)

	category := &entities.Category{
		PK:   fmt.Sprintf("USER#%s", userID),
		SK:   "CATEGORY#new-category",
		Name: "New Category",
	}

	err := sut.CreateCategory(ctx, category)

	assert.Error(t, err)
}
