package api

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/valyala/fasthttp"
	"github.com/wlibo666/common-lib/webutils"
	"github.com/wlibo666/json"
)

type ConfigItems struct {
	AppId     string            `json:"appId"`
	Cluster   string            `json:"cluster"`
	Namespace string            `json:"namespaceName"`
	Items     map[string]string `json:"configurations"`
	Release   string            `json:"releaseKey"`
}

func (items *ConfigItems) ItemsString() string {
	var configItems string
	for key, value := range items.Items {
		line := fmt.Sprintf("%s = %s\n", key, value)
		configItems += line
	}
	return configItems
}

func (items *ConfigItems) WriteToFile(filename string) error {
	fp, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer fp.Close()

	_, err = fp.WriteString(items.ItemsString())
	if err != nil {
		return err
	}
	return nil
}

// 如果指定Apollo配置,则根据指定配置获取信息,否则使用默认配置
func GetConfigItems(apolloConfig ...*ApolloConfig) (*ConfigItems, error) {
	var config *ApolloConfig

	if len(apolloConfig) > 0 {
		config = apolloConfig[0]
	} else {
		config = defaultConfig
	}

	args := make(map[string]string)
	args["t"] = fmt.Sprintf("%d", time.Now().UnixNano())
	requestUrl := fmt.Sprintf("/configs/%s/%s/%s", config.AppId, config.Cluster, config.Namespace)

	req := webutils.NewRequest(config.Server, "GET", requestUrl, args, []byte{}, nil)
	resp := webutils.NewResponse()
	err := fasthttp.DoTimeout(req, resp, 3*time.Second)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("http status code is:%d,content:%s", resp.StatusCode(), resp.Body())
	}
	items := &ConfigItems{
		Items: make(map[string]string),
	}
	err = json.Unmarshal(resp.Body(), items)
	if err != nil {
		return nil, fmt.Errorf("Unmarshal resp body failed,err:%s,body:%s", err.Error(), resp.Body())
	}
	return items, nil
}
