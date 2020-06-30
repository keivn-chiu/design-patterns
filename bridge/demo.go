package bridge

import "fmt"

// bridge the abstract part is separated from its implementation part so that they can be changed independently.

// implementor part
// ImplMsg the interface of implementor part, for the real msg type
type ImplMsg interface {
	Send(to, msg string)
}

// MsgMSM the implementation of ImplMsg
type MsgMSM struct {
}

func (i *MsgMSM) Send(to, msg string) {
	fmt.Printf("msm[ to %s : %s]\n", to, msg)
}
func NewMsgMSM() *MsgMSM {
	return &MsgMSM{}
}

// MsgEmail the implementation of ImplMsg
type MsgEmail struct {
}

func (i *MsgEmail) Send(to, msg string) {
	fmt.Printf("email[ to %s : %s]\n", to, msg)
}
func NewMsgEmail() *MsgEmail {
	return &MsgEmail{}
}

// abstract part
// AbsMsg the interface of abstract part
type AbsMsg interface {
	SendMsg(to, msg string)
}

// CommonMsg the implementation of AbsMsg
type CommonMsg struct {
	from string
	ImplMsg
}

func (m *CommonMsg) SendMsg(to, msg string) {
	fmt.Printf("common msg from: %s\n", m.from)
	m.Send(to, msg)
}
func NewCommonMsg(from string, method ImplMsg) *CommonMsg {
	return &CommonMsg{from, method}
}

// UrgentMsg the implementation of AbsMsg
type UrgentMsg struct {
	from string
	ImplMsg
}

func (m *UrgentMsg) SendMsg(to, msg string) {
	fmt.Printf("urgent msg from: %s\n", m.from)
	m.Send(to, msg)
}
func NewUrgentMsg(from string, method ImplMsg) *UrgentMsg {
	return &UrgentMsg{from, method}
}
