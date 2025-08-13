package repositories_test

import (
	"context"
	"os"
	"testing"

	"github.com/dmarins/my-finances/internal/infrastructure/aws"
	"github.com/dmarins/my-finances/internal/infrastructure/aws/dynamo"
	"github.com/dmarins/my-finances/internal/infrastructure/localstack"
)

var db dynamo.IDb
var failedDb dynamo.IDb
var ctx context.Context = context.Background()

func TestMain(m *testing.M) {
	// Initialize the environment variables
	config := aws.NewConfig()

	// Create a new DynamoDB clients
	db = dynamo.NewDynamoDBClient(config)
	failedDb = dynamo.NewDynamoDBFailedClient(config)

	// Create the DynamoDB tables
	localstack.InitLocalstack(db)

	// Run the tests
	exitCode := m.Run()

	// Exit with the appropriate code
	os.Exit(exitCode)
}
