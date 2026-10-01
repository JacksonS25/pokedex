package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		cleanedText := cleanInput(text)
		switch cleanedText[0] {
		case "exit":
			err := commands["exit"].callback()
			if err != nil {
				fmt.Println("Error:", err)
			}
		case "help":
			err := commands["help"].callback()
			if err != nil {
				fmt.Println("Error:", err)
			}
		default:
			fmt.Println("Unkown command")
		}
	}
}
