package command

import "fmt"

// 抽象命令接口
type Commander interface {
	Execute()
}

// 具体命令类1
type StartCommand struct {
	mb *MotherBoard
}

func NewStartCommand(mb *MotherBoard) *StartCommand {
	return &StartCommand{mb:mb}
}

func (s *StartCommand) Execute()  {
	s.mb.Start()
}

// 具体命令类2
type RebootCommand struct {
	mb *MotherBoard
}

func NewRebootCommand(mb *MotherBoard) *RebootCommand {
	return &RebootCommand{mb:mb}
}

func (r *RebootCommand)Execute()  {
	r.mb.Reboot()
}

// 接收者类
type MotherBoard struct {}

func NewMotherBoard() *MotherBoard {
	return &MotherBoard{}
}

func (*MotherBoard) Start() {
	fmt.Printf("system starting!\n")
}

func (*MotherBoard) Reboot() {
	fmt.Printf("system rebooting\n")
}

// 调用者类
type Box struct {
	button1 Commander
	button2 Commander
}

func NewBox(button1, button2 Commander) *Box {
	return &Box{
		button1: button1,
		button2: button2,
	}
}

func (b *Box) PressButton1() {
	b.button1.Execute()
}


func (b *Box) PressButton2() {
	b.button2.Execute()
}

