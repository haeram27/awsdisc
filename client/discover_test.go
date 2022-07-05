package client

import (
	"testing"
)

func TestDiscoverAll(t *testing.T) {
	j, err := DiscoverAll()
	if err != nil {
		t.Errorf("DiscoverAll() failed ")
	}
	t.Log(PrettyJson(j).String())
}
