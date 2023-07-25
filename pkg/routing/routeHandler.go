package routing

import (
	IDL_Provider "cloudwego-api-gateway/pkg/IDL-provider"
	KClient_Provider "cloudwego-api-gateway/pkg/kitex-client-provider"
	"time"

	"github.com/cloudwego/kitex/client/genericclient"
)

type RouteHandler interface {
	RoutingDistribute(serviceName string) (client genericclient.Client)
}

type DefaultRouteHandler struct {
	cliProvider KClient_Provider.KitexClientProvider
	iProvider   IDL_Provider.IDLProvider
}

func NewDefaultRouteHandler() (res *DefaultRouteHandler) {
	res = &DefaultRouteHandler{}
	res.iProvider = IDL_Provider.NewDefaultIdlProvider()
	res.cliProvider = KClient_Provider.NewDefaultKitexClientProvider()

	// start a new routine to update idl
	go func() {
		for {
			shouldUpdate := res.iProvider.Update()
			res.cliProvider.RegularUpdate()
			if shouldUpdate {
				res.cliProvider.TriggerUpdate(res.iProvider.GetInfo())
			}
			time.Sleep(time.Second * 10)
		}
	}()

	return
}

// 实现这个接口
func (ptr *DefaultRouteHandler) RoutingDistribute(serviceName string) (client genericclient.Client) {
	idlContent := ptr.iProvider.FindIDLByServiceName(serviceName)
	client = ptr.cliProvider.NewGenericClient(serviceName, idlContent)
	return
}
