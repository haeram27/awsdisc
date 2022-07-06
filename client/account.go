package client

import (
	apps "awsdisc/apps"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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

type CicdCreds struct {
	AccessKeyId     string `json:"AccessKeyId"`
	SecretAccessKey string `json:"SecretAccessKey"`
	SessionToken    string `json:"SessionToken,omitempty"`
	Expiration      string `json:"Expiration,omitempty"`
	RoleArn         string `json:"RoleArn,omitempty"`
}

func AwsConfig() *aws.Config {
	onceAwsCfg.Do(func() {
		awscfg = StsAssumeRoleConfigFromFile()
	})

	return &awscfg
}

/*
/tmp/awsuser.json:
{
	"AccessKeyId": "access_key_id",
	"SecretAccessKey": "secret_access_key",
	"SessionToken": "session_token",
	"Expiration": "expiration",
	"RoleArn": "role_arn"
}
*/
func ReadCredentialsFromFile(path string) (CicdCreds, error) {
	if path == "" {
		path = `/tmp/awsuser.json`
	}

	f, err := os.Open(`/tmp/awsuser.json`)
	if err != nil {
		apps.Logs.Error(err)
		return CicdCreds{}, err
	}
	defer f.Close()

	fmt.Println("Successfully Opened users.json")
	j, _ := ioutil.ReadAll(f)

	cred := CicdCreds{}
	err = json.Unmarshal(j, &cred)
	if err != nil {
		apps.Logs.Error(err)
		return CicdCreds{}, err
	}

	return cred, nil
}

func DefaultConfig() (aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		apps.Logs.Error(err)
		return *aws.NewConfig(), err
	}

	return cfg, nil
}

func SharedProfileConfig(profile string) (aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile(profile))
	if err != nil {
		apps.Logs.Error(err)
		return *aws.NewConfig(), err
	}

	return cfg, nil
}

func StaticCredentialConfig(akid string, seckey string, token string) (aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(akid, seckey, token)))
	if err != nil {
		apps.Logs.Error(err)
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
	roleArn : arn of role
*/
func StsAssumeRoleConfig(akId string, secKey string, roleArn string) (aws.Config, error) {
	cfg, err := StaticCredentialConfig(akId, secKey, "")
	if err != nil {
		apps.Logs.Error(err)
		return aws.Config{}, err
	}

	AssumeRoleConfig(&cfg, roleArn)

	return cfg, nil
}

/*
	auth for testing
*/
func StsAssumeRoleConfigFromFile() aws.Config {
	c, err := ReadCredentialsFromFile("")
	if err != nil {
		apps.Logs.Error(err)
		return aws.Config{}
	}

	cfg, err := StsAssumeRoleConfig(c.AccessKeyId, c.SecretAccessKey, c.RoleArn)
	if err != nil {
		apps.Logs.Error(err)
		return aws.Config{}
	}

	return cfg
}
