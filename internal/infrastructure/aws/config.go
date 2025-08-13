package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/dmarins/my-finances/internal/infrastructure/env"
)

type (
	IConfig interface {
		GetAwsConfig() aws.Config
	}

	Config struct {
		AwsConfig aws.Config
	}
)

func NewConfig() IConfig {
	return &Config{
		AwsConfig: buildAwsConfig(),
	}
}

func (c *Config) GetAwsConfig() aws.Config {
	return c.AwsConfig
}

func buildAwsConfig() aws.Config {
	var options []func(*config.LoadOptions) error
	region := env.GetEnvWithDefaultAsString(env.AwsRegion, env.DefaultAwsRegion)
	options = append(options, config.WithRegion(region))

	isLocal := env.GetEnvWithDefaultAsString(env.Env, env.DefaultEnv) == env.DefaultEnv
	if isLocal {
		endpoint := env.GetEnvWithDefaultAsString(env.AwsEndpoint, env.DefaultAwsEndpoint)
		options = append(options,
			config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("test", "test", "")),
			config.WithEndpointResolverWithOptions(
				aws.EndpointResolverWithOptionsFunc(
					func(service, region string, _ ...interface{}) (aws.Endpoint, error) {
						return aws.Endpoint{
							URL:               endpoint,
							SigningRegion:     region,
							HostnameImmutable: true,
						}, nil
					},
				),
			),
		)
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(), options...)
	if err != nil {
		panic("Failed to start AWS session")
	}

	return cfg
}
