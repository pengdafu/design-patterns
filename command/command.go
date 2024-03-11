package main

import "fmt"

// 命令角色
type Command interface {
	Execute()
}

// 接收者角色
type Light struct {
	Name string
}

func (l *Light) TurnOn() {
	fmt.Printf("%s light is turned on\n", l.Name)
}

func (l *Light) TurnOff() {
	fmt.Printf("%s light is turned off\n", l.Name)
}

// 具体命令角色
type TurnOnCommand struct {
	light *Light
}

func (c *TurnOnCommand) Execute() {
	c.light.TurnOn()
}

type TurnOffCommand struct {
	light *Light
}

func (c *TurnOffCommand) Execute() {
	c.light.TurnOff()
}

// 调用者角色
type RemoteControl struct {
	commands []Command
}

func (r *RemoteControl) SetCommand(index int, command Command) {
	r.commands[index] = command
}

func (r *RemoteControl) PressButton(index int) {
	r.commands[index].Execute()
}

func main() {
	light := &Light{"Living Room"}
	turnOnCommand := &TurnOnCommand{light}
	turnOffCommand := &TurnOffCommand{light}

	remoteControl := &RemoteControl{
		commands: make([]Command, 2),
	}

	remoteControl.SetCommand(0, turnOnCommand)
	remoteControl.SetCommand(1, turnOffCommand)

	remoteControl.PressButton(0) // 打开客厅灯
	remoteControl.PressButton(1) // 关闭客厅灯
}
