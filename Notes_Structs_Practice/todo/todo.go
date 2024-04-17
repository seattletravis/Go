package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Note struct {
	Content string `json:"text"`
}

func (todo Todo) Display() {
	fmt.Printf(todo)
}

func (todo Todo) Save() error {
	fileName := "todo.json"

	json, err := json.Marshal(todo)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("invalid input")
	}

	return Note{
		Text: content,
	}, nil
} 
