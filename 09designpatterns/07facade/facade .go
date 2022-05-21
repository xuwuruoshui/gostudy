package facade

import "fmt"

// 外观模式
// 道由阴阳组成

type Yin interface {
	Reduce()
}

type YinImpl struct {
	
}

func (y *YinImpl)Reduce(){
	fmt.Println("阴")
}



type Yang interface {
	Increase()
}

type YangImpl struct {
	
}

func (y *YangImpl)Increase(){
	fmt.Println("阳")
}

type Dao struct {
	yin Yin
	yang Yang
}

func NewDao()  *Dao{
	return &Dao{
		&YinImpl{},
		&YangImpl{},
	}
}

func (c *Dao)Run(){
	c.yin.Reduce()
	c.yang.Increase()
}