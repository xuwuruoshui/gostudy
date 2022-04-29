package builder

import "fmt"

// 构建者模式
// 用户无需关系内部实现，只需要创建BenchiBuilder、BaomaBuilder即可
// 用户手动实现Builder也无需关心内部的组装流程，只关注单独的每一个部件即可

type Builder interface {
	// 组装发动机
	AssembleEngine()
	// 组装车身
	AssembleBody()
	// 组装轮胎
	AssembleTire()
}

type Director struct {
	builder Builder
}

func NewDirector(builder Builder)*Director{
	return &Director{
		builder: builder,
	}
}

// 组装流程
func (d *Director)Construct(){
	d.builder.AssembleEngine()
	d.builder.AssembleBody()
	d.builder.AssembleTire()
	fmt.Println("-----------汽车组装完成----------")
}

// Builder的实现1
type BenchiBuilder struct {
	result string
}

func (b *BenchiBuilder)AssembleEngine(){
	b.result+="1"
}

func (b *BenchiBuilder)AssembleBody(){
	b.result+="2"
}

func (b *BenchiBuilder)AssembleTire(){
	b.result+="3"
}

func (b *BenchiBuilder)GetResult() string{
	return b.result
}

// Builder的实现2
type BaomaBuilder struct {
	result int
}

func (b *BaomaBuilder)AssembleEngine(){
	b.result+=1
}

func (b *BaomaBuilder)AssembleBody(){
	b.result+=2
}

func (b *BaomaBuilder)AssembleTire(){
	b.result+=3
}

func (b *BaomaBuilder)GetResult() int{
	return b.result
}