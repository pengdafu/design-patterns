package main

import "fmt"

// Button 是按钮接口，所有的具体按钮都需要实现 Click 方法
type Button interface {
	Click() string
}

// MacButton 是 Mac 按钮结构体，实现了 Button 接口
type MacButton struct{}

// Click 方法实现了点击 Mac 按钮的逻辑
func (mb *MacButton) Click() string {
	return "Clicked Mac button"
}

// WinButton 是 Windows 按钮结构体，实现了 Button 接口
type WinButton struct{}

// Click 方法实现了点击 Windows 按钮的逻辑
func (wb *WinButton) Click() string {
	return "Clicked Windows button"
}

// GUIFactory 是 GUI 工厂接口，包含创建按钮的方法
type GUIFactory interface {
	CreateButton() Button
}

// MacFactory 是 Mac GUI 工厂结构体，实现了 GUIFactory 接口
type MacFactory struct{}

// CreateButton 方法创建了一个 Mac 按钮对象
func (mf *MacFactory) CreateButton() Button {
	return &MacButton{}
}

// WinFactory 是 Windows GUI 工厂结构体，实现了 GUIFactory 接口
type WinFactory struct{}

// CreateButton 方法创建了一个 Windows 按钮对象
func (wf *WinFactory) CreateButton() Button {
	return &WinButton{}
}

func main() {
	// 使用 Mac 工厂创建 Mac 按钮
	var macFactory GUIFactory
	macFactory = &MacFactory{}
	macButton := macFactory.CreateButton()
	fmt.Println(macButton.Click())

	// 使用 Windows 工厂创建 Windows 按钮
	winFactory := &WinFactory{}
	winButton := winFactory.CreateButton()
	fmt.Println(winButton.Click())
}
