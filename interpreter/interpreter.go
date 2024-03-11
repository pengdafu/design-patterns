package main

import "fmt"

// 抽象表达式
type Expression interface {
	Interpret(context *Context) int
}

// 非终结符表达式
type PlusExpression struct {
	left, right Expression
}

func (p *PlusExpression) Interpret(context *Context) int {
	return p.left.Interpret(context) + p.right.Interpret(context)
}

// 终结符表达式
type NumericExpression struct {
	value int
}

func (n *NumericExpression) Interpret(context *Context) int {
	return n.value
}

// 环境角色
type Context struct {
	// 可以添加一些全局信息
}

func main() {
	// 构造抽象语法树
	expression := &PlusExpression{
		left: &NumericExpression{5},
		right: &PlusExpression{
			left:  &NumericExpression{3},
			right: &NumericExpression{2},
		},
	}

	context := &Context{}
	result := expression.Interpret(context)
	fmt.Println("Result:", result) // 输出: Result: 10
}
