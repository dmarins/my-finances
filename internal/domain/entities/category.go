package entities

type (
	Category struct {
		PK   string `dynamodbav:"PK"`
		SK   string `dynamodbav:"SK"`
		Name string `dynamodbav:"Name"`
	}
)

func NewCategory(name string) *Category {
	return &Category{
		Name: name,
	}
}
