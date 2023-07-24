//manage.thrift
namespace go IDLManage

struct IDLEntity {
    1: string Date;
    2: string Name;
    3: string Description;
}

struct EmptyReq {
    
}

struct NameBasedReq {
    1: string Name(api.query="name");
}

struct QueryResp {
    1: list<IDLEntity> Ls;
}

struct AddReq {
    1: string Name;
    2: string Description;
    3: binary Data;
}

struct CommonResp {
    1: string Message;
}


struct DownloadResp {
    1: binary Data;
}


service ManageService {
    QueryResp SelectAll(1: EmptyReq req) (api.get="/api/getAll");

    QueryResp SelectByName(1: NameBasedReq req) (api.get="/api/search");

    CommonResp AddByName(1: AddReq req) (api.post="/api/add");

    CommonResp DeleteByName(1: NameBasedReq req) (api.delete="/api/delete");

    DownloadResp DownloadByName(1: NameBasedReq req) (api.get="/api/download");
}
