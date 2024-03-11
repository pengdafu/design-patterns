package main

import "fmt"

// 子系统1
type SubSystem1 struct{}

func (s *SubSystem1) Operation1() {
	fmt.Println("SubSystem1 Operation1")
}

func (s *SubSystem1) Operation2() {
	fmt.Println("SubSystem1 Operation2")
}

// 子系统2
type SubSystem2 struct{}

func (s *SubSystem2) Operation1() {
	fmt.Println("SubSystem2 Operation1")
}

func (s *SubSystem2) Operation2() {
	fmt.Println("SubSystem2 Operation2")
}

// 外观角色
type Facade struct {
	sub1 *SubSystem1
	sub2 *SubSystem2
}

func (f *Facade) Operation() {
	f.sub1.Operation1()
	f.sub1.Operation2()
	f.sub2.Operation1()
	f.sub2.Operation2()
}

func main() {
	facade := &Facade{
		sub1: &SubSystem1{},
		sub2: &SubSystem2{},
	}
	facade.Operation()
}
