package prototype

import "fmt"

// 1. Определение прототипа.

// Prototype интерфейс прототипа.
type Prototype interface {
	Clone() Prototype
}

// 2. Реализуем конкретные объекты, которые поддерживают клонирование.

// Document прототип документа.
type Document struct {
	Content   string
	FontSize  int
	FontColor string
}

// Clone реализация метода клонирования.
func (d *Document) Clone() Prototype {
	return &Document{
		Content:   d.Content,
		FontSize:  d.FontSize,
		FontColor: d.FontColor,
	}
}

func (d *Document) Print() {
	fmt.Printf("Content: %s, Font Size: %d, Font Color: %s\n", d.Content, d.FontSize, d.FontColor)
}

func StartPrototypePattern() {
	originalDoc := Document{
		Content:   "original",
		FontSize:  12,
		FontColor: "red",
	}

	cloneDoc := originalDoc.Clone().(*Document)

	cloneDoc.Content = "clone"
	cloneDoc.FontSize = 11
	cloneDoc.FontColor = "black"

	originalDoc.Print()
	cloneDoc.Print()
}
