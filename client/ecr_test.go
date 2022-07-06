package client

import (
	awsutil "awsdisc/client/util"
	"testing"
)

func TestEcrDiscoverAll(t *testing.T) {
	j, err := ECRDescribeRegistryCmd(AwsConfig())
	if err != nil {
		t.Error(err)
	}
	t.Log(awsutil.PrettyJson(j).String())

	j, err = ECRDescribeRepositoriesCmd(AwsConfig())
	if err != nil {
		t.Error(err)
	}
	t.Log(awsutil.PrettyJson(j).String())

}

func TestECRDescribeRegistryCmd(t *testing.T) {
	j, err := ECRDescribeRegistryCmd(AwsConfig())
	if err != nil {
		t.Error(err)
	}
	t.Log(awsutil.PrettyJson(j).String())
}

func TestECRDescribeRepositoriesCmd(t *testing.T) {
	j, err := ECRDescribeRepositoriesCmd(AwsConfig())
	if err != nil {
		t.Error(err)
	}
	t.Log(awsutil.PrettyJson(j).String())

	repos, err := awsutil.JsonPath(j, "$.Repositories[:].RepositoryName")
	if err != nil {
		t.Error(err)
	}
	t.Logf("%v", repos)
}

func TestECRListImagesAll(t *testing.T) {
	ECRListImagesAll(AwsConfig())
}
