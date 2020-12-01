package command

import "testing"

func TestCommand(t *testing.T)  {
	mb := NewMotherBoard()

	button1 := NewStartCommand(mb)
	button2 := NewRebootCommand(mb)

	box := NewBox(button1, button2)
	box.PressButton1()
	box.PressButton2()
}