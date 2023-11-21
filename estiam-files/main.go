// main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"estiam/dictionary"
	"strings"
)

func main() {
	d, err := dictionary.New("dictionary.txt")
	if err != nil {
		fmt.Println("Error creating dictionary:", err)
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter action (add/define/remove/list/exit): ")
		action, _ := reader.ReadString('\n')
		action = strings.TrimSpace(action)

		switch action {
		case "add":
			actionAdd(d, reader)
		case "define":
			actionDefine(d, reader)
		case "remove":
			actionRemove(d, reader)
		case "list":
			actionList(d)
		case "exit":
			err := d.SaveToFile()
			if err != nil {
				fmt.Println("Error saving dictionary:", err)
			}
			os.Exit(0)
		default:
			fmt.Println("Invalid action. Please try again.")
		}
	}
}

func actionAdd(d *dictionary.Dictionary, reader *bufio.Reader) {
	fmt.Print("Enter word: ")
	word, _ := reader.ReadString('\n')
	fmt.Print("Enter definition: ")
	definition, _ := reader.ReadString('\n')

	d.Add(word, definition)
	fmt.Println("Word added successfully.")
}

func actionDefine(d *dictionary.Dictionary, reader *bufio.Reader) {
	fmt.Print("Enter word: ")
	word, _ := reader.ReadString('\n')

	entry, err := d.Get(word)
	if err != nil {
		fmt.Println("Word not found.")
		return
	}

	fmt.Printf("Definition: %s\n", entry)

}

func actionRemove(d *dictionary.Dictionary, reader *bufio.Reader) {
	fmt.Print("Enter word: ")
	word, _ := reader.ReadString('\n')

	d.Remove(word)
	fmt.Println("Word removed successfully.")
}

func actionList(d *dictionary.Dictionary) ([]string, map[string]dictionary.Entry) {
	wordList, entries := d.List()
	if len(wordList) == 0 {
		fmt.Println("Dictionary is empty.")
		return wordList, entries
	}

	fmt.Println("Word list:")
	for _, word := range wordList {
		fmt.Printf("%s: %s", word, entries[word])
	}
	return wordList, entries
}
