package facade

import "fmt"

// 外观模式

type Ying interface {
	Reduce()
}

type YingImpl struct {
	
}

func (y *YingImpl)Reduce(){
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
	ying Ying
	yang Yang
}

func NewDao()  *Dao{
	return &Dao{
		&YingImpl{},
		&YangImpl{},
	}
}

func (c *Dao)Run(){
	c.ying.Reduce()
	c.yang.Increase()
}