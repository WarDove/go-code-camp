package todo

import (
	"encoding/json"
	"fmt"
	"os"
)

type Todo struct {
	Content string `json:"content,omitempty"`
}

func New(content string) *Todo {

	return &Todo{
		Content: content,
	}
}

func (t *Todo) Save() error {

	fileContent, err := json.Marshal(t)
	if err != nil {
		return err
	}
	return os.WriteFile("todo.json", []byte(fileContent), 0644)
}

func (t *Todo) Display() {
	fmt.Printf("Todo: %s\n", t.Content)
}
