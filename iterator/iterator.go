package main

import "fmt"

// 定义一个迭代器接口
type Iterator interface {
	HasNext() bool
	Next() interface{}
}

// 定义一个集合结构体
type Collection struct {
	data []interface{}
}

// 实现迭代器接口
type CollectionIterator struct {
	collection *Collection
	index      int
}

// 创建一个新的迭代器实例
func (c *Collection) CreateIterator() Iterator {
	return &CollectionIterator{
		collection: c,
		index:      0,
	}
}

// 判断是否还有下一个元素可以访问
func (it *CollectionIterator) HasNext() bool {
	return it.index < len(it.collection.data)
}

// 获取下一个元素
func (it *CollectionIterator) Next() interface{} {
	if !it.HasNext() {
		return nil
	}
	value := it.collection.data[it.index]
	it.index++
	return value
}

// 使用示例
func main() {
	collection := &Collection{
		data: []interface{}{"A", "B", "C", "D", "E"},
	}

	iterator := collection.CreateIterator()

	for iterator.HasNext() {
		value := iterator.Next()
		fmt.Println(value)
	}
}
