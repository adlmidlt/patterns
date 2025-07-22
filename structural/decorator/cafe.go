package decorator

import "fmt"

// Beverage интерфейс напитка.
type Beverage interface {
	Cost() float64
	Description() string
}

// Coffee базовый напиток.
type Coffee struct{}

func (c *Coffee) Cost() float64 {
	return 1.5
}

func (c *Coffee) Description() string {
	return "Простой кофе"
}

// BeverageDecorator общая структура декоратора.
type BeverageDecorator struct {
	beverage Beverage
}

func NewBeverageDecorator(b Beverage) *BeverageDecorator {
	return &BeverageDecorator{b}
}

// Cost абстрактная реализация методов интерфейса.
func (bd *BeverageDecorator) Cost() float64 {
	return bd.beverage.Cost()
}

func (bd *BeverageDecorator) Description() string {
	return bd.beverage.Description()
}

type MilkDecorator struct {
	*BeverageDecorator
}

func NewMilkDecorator(b Beverage) *MilkDecorator {
	return &MilkDecorator{NewBeverageDecorator(b)}
}

func (md *MilkDecorator) Cost() float64 {
	return md.beverage.Cost() + 0.5
}

func (md *MilkDecorator) Description() string {
	return fmt.Sprintf("%s с молоком", md.beverage.Description())
}

type SugarDecorator struct {
	*BeverageDecorator
}

func NewSugarDecorator(b Beverage) *SugarDecorator {
	return &SugarDecorator{NewBeverageDecorator(b)}
}

func (sd *SugarDecorator) Cost() float64 {
	return sd.beverage.Cost() + 0.2
}

func (sd *SugarDecorator) Description() string {
	return fmt.Sprintf("%s и сахаром.", sd.beverage.Description())
}

func StartDecoratorPattern() {
	coffee := new(Coffee)
	fmt.Printf("Описание: %s\nСтоимсоть: %.2f руб.\n", coffee.Description(), coffee.Cost())

	milkCoffee := NewMilkDecorator(coffee)
	fmt.Printf("Описание: %s\nСтоимсоть: %.2f руб.\n", milkCoffee.Description(), milkCoffee.Cost())

	sweetMilkCoffee := NewSugarDecorator(milkCoffee)
	fmt.Printf("Описание: %s\nСтоимсоть: %.2f руб.\n", sweetMilkCoffee.Description(), sweetMilkCoffee.Cost())
}
