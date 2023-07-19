// hertz.thrift
namespace go api.layer

struct LayReq {
    1: required string ServiceName (api.path="name");
    2: required string ServiceMethod (api.path="method");
}

struct LayResp {
    1: string RespBody;
}


service LayService {
    LayResp GateWayMethod(1: LayReq request) (api.post="/agw/:name/:method");
}
