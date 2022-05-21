package abstract_factory

import (
	"designpatterns/02abstract-factory/factory"
	"testing"
)

func TestAbstractFactory(t *testing.T) {
	huaweiFactory := factory.NewHuaweiFactory()
	phoneFactory := huaweiFactory.CreatePhoneFactory()
	phoneFactory.Call()
	tvFactory := huaweiFactory.CreateTVFactory()
	tvFactory.Play()

	xiaomiFactory := factory.NewXiaomiFactory()
	phoneFactory1 := xiaomiFactory.CreatePhoneFactory()
	phoneFactory1.Call()
	tvFactory1 := xiaomiFactory.CreateTVFactory()
	tvFactory1.Play()
}
