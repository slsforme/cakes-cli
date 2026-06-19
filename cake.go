package main

import "fmt"

type Form string

const (
	Round       Form = "Круглый"
	Square      Form = "Квадратный"
	HeartShaped Form = "Сердце"
	Rectangular Form = "Прямоугольный"
)

type Compound struct {
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

type Cake struct {
	Name       string     `json:"name"`
	Size       Compound   `json:"size"`
	Taste      Compound   `json:"taste"`
	Decor      []Compound `json:"decor"`
	Form       Form       `json:"form"`
	Amount     int        `json:"amount"`
	totalPrice float32
}

var cakes []*Cake

func CreateCake(name string, size, taste Compound, decor []Compound, form Form, amount int) *Cake {
	cake := &Cake{
		Name:   name,
		Size:   size,
		Taste:  taste,
		Decor:  decor,
		Form:   form,
		Amount: amount,
	}

	cake.calculateTotalPrice()
	cakes = append(cakes, cake)

	return cake
}

func (c *Cake) calculateTotalPrice() {
	total := c.Size.Price + c.Taste.Price
	for _, d := range c.Decor {
		total += d.Price
	}
	c.totalPrice = total * float32(c.Amount)
}

func (c *Cake) TotalPrice() float32 {
	return c.totalPrice
}

func FindCake(name string) int {
	for i := 0; i < len(cakes); i++ {
		if cakes[i].Name == name {
			return i
		}
	}
	return -1
}

func DeleteCake(index int) int {
	if index < 0 || index >= len(cakes) {
		return -1
	}

	cakes = append(cakes[:index], cakes[index+1:]...)

	return index
}

func (c *Cake) ChangeName(name string) {
	c.Name = name
}

func (c *Cake) ChangeSize(size Compound) {
	c.Size = size
	c.calculateTotalPrice()
}

func (c *Cake) ChangeTaste(taste Compound) {
	c.Taste = taste
	c.calculateTotalPrice()
}

func (c *Cake) ChangeDecor(decor []Compound) {
	c.Decor = decor
	c.calculateTotalPrice()
}

func (c *Cake) ChangeForm(form Form) {
	c.Form = form
}

func (c *Cake) ChangeAmount(amount int) {
	c.Amount = amount
	c.calculateTotalPrice()
}

func (c *Cake) IsComplete() bool {
	return c.Name != "" &&
		c.Size.Name != "" &&
		c.Taste.Name != "" &&
		len(c.Decor) > 0 &&
		c.Form != "" &&
		c.Amount > 0
}

func (c *Cake) StringifiedCake() string {
	return fmt.Sprintf(
		"Название: %s\nРазмер: %s\nВкус: %s\nДекор: %s\nФорма: %s\nКоличество: %d",
		c.Name,
		c.Size.Name,
		c.Taste.Name,
		decorString(c.Decor),
		string(c.Form),
		c.Amount,
	)
}

var AvailableSizes = []Compound{
	{Name: "Маленький", Price: 500},
	{Name: "Средний", Price: 800},
	{Name: "Большой", Price: 1200},
}

var AvailableTastes = []Compound{
	{Name: "Шоколад", Price: 200},
	{Name: "Ваниль", Price: 150},
	{Name: "Клубника", Price: 250},
}

var AvailableDecor = []Compound{
	{Name: "Розы", Price: 100},
	{Name: "Шарики", Price: 50},
	{Name: "Свечи", Price: 30},
	{Name: "Ленты", Price: 70},
	{Name: "Фигурки", Price: 150},
}

func AllForms() []Form {
	return []Form{Round, Square, HeartShaped, Rectangular}
}
