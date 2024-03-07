package main

import "fmt"

// Prototype 是原型接口，定义了一个克隆方法
type Prototype interface {
	Clone() Prototype
}

// ConcretePrototype 是具体原型类，实现了 Prototype 接口
type ConcretePrototype struct {
	Name string
}

// Clone 方法复制自身并返回一个新的对象实例
func (cp *ConcretePrototype) Clone() Prototype {
	return &ConcretePrototype{Name: cp.Name}
}

func main() {
	// 创建原型对象
	prototype := &ConcretePrototype{Name: "Prototype"}

	// 克隆原型对象来创建新对象实例
	clone1 := prototype.Clone()
	clone2 := prototype.Clone()

	// 输出新对象实例的内容
	fmt.Println(clone1.(*ConcretePrototype).Name) // 输出：Prototype
	fmt.Println(clone2.(*ConcretePrototype).Name) // 输出：Prototype
}
