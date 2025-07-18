package abstract

import (
	"fmt"
)

// 1. Определяем интерфейсы для окон и кнопок.

// Window интерфейс продукта.
type Window interface {
	Render()
}

// Button интерфейс продукта.
type Button interface {
	OnClick()
}

// 2. Реализуем конкретные классы продуктов для каждой платформы.

// WinWindow конкретный продукт для Windows.
type WinWindow struct{}

func (w *WinWindow) Render() {
	fmt.Println("Rendering Windows window...")
}

// WinButton конкретный продукт для Windows.
type WinButton struct{}

func (b *WinButton) OnClick() {
	fmt.Println("Handling Windows button click...")
}

// MacWindow конкретный продукт для macOS.
type MacWindow struct{}

func (w *MacWindow) Render() {
	fmt.Println("Rendering macOS window...")
}

// MacButton конкретный продукт для macOS.
type MacButton struct{}

func (b *MacButton) OnClick() {
	fmt.Println("Handling macOS button click...")
}

// 3 Определение абстрактной фабрики, которая будет отвечать за создание обоих типов продуктов.

type GUIFactory interface {
	CreateWindow() Window
	CreateButton() Button
}

// 4 Реализует две конкретные фабрики для каждой платформы.

// WinGUIFactory конкретная фабрика для Windows.
type WinGUIFactory struct{}

func (g *WinGUIFactory) CreateWindow() Window {
	return &WinWindow{}
}

func (g *WinGUIFactory) CreateButton() Button {
	return &WinButton{}
}

// MacGUIFactory конкретная фабрика для macOS.
type MacGUIFactory struct{}

func (g *MacGUIFactory) CreateWindow() Window {
	return &MacWindow{}
}

func (g *MacGUIFactory) CreateButton() Button {
	return &MacButton{}
}

func StartAbstractPattern() {
	winFactory := &WinGUIFactory{}
	macFactory := &MacGUIFactory{}

	win := winFactory.CreateWindow()
	btn := winFactory.CreateButton()

	win.Render()
	btn.OnClick()

	win = macFactory.CreateWindow()
	btn = macFactory.CreateButton()

	win.Render()
	btn.OnClick()
}
