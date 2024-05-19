package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
}

func New(title, content string) (*Note, error) {
	if title == "" || content == "" {
		return nil, errors.New("invalid input: Title and Content must not be empty")
	}

	return &Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (n *Note) Save() error {

	fileName := strings.ReplaceAll(n.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	fileContent, err := json.Marshal(n)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, []byte(fileContent), 0644)
}

func (n *Note) Display() {
	fmt.Printf("Title: %s\nContent: %s\nCreated At: %s\n", n.Title, n.Content,
		n.CreatedAt.Format("2006-01-02 15:04:05"))
}
