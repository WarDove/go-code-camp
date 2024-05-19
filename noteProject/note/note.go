package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"
)

type Note struct {
	Title     string
	Content   string
	CreatedAt time.Time
}

func New(title, content string) (*Note, error) {
	if title == "" || content == "" {
		return nil, errors.New("Invalid Input: Title and Content must not be empty.")
	}

	return &Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (n *Note) Encoder(w io.Writer) error {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(n)

	if err != nil {
		return err
	}

	return nil
}

func (n *Note) Display() {
	fmt.Printf("Title: %s\nContent: %s\nCreated At: %s\n", n.Title, n.Content,
		n.CreatedAt.Format("2006-01-02 15:04:05"))
}
