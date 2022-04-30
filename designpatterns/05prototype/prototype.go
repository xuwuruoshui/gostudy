package prototype

// 原型模式
// 简单来说就是,提供一个接口方法,复制一份地址不同,值相同的对象
// 原型管理使用一个map维护

type Cloneable interface {
	Clone()Cloneable
}

// 原型管理器
// 用一个map，key value形式来管理
type CloneableManger struct {
	prototypes map[string]Cloneable
}

func NewPrototypeManager() *CloneableManger{
	return &CloneableManger{
		prototypes: make(map[string]Cloneable),
	}
}

func (p *CloneableManger)Get(name string)Cloneable{
	return p.prototypes[name]
}

func (p *CloneableManger)Set(name string,modle Cloneable){
	p.prototypes[name] = modle
}