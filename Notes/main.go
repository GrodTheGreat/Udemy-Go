package main

import (
	"bufio"
	"fmt"
	"notes/note"
	"os"
	"strings"
)

func main() {
	title, content, err := getNoteData()
	if err != nil {
		return
	}

	userNote, err := note.New(title, content)
	if err != nil {
		return
	}
	userNote.Display()
	err = userNote.Save()
	if err != nil {
		fmt.Println("failed to save note")
		return
	}

	fmt.Println("note saved successfully")
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
