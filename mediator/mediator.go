package main

import "fmt"

// 定义抽象中介者
type Mediator interface {
	Send(sender, receiver, message string)
}

// 定义具体中介者
type ConcreteMediator struct {
	colleagues map[string]Colleague
}

func NewMediator() *ConcreteMediator {
	return &ConcreteMediator{
		colleagues: make(map[string]Colleague),
	}
}

func (m *ConcreteMediator) Add(name string, colleague Colleague) {
	m.colleagues[name] = colleague
}

func (m *ConcreteMediator) Send(sender, receiver, message string) {
	if send, ok := m.colleagues[sender]; ok {
		if receive, ok := m.colleagues[receiver]; ok {
			send.Receive(message)
			receive.Receive(message)
		}
	}
}

// 定义抽象同事类
type Colleague interface {
	Receive(message string)
	GetName() string
}

// 定义具体同事类
type ConcreteColleague struct {
	name     string
	mediator Mediator
}

func NewColleague(name string, mediator Mediator) *ConcreteColleague {
	return &ConcreteColleague{
		name:     name,
		mediator: mediator,
	}
}

func (c *ConcreteColleague) Receive(message string) {
	fmt.Printf("%s received message: %s\n", c.name, message)
}

func (c *ConcreteColleague) GetName() string {
	return c.name
}

func (c *ConcreteColleague) Send(receiver, message string) {
	c.mediator.Send(c.name, receiver, message)
}

// 使用示例
func main() {
	mediator := NewMediator()

	colleague1 := NewColleague("Alice", mediator)
	colleague2 := NewColleague("Bob", mediator)

	mediator.Add(colleague1.GetName(), colleague1)
	mediator.Add(colleague2.GetName(), colleague2)

	colleague1.Send(colleague2.GetName(), "Hello, Bob!")
	colleague2.Send(colleague1.GetName(), "Hi, Alice!")
}
