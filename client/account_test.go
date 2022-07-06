package client

import (
	"testing"
)

func TestStsAssumeRole(t *testing.T) {
	c, err := ReadCredentialsFromFile("")
	if err != nil {
		t.Error(err)
	}

	_, err = StaticCredentialConfig(c.AccessKeyId, c.SecretAccessKey, "")
	if err != nil {
		t.Error(err)
	}

	/*
		creds, err := cfg.Credentials.Retrieve(context.Background())
		if err == nil {
			t.Logf("%+v", creds)
		}

		AssumeRoleConfig(&cfg, c.RoleArn)
		creds, err = cfg.Credentials.Retrieve(context.Background())
		if err == nil {
			t.Logf("%+v", creds)
		}
	*/
}

func TestReadUser(t *testing.T) {
	_, err := ReadCredentialsFromFile("")
	if err != nil {
		t.Error(err)
	}
}
