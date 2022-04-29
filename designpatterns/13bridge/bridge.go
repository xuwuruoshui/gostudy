package bridge

import "fmt"

// 桥接模式
// 有个jb用

type AbstractMessage interface {
	SendMessage(text, to string)
}

// 普通信息
type CommonMessage struct {
	messageImplement MessageImplement
}

func NewCommonMessage(method MessageImplement) AbstractMessage{
	return &CommonMessage{
		method,
	}
}

func  (c *CommonMessage) SendMessage(text, to string){
	c.messageImplement.Send(text,to)
}


// 紧急信息
type UrgencyMessage struct {
	messageImplement MessageImplement
}

func NewUrgencyMessage(method MessageImplement) AbstractMessage{
	return &UrgencyMessage{
		method,
	}
}

func  (u *UrgencyMessage) SendMessage(text, to string){
	u.messageImplement.Send("[UrgencyMessage]:"+text,to)
}


type MessageImplement interface {
	Send(text,to string)
}


// SMS发送
type MessageSMS struct {

}

func NewSMS()MessageImplement{
	return &MessageSMS{}
}

func (m *MessageSMS)Send(text,to string){
	fmt.Println("send sms:",fmt.Sprintf("[%s]",text),"to",to)
}


// Email发送
type MessageEmail struct {
	
}

func NewEmail()MessageImplement{
	return &MessageEmail{}
}

func (m *MessageEmail)Send(text,to string){
	fmt.Println("send email:",fmt.Sprintf("[%s]",text),"to",to)
}