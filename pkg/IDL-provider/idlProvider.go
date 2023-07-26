package idlprovider

import (
	"context"
	"fmt"

	idlmanage "cloudwego-api-gateway/pkg/IDL-provider/client/hertz_gen/IDLManage"
	"cloudwego-api-gateway/pkg/IDL-provider/client/hz_client/manage_service"
)

type IDLProvider interface {
	FindIDLByServiceName(serviceName string) (idlContent string)
}

type defaultIdlProvider struct {
	idlManagement manage_service.Client //与IDL管理服务通信的客户端
}

func NewDefaultIdlProvider() (res *defaultIdlProvider) {
	var err error
	res = &defaultIdlProvider{}
	res.idlManagement, err = manage_service.NewManageServiceClient("http://127.0.0.1:7210")
	if err != nil {
		panic(err)
	}
	return
}

func (ptr *defaultIdlProvider) FindIDLByServiceName(serviceName string) (idlContent string) {
	fmt.Println("FindIDLByServiceName called with serviceName:", serviceName)
	req := idlmanage.NameBasedReq{Name: serviceName}
	_, resp, _ := ptr.idlManagement.DownloadByName(context.Background(), &req)
	if resp.StatusCode() != 200 { // 如果IDL管理服务返回状态码不是200,则返回空字符串
		return
	}
	idlContent = string(resp.Body())
	return
}
