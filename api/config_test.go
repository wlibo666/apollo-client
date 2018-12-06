package api

import (
	"testing"
)

func TestLoadApolloConfig(t *testing.T) {
	config, err := LoadApolloConfig("config.json")
	if err != nil {
		t.Fatalf("loadconfig failed,err:%s", err.Error())
	}
	t.Logf("config:\n%s\n", config)
}
