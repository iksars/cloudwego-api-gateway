package routing

type RouteHandler interface {
	RoutingDistribute(serviceName string, serviceMethod string, reqBody string) (respBody string)
}

//实现这个接口
