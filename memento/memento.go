package main

import "fmt"

// Originator 发起人
type Editor struct {
	content string
}

func (e *Editor) SetContent(content string) {
	e.content = content
}

func (e *Editor) GetContent() string {
	return e.content
}

func (e *Editor) CreateMemento() *Memento {
	return &Memento{content: e.content}
}

func (e *Editor) RestoreMemento(m *Memento) {
	e.content = m.GetContent()
}

// Memento 备忘录
type Memento struct {
	content string
}

func (m *Memento) GetContent() string {
	return m.content
}

// CareTaker 备忘录管理器
type History struct {
	mementos []*Memento
	current  int
}

func (h *History) Backup(m *Memento) {
	h.mementos = append(h.mementos, m)
	h.current++
}

func (h *History) Undo() *Memento {
	if h.current > 0 {
		h.current--
	}
	if h.current >= 1 {
		return h.mementos[h.current-1]
	}
	return nil
}

func (h *History) Redo() *Memento {
	if h.current < len(h.mementos) {
		it := h.mementos[h.current]
		h.current++
		return it
	}
	return nil
}

// 使用示例
func main() {
	editor := &Editor{content: "Initial content"}
	history := &History{}

	// 进行一些编辑操作并保存备忘录
	editor.SetContent("First edit")
	history.Backup(editor.CreateMemento())

	editor.SetContent("Second edit")
	history.Backup(editor.CreateMemento())

	editor.SetContent("Third edit")
	history.Backup(editor.CreateMemento())

	// 撤销操作
	m := history.Undo()
	editor.RestoreMemento(m)
	fmt.Println("After undo:", editor.GetContent()) // Output: After undo: Second edit

	// 重做操作
	m = history.Redo()
	editor.RestoreMemento(m)
	fmt.Println("After redo:", editor.GetContent()) // Output: After redo: Third edit
}
