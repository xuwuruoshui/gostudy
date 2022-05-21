package builder

import (
	"fmt"
	"testing"
)

func TestBuilder(t *testing.T){
	builder := &BenchiBuilder{}
	director := NewDirector(builder)
	director.Construct()
	fmt.Println(builder.GetResult())

	builder2 := &BaomaBuilder{}
	director2 := NewDirector(builder2)
	director2.Construct()
	fmt.Println(builder2.GetResult())
}
