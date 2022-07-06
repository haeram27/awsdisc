package client

import (
	awsutil "awsdisc/client/util"
	"testing"
)

func TestDiscoverAll(t *testing.T) {
	j, err := DiscoverAll()
	if err != nil {
		t.Error(err)
	}
	t.Log(awsutil.PrettyJson(j).String())
}
