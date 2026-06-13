package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

func main() {
	if err := keyboard.Open(); err != nil {
		fmt.Println(err)
	}
	defer keyboard.Close()

	fmt.Print("\033[2J\033[H")

	MainMenu()
}
