package adapter

// 适配器模式
// ChargerAdapter接口实现的变量中包含一个ChargerINF接口，ChargerINF接口有众多实现。调用者依然使用原来的Charger

// 充电器
type Charger interface {
	Run()string
}

type ChargerAdapter struct {
	ChargerINF
}

// 具体的适配
func (a *ChargerAdapter)Run()string{
	return a.Msg()+"充电"
}

func NewAdapter(inf ChargerINF)Charger{
	return &ChargerAdapter{
		ChargerINF:inf,
	}
}

// 充电的接口
type ChargerINF interface {
	Msg() string
}

// TypeC接口
type  TypeC struct {
	
}

func NewTypeC()ChargerINF{
	return &TypeC{}
}

func (a *TypeC)Msg() string{
	return "TypC接口:"
}

// 苹果的Lightning
type Lightning struct {
	
}
func NewLightning()ChargerINF{
	return &Lightning{}
}

func (a *Lightning)Msg() string{
	return "Lightning接口:"
}

