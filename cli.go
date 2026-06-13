package main

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/eiannone/keyboard"
)

var ArrowStarterCursor = 6

func clearScreen() {
	fmt.Print("\033[2J\033[H")
}

func moveArrowDown(cursor, max int) int {
	if cursor > ArrowStarterCursor+max {
		return cursor
	}

	moveTo(cursor, 1)
	fmt.Printf("  ")

	cursor++
	moveTo(cursor, 1)
	fmt.Print("->")
	return cursor
}

func moveArrowUp(cursor int) int {
	if cursor < ArrowStarterCursor {
		return cursor
	}

	moveTo(cursor, 1)
	fmt.Printf("  ")
	cursor--
	moveTo(cursor, 1)
	fmt.Printf("->")
	return cursor
}

func printBox(title string) {
	width := utf8.RuneCountInString(title)

	top := "╔" + strings.Repeat("=", width+5) + "╗"
	mid := "║  " + title + "  ║"
	low := "╚" + strings.Repeat("=", width+5) + "╝"

	fmt.Println(top)
	fmt.Println(mid)
	fmt.Println(low)
}

func printActions(actions []string) {

	for i := 0; i < len(actions); i++ {
		if i == 0 {
			actions[i] = "-> " + actions[i]
			continue
		}

		actions[i] = "   " + actions[i]
	}

	result := strings.Join(actions, "\n")

	fmt.Println("Выберите действие.")
	fmt.Println(result)
}

func moveTo(row, col int) {
	fmt.Printf("\033[%d;%dH", row, col)
}

func printMenu(actions []string, title string, cursor int) {
	clearScreen()
	printBox(title)
	printActions(actions)

	for {
		_, key, err := keyboard.GetKey()

		if err != nil {
			fmt.Println(err)
		}

		if key == keyboard.KeyArrowDown {
			cursor = moveArrowDown(cursor, len(actions)/2-1)
			continue
		}

		if key == keyboard.KeyArrowUp {
			cursor = moveArrowUp(cursor)
			continue
		}

		if key == keyboard.KeyEnter {
			clearScreen()
			switch cursor - len(actions) {
			case 0:
				createCakeMenu()
			case 1:
				changeCakeMenu()
			case 2:
				deleteCakeMenu()
			case 3:
				createOrderMenu()
			}
			return
		}

		if key == keyboard.KeyEsc {
			break
		}
	}
}

func MainMenu() {
	title := "Главное меню"
	var actions = []string{
		"1. Создать торт",
		"2. Изменить торт",
		"3. Удалить торт",
		"4. Сделать заказ",
		"Нажмите Escape, чтобы выйти",
	}

	cursor := 5

	printMenu(actions, title, cursor)
}
