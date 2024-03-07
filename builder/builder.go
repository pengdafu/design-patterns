package main

import "fmt"

// Product 产品角色
type Computer struct {
	cpu     string
	ram     string
	gpu     string
	storage string
}

// Builder 抽象建造者
type Builder interface {
	SetCPU(cpu string) Builder
	SetRAM(ram string) Builder
	SetGPU(gpu string) Builder
	SetStorage(storage string) Builder
	Build() Computer
}

// ConcreteBuilder 具体建造者
type ComputerBuilder struct {
	computer Computer
}

func (b *ComputerBuilder) SetCPU(cpu string) Builder {
	b.computer.cpu = cpu
	return b
}

func (b *ComputerBuilder) SetRAM(ram string) Builder {
	b.computer.ram = ram
	return b
}

func (b *ComputerBuilder) SetGPU(gpu string) Builder {
	b.computer.gpu = gpu
	return b
}

func (b *ComputerBuilder) SetStorage(storage string) Builder {
	b.computer.storage = storage
	return b
}

func (b *ComputerBuilder) Build() Computer {
	return b.computer
}

// Director 指挥者
type Director struct {
	builder Builder
}

func NewDirector(b Builder) *Director {
	return &Director{builder: b}
}

func (d *Director) Construct() Computer {
	return d.builder.
		SetCPU("Intel Core i7").
		SetRAM("32GB").
		SetGPU("Nvidia GeForce RTX 3080").
		SetStorage("1TB SSD").
		Build()
}

func main() {
	builder := &ComputerBuilder{}
	director := NewDirector(builder)

	// 构建一台计算机
	computer := director.Construct()
	fmt.Println(computer)
}
