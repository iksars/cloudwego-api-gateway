package kitexclientprovider

type KitexClientProvider interface {
	NewGenericClient(serviceName string, idlContent string) (client interface{})
}
