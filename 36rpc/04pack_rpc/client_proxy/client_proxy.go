package client_proxy

import (
	"net/rpc"
	"pack_rpc/client/entity"
	"pack_rpc/handler"
)

type RectService struct {
	*rpc.Client
}

func NewRectService(protol string, addr string) *RectService {
	client, err := rpc.Dial(protol, addr)
	if err != nil {
		panic(err)
	}
	return &RectService{Client: client}
}

// RPC服务端方法，求矩形面积
func (r *RectService) Area(p entity.Params, result *int) error {
	err := r.Call(handler.RectServiceName+".Area", p, &result)
	return err
}
