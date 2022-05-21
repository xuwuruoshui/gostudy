package normal_factory

import "testing"

func TestFactory(t *testing.T) {
	phone := NewPhoneFactory("x")
	phone.Call()
}
