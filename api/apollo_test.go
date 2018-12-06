package api

import (
	"testing"
)

func TestGetConfigItems(t *testing.T) {
	config, err := LoadApolloConfig("config.json")
	if err != nil {
		t.Fatalf("loadconfig failed,err:%s", err.Error())
	}
	t.Logf("config:\n%s\n", config)

	data, err := GetConfigItems(config)
	if err != nil {
		t.Fatalf("get config failed,err:%s", err.Error())
	}
	t.Logf("config:%v", data.Items)

	err = data.WriteToFile("/tmp/config.ini")
	if err != nil {
		t.Fatalf("write failed,err:%s", err.Error())
	}
}
