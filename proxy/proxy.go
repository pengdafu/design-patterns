package main

import "fmt"

// 抽象主题角色
type Image interface {
	Display()
	Exist() bool
	SetData(data []byte)
}

// 真实主题角色
type RealImage struct {
	filename string
	data     []byte
}

func (r *RealImage) Display() {
	fmt.Println("Displaying:", r.filename)
}

func (r *RealImage) Exist() bool {
	return r.data != nil
}

func (r *RealImage) SetData(data []byte) {
	r.data = data
	fmt.Println("Data has been set for:", r.filename)
}

// 代理角色
type ProxyImage struct {
	realImage *RealImage
	filename  string
}

func (p *ProxyImage) Display() {
	if p.realImage == nil {
		p.realImage = &RealImage{filename: p.filename}
		// 从文件加载图像数据
		p.realImage.SetData([]byte("Image data"))
	}
	p.realImage.Display()
}

func (p *ProxyImage) Exist() bool {
	return p.realImage != nil
}

func (p *ProxyImage) SetData(data []byte) {
	if p.realImage == nil {
		p.realImage = &RealImage{filename: p.filename}
	}
	p.realImage.SetData(data)
}

func main() {
	image := &ProxyImage{filename: "image.png"}
	image.Display() // 首次调用时会加载图像数据
	image.Display() // 再次调用时不需要加载数据
}
