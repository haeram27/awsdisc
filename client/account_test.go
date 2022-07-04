package client

import (
	"awsdisc/apps"
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
)

func TestStsAssumeRole(t *testing.T) {
	cfg, err := StaticCredentialConfig("", "", "")
	if err != nil {
		apps.Logs.Error("StaticCredentialConfig error: ", err.Error())
	}

	creds, err := cfg.Credentials.Retrieve(context.Background())
	if err == nil {
		apps.Logs.Debug(fmt.Sprintf("%+v", creds))
	}

	AssumeRoleConfig(&cfg, ``)
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
