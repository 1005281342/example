package main

import (
	"log"
	"time"

	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/env/config"
	"github.com/apolloconfig/agollo/v4/storage"
)

func main() {
	Init()

	var conf = ApolloConfig()
	for i := 0; i < 100; i++ {
		log.Println(conf.GetStringValue("key", "xxx"))
		time.Sleep(time.Second)
	}
}

var (
	apolloClient agollo.Client
	namespace    = "application"
)

func ApolloConfig() *storage.Config {
	return apolloClient.GetConfig(namespace)
}

func Init() {
	c := &config.AppConfig{
		AppID: "bello",
		//Cluster:        "pro",
		Cluster:        "default",
		IP:             "http://127.0.0.1:9080",
		NamespaceName:  namespace,
		IsBackupConfig: true,
		Secret:         "",
	}

	var err error
	apolloClient, err = agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return c, nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
