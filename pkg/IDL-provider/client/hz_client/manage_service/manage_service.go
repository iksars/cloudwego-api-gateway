// Code generated by hertz generator.

package manage_service

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/protocol"
	IDLManage "github.com/iksars/cloudwego-api-gateway/pkg/IDL-provider/client/hertz_gen/IDLManage"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
)

type Client interface {
	SelectAll(context context.Context, req *IDLManage.EmptyReq, reqOpt ...config.RequestOption) (resp *IDLManage.QueryResp, rawResponse *protocol.Response, err error)

	SelectByName(context context.Context, req *IDLManage.NameBasedReq, reqOpt ...config.RequestOption) (resp *IDLManage.QueryResp, rawResponse *protocol.Response, err error)

	AddByName(context context.Context, req *IDLManage.AddReq, reqOpt ...config.RequestOption) (resp *IDLManage.CommonResp, rawResponse *protocol.Response, err error)

	DeleteByName(context context.Context, req *IDLManage.NameBasedReq, reqOpt ...config.RequestOption) (resp *IDLManage.CommonResp, rawResponse *protocol.Response, err error)

	DownloadByName(context context.Context, req *IDLManage.NameBasedReq, reqOpt ...config.RequestOption) (resp *IDLManage.DownloadResp, rawResponse *protocol.Response, err error)
}

type ManageServiceClient struct {
	client *cli
}

func NewManageServiceClient(hostUrl string, ops ...Option) (Client, error) {
	opts := getOptions(append(ops, withHostUrl(hostUrl))...)
	cli, err := newClient(opts)
	if err != nil {
		return nil, err
	}
	return &ManageServiceClient{
		client: cli,
	}, nil
}

func (s *ManageServiceClient) SelectAll(context context.Context, req *IDLManage.EmptyReq, reqOpt ...config.RequestOption) (resp *IDLManage.QueryResp, rawResponse *protocol.Response, err error) {
	httpResp := &IDLManage.QueryResp{}
	ret, err := s.client.r().
		setContext(context).
		setQueryParams(map[string]interface{}{}).
		setPathParams(map[string]string{}).
		setHeaders(map[string]string{}).
		setFormParams(map[string]string{}).
		setFormFileParams(map[string]string{}).
		setBodyParam(req).
		setRequestOption(reqOpt...).
		setResult(httpResp).
		execute("GET", "/api/getAll")
	if err != nil {
		return nil, nil, err
	}

	resp = httpResp
	rawResponse = ret.rawResponse
	return resp, rawResponse, nil
}

func (s *ManageServiceClient) SelectByName(context context.Context, req *IDLManage.NameBasedReq, reqOpt ...config.RequestOption) (resp *IDLManage.QueryResp, rawResponse *protocol.Response, err error) {
	httpResp := &IDLManage.QueryResp{}
	ret, err := s.client.r().
		setContext(context).
		setQueryParams(map[string]interface{}{
			"name": req.GetName(),
		}).
		setPathParams(map[string]string{}).
		setHeaders(map[string]string{}).
		setFormParams(map[string]string{}).
		setFormFileParams(map[string]string{}).
		setBodyParam(req).
		setRequestOption(reqOpt...).
		setResult(httpResp).
		execute("GET", "/api/search")
	if err != nil {
		return nil, nil, err
	}

	resp = httpResp
	rawResponse = ret.rawResponse
	return resp, rawResponse, nil
}

func (s *ManageServiceClient) AddByName(context context.Context, req *IDLManage.AddReq, reqOpt ...config.RequestOption) (resp *IDLManage.CommonResp, rawResponse *protocol.Response, err error) {
	httpResp := &IDLManage.CommonResp{}
	ret, err := s.client.r().
		setContext(context).
		setQueryParams(map[string]interface{}{}).
		setPathParams(map[string]string{}).
		setHeaders(map[string]string{}).
		setFormParams(map[string]string{}).
		setFormFileParams(map[string]string{}).
		setBodyParam(req).
		setRequestOption(reqOpt...).
		setResult(httpResp).
		execute("POST", "/api/add")
	if err != nil {
		return nil, nil, err
	}

	resp = httpResp
	rawResponse = ret.rawResponse
	return resp, rawResponse, nil
}

func (s *ManageServiceClient) DeleteByName(context context.Context, req *IDLManage.NameBasedReq, reqOpt ...config.RequestOption) (resp *IDLManage.CommonResp, rawResponse *protocol.Response, err error) {
	httpResp := &IDLManage.CommonResp{}
	ret, err := s.client.r().
		setContext(context).
		setQueryParams(map[string]interface{}{
			"name": req.GetName(),
		}).
		setPathParams(map[string]string{}).
		setHeaders(map[string]string{}).
		setFormParams(map[string]string{}).
		setFormFileParams(map[string]string{}).
		setBodyParam(req).
		setRequestOption(reqOpt...).
		setResult(httpResp).
		execute("DELETE", "/api/delete")
	if err != nil {
		return nil, nil, err
	}

	resp = httpResp
	rawResponse = ret.rawResponse
	return resp, rawResponse, nil
}

func (s *ManageServiceClient) DownloadByName(context context.Context, req *IDLManage.NameBasedReq, reqOpt ...config.RequestOption) (resp *IDLManage.DownloadResp, rawResponse *protocol.Response, err error) {
	httpResp := &IDLManage.DownloadResp{}
	ret, err := s.client.r().
		setContext(context).
		setQueryParams(map[string]interface{}{
			"name": req.GetName(),
		}).
		setPathParams(map[string]string{}).
		setHeaders(map[string]string{}).
		setFormParams(map[string]string{}).
		setFormFileParams(map[string]string{}).
		setBodyParam(req).
		setRequestOption(reqOpt...).
		setResult(httpResp).
		execute("GET", "/api/download")
	if err != nil {
		return nil, nil, err
	}

	resp = httpResp
	rawResponse = ret.rawResponse
	return resp, rawResponse, nil
}

var defaultClient, _ = NewManageServiceClient("")

func ConfigDefaultClient(ops ...Option) (err error) {
	defaultClient, err = NewManageServiceClient("", ops...)
	return
}

func SelectAll(context context.Context, req *IDLManage.EmptyReq, reqOpt ...config.RequestOption) (resp *IDLManage.QueryResp, rawResponse *protocol.Response, err error) {
	return defaultClient.SelectAll(context, req, reqOpt...)
}

func SelectByName(context context.Context, req *IDLManage.NameBasedReq, reqOpt ...config.RequestOption) (resp *IDLManage.QueryResp, rawResponse *protocol.Response, err error) {
	return defaultClient.SelectByName(context, req, reqOpt...)
}

func AddByName(context context.Context, req *IDLManage.AddReq, reqOpt ...config.RequestOption) (resp *IDLManage.CommonResp, rawResponse *protocol.Response, err error) {
	return defaultClient.AddByName(context, req, reqOpt...)
}

func DeleteByName(context context.Context, req *IDLManage.NameBasedReq, reqOpt ...config.RequestOption) (resp *IDLManage.CommonResp, rawResponse *protocol.Response, err error) {
	return defaultClient.DeleteByName(context, req, reqOpt...)
}

func DownloadByName(context context.Context, req *IDLManage.NameBasedReq, reqOpt ...config.RequestOption) (resp *IDLManage.DownloadResp, rawResponse *protocol.Response, err error) {
	return defaultClient.DownloadByName(context, req, reqOpt...)
}
