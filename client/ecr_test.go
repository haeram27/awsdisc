package client

import (
	awsutil "awsdisc/client/util"
	"encoding/json"
	"testing"
)

func TestEcrDiscoverAll(t *testing.T) {
	var jsonBlob []byte
	var result interface{}

	result, err := ECRDescribeRegistryCmd(AwsConfig())
	if err != nil {
		t.Error(err)
	} else {
		jsonBlob, err = json.Marshal(result)
		if err != nil {
			t.Error(err)
		}
	}
	t.Log(awsutil.PrettyJson(jsonBlob).String())

	result, err = ECRDescribeRepositoriesCmd(AwsConfig())
	if err != nil {
		t.Error(err)
	} else {
		jsonBlob, err = json.Marshal(result)
		if err != nil {
			t.Error(err)
		}
	}
	t.Log(awsutil.PrettyJson(jsonBlob).String())
}

func TestECRDescribeRegistryCmd(t *testing.T) {
	var jsonBlob []byte
	var result interface{}

	result, err := ECRDescribeRegistryCmd(AwsConfig())
	if err != nil {
		t.Error(err)
	} else {
		jsonBlob, err = json.Marshal(result)
		if err != nil {
			t.Error(err)
		}
	}
	t.Log(awsutil.PrettyJson(jsonBlob).String())
}

func TestECRDescribeRepositoriesCmd(t *testing.T) {
	var jsonBlob []byte
	var result interface{}

	result, err := ECRDescribeRepositoriesCmd(AwsConfig())
	if err != nil {
		t.Error(err)
	} else {
		jsonBlob, err = json.Marshal(result)
		if err != nil {
			t.Error(err)
		}
	}
	t.Log(awsutil.PrettyJson(jsonBlob).String())

	repos, err := awsutil.JsonPath(jsonBlob, "$.Repositories[:].RepositoryName")
	if err != nil {
		t.Error(err)
	}
	t.Logf("%v", repos)
}

func TestECRListImagesAll(t *testing.T) {
	ECRListImagesAll(AwsConfig())
}
