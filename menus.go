package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/eiannone/keyboard"
)

var items = []string{
	"Название",
	"Размер",
	"Вкус",
	"Декор",
	"Форма",
	"Количество",
	"Создать заказ",
	"Нажмите Escape, чтобы вернуться обратно",
}

func buildItems(cake *Cake) []string {
	return []string{
		"Название: " + cake.Name,
		"Размер: " + cake.Size.Name,
		"Вкус: " + cake.Taste.Name,
		"Декор: " + decorString(cake.Decor),
		"Форма: " + string(cake.Form),
		"Количество: " + strconv.Itoa(cake.Amount),
		"Нажмите Escape, чтобы вернуться обратно",
	}
}

func decorString(decor []Compound) string {
	if len(decor) == 0 {
		return ""
	}
	names := make([]string, len(decor))
	for i, d := range decor {
		names[i] = d.Name
	}
	return strings.Join(names, ", ")
}

func createCakeMenu() {
	cake := &Cake{}

	for {
		items := buildItems(cake)

		selected := printMenu(items, "Создание торта")

		if selected == -1 {
			return
		}

		switchItem(selected, cake)
	}
}

func changeCakeMenu() {
	printMenu([]string{"Заглушка", "Назад"}, "Изменение торта")
}

func deleteCakeMenu() {
	printMenu([]string{"Заглушка", "Назад"}, "Удаление торта")
}

func createOrderMenu() {
	printMenu([]string{"Заглушка", "Назад"}, "Создание заказа")
}

func switchItem(selectedItemIndex int, cake *Cake) {
	switch selectedItemIndex {
	case 0:
		name := inputString("Введите название для торта")
		cake.Name = name
	case 1:
		id := printMenu(getStringifiedData(AvailableSizes), "Выбор размера")
		cake.Size = AvailableSizes[id]
	case 2:
		id := printMenu(getStringifiedData(AvailableTastes), "Выбор вкуса")
		cake.Taste = AvailableTastes[id]
	case 3:
		id := printMenu(getStringifiedData(AvailableDecor), "Выбор декора")
		cake.Decor = []Compound{AvailableDecor[id]} // TODO: тут можно будет выбрать несколько вариантов декора
	case 4:
		id := printMenu(getStringifiedData(AllForms()), "Выбор формы")
		cake.Form = AllForms()[id]
	case 5:
		amount := inputInt("Введите необходимое количество тортов: ")
		cake.Amount = amount
	case 6:
		// TODO: реализовать кнопку для создания заказа - торта
	}
}

func getStringifiedData(i interface{}) []string {
	var actions []string

	fmt.Printf("%T", i)

	switch val := i.(type) {
	case []Compound:
		for _, value := range val {
			actions = append(actions, fmt.Sprintf("%s - %.2f рублей", value.Name, value.Price))
		}
	case []Form:
		for _, name := range val {
			actions = append(actions, string(name))
		}
	default:
		return []string{"Информация по данной категории отсутствует"}
	}

	return actions
}

func inputString(clue string) string {
	keyboard.Close()
	defer keyboard.Open()
	fmt.Print(clue)
	var input string
	fmt.Scan(&input)
	return input
}

func inputInt(clue string) int {
	keyboard.Close()
	defer keyboard.Open()
	for {
		fmt.Print(clue)

		var input string
		fmt.Scan(&input)

		n, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(ErrNotNumber)
			continue
		}

		if n < 0 {
			fmt.Printf("%s\n", ErrOutOfRange)
			continue
		}

		return n
	}
}
