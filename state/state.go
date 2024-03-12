package main

import "fmt"

// State 抽象状态接口
type State interface {
	InsertMoney(c *Context, money int) State
	EjectMoney(c *Context) State
	DispenseProduct(c *Context) State
}

// Context 上下文
type Context struct {
	currentState State
	balance      int
}

func NewContext() *Context {
	return &Context{currentState: NewNoMoneyState()}
}

func (c *Context) SetState(s State) {
	c.currentState = s
}

func (c *Context) InsertMoney(money int) {
	c.currentState = c.currentState.InsertMoney(c, money)
}

func (c *Context) EjectMoney() {
	c.currentState = c.currentState.EjectMoney(c)
}

func (c *Context) DispenseProduct() {
	c.currentState = c.currentState.DispenseProduct(c)
}

// ConcreteState 具体状态实现

type NoMoneyState struct{}

func NewNoMoneyState() *NoMoneyState {
	return &NoMoneyState{}
}

func (s *NoMoneyState) InsertMoney(c *Context, money int) State {
	c.balance += money
	return NewHasMoneyState()
}

func (s *NoMoneyState) EjectMoney(c *Context) State {
	fmt.Println("No money to eject.")
	return s
}

func (s *NoMoneyState) DispenseProduct(c *Context) State {
	fmt.Println("Please insert money first.")
	return s
}

type HasMoneyState struct{}

func NewHasMoneyState() *HasMoneyState {
	return &HasMoneyState{}
}

func (s *HasMoneyState) InsertMoney(c *Context, money int) State {
	c.balance += money
	return s
}

func (s *HasMoneyState) EjectMoney(c *Context) State {
	fmt.Printf("Ejecting %d cents.\n", c.balance)
	c.balance = 0
	return NewNoMoneyState()
}

func (s *HasMoneyState) DispenseProduct(c *Context) State {
	if c.balance >= 50 {
		fmt.Println("Dispensing product.")
		c.balance -= 50
		if c.balance == 0 {
			return NewNoMoneyState()
		} else {
			return s
		}
	} else {
		fmt.Println("Not enough money to dispense product.")
		return s
	}
}

// 使用示例
func main() {
	context := NewContext()

	context.DispenseProduct() // Please insert money first.
	context.InsertMoney(30)   // 余额为30
	context.DispenseProduct() // Not enough money to dispense product.
	context.InsertMoney(20)   // 余额为50
	context.DispenseProduct() // Dispensing product.
	context.EjectMoney()      // Ejecting 0 cents.
}
