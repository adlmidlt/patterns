package main

import (
	"fmt"
	"patterns/generative/builder"
	"patterns/generative/factory"
)

func main() {
	fmt.Println("Hello patterns!")

	fmt.Println("Generative patterns ---> Builder")
	builder.StartBuilderPattern()

	fmt.Println("Generative patterns ---> Factory")
	factory.StartFactoryPattern()
}
