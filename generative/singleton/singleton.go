package singleton

import "fmt"

// Logger сущность синглтона, хранит общие настройки и методы логирования.
type Logger struct {
	level string
}

var instance *Logger

func GetInstance() *Logger {
	if instance == nil {
		instance = &Logger{"INFO"}
		fmt.Println("Создан новый экземпляр")

		return instance
	}
	fmt.Println("Возвращен существующий экземпляр")

	return instance
}

func (l *Logger) Log(level string) {
	fmt.Printf("[%s]: %s\n", l.level, level)
}

func StartSingletonPattern() {
	logger := GetInstance()
	logger.Log("Это первое сообщение")

	anotherLogger := GetInstance()
	anotherLogger.Log("Это второе сообщение")

	// Изменение уровня логирования влияет на оба экземпляра, потому что один и тот же объект.
	anotherLogger.level = "DEBUG"
	logger.Log("Это третье сообщение")
}
