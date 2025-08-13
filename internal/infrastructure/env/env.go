package env

import (
	"os"
	"strconv"
)

const (
	HeaderContentType = "Content-Type"
	ContentTypeJSON   = "application/json;charset=utf8"

	Env        = "ENV"
	DefaultEnv = "dev"

	// DYNAMODB
	DynamoTable = "MyFinances"

	// AWS
	AwsAccount    = "AWS_ACCOUNT"
	AwsEndpoint   = "AWS_ENDPOINT"
	AwsRegion     = "AWS_REGION"
	AWSDisableTLS = "AWS_DISABLE_TLS"

	DefaultAwsAccount    = "000000000000"
	DefaultAwsEndpoint   = "http://localhost:4566"
	DefaultAwsRegion     = "sa-east-1"
	DefaultAWSDisableTLS = false
)

func GetEnvWithDefaultAsString(envKey string, defaultVal string) string {
	val := os.Getenv(envKey)
	if val == "" {
		return defaultVal
	}

	return val
}

func GetEnvWithDefaultAsInt(envKey string, defaultVal int) int {
	val := os.Getenv(envKey)
	if val == "" {
		return defaultVal
	}

	intValue, err := strconv.Atoi(val)
	if err != nil {
		return defaultVal
	}

	return intValue
}

func GetEnvWithDefaultAsBoolean(envKey string, defaultVal bool) bool {
	val := os.Getenv(envKey)
	result, err := strconv.ParseBool(val)
	if err != nil {
		return defaultVal
	}

	return result
}
