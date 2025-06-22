package main

import (
	"bufio"
	"fmt"
	"notes/note"
	"notes/todo"
	"os"
	"strings"
)

type saver interface {
	Save() error
}

type displayer interface {
	Display()
}

type outputtable interface {
	saver
	displayer
}

func main() {
	title, content, err := getNoteData()
	if err != nil {
		return
	}

	todoText := getTodoData()

	userNote, err := note.New(title, content)
	if err != nil {
		return
	}

	userTodo, err := todo.New(todoText)
	if err != nil {
		return
	}

	err = outputData(userNote)
	if err != nil {
		return
	}

	fmt.Println("note saved successfully")

	err = outputData(userTodo)
	if err != nil {
		return
	}

	fmt.Println("todo saved successfully")
}

func add[T int | float64 | string](a, b T) T {
	return a + b
}

func printSomething(value interface{}) {
	// interface{} === any
	intValue, ok := value.(int)
	if ok {
		fmt.Println(intValue + 1)
		fmt.Println("Integer: ", intValue)
		return
	}

	floatValue, ok := value.(float64)
	if ok {
		fmt.Println(floatValue + 1)
		fmt.Println("Integer: ", floatValue)
		return
	}

	stringValue, ok := value.(string)
	if ok {
		fmt.Println("Integer: ", stringValue)
		return
	}

	// alternative

	switch value.(type) {
	case int:
		fmt.Println("Integer: ", value)
	case float64:
		fmt.Println("Float: ", value)
	case string:
		fmt.Println("String: ", value)
	}

	fmt.Println(value)
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("failed to save data")
		return err
	}

	return nil
}

func getTodoData() string {
	text := getUserInput("Todo Text: ")

	return text
}

func getNoteData() (string, string, error) {
	title := getUserInput("Note Title: ")
	content := getUserInput("Note Content: ")

	return title, content, nil
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSpace(text)

	return text
}
