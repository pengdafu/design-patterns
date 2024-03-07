package main

import "fmt"

// 目标接口
type Target interface {
	Request() string
}

// 被适配者
type Adaptee struct{}

func (a *Adaptee) SpecificRequest() string {
	return "Adaptee Request"
}

// 适配器
type Adapter struct {
	adaptee *Adaptee
}

func (a *Adapter) Request() string {
	return a.adaptee.SpecificRequest()
}

func main() {
	adaptee := &Adaptee{}
	adapter := &Adapter{adaptee: adaptee}

	target := adapter

	response := target.Request()
	fmt.Println(response)
	// 输出: Adaptee Request
}
