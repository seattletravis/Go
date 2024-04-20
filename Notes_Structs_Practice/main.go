package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}

// type displayer interface {
// 	Display() 
// }

type outputable interface{
	saver
	Display()
}

// type outputable interface {
// 	Save() error
// 	Display()
// }

func main(){
printSomething(1)
printSomething(1.5)
printSomething("Hello")

	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)

	if err != nil {
		return
	}

	outputData(userNote)
}

func printSomething(value interface{}){
	intVal, isTypeInt := value.(int)
	if  isTypeInt {
		fmt.Println("Integer: ", intVal)
		return
	}
	stringVal, isTypeString := value.(string)
	if isTypeString {
		fmt.Println("String: ", stringVal)
		return
	}
	floatVal, isTypeVal := value.(float64)
	if isTypeVal {
		fmt.Println("Float: ", floatVal)
		return
	}
	// switch value.(type){
	// case int:
	// 	fmt.Println("Integer:", value)
	// case float64:
	// 	fmt.Println("Float:", value)
	// case string:
	// 	fmt.Println(value)
	// }
}

func outputData(data outputable) error{
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return err
	}

	fmt.Println("Saving the note succeeded!")
	return nil
}


func getNoteData() (string, string) {
	title := getUserInput("Note title:")

	content := getUserInput("Note content:")

	return title, content

}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")


	return text

}