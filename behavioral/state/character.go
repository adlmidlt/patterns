package state

import (
	"fmt"
)

const maxHPWarrior = 100

// CharacterState интерфейс состояния персонажа.
type CharacterState interface {
	OnAttack(character *Character)
}

type Character struct {
	name           string
	healthPoints   int
	state          CharacterState
	isDead         bool
	maxHealthPoint int
}

// NewCharacter создание нового персонажа.
func NewCharacter(name string, maxHealthPoint int) *Character {
	character := &Character{name: name, healthPoints: maxHealthPoint, maxHealthPoint: maxHealthPoint}
	character.setState(&RestingState{})
	return character
}

func (ch *Character) setState(state CharacterState) {
	ch.state = state
	fmt.Printf("%s теперь находится в состоянии: %T\n", ch.name, state)
}

// Attack атака противника.
func (ch *Character) Attack() {
	if !ch.isDead {
		ch.state.OnAttack(ch)
	} else {
		fmt.Printf("%s мертв и не может атаковать!\n", ch.name)
	}
}

// RestingState поведение отдыха.
type RestingState struct{}

func (r *RestingState) OnAttack(ch *Character) {
	fmt.Printf("%s бросился в бой! (%v HP)\n", ch.name, ch.healthPoints)
	ch.setState(&FightingState{})
}

// FightingState поведение боя.
type FightingState struct{}

func (f *FightingState) OnAttack(ch *Character) {
	fmt.Printf("%s контратаковал врага! (%v HP)\n", ch.name, ch.healthPoints)
}

// WoundedState раненое состояние.
type WoundedState struct{}

func (w *WoundedState) OnAttack(ch *Character) {
	fmt.Printf("%s ранен и не может атаковать эффективно... (%v HP)\n", ch.name, ch.healthPoints)
}

func StartStatePattern() {
	player := NewCharacter("Воин", maxHPWarrior)

	for i := 0; i < 10; i++ {
		player.Attack()

		switch player.healthPoints {
		case 50:
			player.setState(&WoundedState{}) // Переход в раненое состояние
			fallthrough
		default:
			player.healthPoints -= 10 // Получил повреждения
		}

		if player.healthPoints <= 0 {
			player.isDead = true
			break
		}
	}
	fmt.Println("Игра закончена.")
}
