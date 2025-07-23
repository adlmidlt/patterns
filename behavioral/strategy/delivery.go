package strategy

import "fmt"

// DeliveryStrategy общий интерфейс стратегии доставки.
type DeliveryStrategy interface {
	CalculateCost(weight float64) float64
	GetDescription() string
}

// StandardDelivery стандартная доставка.
type StandardDelivery struct{}

func (s *StandardDelivery) CalculateCost(weight float64) float64 {
	if weight <= 1 {
		return 100
	}

	return 100 + (weight-1)*50
}

func (s *StandardDelivery) GetDescription() string {
	return "Стандартная доставка"
}

// ExpressDelivery быстрая доставка.
type ExpressDelivery struct{}

func (e *ExpressDelivery) CalculateCost(weight float64) float64 {
	if weight <= 1 {
		return 200
	}

	return 200 + (weight-1)*80
}
func (e *ExpressDelivery) GetDescription() string {
	return "Быстрая  доставка"
}

// PickupDelivery самовывоз.
type PickupDelivery struct{}

func (p *PickupDelivery) CalculateCost(weight float64) float64 {
	return 0
}

func (p *PickupDelivery) GetDescription() string {
	return "Самовывоз"
}

// OrderService контекст использования стратегии.
type OrderService struct {
	strategy DeliveryStrategy
}

func NewOrderService(strategy DeliveryStrategy) *OrderService {
	return &OrderService{strategy: strategy}
}

// CalculateShippingCost рассчитать стоимость доставки.
func (o *OrderService) CalculateShippingCost(weight float64) float64 {
	return o.strategy.CalculateCost(weight)
}

func (o *OrderService) GetDeliveryMethod() string {
	return o.strategy.GetDescription()
}

func StartStrategyPattern() {
	orderWeight := 2.5

	standardDelivery := new(StandardDelivery)
	expressDelivery := new(ExpressDelivery)
	pickupDelivery := new(PickupDelivery)

	osStandard := NewOrderService(standardDelivery)
	fmt.Printf("Метод доставки: %s, стоимость: %.2f рублей\n", osStandard.GetDeliveryMethod(), osStandard.CalculateShippingCost(orderWeight))

	osExpress := NewOrderService(expressDelivery)
	fmt.Printf("Метод доставки: %s, стоимость: %.2f рублей\n", osExpress.GetDeliveryMethod(), osExpress.CalculateShippingCost(orderWeight))

	osPickup := NewOrderService(pickupDelivery)
	fmt.Printf("Метод доставки: %s, стоимость: %.2f рублей\n", osPickup.GetDeliveryMethod(), osPickup.CalculateShippingCost(orderWeight))
}
