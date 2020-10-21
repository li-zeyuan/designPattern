package bridge

import "fmt"

/*
示例：实现SMS/Email发送 普通、紧急消息
两种发送方式：SMS、Email
两种类型消息：普通消息、紧急消息
 */

// 抽象化（Abstraction）
type AbstractMessage interface {
	SendMessage(text, to string)
}

// 实现化（Implementor）
type MessageImplementer interface {
	Send(text, to string)
}

// 具体实现化A（Concrete Implementor）
type MessageSMS struct {

}

func ViaSMS()MessageImplementer  {
	return &MessageSMS{}
}

func (s *MessageSMS)Send(text, to string)  {
	fmt.Printf("send %s to %s via 【SMS】", text, to)
}

// 具体实现化B（Concrete Implementor）
type MessageEmail struct {

}

func ViaEmail()MessageImplementer  {
	return &MessageEmail{}
}

func (e *MessageEmail)Send(text, to string)  {
	fmt.Printf("send %s to %s via 【email】", text, to)
}

// ===================
// 扩张抽象A（Refined Abstraction）
type CommonMessage struct {
	method MessageImplementer
}

func NewCommonMessage(method MessageImplementer) *CommonMessage {
	return &CommonMessage{
		method:method,
	}
}

func (c *CommonMessage)SendMessage(text, to  string) {
	c.method.Send(fmt.Sprintf("【common_msg】: %s", text), to)
}

// 扩张抽象B（Refined Abstraction）
type UrgencyMessage struct {
	method MessageImplementer
}

func NewUrgencyMessage(method MessageImplementer) *UrgencyMessage {
	return &UrgencyMessage{
		method:method,
	}
}

func (u *UrgencyMessage)SendMessage(text, to string)  {
	u.method.Send(fmt.Sprintf("【urgency_msg】: %s", text), to)
}

