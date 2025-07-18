package abstract

import "fmt"

type HumanCat interface {
	Voice()
}

type HumanDog interface {
	Voice()
}

type Cat struct{}

func (*Cat) Voice() {
	fmt.Println("Mao mao ...")
}

type Dog struct{}

func (*Dog) Voice() {
	fmt.Println("Gav gav ...")
}

type AnimalsFactory interface {
	CreateCat() HumanCat
	CreateDog() HumanDog
}

type HumanCatFactory struct{}

func (h *HumanCatFactory) CreateCat() HumanCat {
	return &Cat{}
}

type HumanDogFactory struct{}

func (h *HumanDogFactory) CreateDog() HumanDog {
	return &Dog{}
}

func StartAbstractExamplePattern() {
	catFactory := HumanCatFactory{}
	dogFactory := HumanDogFactory{}

	cat := catFactory.CreateCat()
	dog := dogFactory.CreateDog()

	cat.Voice()
	dog.Voice()
}
