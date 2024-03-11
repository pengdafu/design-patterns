package main

import "fmt"

// Handler 抽象处理者
type Handler interface {
	HandleRequest(request int) bool
	SetSuccessor(successor Handler)
}

// ConcreteHandler 具体处理者
type ConcreteHandler struct {
	successor   Handler
	handlerType int
}

func (h *ConcreteHandler) HandleRequest(request int) bool {
	if request == h.handlerType {
		fmt.Printf("Request %d handled by handler %d\n", request, h.handlerType)
		return true
	} else if h.successor != nil {
		return h.successor.HandleRequest(request)
	}
	return false
}

func (h *ConcreteHandler) SetSuccessor(successor Handler) {
	h.successor = successor
}

func main() {
	handler1 := &ConcreteHandler{handlerType: 1}
	handler2 := &ConcreteHandler{handlerType: 2}
	handler3 := &ConcreteHandler{handlerType: 3}

	handler1.SetSuccessor(handler2)
	handler2.SetSuccessor(handler3)

	requests := []int{2, 3, 1, 4}
	for _, request := range requests {
		handler1.HandleRequest(request)
	}
}
