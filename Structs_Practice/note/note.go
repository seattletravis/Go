package Note

import (
	"time"
)

type Note struct {
	title string
	content string
	createdAt time.Time
}

func New(title, content string) Note {
	return Note{
		title: title,
		content: content,
		createdAt: time.Now(),
	}

} 
