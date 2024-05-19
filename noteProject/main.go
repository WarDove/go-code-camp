package main

import (
	"bufio"
	"fmt"
	"github.com/wardove/noteProject/note"
	"log"
	"os"
	"strings"
)

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

func getNodeData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")
	return title, content
}

func main() {

	title, content := getNodeData()

	userNote, err := note.New(title, content)

	if err != nil {
		log.Fatalln(err)
	}

	userNote.Display()
	userNoteJson, err := os.OpenFile("userNote.json", os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalln(err)
	}

	defer userNoteJson.Close()

	err = userNote.Encoder(userNoteJson)

	if err != nil {
		log.Fatalln(err)
	}
}
