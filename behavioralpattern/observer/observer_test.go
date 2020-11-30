package observer

import "testing"

func TestObserver(t *testing.T)  {
	subject := new(Subject)

	reader1 := NewReader1("reader1")
	reader2 := NewReader1("reader2")
	reader3 := NewReader1("reader3")

	subject.Attach(reader1)
	subject.Attach(reader2)
	subject.Attach(reader3)

	subject.UpdateContext("lizeyuan")
}