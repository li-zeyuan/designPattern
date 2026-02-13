package observer

import "testing"

func TestObserver(t *testing.T) {
	subject := NewConcreteSubject()

	observer1 := &ConcreteObserver{ID: "1"}
	observer2 := &ConcreteObserver{ID: "2"}

	subject.Register(observer1)
	subject.Register(observer2)

	subject.NotifyObservers("Hello, Observers!")

	// subject.Unregister(observer1)
}
