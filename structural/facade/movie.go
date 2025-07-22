package facade

import "fmt"

// Projector отдельные компоненты домашней развлекательной системы.
type Projector struct{}

func (p *Projector) On()  { fmt.Println("Проектор включен.") }
func (p *Projector) Off() { fmt.Println("Проектор выключен.") }

type Screen struct{}

func (s *Screen) Down() { fmt.Println("Экран опущен.") }
func (s *Screen) Up()   { fmt.Println("Экран поднять.") }

type SoundSystem struct{}

func (ss *SoundSystem) On()  { fmt.Println("Звук включен") }
func (ss *SoundSystem) Off() { fmt.Println("Звук выключен") }
func (ss *SoundSystem) SetVolume(volume int) {
	fmt.Printf("Уровень громкости установлен на %d.\n", volume)
}

type Lighting struct{}

func (l *Lighting) DimLights(level int) {
	fmt.Printf("Освещение уменьшено до уровня %d.\n", level)
}

// HomeTheaterFacade фасад домашнего кинотеатра.
type HomeTheaterFacade struct {
	projector   *Projector
	screen      *Screen
	soundSystem *SoundSystem
	lighting    *Lighting
}

func NewHomeTheaterFacade(projector *Projector, screen *Screen, system *SoundSystem, lighting *Lighting) *HomeTheaterFacade {
	return &HomeTheaterFacade{projector, screen, system, lighting}
}

// WatchMovie метод фасада для начала просмотра фильма.
func (ht *HomeTheaterFacade) WatchMovie(movie string) {
	fmt.Println("Готовимся смотреть фильм:", movie)
	ht.lighting.DimLights(10)
	ht.screen.Down()
	ht.projector.On()
	ht.soundSystem.On()
	ht.soundSystem.SetVolume(7)
	fmt.Println("Фильм запущен")
}

func (ht *HomeTheaterFacade) EndMovie() {
	fmt.Println("Заканчиваем просмотр.")
	ht.projector.Off()
	ht.screen.Up()
	ht.soundSystem.Off()
	ht.lighting.DimLights(100)
	fmt.Println("Просмотр завершен.")
}

func StartFacadePattern() {
	projector := new(Projector)
	screen := new(Screen)
	soundSystem := new(SoundSystem)
	lighting := new(Lighting)

	facade := NewHomeTheaterFacade(projector, screen, soundSystem, lighting)

	facade.WatchMovie("Матрица")
	fmt.Println("------------------------")
	facade.EndMovie()
}
