package observer

import "fmt"


// ==================抽象发布主题接口========
type Subject interface {
	Register(observer Observer)
	Unregister(observer Observer)
	NotifyObservers(data string)
}

// ===================具体发布主题对象=============
type ConcreteSubject struct {
	observers map[string]Observer // 观察者列表
}

func NewConcreteSubject() Subject {
	return &ConcreteSubject{
		observers: make(map[string]Observer),
	}
}

func (s *ConcreteSubject) Register(observer Observer) {
	observerID := fmt.Sprintf("%p", observer)
	s.observers[observerID] = observer
}

func (s *ConcreteSubject) Unregister(observer Observer) {
	observerID := fmt.Sprintf("%p", observer)
	delete(s.observers, observerID)
}

func (s *ConcreteSubject) NotifyObservers(data string) {
	for _, observer := range s.observers {
		observer.Notify(data)
	}
}

// ================抽象观察者接口================
type Observer interface {
	Notify(data string)
}

// ===============具体观察者对象=================
type ConcreteObserver struct {
	ID   string
	Data string
}

func (o *ConcreteObserver) Notify(data string) {
	fmt.Printf("Observer %s received: %s\n", o.ID, data)
	o.Data = data
}
