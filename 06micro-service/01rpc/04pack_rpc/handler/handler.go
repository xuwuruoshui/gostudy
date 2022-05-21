package handler

import "pack_rpc/server/entity"

// 具体业务

const RectServiceName = "handler/RectService"

type RectService struct{}

func (r *RectService) Area(p entity.Params, result *int) error {
	*result = p.Height * p.Width
	return nil
}
