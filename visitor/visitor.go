package main

import "fmt"

// Component 表示电脑组件的接口
type Component interface {
	Accept(Visitor)
}

// Visitor 接口声明了访问者的方法
type Visitor interface {
	VisitCPU(cpu *CPU)
	VisitMemory(memory *Memory)
	VisitHardDrive(hardDrive *HardDrive)
}

// CPU 是一个具体的组件
type CPU struct{}

func (c *CPU) Accept(v Visitor) {
	v.VisitCPU(c)
}

// Memory 是另一个具体的组件
type Memory struct{}

func (c *Memory) Accept(v Visitor) {
	v.VisitMemory(c)
}

// HardDrive 是另一个具体的组件
type HardDrive struct{}

func (c *HardDrive) Accept(v Visitor) {
	v.VisitHardDrive(c)
}

// UpdateVisitor 是用于更新组件驱动的访问者
type UpdateVisitor struct{}

func (u *UpdateVisitor) VisitCPU(cpu *CPU) {
	fmt.Println("Updating CPU driver")
}

func (u *UpdateVisitor) VisitMemory(memory *Memory) {
	fmt.Println("Updating Memory driver")
}

func (u *UpdateVisitor) VisitHardDrive(hardDrive *HardDrive) {
	fmt.Println("Updating Hard Drive driver")
}

// DetectionVisitor 是用于进行硬件检测的访问者
type DetectionVisitor struct{}

func (d *DetectionVisitor) VisitCPU(cpu *CPU) {
	fmt.Println("Detecting CPU")
}

func (d *DetectionVisitor) VisitMemory(memory *Memory) {
	fmt.Println("Detecting Memory")
}

func (d *DetectionVisitor) VisitHardDrive(hardDrive *HardDrive) {
	fmt.Println("Detecting Hard Drive")
}

func main() {
	components := []Component{&CPU{}, &Memory{}, &HardDrive{}}

	updateVisitor := &UpdateVisitor{}
	detectionVisitor := &DetectionVisitor{}

	for _, component := range components {
		component.Accept(updateVisitor)
		component.Accept(detectionVisitor)
	}
}
