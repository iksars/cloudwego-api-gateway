// Code generated by hertz generator. DO NOT EDIT.

package IDLManage

import (
	IDLManage "github.com/iksars/cloudwego-api-gateway/pkg/IDL-Management/hertz-back/biz/handler/IDLManage"
	"github.com/cloudwego/hertz/pkg/app/server"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_api := root.Group("/api", _apiMw()...)
		_api.POST("/add", append(_addbynameMw(), IDLManage.AddByName)...)
		_api.DELETE("/delete", append(_deletebynameMw(), IDLManage.DeleteByName)...)
		_api.GET("/download", append(_downloadbynameMw(), IDLManage.DownloadByName)...)
		_api.GET("/getAll", append(_selectallMw(), IDLManage.SelectAll)...)
		_api.GET("/search", append(_selectbynameMw(), IDLManage.SelectByName)...)
	}
}
