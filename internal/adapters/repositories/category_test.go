package repositories_test

import (
	"testing"

	"github.com/dmarins/my-finances/internal/adapters/repositories"
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
