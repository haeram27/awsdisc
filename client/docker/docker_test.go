package dkr

import (
	aws "awsdisc/client"
	"testing"
)

func TestPullImage(t *testing.T) {
	tok, err := aws.ECRGetAuthorizationTokenCmd(aws.AwsConfig())
	if err != nil {
		t.Fatal(err)
	}

	// decoded, err := base64.StdEncoding.DecodeString(tok)
	// if err != nil {
	// 	return "", err
	// }

	PullImage("797216966998.dkr.ecr.ap-northeast-2.amazonaws.com", "797216966998.dkr.ecr.ap-northeast-2.amazonaws.com/cicd-sec/alpine@sha256:4ff3ca91275773af45cb4b0834e12b7eb47d1c18f770a0b151381cd227f4c253", tok)
}
