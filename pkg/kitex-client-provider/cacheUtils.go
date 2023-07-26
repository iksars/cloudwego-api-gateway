package kitexclientprovider

import (
	"container/list"
	"sync"

	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
)

// LRUcache能以o(1)的复杂度执行get() & put()
type LRUcache struct {
	data     map[string]*list.Element // hashMap, key是serviceName, value是ls的*Element, *Element.Value是*cacheData
	ls       *list.List               // 双向链表,用于记录缓存的使用情况,ls.Front()是最近使用的,ls.Back()是最久未使用的
	capacity int                      // 缓存容量
	lock     sync.RWMutex             // 读写锁
}

type cacheData struct {
	serviceName string
	client      *genericclient.Client
	provider    *generic.ThriftContentProvider
}

func newLRUcache() *LRUcache {
	return &LRUcache{
		data:     make(map[string]*list.Element),
		ls:       list.New().Init(),
		capacity: 100,
	}
}

func (d *LRUcache) get(k string) cacheData {
	d.lock.RLock()
	defer d.lock.RUnlock()
	if v, ok := d.data[k]; ok {
		if ptr, ok := v.Value.(*cacheData); ok {
			d.ls.MoveToFront(v)
			return *ptr
		}
	}
	return cacheData{}
}

func (d *LRUcache) put(k string, v *cacheData) {
	d.lock.Lock()
	defer d.lock.Unlock()
	if d.ls.Len() > d.capacity {
		e := d.ls.Remove(d.ls.Back())
		ptr, _ := e.(*cacheData)
		delete(d.data, ptr.serviceName)
	}
	e := d.ls.PushFront(v)
	d.data[k] = e
}

func (d *LRUcache) delete(k string) {
	e := d.data[k]
	d.ls.Remove(e)
	delete(d.data, k)
}
