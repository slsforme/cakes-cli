package main

var items = []string{
	"Название",
	"Размер",
	"Вкус",
	"Декор",
	"Форма",
	"Количество",
	"Цена",
}

func createCakeMenu() {
	cursor := 5

	printMenu(items, "Создание торта", cursor)

	// type Cake struct {
	// 	Name   string
	// 	Size   string
	// 	Taste  string
	// 	Decor  []string
	// 	Form   Form
	// 	Amount int
	// 	Price  float32
	// }

}

func changeCakeMenu() {
	printBox("Изменение торта")

}

func deleteCakeMenu() {
	printBox("Удаление торта")

}

func createOrderMenu() {
	printBox("Создание заказа")

}
