package command

import "fmt"

// Задача: Есть простая текстовая редакторская команда: вставка текста, удаление текста и отмена последней операции.
// Каждая операция оформляется как отдельная команда.

const positionInText = 6

// Command набор команд интерфейса.
type Command interface {
	Execute()
	Undo()
}

// TextEditor редактор текста.
type TextEditor struct {
	text string
}

// InsertTextCommand команда вставки текста.
type InsertTextCommand struct {
	editor       *TextEditor
	position     int
	insertedText string
	previousText string
}

func NewInsertTextCommand(editor *TextEditor, position int, text string) *InsertTextCommand {
	return &InsertTextCommand{
		editor:       editor,
		position:     position,
		insertedText: text,
	}
}

func (cmd *InsertTextCommand) Execute() {
	cmd.previousText = cmd.editor.text
	firstPart := cmd.editor.text[:cmd.position]
	secondPart := cmd.editor.text[cmd.position:]
	cmd.editor.text = firstPart + cmd.insertedText + secondPart
	fmt.Printf("Текст после вставки: %q\n", cmd.editor.text)
}

func (cmd *InsertTextCommand) Undo() {
	cmd.editor.text = cmd.previousText
	fmt.Printf("Текст после отмены: %q\n", cmd.editor.text)
}

// DeleteTextCommand команда для удаления текста.
type DeleteTextCommand struct {
	editor        *TextEditor
	startPosition int // startPosition стартовая позиция в тексте
	endPosition   int // endPosition конечная позиция в тексте
	previousText  string
}

func NewDeleteTextCommand(editor *TextEditor, startPosition, endPosition int) *DeleteTextCommand {
	return &DeleteTextCommand{
		editor:        editor,
		startPosition: startPosition,
		endPosition:   endPosition,
	}
}

func (cmd *DeleteTextCommand) Execute() {
	cmd.previousText = cmd.editor.text
	cmd.editor.text = cmd.editor.text[:cmd.startPosition] + cmd.editor.text[cmd.endPosition:]
	fmt.Printf("Текст после удаления: %q\n", cmd.editor.text)
}

func (cmd *DeleteTextCommand) Undo() {
	cmd.editor.text = cmd.previousText
	fmt.Printf("Текст после отмены: %q\n", cmd.editor.text)
}

// History история команд.
type History struct {
	commands []Command
}

func (h *History) Push(cmd Command) {
	h.commands = append(h.commands, cmd)
}

func (h *History) Pop() Command {
	if len(h.commands) > 0 {
		lastIndex := len(h.commands) - 1
		cmd := h.commands[lastIndex]
		h.commands = h.commands[:lastIndex]

		return cmd
	}

	return nil
}

type EditorUI struct {
	editor  *TextEditor
	history *History
}

func NewEditorUI(editor *TextEditor) *EditorUI {
	return &EditorUI{
		editor:  editor,
		history: &History{},
	}
}

func (ui *EditorUI) Perform(command Command) {
	command.Execute()
	ui.history.Push(command)
}

func (ui *EditorUI) UndoLastAction() {
	cmd := ui.history.Pop()
	if cmd != nil {
		cmd.Undo()
	}
}

func StartCommandPattern() {
	editor := &TextEditor{"Hello World"}
	ui := NewEditorUI(editor)

	ui.Perform(NewInsertTextCommand(editor, 6, "Golang "))
	ui.Perform(NewDeleteTextCommand(editor, 6, 13))
	ui.UndoLastAction()
	ui.UndoLastAction()
}
