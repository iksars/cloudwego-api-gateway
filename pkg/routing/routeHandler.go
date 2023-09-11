package routing

import (
	KClient_Provider "github.com/iksars/cloudwego-api-gateway/pkg/kitex-client-provider"

	"github.com/cloudwego/kitex/client/genericclient"
)

type RouteHandler interface {
	RoutingDistribute(serviceName string) (client genericclient.Client)
}

type defaultRouteHandler struct {
	cliProvider KClient_Provider.KitexClientProvider
}

func NewDefaultRouteHandler() (res *defaultRouteHandler) {
	res = &defaultRouteHandler{}
	res.cliProvider = KClient_Provider.NewDefaultKitexClientProvider()
	return
}

func (ptr *defaultRouteHandler) RoutingDistribute(serviceName string) (client genericclient.Client) {
	client = ptr.cliProvider.NewGenericClient(serviceName)
	return
}
