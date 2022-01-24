package server_proxy

import (
	"net/rpc"
	"pack_rpc/handler"
	"pack_rpc/server/entity"
)

type RectServicer interface {
	Area(p entity.Params, result *int) error
}

func RegisterRectService(rectService RectServicer) error {
	err := rpc.RegisterName(handler.RectServiceName, rectService)
	return err
}
