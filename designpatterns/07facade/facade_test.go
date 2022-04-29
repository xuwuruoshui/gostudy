package facade

import "testing"

func TestFacade(t *testing.T){
	NewDao().Run()
}