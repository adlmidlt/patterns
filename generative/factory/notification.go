package factory

import "fmt"

// Notifier 1. Общий интерфейс для всех продуктов.
type Notifier interface {
	Send(message string) error
}

// EmailNotifier 2. Конкретная реализация.
type EmailNotifier struct{}

func (n *EmailNotifier) Send(message string) error {
	fmt.Printf("Sending email: %s\n", message)

	return nil
}

type SMSNotifier struct{}

func (s *SMSNotifier) Send(message string) error {
	fmt.Printf("Sending SMS: %s\n", message)

	return nil
}

// CreateNotifier 3.Создание фабрики.
func CreateNotifier(notifierType string) (Notifier, error) {
	switch notifierType {
	case "email":
		return &EmailNotifier{}, nil
	case "sms":
		return &SMSNotifier{}, nil
	default:
		return nil, fmt.Errorf("not supported notifier type: %s", notifierType)
	}
}

func StartFactory1Pattern() {
	notifier, err := CreateNotifier("sms")
	if err != nil {
		fmt.Printf("Error creating notifier: %s\n", err)
	}
	err = notifier.Send("Hello!")
}
