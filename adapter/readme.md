适配器模式(Adapter Pattern)是一种结构型设计模式,它通过组合的方式,将一个接口转换为客户端所期望的另一种接口,从而使原本由于接口不兼容而无法工作的类能够协同工作。适配器模式的作用就像它的名字一样,是一种适配的角色,使两个原本不兼容的接口协同工作。

适配器模式主要包含以下四个角色:

1. **目标接口(Target)**: 客户端所期望的接口,目标接口可以是具体的或抽象的,只要有客户端使用即可。
2. **被适配者(Adaptee)**: 需要被适配的类,即待适配的接口。
3. **适配器(Adapter)**: 适配器类,它继承或引用了被适配者类,实现了目标接口,这是适配器模式的核心。
4. **客户端(Client)**: 调用适配器转化之后的目标接口。

适配器模式分为两种:

1. **类适配器模式**: 适配器通过继承将被适配者的接口转换为目标接口。
2. **对象适配器模式**: 适配器通过组合将被适配者的接口转换为目标接口。

下面是一个使用Go语言实现对象适配器模式的示例:

```go
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
```

在这个例子中:

- `Target`是目标接口,定义了客户端所期望的接口。
- `Adaptee`是被适配者,拥有一个不同接口的`SpecificRequest()`方法。
- `Adapter`是适配器,通过组合的方式包含了`Adaptee`实例,并实现了`Target`接口的`Request()`方法,将`SpecificRequest()`适配为目标接口。
- 在`main()`函数中,我们通过适配器实例来调用`Request()`方法,实现了对不同接口的适配。

适配器模式的优点:

1. 让两个原本接口不兼容的类可以协同工作。
2. 增加了代码的复用性和可扩展性。
3. 遵循了"单一职责原则"和"开放-封闭原则"。

适配器模式的缺点:

1. 会增加系统的复杂性,需要引入一个新的适配器类。
2. 由于使用了组合或继承的关系,会导致系统可扩展性降低。

适配器模式常见应用场景:

1. 需要使用一个已存在的类,但其接口不符合要求时。
2. 需要创建可复用的类,同时又希望这些类可以被不同接口使用时。
3. 对于遗留系统,在不改变已有接口的前提下,实现新接口时。

总之,适配器模式通过引入一个适配器类,可以将不兼容的接口包装在一起,为客户端提供统一的接口。它遵循了"开放-封闭原则"和"单一职责原则",有助于代码复用和可扩展性,是解决接口不兼容问题的常用模式。

### 为什么需要适配器？

举一个适配器模式的实际应用场景例子:

假设我们有一个第三方提供的旧版本外部库,它提供了一系列用于读写文件的方法,但接口已经过时且不符合当前项目的要求。例如,它提供了一个`LegacyFileIO`类,有`readFile()`和`writeFile()`等方法:

```go
type LegacyFileIO struct{}

func (l *LegacyFileIO) readFile(filePath string) ([]byte, error) {
    // 旧版本读取文件的代码
    return []byte{}, nil
}

func (l *LegacyFileIO) writeFile(filePath string, data []byte) error {
    // 旧版本写入文件的代码
    return nil
}
```

而在我们的项目中,我们期望使用一个更加现代化的`FileIO`接口,它定义了`ReadFile()`和`WriteFile()`方法:

```go
type FileIO interface {
    ReadFile(filePath string) ([]byte, error)
    WriteFile(filePath string, data []byte) error
}
```

此时,我们可以使用适配器模式来解决接口不兼容的问题。我们创建一个`FileIOAdapter`适配器,它实现了`FileIO`接口,内部组合了`LegacyFileIO`实例:

```go
type FileIOAdapter struct {
    legacyFileIO *LegacyFileIO
}

func (a *FileIOAdapter) ReadFile(filePath string) ([]byte, error) {
    return a.legacyFileIO.readFile(filePath)
}

func (a *FileIOAdapter) WriteFile(filePath string, data []byte) error {
    return a.legacyFileIO.writeFile(filePath, data)
}
```

然后,在我们的客户端代码中,我们就可以使用`FileIO`接口,而不需要直接依赖`LegacyFileIO`:

```go
func main() {
    legacyFileIO := &LegacyFileIO{}
    fileIOAdapter := &FileIOAdapter{legacyFileIO: legacyFileIO}

    fileIO := fileIOAdapter // 使用适配器实例

    data, err := fileIO.ReadFile("file.txt")
    if err != nil {
        // 处理错误
    }

    err = fileIO.WriteFile("file.txt", data)
    if err != nil {
        // 处理错误
    }
}
```

在这个例子中:

- `LegacyFileIO`是被适配者,提供了旧版本的文件IO接口。
- `FileIO`是目标接口,客户端所期望的新版本文件IO接口。
- `FileIOAdapter`是适配器,它将`LegacyFileIO`的接口适配为`FileIO`接口。
- 在`main()`函数中,我们通过适配器实例`fileIOAdapter`来使用`FileIO`接口,从而间接访问`LegacyFileIO`的功能。

通过使用适配器模式,我们可以在不修改客户端代码和第三方库代码的情况下,将不兼容的接口进行适配,使它们能够协同工作。这提高了代码的可维护性和可扩展性,符合"开放-封闭原则"。