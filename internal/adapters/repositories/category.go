package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/dmarins/my-finances/internal/domain/entities"
	"github.com/dmarins/my-finances/internal/domain/repositories"
	"github.com/dmarins/my-finances/internal/infrastructure/aws/dynamo"
	"github.com/dmarins/my-finances/internal/infrastructure/env"
)

type CategoryRepository struct {
	Db dynamo.IDb
}

func NewCategoryRepository(db dynamo.IDb) repositories.ICategoryRepository {
	return &CategoryRepository{
		Db: db,
	}
}

func (r *CategoryRepository) GetByName(ctx context.Context, userID, name string) (*entities.Category, error) {
	result, err := r.Db.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(env.DynamoTable),
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{Value: fmt.Sprintf("USER#%s", userID)},
			"SK": &types.AttributeValueMemberS{Value: fmt.Sprintf("CATEGORY#%s", name)},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar categoria: %w", err)
	}

	if result.Item == nil {
		return nil, nil
	}

	var category entities.Category
	err = attributevalue.UnmarshalMap(result.Item, &category)
	if err != nil {
		return nil, fmt.Errorf("erro ao converter categoria do DynamoDB: %w", err)
	}

	return &category, nil
}

func (r *CategoryRepository) CreateCategory(ctx context.Context, category *entities.Category) error {
	item, err := attributevalue.MarshalMap(category)
	if err != nil {
		return fmt.Errorf("erro ao converter categoria para DynamoDB: %w", err)
	}

	_, err = r.Db.PutItem(ctx, &dynamodb.PutItemInput{
		TableName:           aws.String(env.DynamoTable),
		Item:                item,
		ConditionExpression: aws.String("attribute_not_exists(PK) AND attribute_not_exists(SK)"),
	})
	if err != nil {
		var conditionalCheckFailed *types.ConditionalCheckFailedException
		if errors.As(err, &conditionalCheckFailed) {
			return errors.New("categoria já existe")
		}

		return fmt.Errorf("erro ao criar categoria: %w", err)
	}

	return nil
}
