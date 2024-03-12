中介者模式是一种行为型设计模式,用于定义对象之间的复杂通信和交互行为。它通过引入一个中介对象,来封装一系列对象之间的交互,使对象之间不需要显式地相互引用,从而使它们可以松散耦合。当一个对象发生改变时,只需要通知中介对象,由中介对象来协调其他对象的行为。

中介者模式有以下几个重要角色:

1. **抽象中介者(Mediator)**:定义了中介者的接口,它是各同事对象和中介者之间的接口。
2. **具体中介者(ConcreteMediator)**:实现了抽象中介者接口,协调和处理同事对象之间的交互关系。
3. **抽象同事类(Colleague)**:定义了同事类的接口,它是同事对象和中介者之间的接口。
4. **具体同事类(ConcreteColleague)**:实现了抽象同事类接口,每个具体同事类都知道它的中介者对象。

下面是一个使用Go语言实现中介者模式的示例:

```go
// 定义抽象中介者
type Mediator interface {
    Send(sender, receiver, message string)
}

// 定义具体中介者
type ConcreteMediator struct {
    colleagues map[string]Colleague
}

func NewMediator() *ConcreteMediator {
    return &ConcreteMediator{
        colleagues: make(map[string]Colleague),
    }
}

func (m *ConcreteMediator) Add(name string, colleague Colleague) {
    m.colleagues[name] = colleague
}

func (m *ConcreteMediator) Send(sender, receiver, message string) {
    if send, ok := m.colleagues[sender]; ok {
        if receive, ok := m.colleagues[receiver]; ok {
            send.Receive(message)
            receive.Receive(message)
        }
    }
}

// 定义抽象同事类
type Colleague interface {
    Receive(message string)
    GetName() string
}

// 定义具体同事类
type ConcreteColleague struct {
    name     string
    mediator Mediator
}

func NewColleague(name string, mediator Mediator) *ConcreteColleague {
    return &ConcreteColleague{
        name:     name,
        mediator: mediator,
    }
}

func (c *ConcreteColleague) Receive(message string) {
    fmt.Printf("%s received message: %s\n", c.name, message)
}

func (c *ConcreteColleague) GetName() string {
    return c.name
}

func (c *ConcreteColleague) Send(receiver, message string) {
    c.mediator.Send(c.name, receiver, message)
}

// 使用示例
func main() {
    mediator := NewMediator()

    colleague1 := NewColleague("Alice", mediator)
    colleague2 := NewColleague("Bob", mediator)

    mediator.Add(colleague1.GetName(), colleague1)
    mediator.Add(colleague2.GetName(), colleague2)

    colleague1.Send(colleague2.GetName(), "Hello, Bob!")
    colleague2.Send(colleague1.GetName(), "Hi, Alice!")
}
```

在这个示例中:

1. `Mediator` 接口定义了中介者的方法 `Send`。
2. `ConcreteMediator` 结构体实现了 `Mediator` 接口,维护了一个同事对象映射表,用于存储和管理所有同事对象。
3. `Colleague` 接口定义了同事类的方法 `Receive` 和 `GetName`。
4. `ConcreteColleague` 结构体实现了 `Colleague` 接口,每个实例都持有一个中介者对象的引用。

在 `main` 函数中,我们创建了一个 `ConcreteMediator` 实例,然后创建两个 `ConcreteColleague` 实例,并将它们添加到中介者中。接着,通过调用 `Send` 方法,两个同事对象就可以相互发送消息了,而不需要直接引用对方。

通过引入中介者对象,各个同事对象之间的直接耦合被解除,对象之间的通信被中介者对象所封装,从而使得对象之间的关系更加松散,易于扩展和维护。