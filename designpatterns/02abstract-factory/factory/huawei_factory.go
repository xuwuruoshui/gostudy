package factory

import (
	"designpatterns/02abstract-factory/entity"
	"designpatterns/02abstract-factory/features"
)

type HuaweiFactory struct {
}

func (s *HuaweiFactory) CreatePhoneFactory() features.PhoneFeatures {
	return &entity.Phone{"手机", "华为", "meta40"}
}
func (s *HuaweiFactory) CreateTVFactory() features.TVFeatures {
	return &entity.TV{"电视", "华为", "xxxx"}
}

func NewHuaweiFactory() *HuaweiFactory {
	return &HuaweiFactory{}
}
