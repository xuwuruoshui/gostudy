package factory

import (
	"abstract-factory/entity"
	"abstract-factory/features"
)

type XiaomiFactory struct {
}

func (s *XiaomiFactory) CreatePhoneFactory() features.PhoneFeatures {
	return &entity.Phone{"手机", "小米", "12"}
}
func (s *XiaomiFactory) CreateTVFactory() features.TVFeatures {
	return &entity.TV{"电视", "小米", "xxxx"}
}

func NewXiaomiFactory() *XiaomiFactory {
	return &XiaomiFactory{}
}
