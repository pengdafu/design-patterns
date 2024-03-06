## 单例模式详细说明

### 概述

单例模式是一种设计模式，确保一个类只有一个实例，并提供一个全局访问点供外部访问。它主要用于控制对共享资源的访问，如配置管理器、数据库连接池等。单例模式能够减少不必要的内存开销，避免对共享资源的多重占用，同时也可以提高访问效率。

### 实现方式

在Go语言中，实现单例模式通常有两种方式：懒汉式（延迟初始化）和饿汉式。

#### 懒汉式（延迟初始化）

懒汉式单例模式指的是单例实例在第一次被请求时才创建。这种方式的主要优点是实现了按需加载，但需要考虑线程安全问题。

**实现示例**:

```go
package singleton

import (
    "sync"
)

type Singleton struct{}

var instance *Singleton
var once sync.Once

func GetInstance() *Singleton {
    once.Do(func() {
        instance = &Singleton{}
    })
    return instance
}
```

在这个示例中，`sync.Once`的`Do`方法保证了无论多少次调用`GetInstance`，实例化代码只执行一次，确保了线程安全。

#### 饿汉式

饿汉式单例模式指的是单例实例在程序启动时即被创建。这种方式的主要优点是实现简单，由于Go语言包的初始化（包括`init`函数的调用）是并发安全的，因此它也是线程安全的。缺点是不管实例最终是否被用到，它都会被创建。

**实现示例**:

```go
package singleton

type Singleton struct{}

var instance = &Singleton{}

//var instance *Singleton
//func init() {
//	instance = &Singleton{}
//}

func GetInstance() *Singleton {
    return instance
}
```

在这个示例中，单例`instance`在包级别直接初始化，确保了它的全局唯一性和线程安全性。

### 选择建议

- **懒汉式**适用于单例对象初始化开销较大，或初始化时机有特殊要求的场景，因为它支持延迟初始化。
- **饿汉式**适用于单例对象占用资源少，对启动性能影响不大的场景，因为它简化了实现并且天然线程安全。

### 注意事项

单例模式虽然有很多优点，但也不应过度使用。滥用单例模式可能导致系统中的模块高度耦合，增加代码的复杂度和测试难度。在使用单例模式时，需要权衡其带来的便利与潜在的问题。
