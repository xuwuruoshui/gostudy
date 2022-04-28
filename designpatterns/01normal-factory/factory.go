package normal_factory

import "fmt"

type PhoneFactory interface {
	Call()
}

// 华为
type Huawei struct {
	name    string
	version string
}

func (h *Huawei) Call() {
	fmt.Println(h.name, h.version)
}

// 小米
type Xiaomi struct {
	name    string
	version string
}

func (h *Xiaomi) Call() {
	fmt.Println(h.name, h.version)
}

func NewPhoneFactory(name string) PhoneFactory {
	switch name {
	case "h":
		return &Huawei{name: "华为", version: "p50"}
	case "x":
		return &Xiaomi{name: "小米", version: "12"}
	default:
		return &Huawei{name: "华为", version: "p40"}
	}
}
