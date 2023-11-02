package awssdk

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/ruelala/aws-mfa/pkg/utils"
)

func build_config(profile string) aws.Config {
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithSharedConfigProfile(profile),
		config.WithSharedConfigFiles(
			[]string{
				config.DefaultSharedConfigFilename(),
				config.DefaultSharedCredentialsFilename(),
			}),
		config.WithClientLogMode(aws.LogRetries|aws.LogRequest),
	)
	utils.Panic(err)
	return cfg
}

func IAMClient(profile string) *iam.Client {
	cfg := build_config(profile)
	cli_opt := iam.Options{
		Credentials: cfg.Credentials,
		Region:      "us-east-1",
	}
	client := iam.New(cli_opt)
	return client
}

func STSClient(profile string) *sts.Client {
	cfg := build_config(profile)
	cli_opt := sts.Options{
		Credentials: cfg.Credentials,
		Region:      "us-east-1",
	}
	client := sts.New(cli_opt)
	return client
}
