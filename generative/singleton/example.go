package singleton

import (
	"fmt"
	"sync"
)

type ExampleSingleton struct{}

var (
	example *ExampleSingleton
	once    sync.Once // once гарантирует, что инициализация выполнится только один раз.
)

// GetExampleInstance возвращает единственный экземпляр Singleton.
func GetExampleInstance() *ExampleSingleton {
	once.Do(func() {
		example = &ExampleSingleton{}
	})

	return example
}

func StartSingleton1Pattern() {
	s1 := GetExampleInstance()
	s2 := GetExampleInstance()

	// проверяем что это один и тот же объект.
	fmt.Printf("Объекты равны: %+v\n", s1 == s2)
}
