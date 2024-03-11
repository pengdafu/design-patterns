命令模式(Command Pattern)是一种行为型设计模式,它将请求封装为一个对象,从而允许你使用不同的请求对客户端进行参数化,对请求进行排队或记录请求日志,以及支持可撤销的操作。

命令模式涉及以下几个角色:

1. **Command(命令角色)**: 声明执行操作的接口。
2. **ConcreteCommand(具体命令角色)**: 定义一个对象和与其相关的执行动作的绑定关系。调用具体命令的`Execute`方法,将导致相应的操作发生。
3. **Client(客户端角色)**: 创建具体命令对象,并设置接收者。
4. **Invoker(调用者角色)**: 负责调用命令对象执行请求。
5. **Receiver(接收者角色)**: 执行与命令有关的操作。任何类都可能成为一个接收者。

下面是一个使用 Go 语言实现的命令模式示例,模拟了一个简单的智能家居系统:

```go
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
```

在这个示例中:

- `Command`是命令角色接口,定义了执行操作的方法。
- `Light`是接收者角色,实现了开关灯的操作。
- `TurnOnCommand`和`TurnOffCommand`是具体命令角色,分别封装了开灯和关灯的操作。
- `RemoteControl`是调用者角色,负责存储命令对象并调用它们的`Execute`方法。
- 在`main`函数中,我们创建了一个`Light`对象和相应的命令对象,然后通过`RemoteControl`对象控制灯的开关。

通过使用命令模式,我们可以将请求封装为对象,从而实现了请求的参数化。这种方式可以让我们很容易地设计一个可撤销的操作,只需记录命令对象并在需要时执行相反的命令即可。此外,命令模式还支持请求的排队和日志记录等功能。

命令模式的优点:

1. 将请求与执行操作的对象解耦,增加了灵活性。
2. 方便实现对请求的撤销和重做。
3. 支持请求的排队、日志记录和持久化等功能。
4. 可以将多个命令组合成一个复合命令。

缺点:

1. 可能会导致系统中存在大量具体命令类。
2. 增加了内存消耗,因为每个命令都需要一个具体的命令对象。

命令模式通常应用于以下场景:

- 需要将请求与执行操作的对象解耦。
- 需要支持请求的撤销和重做操作。
- 需要支持请求的排队、日志记录和持久化等功能。
- 需要将多个操作组合成一个复合操作。

总之,命令模式通过将请求封装为对象,实现了请求的参数化,使得请求发送者与请求接收者之间解耦,增加了系统的灵活性。在合适的场景下使用命令模式,可以使代码更加清晰、可维护,并增强了系统的功能。

redis 就是一个使用命令模式很好的例子。