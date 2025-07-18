package factory

import "fmt"

// Transport общий интерфейс транспортного средства.
type Transport interface {
	deliver() string
}

// Truck грузовик.
type Truck struct{}

func (t *Truck) deliver() string {
	return "Доставка грузовиком"
}

// Bike велосипед.
type Bike struct{}

func (b *Bike) deliver() string {
	return "Доставка велосипедом"
}

// Bus пример добавления нового транспорта.
type Bus struct{}

func (b *Bus) deliver() string {
	return "Доставка автобусом"
}

// TransportFactory интерфейс фабрики.
type TransportFactory interface {
	createTransport() Transport
}

// TruckFactory "фабрика" грузовика.
type TruckFactory struct{}

func (t *TruckFactory) createTransport() Transport {
	return &Truck{}
}

// BikeFactory "фабрика" велосипеда.
type BikeFactory struct{}

func (b *BikeFactory) createTransport() Transport {
	return &Bike{}
}

type BusFactory struct{}

func (b *BusFactory) createTransport() Transport {
	return &Bus{}
}

func StartFactoryPattern() {
	truckFactory := &TruckFactory{}
	bikeFactory := &BikeFactory{}
	busFactory := &BusFactory{}

	truck := truckFactory.createTransport()
	bike := bikeFactory.createTransport()
	bus := busFactory.createTransport()

	fmt.Printf("%s\n", truck.deliver())
	fmt.Printf("%s\n", bike.deliver())
	fmt.Printf("%s\n", bus.deliver())
}
