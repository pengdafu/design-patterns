package main

import "fmt"

// Implementor 定义了实现类的接口
type OperatingSystem interface {
	Render()
}

// ConcreteImplementor 是具体的实现类
type Windows struct{}

func (w *Windows) Render() {
	fmt.Println("Rendering on Windows")
}

type MacOS struct{}

func (m *MacOS) Render() {
	fmt.Println("Rendering on MacOS")
}

type Linux struct{}

func (l *Linux) Render() {
	fmt.Println("Rendering on Linux")
}

// Abstraction 定义了抽象类的接口
type AbstractGUI struct {
	os OperatingSystem
}

func (gui *AbstractGUI) Render() {
	gui.os.Render()
}

func (gui *AbstractGUI) SetOperatingSystem(os OperatingSystem) {
	gui.os = os
}

// RefinedAbstraction 扩展了抽象类的接口
type RefinedGUI struct {
	AbstractGUI
	theme string // 新增一个主题属性
}

func NewRefinedGUI(os OperatingSystem, theme string) *RefinedGUI {
	refinedGUI := &RefinedGUI{theme: theme}
	refinedGUI.SetOperatingSystem(os)
	return refinedGUI
}

func (gui *RefinedGUI) Render() {
	fmt.Printf("Rendering with %s theme...\n", gui.theme)
	gui.AbstractGUI.Render()
}

func main() {
	windowsGUI := NewRefinedGUI(&Windows{}, "Classic")
	windowsGUI.Render() // Rendering with Classic theme...
	// Rendering on Windows

	macGUI := NewRefinedGUI(&MacOS{}, "Modern")
	macGUI.Render() // Rendering with Modern theme...
	// Rendering on MacOS
}
