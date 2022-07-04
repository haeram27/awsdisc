package client

import (
	"awsdisc/apps"
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
)

func TestStsAssumeRole(t *testing.T) {
	c, err := ReadCredentialsFromFile("")
	if err != nil {
		apps.Logs.Error("failed to get credentials from file: ", err.Error())
	}

	cfg, err := StaticCredentialConfig(c.AccessKeyId, c.SecretAccessKey, "")
	if err != nil {
		apps.Logs.Error("StaticCredentialConfig error: ", err.Error())
	}

	creds, err := cfg.Credentials.Retrieve(context.Background())
	if err == nil {
		apps.Logs.Debug(fmt.Sprintf("%+v", creds))
	}

	AssumeRoleConfig(&cfg, c.RoleArn)
	creds, err = cfg.Credentials.Retrieve(context.Background())
	if err == nil {
		apps.Logs.Debug(fmt.Sprintf("%+v", creds))
	}
}

func TestAwsConfig(t *testing.T) {
	cfg := aws.NewConfig()
	apps.Logs.Debug(fmt.Sprintf("%+v", cfg))
	if cfg.Credentials == nil {
		apps.Logs.Debug("credentials are nil!!!")
	}
}

func TestReadUser(t *testing.T) {
	ReadCredentialsFromFile("")
}
