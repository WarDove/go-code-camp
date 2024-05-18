package note

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

func New(title, content string) (*Note, error) {
	if title == "" || content == "" {
		return nil, errors.New("Invalid Input: Title and Content must not be empty.")
	}

	return &Note{
		title:     title,
		content:   content,
		createdAt: time.Now(),
	}, nil
}

func (n *Note) Display() {
	fmt.Printf("Title: %s\nContent: %s\nCreated At: %s\n", n.title, n.content,
		n.createdAt.Format("2006-01-02 15:04:05"))
}
