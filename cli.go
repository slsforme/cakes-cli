package main

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/eiannone/keyboard"
)

func clearScreen() {
	fmt.Print("\033[2J\033[H")
}

func printBox(title string) {
	width := utf8.RuneCountInString(title)
	border := strings.Repeat("=", width+4)

	fmt.Println("╔" + border + "╗")
	fmt.Println("║  " + title + "  ║")
	fmt.Println("╚" + border + "╝")
}

func render(actions []string, title string, selected int) {
	clearScreen()
	printBox(title)
	fmt.Println("Выберите действие.")

	for i, a := range actions {
		if i == selected {
			fmt.Println("-> " + a)
		} else {
			fmt.Println("   " + a)
		}
	}
}

func printMenu(actions []string, title string) int {
	selected := 0

	for {
		render(actions, title, selected)

		_, key, err := keyboard.GetKey()
		if err != nil {
			fmt.Println(err)
		}

		switch key {
		case keyboard.KeyArrowDown:
			if selected < len(actions)-1 {
				selected++
			}
		case keyboard.KeyArrowUp:
			if selected > 0 {
				selected--
			}
		case keyboard.KeyEnter:
			return selected
		case keyboard.KeyEsc:
			return -1
		}
	}
}

func MainMenu() {
	actions := []string{
		"1. Создать торт",
		"2. Изменить торт",
		"3. Удалить торт",
		"4. Сделать заказ",
		"Нажмите Escape, чтобы выйти",
	}

	for {
		selected := printMenu(actions, "Главное меню")

		switch selected {
		case 0:
			createCakeMenu()
		case 1:
			changeCakeMenu()
		case 2:
			deleteCakeMenu()
		case 3:
			createOrderMenu()
		case -1, 4:
			clearScreen()
			fmt.Println("До свидания!")
			return
		}
	}
}
