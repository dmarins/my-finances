package localstack

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/dmarins/my-finances/internal/infrastructure/aws/dynamo"
	"github.com/dmarins/my-finances/internal/infrastructure/env"
	"github.com/dmarins/my-finances/internal/infrastructure/path"
)

func InitLocalstack(db dynamo.IDb) {
	CreateDynamoTables(db)
}

func CreateDynamoTables(db dynamo.IDb) {
	if !TableExists(db, env.DynamoTable) {
		CreateDynamoTable(db)
		PutItemsInTable(db, "scripts/test/categories.json")
	}
}

func TableExists(db dynamo.IDb, tableName string) bool {
	_, err := db.GetClient().DescribeTable(context.TODO(),
		&dynamodb.DescribeTableInput{
			TableName: aws.String(tableName),
		})
	return err == nil
}

func CreateDynamoTable(db dynamo.IDb) {
	var tableInput *dynamodb.CreateTableInput = BuildTableInput()
	_, err := db.GetClient().CreateTable(context.TODO(), tableInput)
	if err != nil {
		log.Print("Got error calling CreateTable: ", err)
		return
	}
}

func PutItemsInTable(db dynamo.IDb, file string) {
	baseDir, _ := path.GetProjectRoot()
	content, err := os.ReadFile(filepath.Clean(filepath.Join(baseDir, file)))
	if err != nil {
		log.Print("Error reading json to put item: ", err)
		return
	}

	lm := make([]map[string]any, 0)
	err = json.Unmarshal(content, &lm)
	if err != nil {
		log.Print("Error parsing json to put item: ", err)
		return
	}

	for _, m := range lm {
		item, err := attributevalue.MarshalMap(m)
		if err != nil {
			log.Print("Error parsing map: ", err)
			return
		}
		_, err = db.GetClient().PutItem(context.TODO(), &dynamodb.PutItemInput{
			TableName: aws.String(env.DynamoTable),
			Item:      item,
		})
		if err != nil {
			log.Print("Error put item: ", err)
			return
		}
	}
}

func BuildTableInput() *dynamodb.CreateTableInput {
	return &dynamodb.CreateTableInput{
		TableName: aws.String(env.DynamoTable),
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("PK"),
				AttributeType: types.ScalarAttributeTypeS,
			},
			{
				AttributeName: aws.String("SK"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("PK"),
				KeyType:       types.KeyTypeHash,
			},
			{
				AttributeName: aws.String("SK"),
				KeyType:       types.KeyTypeRange,
			},
		},
		ProvisionedThroughput: &types.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(2),
			WriteCapacityUnits: aws.Int64(2),
		},
	}
}
