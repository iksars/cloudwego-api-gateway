package idlprovider

import (
	IDL_Management "cloudwego-api-gateway/pkg/IDL-Management"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type IDLProvider interface {
	FindIDLByServiceName(serviceName string) (idlContent string)
	insertIDLIntoCache(serviceName string, content string)
	Update() bool
	GetInfo() (string, string)
}

//todo: IDL路径及内容获取，IDL热更新

type IdlProvider struct {
	IdlCache      *IDLCache
	IdlManagement IDL_Management.IdlManagement
}

type IDLCache struct {
	serviceName  string
	content      string
	IDLRecorder  map[string]string
	TimeRecorder []string
	count        int
	pos          int

	db *gorm.DB
}

type IDLRecorder struct {
	gorm.Model
	serviceName string
	content     string
}

func NewDefaultIdlProvider() (res *IdlProvider) {
	res = &IdlProvider{
		IdlCache: &IDLCache{
			serviceName:  "",
			content:      "",
			IDLRecorder:  make(map[string]string, 100),
			TimeRecorder: make([]string, 100),
			count:        0,
			pos:          0,
			db:           nil,
		},
	}
	var err error
	res.IdlCache.db, err = gorm.Open(sqlite.Open("foo.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = res.IdlCache.db.AutoMigrate(&IDLRecorder{})
	return
}

func (idlProvider *IdlProvider) FindIDLByServiceName(serviceName string) (idlContent string) {
	idlProvider.IdlCache.serviceName = serviceName

	var idlRecorder IDLRecorder
	res := idlProvider.IdlCache.db.First(&idlRecorder, "serviceName = ?", idlProvider.IdlCache.serviceName)
	if res.Error != nil {
		log.Println("IDL Not Found!!!")
	}

	idlProvider.insertIDLIntoCache(serviceName, idlRecorder.content)
	return idlProvider.IdlCache.IDLRecorder[serviceName]
}

func (idlProvider *IdlProvider) Update() bool {

	if idlProvider.IdlCache.serviceName == "" {
		return false
	}

	var idlRecorder IDLRecorder
	res := idlProvider.IdlCache.db.First(&idlRecorder, "serviceName = ?", idlProvider.IdlCache.serviceName)
	if res.Error != nil {
		log.Println("IDL Not Found!!!")
	}

	if idlRecorder.content != idlProvider.IdlCache.content {
		idlProvider.insertIDLIntoCache(idlRecorder.serviceName, idlRecorder.content)
		// kitexClient更新客户端
		return true
	}
	return false
}

func (idlProvider *IdlProvider) insertIDLIntoCache(serviceName string, content string) {
	if idlProvider.IdlCache.IDLRecorder[serviceName] == "" {
		if idlProvider.IdlCache.count < 100 {
			idlProvider.IdlCache.TimeRecorder[idlProvider.IdlCache.count] = serviceName
			idlProvider.IdlCache.count++
		} else {
			delete(idlProvider.IdlCache.IDLRecorder, idlProvider.IdlCache.TimeRecorder[idlProvider.IdlCache.pos])
			idlProvider.IdlCache.TimeRecorder[idlProvider.IdlCache.pos] = serviceName
			idlProvider.IdlCache.pos++
			if idlProvider.IdlCache.pos >= 100 {
				idlProvider.IdlCache.pos = 0
			}
		}
	}
	idlProvider.IdlCache.IDLRecorder[serviceName] = content
	idlProvider.IdlCache.content = idlProvider.IdlCache.IDLRecorder[serviceName]
	return
}

func (idlProvider *IdlProvider) GetInfo() (string, string) {
	return idlProvider.IdlCache.serviceName, idlProvider.IdlCache.content
}
