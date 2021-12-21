package registry

import "time"

type Options struct {
	// 地址
	Addres []string

	// 超时时间
	Timeout time.Duration

	//心跳时间
	HeartBeat int64

	//注册地址
	RegistryPath string
}

// 使用选择模式
type Option func(opts *Options)

func WithAddresses(addrs []string) Option {
	return func(opts *Options) {
		opts.Addres = addrs
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(opts *Options) {
		opts.Timeout = timeout
	}
}

func WithHeartBeat(heartbeat int64) Option {
	return func(opts *Options) {
		opts.HeartBeat = heartbeat
	}

}

func WithRegistryPath(registrypath string) Option {
	return func(opts *Options) {
		opts.RegistryPath = registrypath
	}
}
