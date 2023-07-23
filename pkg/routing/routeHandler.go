package routing

import (
	IDL_Provider "cloudwego-api-gateway/pkg/IDL-provider"
	KClient_Provider "cloudwego-api-gateway/pkg/kitex-client-provider"
	"github.com/cloudwego/kitex/client/genericclient"
)

type RouteHandler interface {
	RoutingDistribute(serviceName string, serviceMethod string) (client genericclient.Client)
}

//实现这个接口
func RoutingDistribute(serviceName string, serviceMethod string) (client genericclient.Client) {
	idlContent := IDL_Provider.IdlProvider.FindIDLByServiceName(IDL_Provider.IdlProvider{}, serviceName)
	cli := KClient_Provider.KitexClientProvider.NewGenericClient(&KClient_Provider.DefaultKitexClientProvider{}, serviceName, idlContent)

	return cli
}
