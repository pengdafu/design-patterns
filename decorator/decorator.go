package main

import "fmt"

// Component 抽象构件角色
type Beverage interface {
	getDescription() string
	cost() float64
}

// ConcreteComponent 具体构件角色
type Espresso struct{}

func (e *Espresso) getDescription() string {
	return "Espresso"
}
func (e *Espresso) cost() float64 {
	return 1.99
}

// Decorator 装饰角色
type CondimentDecorator struct {
	beverage Beverage
}

func (c *CondimentDecorator) getDescription() string {
	return c.beverage.getDescription()
}
func (c *CondimentDecorator) cost() float64 {
	return c.beverage.cost()
}

// ConcreteDecorator 具体装饰角色 go中没有继承,所以这里使用组合
type Mocha struct {
	CondimentDecorator
}

func (m *Mocha) getDescription() string {
	return m.CondimentDecorator.getDescription() + ", Mocha"
}
func (m *Mocha) cost() float64 {
	return m.CondimentDecorator.cost() + 0.20
}

func main() {
	beverage := &Espresso{}
	println(beverage.getDescription() + " $" + fmt.Sprintf("%.2f", beverage.cost()))
	// 输出: Espresso $1.99

	mocha := &Mocha{CondimentDecorator{beverage}}
	println(mocha.getDescription() + " $" + fmt.Sprintf("%.2f", mocha.cost()))
	// 输出: Espresso, Mocha $2.19
}
