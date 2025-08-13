package models

type (
	CategoryCreateInput struct {
		Name string `json:"name" validate:"required,max=100"`
	}

	CategoryOutput struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
)
