package builder

import "fmt"

// Car структура автомобиля.
type Car struct {
	Brand       string
	Color       string
	EnginePower int
}

// CarBuilder определяет методы для задания характеристик автомобиля.
type CarBuilder interface {
	SetBrand(brand string) CarBuilder
	SetColor(color string) CarBuilder
	SetEnginePower(engine int) CarBuilder
	Build() Car
}

// carBuilder конкретная реализация CarBuilder.
type carBuilder struct {
	brand       string
	color       string
	enginePower int
}

// NewCarBuilder конструктор.
func NewCarBuilder() CarBuilder {
	return &carBuilder{}
}

// SetBrand установка бренда машины.
func (cb *carBuilder) SetBrand(brand string) CarBuilder {
	cb.brand = brand

	return cb
}

// SetColor установка цвета машины.
func (cb *carBuilder) SetColor(color string) CarBuilder {
	cb.color = color

	return cb
}

// SetEnginePower установка мощности автомобиля.
func (cb *carBuilder) SetEnginePower(engine int) CarBuilder {
	cb.enginePower = engine

	return cb
}

// Build завершающий метод, которые возвращает объект Car.
func (cb *carBuilder) Build() Car {
	return Car{
		Brand:       cb.brand,
		Color:       cb.color,
		EnginePower: cb.enginePower,
	}
}

func StartBuilderPattern() {
	build := NewCarBuilder()
	car1 := build.SetBrand("Tesla").SetColor("красный").SetEnginePower(300).Build()
	car2 := build.SetBrand("Ferrari").SetEnginePower(1200).Build()

	fmt.Printf("Автомобиль марки: %s, цвет: %s, мощность двигателя: %d л.с.\n", car1.Brand, car1.Color, car1.EnginePower)
	fmt.Printf("Автомобиль марки: %s, цвет: неизвестный, мощность двигателя: %d л.с.\n", car2.Brand, car2.EnginePower)
}
