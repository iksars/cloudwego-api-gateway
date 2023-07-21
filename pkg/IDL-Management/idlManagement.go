package IDL_Management

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type IDLRecorder struct {
	gorm.Model
	ServiceName string
	Content     string
}

type IDLManagement interface {
	IDLManagementInit(name string)
	Insert(serviceName string, content string)
	Delete(serviceName string)
	Update(serviceName string, content string)
	Search(serviceName string) IDLRecorder
}

type IdlManagement struct {
	db  *gorm.DB
	err error
}

func (i *IdlManagement) IDLManagementInit(name string) {
	i.db, i.err = gorm.Open(sqlite.Open("foo.db"), &gorm.Config{})
	if i.err != nil {
		panic(i.err)
	}
	i.err = i.db.AutoMigrate(&IDLRecorder{})
	if i.err != nil {
		panic(i.err)
	}

}

func (i *IdlManagement) Insert(serviceName string, content string) {
	i.db.Create(&IDLRecorder{
		ServiceName: serviceName,
		Content:     content,
	})
}

func (i *IdlManagement) Delete(serviceName string) {
	i.db.Delete(&IDLRecorder{}, "serviceName = ?", serviceName)
}

func (i *IdlManagement) Update(serviceName string, content string) {
	rec := i.Search(serviceName)
	if rec.ServiceName != "" {
		i.db.Model(&rec).Update("content", content)
	}
}

func (i *IdlManagement) Search(serviceName string) IDLRecorder {
	var idlRecorder IDLRecorder
	res := i.db.First(&idlRecorder, "serviceName = ?", serviceName)
	if res.Error != nil {
		log.Println("IDL Not Found!!!")
	}
	return idlRecorder
}
