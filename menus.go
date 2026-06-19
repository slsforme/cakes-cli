package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/eiannone/keyboard"
)

func buildItems(cake *Cake) []string {
	return []string{
		"Название: " + cake.Name,
		"Размер: " + cake.Size.Name,
		"Вкус: " + cake.Taste.Name,
		"Декор: " + decorString(cake.Decor),
		"Форма: " + string(cake.Form),
		"Количество: " + strconv.Itoa(cake.Amount),
		"Создать заказ",
		"Нажмите Escape, чтобы вернуться обратно",
	}
}

func buildEditItems(cake *Cake) []string {
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

func buildCakes() []string {
	var stringifiedCakes []string

	for i := 0; i < len(cakes); i++ {
		stringifiedCakes = append(stringifiedCakes, fmt.Sprintf("%d. Торт \"%s\"", i+1, cakes[i].Name))
	}

	return stringifiedCakes
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

func buildOrderItems(order *Order) []string {
	return []string{
		"Клиент: " + order.Customer,
		"Добавить торт в заказ",
		"Состав заказа: " + orderItemsString(order.Items),
		"Оформить заказ",
		"Нажмите Escape, чтобы вернуться обратно",
	}
}

func orderItemsString(items []*Cake) string {
	if len(items) == 0 {
		return "пусто"
	}
	names := make([]string, len(items))
	for i, c := range items {
		names[i] = c.Name
	}
	return strings.Join(names, ", ")
}

func createOrderMenu() {
	order := &Order{}

	for {
		items := buildOrderItems(order)

		selected := printMenu(items, "Создание заказа")

		if selected == -1 {
			return
		}

		if switchOrderItem(selected, order) {
			return
		}
	}
}

func switchOrderItem(selectedItemIndex int, order *Order) bool {
	switch selectedItemIndex {
	case 0:
		order.Customer = inputString("Введите имя клиента: ")

	case 1:
		cake := chooseCakeMenu()
		if cake != nil {
			order.Items = append(order.Items, cake)
		}

	case 2:
		// pass

	case 3:
		clearScreen()
		if order.IsComplete() {
			orders = append(orders, order)
			fmt.Println("Заказ оформлен!")
			printOrderSummary(order)
			fmt.Println("\nНажмите любую клавишу...")
			keyboard.GetKey()
			return true
		}
		fmt.Println("Заказ не готов:", strings.Join(order.MissingFields(), ", "))
		fmt.Println("\nНажмите любую клавишу...")
		keyboard.GetKey()
	}
	return false
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
	chosenCake := chooseCakeMenu()

	if chosenCake != nil {
		editCakeMenu(chosenCake)
	}
}

func deleteCakeMenu() {
	chosenCake := chooseCakeMenu()
	if chosenCake == nil {
		return
	}

	index := -1
	for i := range cakes {
		if cakes[i] == chosenCake {
			index = i
			break
		}
	}

	if index == -1 {
		return
	}

	clearScreen()
	fmt.Printf("Удалить торт \"%s\"?\n", chosenCake.Name)

	confirm := printMenu([]string{"Да, удалить", "Нет, отмена"}, "Подтверждение")

	if confirm == 0 {
		cakes = append(cakes[:index], cakes[index+1:]...)
		clearScreen()
		fmt.Println("Торт удалён.")
		fmt.Println("\nНажмите любую клавишу...")
		keyboard.GetKey()
	}
}

func chooseCakeMenu() *Cake {
	for {
		if len(cakes) == 0 {
			clearScreen()
			fmt.Println("Список тортов пуст.")
			fmt.Println("\nНажмите любую клавишу...")
			keyboard.GetKey()
			return nil
		}

		stringifiedCakes := buildCakes()
		stringifiedCakes = append(stringifiedCakes, "Нажмите Escape, чтобы вернуться обратно")

		selected := printMenu(stringifiedCakes, "Выбор торта")

		if selected == -1 || selected == len(cakes) {
			return nil
		}

		return cakes[selected]
	}
}

func editCakeMenu(cake *Cake) {
	for {
		items := buildEditItems(cake)

		selected := printMenu(items, "Изменение торта: "+cake.Name)

		if selected == -1 {
			return
		}

		switchItem(selected, cake)
	}
}
func switchItem(selectedItemIndex int, cake *Cake) {
	switch selectedItemIndex {
	case 0:
		name := inputString("Введите название для торта: ")
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
		clearScreen()

		if cake.IsComplete() {
			cakes = append(cakes, cake)
			fmt.Println("Заказ успешно создан!")
		} else {
			fmt.Println("Торт не был сделан до конца.")
		}

		keyboard.GetKey()
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
