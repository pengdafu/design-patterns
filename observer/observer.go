package main

import "fmt"

// Subject 主题接口
type Subject interface {
	Register(observer Observer)
	Deregister(observer Observer)
	NotifyAll()
}

// Observer 观察者接口
type Observer interface {
	Update(subject Subject)
}

// ConcreteSubject 具体主题
type ConcreteSubject struct {
	observers []Observer
	state     string
}

func (s *ConcreteSubject) Register(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *ConcreteSubject) Deregister(observer Observer) {
	for i, ob := range s.observers {
		if ob == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *ConcreteSubject) NotifyAll() {
	for _, observer := range s.observers {
		observer.Update(s)
	}
}

func (s *ConcreteSubject) SetState(state string) {
	s.state = state
	s.NotifyAll()
}

func (s *ConcreteSubject) GetState() string {
	return s.state
}

// ConcreteObserver 具体观察者
type ConcreteObserver struct {
	id    string
	state string
}

func (o *ConcreteObserver) Update(subject Subject) {
	s, ok := subject.(*ConcreteSubject)
	if ok {
		o.state = s.GetState()
		fmt.Printf("Observer %s state updated to %s\n", o.id, o.state)
	}
}

// 使用示例
func main() {
	subject := &ConcreteSubject{state: "Initial state"}

	observer1 := &ConcreteObserver{id: "Observer 1"}
	observer2 := &ConcreteObserver{id: "Observer 2"}

	subject.Register(observer1)
	subject.Register(observer2)

	subject.SetState("New state")
	// Output:
	// Observer Observer 1 state updated to New state
	// Observer Observer 2 state updated to New state

	subject.Deregister(observer2)

	subject.SetState("Another state")
	// Output:
	// Observer Observer 1 state updated to Another state
}
