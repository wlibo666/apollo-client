package api

import (
	"encoding/json"
	"io/ioutil"

	"github.com/wlibo666/common-lib/utils"
	"golangsrc/fmt"
)

const (
	ENV_SERVER    = "APOLLO_SERVER"
	ENV_APPID     = "APOLLO_APPID"
	ENV_CLUSTER   = "APOLLO_CLUSTER"
	ENV_NAMESPACE = "APOLLO_NAMESPACE"
)

type ApolloConfig struct {
	Server    string `json:"server"`
	AppId     string `json:"appId"`
	Cluster   string `json:"cluster"`
	Namespace string `json:"namespace"`
}

var (
	defaultConfig = &ApolloConfig{
		Cluster:   "default",
		Namespace: "application",
	}
)

func InitDefaultApolloConfig(server, appId, cluster, namespace string) {
	defaultConfig.Server = server
	defaultConfig.AppId = appId
	defaultConfig.Cluster = cluster
	defaultConfig.Namespace = namespace
}

func LoadApolloConfig(file string) (*ApolloConfig, error) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	tmpConfig := &ApolloConfig{}
	err = json.Unmarshal(content, tmpConfig)
	if err != nil {
		return nil, err
	}
	return tmpConfig, nil
}

func GetDefaultApolloConfig() *ApolloConfig {
	return defaultConfig
}

func GetApolloConfig() *ApolloConfig {
	if defaultConfig.Server != "" && defaultConfig.AppId != "" &&
		defaultConfig.Cluster != "" && defaultConfig.Namespace != "" {
		return defaultConfig
	}

	envConfig := &ApolloConfig{
		Server:    utils.GetEnvDef(ENV_SERVER, ""),
		AppId:     utils.GetEnvDef(ENV_APPID, ""),
		Cluster:   utils.GetEnvDef(ENV_CLUSTER, ""),
		Namespace: utils.GetEnvDef(ENV_NAMESPACE, ""),
	}
	if envConfig.Cluster == "" {
		envConfig.Cluster = "default"
	}
	if envConfig.Namespace == "" {
		envConfig.Namespace = "application"
	}
	return envConfig
}

func (config *ApolloConfig) String() string {
	data, err := json.Marshal(config)
	if err != nil {
		return err.Error()
	}
	return string(data)
}

func (config *ApolloConfig) Check() error {
	if config.Server == "" {
		return fmt.Errorf("lost param: server")
	}
	if config.AppId == "" {
		return fmt.Errorf("lost param: appid")
	}
	if config.Cluster == "" {
		return fmt.Errorf("lost param: cluster")
	}
	if config.Namespace == "" {
		return fmt.Errorf("lost param: namespace")
	}
	return nil
}
