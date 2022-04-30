package adapter

import (
	"fmt"
	"testing"
)

func TestAdapter(t *testing.T){
	adapter := NewAdapter(NewTypeC())
	msg := adapter.Run()
	fmt.Println(msg)

	adapter2 := NewAdapter(NewLightning())
	msg2 := adapter2.Run()
	fmt.Println(msg2)
}
