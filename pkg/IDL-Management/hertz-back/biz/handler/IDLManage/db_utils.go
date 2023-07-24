package IDLManageService

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type IDLRecorder struct {
	gorm.Model
	ServiceName string
	Description string
	Content     string
}

func InitDB() (db *gorm.DB) {
	db, err := gorm.Open(sqlite.Open("foo.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// drop table
	db.Migrator().DropTable(IDLRecorder{})
	// create table
	err = db.Migrator().CreateTable(IDLRecorder{})
	if err != nil {
		panic(err)
	}
	return
}
