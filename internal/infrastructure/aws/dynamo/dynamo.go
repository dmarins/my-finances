package dynamo

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/dmarins/my-finances/internal/infrastructure/aws"
	"github.com/dmarins/my-finances/internal/infrastructure/env"
)

type (
	IDb interface {
		PutItem(ctx context.Context, params any) (any, error)
		UpdateItem(ctx context.Context, params any) (any, error)
		QueryItems(ctx context.Context, params any) (*dynamodb.QueryOutput, error)
		GetItem(ctx context.Context, params *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error)
		DeleteItem(ctx context.Context, params any) (any, error)
		GetClient() *dynamodb.Client
	}

	Db struct {
		client *dynamodb.Client
	}

	Cursor struct {
		PageLimit        int32
		Cursor           *string
		ScanIndexForward bool
	}
)

func NewDynamoDBClient(config aws.IConfig) IDb {
	awsCfg := config.GetAwsConfig()

	return &Db{
		client: dynamodb.NewFromConfig(awsCfg, func(options *dynamodb.Options) {
			options.EndpointOptions.DisableHTTPS = env.GetEnvWithDefaultAsBoolean(env.AWSDisableTLS, env.DefaultAWSDisableTLS)
		}),
	}
}

func NewDynamoDBFailedClient(config aws.IConfig) IDb {
	awsCfg := config.GetAwsConfig()

	return &Db{
		client: dynamodb.NewFromConfig(awsCfg, func(options *dynamodb.Options) {
			options.Region = "sbrubles"
			options.EndpointOptions.DisableHTTPS = true
		}),
	}
}

func (s *Db) PutItem(ctx context.Context, params any) (any, error) {
	result, err := s.client.PutItem(ctx, params.(*dynamodb.PutItemInput))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Db) GetItem(ctx context.Context, params *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	result, err := s.client.GetItem(ctx, params)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Db) QueryItems(ctx context.Context, params any) (*dynamodb.QueryOutput, error) {
	result, err := s.client.Query(ctx, params.(*dynamodb.QueryInput))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Db) DeleteItem(ctx context.Context, params any) (any, error) {
	result, err := s.client.DeleteItem(ctx, params.(*dynamodb.DeleteItemInput))
	if err != nil {
		return nil, err
	}

	return result, err
}

func (s *Db) UpdateItem(ctx context.Context, params any) (any, error) {
	result, err := s.client.UpdateItem(ctx, params.(*dynamodb.UpdateItemInput))
	if err != nil {
		return nil, err
	}

	return result, err
}

func (s *Db) GetClient() *dynamodb.Client {
	return s.client
}
