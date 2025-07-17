package main

import (
	"fmt"
	"patterns/generative/builder"
)

func main() {
	fmt.Println("Hello patterns!")

	fmt.Println("Generative patterns ---> Builder")

	build := builder.NewCarBuilder()
	car1 := build.SetBrand("Tesla").SetColor("красный").SetEnginePower(300).Build()
	car2 := build.SetBrand("Ferrari").SetEnginePower(1200).Build()

	fmt.Printf("Автомобиль марки: %s, цвет: %s, мощность двигателя: %d л.с.\n", car1.Brand, car1.Color, car1.EnginePower)
	fmt.Printf("Автомобиль марки: %s, цвет: неизвестный, мощность двигателя: %d л.с.\n", car2.Brand, car2.EnginePower)
}
