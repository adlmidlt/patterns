package main

import (
	"fmt"
	"patterns/generative/abstract"
	"patterns/generative/builder"
	"patterns/generative/factory"
	"patterns/generative/prototype"
	"patterns/generative/singleton"
	"patterns/structural/adapter"
	"patterns/structural/decorator"
)

func main() {
	fmt.Println("Hello patterns!")

	// GENERATIVE
	fmt.Println("Generative patterns ---> Builder")
	builder.StartBuilderPattern()

	fmt.Println("Generative patterns ---> Factory")
	factory.StartFactoryPattern()

	fmt.Println("Generative patterns ---> Abstract")
	abstract.StartAbstractPattern()
	fmt.Println("Generative patterns ---> Abstract ---> Example")
	abstract.StartAbstractExamplePattern()

	fmt.Println("Generative patterns ---> Prototype")
	prototype.StartPrototypePattern()

	fmt.Println("Generative patterns ---> Singleton")
	singleton.StartSingletonPattern()

	// STRUCTURAL
	fmt.Println("Structural patterns ---> Adapter")
	adapter.StartAdapterPattern()

	fmt.Println("Structural patterns ---> Decorator")
	decorator.StartDecoratorPattern()
}
