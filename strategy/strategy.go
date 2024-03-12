package main

import "fmt"

// Strategy 抽象策略接口
type PaymentStrategy interface {
	Pay(amount int)
}

// Context 上下文
type PaymentContext struct {
	strategy PaymentStrategy
}

func NewPaymentContext(strategy PaymentStrategy) *PaymentContext {
	return &PaymentContext{strategy: strategy}
}

func (c *PaymentContext) SetStrategy(strategy PaymentStrategy) {
	c.strategy = strategy
}

func (c *PaymentContext) Pay(amount int) {
	c.strategy.Pay(amount)
}

// ConcreteStrategy 具体策略实现

type CashPayment struct{}

func (*CashPayment) Pay(amount int) {
	fmt.Printf("Paid %d using cash\n", amount)
}

type CreditCardPayment struct {
	name   string
	number string
}

func (c *CreditCardPayment) Pay(amount int) {
	fmt.Printf("Paid %d using credit card %s (%s)\n", amount, c.name, c.number)
}

type PaypalPayment struct {
	email string
}

func (p *PaypalPayment) Pay(amount int) {
	fmt.Printf("Paid %d using Paypal account %s\n", amount, p.email)
}

// 使用示例
func main() {
	cash := &CashPayment{}
	creditCard := &CreditCardPayment{
		name:   "John Doe",
		number: "1234567890123456",
	}
	paypal := &PaypalPayment{
		email: "john@example.com",
	}

	context := NewPaymentContext(cash)
	context.Pay(100) // Paid 100 using cash

	context.SetStrategy(creditCard)
	context.Pay(200) // Paid 200 using credit card John Doe (1234567890123456)

	context.SetStrategy(paypal)
	context.Pay(300) // Paid 300 using Paypal account john@example.com
}
