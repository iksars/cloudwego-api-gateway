// Code generated by hertz generator.

package IDLManageService

import (
	IDLManage "cloudwego-api-gateway/pkg/IDL-Management/hertz-back/biz/model/IDLManage"
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"gorm.io/gorm"
)

var db *gorm.DB = InitDB()

// SelectAll .
// @router /api/getAll [GET]
func SelectAll(ctx context.Context, c *app.RequestContext) {
	var err error
	var req IDLManage.EmptyReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	var recoreds []IDLRecorder

	// Get all records
	result := db.Find(&recoreds)
	// SELECT * FROM users;
	if result.Error != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(IDLManage.QueryResp)

	resp.Ls = make([]*IDLManage.IDLEntity, len(recoreds))
	for i, v := range recoreds {
		resp.Ls[i] = RecorderToEntity(&v)
	}

	c.JSON(consts.StatusOK, resp)
}

// SelectByName .
// @router /api/search [GET]
func SelectByName(ctx context.Context, c *app.RequestContext) {
	var err error
	var req IDLManage.NameBasedReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	var recorder IDLRecorder

	result := db.Where("service_name = ?", req.Name).First(&recorder)
	if result.Error != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(IDLManage.QueryResp)

	resp.Ls = append(resp.Ls, RecorderToEntity(&recorder))

	c.JSON(consts.StatusOK, resp)
}

// AddByName .
// @router /api/add [POST]
func AddByName(ctx context.Context, c *app.RequestContext) {
	var err error
	var req IDLManage.AddReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	avatarFile := form.File["file"][0]

	var recorder IDLRecorder
	recorder.ServiceName = strings.Split(avatarFile.Filename, ".")[0]
	recorder.Description = "待定"

	//如果有重名的，返回错误
	var count int64
	db.Model(&IDLRecorder{}).Where("service_name = ?", recorder.ServiceName).Count(&count)
	if count > 0 {
		c.String(415, "Service name already exists")
		return
	}

	file, err := avatarFile.Open()
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	defer file.Close()

	fmt.Println(avatarFile.Size)

	var buffer []byte = make([]byte, 1024)
	var content []byte = make([]byte, 0, avatarFile.Size)
	var index int = 0

	for {
		n, err := file.ReadAt(buffer, int64(index))
		content = append(content, buffer[:n]...)
		index += n

		if err != nil {
			if err == io.EOF {
				break
			}
			c.String(consts.StatusBadRequest, err.Error())
			return
		}
	}

	recorder.Content = string(content)

	result := db.Create(&recorder)
	if result.Error != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(IDLManage.CommonResp)

	resp.Message = "Add success"

	c.JSON(consts.StatusOK, resp)
}

// DeleteByName .
// @router /api/delete [DELETE]
func DeleteByName(ctx context.Context, c *app.RequestContext) {
	var err error
	var req IDLManage.NameBasedReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	fmt.Println(req.Name)

	result := db.Where("service_name = ?", req.Name).Delete(&IDLRecorder{})
	if result.Error != nil {
		c.JSON(consts.StatusBadRequest, err.Error())
		return
	}
	fmt.Println(result.RowsAffected)

	resp := new(IDLManage.CommonResp)
	resp.Message = "Delete success"

	c.JSON(consts.StatusOK, resp)
}

// DownloadByName .
// @router /api/download [GET]
func DownloadByName(ctx context.Context, c *app.RequestContext) {
	fmt.Println("download query from ", c.RemoteAddr())
	var err error
	var req IDLManage.NameBasedReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	var recorder IDLRecorder

	result := db.Where("service_name = ?", req.Name).First(&recorder)
	if result.Error != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	c.Header("content-disposition", "attachment; filename="+recorder.ServiceName+".thrift")
	c.Header("content-type", "text/plain")

	c.String(consts.StatusOK, recorder.Content)
}
