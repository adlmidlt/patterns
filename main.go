package main

import (
	"fmt"
	"patterns/behavioral/command"
	"patterns/behavioral/mediator"
	"patterns/behavioral/state"
	"patterns/behavioral/strategy"
	"patterns/generative/abstract"
	"patterns/generative/builder"
	"patterns/generative/factory"
	"patterns/generative/prototype"
	"patterns/generative/singleton"
	"patterns/structural/adapter"
	"patterns/structural/bridge"
	"patterns/structural/decorator"
	"patterns/structural/facade"
)

func main() {
	fmt.Println("Hello patterns!")

	// GENERATIVE
	fmt.Println("Generative patterns ---> Builder")
	builder.StartBuilderPattern()
	builder.StartBuilder1Pattern()

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
	singleton.StartSingleton1Pattern()

	// STRUCTURAL
	fmt.Println("Structural patterns ---> Adapter")
	adapter.StartAdapterPattern()

	fmt.Println("Structural patterns ---> Decorator")
	decorator.StartDecoratorPattern()

	fmt.Println("Structural patterns ---> Facade")
	facade.StartFacadePattern()

	fmt.Println("Structural patterns ---> Bridge")
	bridge.StartBridgePattern()

	// BEHAVIORAL
	fmt.Println("Behavioral patterns ---> Strategy")
	strategy.StartStrategyPattern()

	fmt.Println("Behavioral patterns ---> State")
	state.StartStatePattern()

	fmt.Println("Behavioral patterns ---> Command")
	command.StartCommandPattern()

	fmt.Println("Behavioral patterns ---> Mediator")
	mediator.StartMediatorPattern()

}
