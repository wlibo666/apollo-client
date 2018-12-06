package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/wlibo666/apollo-client/api"
)

var (
	configFile = flag.String("configFile", "", "--configFile apollo config json file(ingore other params except output)")
	server     = flag.String("server", "", "--server addr,apollo server addr,eg:127.0.0.1:80")
	appId      = flag.String("appid", "", "--appid id, apollo app'id")
	cluster    = flag.String("cluster", "default", "--cluster clusterName")
	namespace  = flag.String("namespace", "application", "--application namespaceName")
	output     = flag.String("outfile", "", "--outfile file that store config items")
)

// 如果指定Apollo配置文件,则从配置文件加载配置,否则从指定参数内获取
func getApolloConfig() (*api.ApolloConfig, error) {
	if *configFile != "" {
		tmpConfig, err := api.LoadApolloConfig(*configFile)
		if err != nil {
			return nil, err
		}
		return tmpConfig, nil
	}
	api.InitDefaultApolloConfig(*server, *appId, *cluster, *namespace)
	return api.GetDefaultApolloConfig(), nil
}

// ./apollo-client --server=172.16.17.6:8080 --appid=apollo-example --cluster=DEV --namespace=application --outfile test.ini
func main() {
	flag.Parse()

	config, err := getApolloConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "getApolloConfig failed,err:%s\n", err.Error())
		os.Exit(1)
	}
	err = config.Check()
	if err != nil {
		fmt.Fprintf(os.Stderr, "config check failed,err:%s\n", err.Error())
		os.Exit(2)
	}
	// 获取所有配置项
	configItems, err := api.GetConfigItems(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "GetConfigItems failed,err:%s\n", err.Error())
		os.Exit(3)
	}
	// 若没有指定输出文件,则打印到标准输出
	if *output == "" {
		fmt.Fprintf(os.Stdout, "%s", configItems.ItemsString())
		return
	}
	// 将配置输出到指定文件
	err = configItems.WriteToFile(*output)
	if err != nil {
		fmt.Fprintf(os.Stderr, "WriteToFile failed,err:%s\n", err.Error())
	}
}
