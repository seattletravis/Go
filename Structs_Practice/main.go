package main

import (
	"fmt"
	"example.com/note/note"
)

func main(){
	title, content := getNoteData()

	userNote, err := Note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()

}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")

	content := getUserInput("Note content: ")

	return title, content

}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	var value string
	fmt.Scanln(&value)



	return value

}