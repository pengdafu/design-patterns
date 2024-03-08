装饰器模式的核心思想是:动态地给一个对象增加一些额外的职责或行为。它通过创建一个装饰类,用来包裹原有对象,并对其进行扩展,这样就可以在运行时选择性地为对象加上新的行为。

这种设计模式涉及到四个角色:

1. **Component(抽象构件角色)**
   定义一个对象的接口,可以是抽象类或接口,为具体构件和装饰类对象的共同接口/超类。

2. **ConcreteComponent(具体构件角色)**
   是基本对象,Component 角色的具体实现。

3. **Decorator(装饰角色)**
   持有一个构件对象,并实现和它相同的接口。它可以在构件的方法执行前后添加自己的行为。

4. **ConcreteDecorator(具体装饰角色)**
   具体的装饰对象,起到给Component 对象增加职责的作用。

我们通过一个示例来理解:

```go
// Component 抽象构件角色
type Beverage interface {
    getDescription() string
    cost() float64
}

// ConcreteComponent 具体构件角色
type Espresso struct{}

func (e *Espresso) getDescription() string {
    return "Espresso"
}
func (e *Espresso) cost() float64 {
    return 1.99
}

// Decorator 装饰角色
type CondimentDecorator struct {
    beverage Beverage
}

func (c *CondimentDecorator) getDescription() string {
    return c.beverage.getDescription()
}
func (c *CondimentDecorator) cost() float64 {
    return c.beverage.cost()
}

// ConcreteDecorator 具体装饰角色 go中没有继承,所以这里使用组合
type Mocha struct {
    CondimentDecorator
}

func (m *Mocha) getDescription() string {
    return m.CondimentDecorator.getDescription() + ", Mocha"
}
func (m *Mocha) cost() float64 {
    return m.CondimentDecorator.cost() + 0.20
}
```

这里模拟一个咖啡店点单的场景:

- `Beverage`是抽象构件,定义了获取描述和价格的方法
- `Espresso`是具体构件,一杯基本的浓缩咖啡
- `CondimentDecorator`是装饰角色,持有一个`Beverage`对象
- `Mocha`是具体装饰角色,给饮料增加了"摩卡"的行为和价格

使用:

```go
beverage := &Espresso{}
println(beverage.getDescription() + " $" + fmt.Sprintf("%.2f", beverage.cost())) 
// 输出: Espresso $1.99

mocha := &Mocha{CondimentDecorator{beverage}}
println(mocha.getDescription() + " $" + fmt.Sprintf("%.2f", mocha.cost()))
// 输出: Espresso, Mocha $2.19
```

通过这个例子可以看出:

1. `Espresso`作为基础构件存在
2. `Mocha`在不修改`Espresso`的情况下,扩展了它的功能,使其新增了描述和价格

装饰器模式的关键是组合和委托。装饰类组合了被装饰对象,并委托给它完成自身无法实现的操作,而在装饰对象中加入了扩展的新行为。通过串联多个装饰类,就可以对对象进行多次行为扩展。