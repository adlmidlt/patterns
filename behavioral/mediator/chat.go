package mediator

import "fmt"

// Приложение чата, где пользователи обмениваются сообщениями. Без медиатора каждый участник знал бы о существовании
// остальных участников и общался бы с ними напрямую. Медиатор позволяет участникам отправлять сообщения через
// центральную точку (чат-комнату), тем самым устраняя прямые зависимости.

// ChatRoom общие методы для общения.
type ChatRoom interface {
	SendMessage(sender, message string)
	AddParticipant(participant Participant)
}

// Participant участник чата.
type Participant interface {
	SetChatRoom(chatRoom ChatRoom)
	Send(message string)
	Receive(sender, message string)
	String() string
}

// chatRoomImpl реализация комнаты чата.
type chatRoomImpl struct {
	participants map[string]Participant
}

func NewChatRoom() ChatRoom {
	return &chatRoomImpl{
		participants: make(map[string]Participant),
	}
}

func (cr *chatRoomImpl) SendMessage(sender, message string) {
	for name, participant := range cr.participants {
		if name != sender {
			participant.Receive(sender, message)
		}
	}
}

func (cr *chatRoomImpl) AddParticipant(participant Participant) {
	cr.participants[participant.String()] = participant
}

type Human struct {
	name     string
	chatRoom ChatRoom
}

func NewHuman(name string) *Human {
	return &Human{name: name}
}

func (h *Human) String() string {
	return h.name
}

func (h *Human) SetChatRoom(chatRoom ChatRoom) {
	h.chatRoom = chatRoom
}

func (h *Human) Send(message string) {
	h.chatRoom.SendMessage(h.name, message)
}

func (h *Human) Receive(sender, message string) {
	fmt.Printf("%s получил сообщение от %s: %s\n", h.name, sender, message)
}

type Bot struct {
	name     string
	chatRoom ChatRoom
}

func NewBoot(name string) *Bot {
	return &Bot{name: name}
}

func (b *Bot) String() string {
	return b.name
}

func (b *Bot) SetChatRoom(chatRoom ChatRoom) {
	b.chatRoom = chatRoom
}

func (b *Bot) Send(message string) {
	b.chatRoom.SendMessage(b.name, message)
}

func (b *Bot) Receive(sender, message string) {
	fmt.Printf("%s получил сообщение от %s: %s\n", b.name, sender, message)
}

func StartMediatorPattern() {
	room := NewChatRoom()

	user1 := NewHuman("User1")
	user2 := NewHuman("User2")
	bot := NewBoot("Support")

	room.AddParticipant(user1)
	room.AddParticipant(user2)
	room.AddParticipant(bot)

	user1.SetChatRoom(room)
	user2.SetChatRoom(room)
	bot.SetChatRoom(room)

	user1.Send("Привет всем!")
	user2.Send("Привет User1!")
	bot.Send("Привет, ребят!")
}
