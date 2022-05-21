package abstract_factory

import "designpatterns/02abstract-factory/features"

// 抽象工厂模式
// 多级工厂，总工厂生产子工厂，子工厂生产对应的对象

type MachineFatory interface {
	CreatePhoneFactory() features.PhoneFeatures
	CreateTVFactory() features.TVFeatures
}
