package kitexclientprovider

import (
	idlprovider "cloudwego-api-gateway/pkg/IDL-provider"
	"fmt"

	cli "github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/generic"
	etcd "github.com/kitex-contrib/registry-etcd"

	"sync"
	"time"
)

type KitexClientProvider interface {
	NewGenericClient(serviceName string) (client genericclient.Client)
}

type cache struct {
	data     map[string]*cacheData // 缓存数据
	count    int                   // 缓存数量
	capacity int                   // 缓存容量
	lock     sync.RWMutex          // 读写锁
}

type cacheData struct {
	client   *genericclient.Client
	provider *generic.ThriftContentProvider
	lastUse  time.Time
}

type defaultKitexClientProvider struct {
	register       discovery.Resolver      // 注册中心,服务发现
	idlProvider    idlprovider.IDLProvider // idl提供者
	cache          *cache                  // 缓存
	updateInterval time.Duration           // cache更新间隔
}

func NewDefaultKitexClientProvider() (res *defaultKitexClientProvider) {
	var err error
	res = &defaultKitexClientProvider{}
	res.register, err = etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		panic(err)
	}
	res.cache = &cache{data: make(map[string]*cacheData), capacity: 30}
	res.idlProvider = idlprovider.NewDefaultIdlProvider()
	res.updateInterval = 10 * time.Second
	// 启动一个goroutine,定时刷新client缓存
	go res.dynamicClientCacheRefresh()
	return
}

func (ptr *defaultKitexClientProvider) NewGenericClient(serviceName string) (client genericclient.Client) {
	// 如果缓存中有,直接返回
	if cacheData := ptr.cache.get(serviceName); cacheData.client != nil {
		return *cacheData.client
	}

	// 如果缓存中没有,从idlProvider中获取IDL内容,并且生成client
	idlContent := ptr.idlProvider.FindIDLByServiceName(serviceName)
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

	ptr.cache.set(serviceName, &cacheData{client: &client, provider: p})
	return
}

type chObj struct {
	serviceName string
	idlContent  string
}

func (ptr *defaultKitexClientProvider) dynamicClientCacheRefresh() {
	for {
		time.Sleep(ptr.updateInterval)
		ptr.cache.lock.Lock()
		fmt.Println("dynamicClientCacheRefresh lock")
		var ch chan chObj = make(chan chObj, ptr.cache.count)
		var taskCount int
		for k := range ptr.cache.data {
			go ptr.fetchIDLContent(k, ch)
			taskCount++
		}
		for i := 0; i < taskCount; i++ {
			chObj := <-ch
			if chObj.idlContent != "" {
				ptr.cache.data[chObj.serviceName].provider.UpdateIDL(chObj.idlContent, map[string]string{})
				fmt.Println("update idlContent for serviceName:", chObj.serviceName, "done")
			} else {
				delete(ptr.cache.data, chObj.serviceName)
				ptr.cache.count--
			}
		}
		fmt.Println("dynamicClientCacheRefresh unlock")
		ptr.cache.lock.Unlock()
	}
}

func (ptr *defaultKitexClientProvider) fetchIDLContent(serviceName string, ch chan chObj) {
	idlContent := ptr.idlProvider.FindIDLByServiceName(serviceName)
	ch <- chObj{serviceName: serviceName, idlContent: idlContent}
}

func (d *cache) get(k string) cacheData {

	d.lock.RLock()
	defer d.lock.RUnlock()
	if v, ok := d.data[k]; ok {
		v.lastUse = time.Now()
		return *v
	}
	return cacheData{}
}

func (d *cache) set(k string, v *cacheData) {
	d.lock.Lock()
	defer d.lock.Unlock()
	if d.count > d.capacity {
		replaceKey := findReplaceCache(d.data)
		delete(d.data, replaceKey)
	} else {
		d.count++
	}
	v.lastUse = time.Now()
	d.data[k] = v
}

func findReplaceCache(m map[string]*cacheData) string {
	var minTime time.Time
	var minKey string
	for k, v := range m {
		if minTime.IsZero() || v.lastUse.Before(minTime) {
			minTime = v.lastUse
			minKey = k
		}
	}
	return minKey
}
