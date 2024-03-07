package main

// Product 定义抽象产品接口
type Product interface {
	Use()
}

// ConcreteProductA 是一个具体产品
type ConcreteProductA struct{}

func (p *ConcreteProductA) Use() {
	println("Using ConcreteProductA")
}

// ConcreteProductB 是另一个具体产品
type ConcreteProductB struct{}

func (p *ConcreteProductB) Use() {
	println("Using ConcreteProductB")
}

// Creator 定义抽象工厂接口
type Creator interface {
	AnOperation()
	FactoryMethod() Product
}

// ConcreteCreatorA 是一个具体工厂
type ConcreteCreatorA struct{}

func (c *ConcreteCreatorA) AnOperation() {
	product := c.FactoryMethod()
	product.Use()
}

func (c *ConcreteCreatorA) FactoryMethod() Product {
	return &ConcreteProductA{}
}

// ConcreteCreatorB 是另一个具体工厂
type ConcreteCreatorB struct{}

func (c *ConcreteCreatorB) AnOperation() {
	product := c.FactoryMethod()
	product.Use()
}

func (c *ConcreteCreatorB) FactoryMethod() Product {
	return &ConcreteProductB{}
}

func main() {
	creator1 := &ConcreteCreatorA{}
	creator1.AnOperation()

	creator2 := &ConcreteCreatorB{}
	creator2.AnOperation()
}
