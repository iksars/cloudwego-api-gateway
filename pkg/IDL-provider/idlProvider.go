package idlprovider

import (
	IDL_Management "cloudwego-api-gateway/pkg/IDL-Management"
	"time"
)

type IDLProvider interface {
	FindIDLByServiceName(serviceName string) (idlContent string)
	Update()
}

//todo: IDL路径及内容获取，IDL热更新

type IdlProvider struct {
	serviceName   string
	content       string
	idlManagement IDL_Management.IdlManagement
}

type IDLCache struct {
	IDLRecorder  map[string]string
	TimeRecorder []string
	count        int
	pos          int
}

var IdlCache = IDLCache{
	IDLRecorder:  make(map[string]string, 100),
	TimeRecorder: make([]string, 100),
	count:        0,
	pos:          0,
}

func (idlProvider *IdlProvider) FindIDLByServiceName(serviceName string) (idlContent string) {
	if IdlCache.IDLRecorder[serviceName] == "" {
		if IdlCache.count < 100 {
			IdlCache.IDLRecorder[serviceName] = idlProvider.idlManagement.Search(serviceName).Content
			IdlCache.TimeRecorder[IdlCache.count] = serviceName
			IdlCache.count++
		} else {
			delete(IdlCache.IDLRecorder, IdlCache.TimeRecorder[IdlCache.pos])
			IdlCache.IDLRecorder[serviceName] = idlProvider.idlManagement.Search(serviceName).Content
			IdlCache.pos++
			if IdlCache.pos >= 100 {
				IdlCache.pos = 0
			}
		}
	}
	return IdlCache.IDLRecorder[serviceName]
}

func (idlProvider *IdlProvider) Update() {
	for {
		if idlProvider.content == "" {
			idlProvider.content = idlProvider.idlManagement.Search(idlProvider.serviceName).Content
			// 调用kitexClient的方法更新客户端
		} else {
			newContent := idlProvider.idlManagement.Search(idlProvider.serviceName).Content
			if newContent != idlProvider.content {
				// 调用kitexClient的方法更新客户端

			}
		}
		time.Sleep(time.Second * 10)
	}
}
