package registry

import "context"

type Reigistry interface {

	// 注册中心的名字
	Name() string

	// 初始化
	Init(ctx context.Context, option ...Option) (err error)

	// 服务注册
	Register(ctx context.Context, service *Service) (err error)

	// 取消注册
	UnRegister(ctx context.Context, service *Service) (err error)

	// 服务发现
	GetService(ctx context.Context,name string) (service *Service,err error)
}
