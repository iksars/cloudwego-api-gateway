package routing

import (
	IDL_Provider "cloudwego-api-gateway/pkg/IDL-provider"
	KClient_Provider "cloudwego-api-gateway/pkg/kitex-client-provider"

	"github.com/cloudwego/kitex/client/genericclient"
)

type RouteHandler interface {
	RoutingDistribute(serviceName string) (client genericclient.Client)
}

type DefaultRouteHandler struct {
	cliProvider KClient_Provider.KitexClientProvider
	iProvider   IDL_Provider.IdlProvider
}

func NewDefaultRouteHandler() (res *DefaultRouteHandler) {
	res = &DefaultRouteHandler{}
	res.iProvider = IDL_Provider.NewDefaultIdlProvider()
	res.cliProvider = KClient_Provider.NewDefaultKitexClientProvider()
	return
}

// 实现这个接口
func (ptr *DefaultRouteHandler) RoutingDistribute(serviceName string) (client genericclient.Client) {
	idlContent := ptr.iProvider.FindIDLByServiceName(serviceName)
	client = ptr.cliProvider.NewGenericClient(serviceName, idlContent)
	return
}
