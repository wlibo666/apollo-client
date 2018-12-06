# apollo-client
client that get config from apollo server

## usage
```
Usage of ./apollo-client:
  -appid string
    	--appid id, apollo app'id
  -cluster string
    	--cluster clusterName (default "default")
  -configFile string
    	--configFile apollo config json file(ingore other params except output)
  -namespace string
    	--application namespaceName (default "application")
  -outfile string
    	--outfile file that store config items
  -server string
    	--server addr,apollo server addr,eg:127.0.0.1:80
```
eg: ./apollo-client --server=172.16.17.6:8080 --appid=apollo-example --cluster=DEV --namespace=application --outfile config.ini
