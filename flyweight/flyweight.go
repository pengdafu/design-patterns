package main

import "fmt"

// 享元对象
type Flyweight interface {
	Render(extrinsicState string)
}

// 具体享元对象
type ConcreteFlyweight struct {
	intrinsicState string // 内在状态
}

func (c *ConcreteFlyweight) Render(extrinsicState string) {
	fmt.Printf("Rendered with %s and %s\n", c.intrinsicState, extrinsicState)
}

// 享元工厂
type FlyweightFactory struct {
	pool map[string]Flyweight
}

func NewFlyweightFactory() *FlyweightFactory {
	return &FlyweightFactory{
		pool: make(map[string]Flyweight),
	}
}

func (f *FlyweightFactory) GetFlyweight(intrinsicState string) Flyweight {
	if flyweight, ok := f.pool[intrinsicState]; ok {
		return flyweight
	}
	flyweight := &ConcreteFlyweight{intrinsicState: intrinsicState}
	f.pool[intrinsicState] = flyweight
	return flyweight
}

func main() {
	factory := NewFlyweightFactory()

	flyweight1 := factory.GetFlyweight("Bold")
	flyweight1.Render("Hello")

	flyweight2 := factory.GetFlyweight("Italic")
	flyweight2.Render("World")

	// 重用已有的享元对象
	sharedFlyweight := factory.GetFlyweight("Bold")
	sharedFlyweight.Render("Hello again")
}
