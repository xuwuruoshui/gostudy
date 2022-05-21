package bridge

import "testing"

func TestBridge(t *testing.T){
	message := NewCommonMessage(NewSMS())
	message.SendMessage("Hello World","Herry")

	urgencyMessage := NewUrgencyMessage(NewEmail())
	urgencyMessage.SendMessage("有内鬼终止交易","Jack")
}
