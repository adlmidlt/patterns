package adapter

import "fmt"

// OldService старый сервис, который имеет неподходящий нас интерфейс.
type OldService interface {
	GetOldData() string
}

// OldServiceImpl структура, реализующая старый интерфейс.
type OldServiceImpl struct{}

func (o *OldServiceImpl) GetOldData() string {
	return "Данные в старом формате"
}

// NewService новый интерфейс, который мы хотим поддерживать.
type NewService interface {
	FetchData() string
}

// ServiceAdapter адаптер для соединения старого интерфейса в новый.
type ServiceAdapter struct {
	oldService OldService
}

// FetchData адаптер преобразующий старый интерфейс в новый.
func (sa *ServiceAdapter) FetchData() string {
	data := sa.oldService.GetOldData()

	return fmt.Sprintf("Преобразование значения: %s", data)
}

func StartAdapterPattern() {
	// Создаем объект старого сервиса.
	oldService := new(OldServiceImpl)

	// Создаем адаптер, для связывания старого и нового интерфейса.
	adapter := &ServiceAdapter{oldService}

	// Получение данных через новый интерфейс.
	result := adapter.FetchData()
	fmt.Println(result)
}
