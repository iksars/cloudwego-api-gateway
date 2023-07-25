package kitexclientprovider

import (
	cli "github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/generic"
	etcd "github.com/kitex-contrib/registry-etcd"
	"time"
)

type KitexClientProvider interface {
	NewGenericClient(serviceName string, idlContent string) (client genericclient.Client)
	RegularUpdate()
	TriggerUpdate(serviceName string, idlContent string)
}

type DefaultKitexClientProvider struct {
	register     discovery.Resolver              // 注册中心,服务发现
	cache        map[string]genericclient.Client // 缓存
	providers    map[string]*generic.ThriftContentProvider
	timeRecorder map[string]time.Time
}

func NewDefaultKitexClientProvider() (res *DefaultKitexClientProvider) {
	var err error
	res = &DefaultKitexClientProvider{}
	res.cache = make(map[string]genericclient.Client)
	res.register, err = etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		panic(err)
	}
	// 启动一个goroutine,定时刷新client缓存
	return
}

func (ptr *DefaultKitexClientProvider) NewGenericClient(serviceName string, idlContent string) (client genericclient.Client) {
	var ok bool
	if client, ok = ptr.cache[serviceName]; ok {
		return
	}

	p, err := generic.NewThriftContentProvider(idlContent, map[string]string{})
	if err != nil {
		panic(err)
	}

	g, err := generic.JSONThriftGeneric(p)
	if err != nil {
		panic(err)
	}
	client, err = genericclient.NewClient(serviceName, g, cli.WithResolver(ptr.register))
	if err != nil {
		panic(err)
	}
	ptr.cache[serviceName] = client
	ptr.providers[serviceName] = p
	return
}

func (ptr *DefaultKitexClientProvider) RegularUpdate() {
	// 从注册中心获取所有服务
	// 与缓存中的服务列表比较
	// 有新服务则创建client,并加入缓存
	// 有服务下线则从缓存中删除
	deleteList := make([]string, len(ptr.cache))
	deleteLen := 1

	for serviceName, _ := range ptr.cache {
		currentTime := time.Now()
		if currentTime.Sub(ptr.timeRecorder[serviceName]) >= time.Hour*24 {
			deleteList[deleteLen] = serviceName
			deleteLen++
		}
	}
	for _, serviceName := range deleteList {
		delete(ptr.cache, serviceName)
		delete(ptr.providers, serviceName)
		delete(ptr.timeRecorder, serviceName)
	}
}

func (ptr *DefaultKitexClientProvider) TriggerUpdate(serviceName string, idlContent string) {
	if _, ok := ptr.providers[serviceName]; !ok {
		panic("Service Unfounded Error")
	}
	err := ptr.providers[serviceName].UpdateIDL(idlContent, map[string]string{})
	if err != nil {
		panic(err)
	}
	return
}
