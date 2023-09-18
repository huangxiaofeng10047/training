package main

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"time"
)

func main() {
	//sc := []constant.ServerConfig{
	//
	//	*constant.NewServerConfig("nacos.kubegems.io", 30498, constant.WithContextPath("/nacos")),
	//}
	//
	//cc := constant.ClientConfig{
	//	NamespaceId:         "2cfea28330290809ac7ebc7747fa7a6f3672cba5", //namespace id
	//	TimeoutMs:           5000,
	//	NotLoadCacheAtStart: true,
	//	LogDir:              "/tmp/nacos/log",
	//	CacheDir:            "/tmp/nacos/cache",
	//	LogLevel:            "debug",
	//}
	sc := []constant.ServerConfig{
		*constant.NewServerConfig("nacos.kubegems.io", 30008),
	}

	//create ClientConfig
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId("2cfea28330290809ac7ebc7747fa7a6f3672cba5"),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("/tmp/nacos/log"),
		constant.WithCacheDir("/tmp/nacos/cache"),
		constant.WithLogLevel("debug"),
	)

	// create config client
	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)

	if err != nil {
		panic(err)
	}
	_, err = client.PublishConfig(vo.ConfigParam{
		DataId:  "test-data",
		Group:   "test-group",
		Content: "hello world!",
	})
	_, err = client.PublishConfig(vo.ConfigParam{
		DataId:  "test-data-2",
		Group:   "test-group",
		Content: "hello world!",
	})
	if err != nil {
		fmt.Printf("PublishConfig err:%+v \n", err)
	}
	time.Sleep(1 * time.Second)
	content, err := client.GetConfig(vo.ConfigParam{
		DataId: "test",
		Group:  "nacos",
	})
	fmt.Println("GetConfig,config :" + content)
	if err != nil {
		fmt.Println(err)
	}
	err = client.ListenConfig(vo.ConfigParam{
		DataId: "test",
		Group:  "nacos",
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("config changed group:" + group + ", dataId:" + dataId + ", content:" + data)
		},
	})
	time.Sleep(300 * time.Second)
}
