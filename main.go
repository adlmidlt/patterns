package main

import (
	"fmt"
	"patterns/generative/abstract"
	"patterns/generative/builder"
	"patterns/generative/factory"
)

func main() {
	fmt.Println("Hello patterns!")

	fmt.Println("Generative patterns ---> Builder")
	builder.StartBuilderPattern()

	fmt.Println("Generative patterns ---> Factory")
	factory.StartFactoryPattern()

	fmt.Println("Generative patterns ---> Abstract")
	abstract.StartAbstractPattern()
	fmt.Println("Generative patterns ---> Abstract ---> Example")
	abstract.StartAbstractExamplePattern()
}
