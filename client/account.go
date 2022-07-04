package client

import (
	"awsdisc/apps"
	"context"
	"fmt"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

// for singleton of aws.Config
var awscfg aws.Config
var onceAwsCfg sync.Once

func AwsConfig() *aws.Config {
	onceAwsCfg.Do(func() {
		awscfg = TestStsAssumeRoleConfig()
	})
	apps.Logs.Error(fmt.Sprintf("awscfg is created: %p", &awscfg))
	return &awscfg
}

func DefaultConfig() (aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		apps.Logs.Error("configuration error: ", err.Error())
		return *aws.NewConfig(), err
	}

	return cfg, nil
}

func SharedProfileConfig(profile string) (aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile(profile))
	if err != nil {
		apps.Logs.Error("configuration error: ", err.Error())
		return *aws.NewConfig(), err
	}

	return cfg, nil
}

func StaticCredentialConfig(akid string, seckey string, token string) (aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(akid, seckey, token)))
	if err != nil {
		apps.Logs.Error("configuration error: ", err.Error())
		return *aws.NewConfig(), err
	}

	return cfg, nil
}

/*
	AssumeRoleXXXConfig() makes aws.Config performed assume-role process.
	stsUserCfg can be made by  DefaultConfig(), SharedProfileConfig(), StaticCredentialConfig()
*/
func AssumeRoleConfig(stsUserCfg *aws.Config, roleArn string) {
	stsSvc := sts.NewFromConfig(*stsUserCfg)
	creds := stscreds.NewAssumeRoleProvider(stsSvc, roleArn)

	stsUserCfg.Credentials = aws.NewCredentialsCache(creds)
}

func AssumeRoleCustomMFAConfig(stsUserCfg *aws.Config, roleArn string, mfaSerialNumber string, mfaToken string) {
	staticTokenProvider := func() (string, error) {
		return mfaToken, nil
	}

	creds := stscreds.NewAssumeRoleProvider(sts.NewFromConfig(*stsUserCfg), roleArn, func(o *stscreds.AssumeRoleOptions) {
		o.SerialNumber = aws.String(mfaSerialNumber)
		o.TokenProvider = staticTokenProvider
	})

	stsUserCfg.Credentials = aws.NewCredentialsCache(creds)
}

/*
	akId : user's Access Key ID
	secKey : user's Secret Access Key
	roleArn :
*/
func StsAssumeRoleConfig(akId string, secKey string, roleArn string) (aws.Config, error) {
	cfg, err := StaticCredentialConfig(akId, secKey, "")
	if err != nil {
		apps.Logs.Error("StaticCredentialConfig error: ", err.Error())
		return aws.Config{}, err
	}

	AssumeRoleConfig(&cfg, roleArn)

	return cfg, nil
}

/*
	Example of Usage
*/
func TestStsAssumeRoleConfig() aws.Config {
	cfg, err := StsAssumeRoleConfig("", "", ``)
	if err != nil {
		apps.Logs.Error("StsAssumeRoleConfig error: ", err.Error())
		return aws.Config{}
	}

	return cfg
}
