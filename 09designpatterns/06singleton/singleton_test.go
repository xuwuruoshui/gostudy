package singleton

import (
	"fmt"
	"testing"
)

func TestSingleton(t *testing.T){
	instance1 := GetInstance()
	instance2 := GetInstance()
	fmt.Println(instance1)
	fmt.Println(instance1)
	fmt.Println(instance1==instance2)
	
}
