package registry

import (
	"context"
	"fmt"
	"sync"
)

// 管理者
type PluginMgr struct {
	// 插件选择
	plugins map[string]Reigistry
	lock sync.Mutex
}

var (
	pluginMgr = &PluginMgr{
		plugins: make(map[string]Reigistry),
	}

)


// 1.注册插件
func (p *PluginMgr)RegisterPlugin(plugin Reigistry)(err error){
	p.lock.Lock()
	defer p.lock.Unlock()
	// 先去看里面有没有
	_,ok := p.plugins[plugin.Name()]
	if ok{
		err = fmt.Errorf("registry plugin exist")
		return
	}
	p.plugins[plugin.Name()] = plugin
	return
}

// 2.注册中心初始化
func (p *PluginMgr)InitRegistry(ctx context.Context,name string,opts ...Option)(registry Reigistry,err error){
	p.lock.Lock()
	defer p.lock.Unlock()
	plugin,ok := p.plugins[name]
	if !ok{
		err = fmt.Errorf("plugin:%s not exist",name)
		return
	}

	registry = plugin
	err = plugin.Init(ctx,opts...)
	return
}