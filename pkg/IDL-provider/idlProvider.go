package idlprovider

type IDLProvider interface {
	FindIDLByServiceName(serviceName string) (idlContent string)
}
