package main

import (
	"bufio"
	"fmt"
	"github.com/wardove/noteProject/note"
	"github.com/wardove/noteProject/todo"
	"log"
	"os"
	"strings"
)

// if an interface has a single method, it is named by the method name plus the "er" suffix
type saver interface {
	Save() error
}

type outputable interface {
	saver
	Display()
}

func getUserInput(prompt string) string {

	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		panic(err)
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r") // for windows
	return text
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")
	return title, content
}

func saveData(data saver) error {
	return data.Save()
}

func outputData(data outputable) {
	data.Display()
}

func main() {

	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")
	userTodo := todo.New(todoText)
	userNote, err := note.New(title, content)

	if err != nil {
		log.Fatalln(err)
	}

	outputData(userNote)
	outputData(userTodo)

	err = saveData(userNote)

	if err != nil {
		log.Fatalln(err)
	}

	err = saveData(userTodo)

	if err != nil {
		log.Fatalln(err)
	}
}
