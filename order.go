package main

import "fmt"

type Order struct {
	Customer string  `json:"customer"`
	Items    []*Cake `json:"items"`
}

var orders []*Order

func (o *Order) IsComplete() bool {
	return o.Customer != "" && len(o.Items) > 0
}

func (o *Order) MissingFields() []string {
	var missing []string
	if o.Customer == "" {
		missing = append(missing, "имя клиента")
	}
	if len(o.Items) == 0 {
		missing = append(missing, "хотя бы один торт")
	}
	return missing
}

func (o *Order) TotalPrice() float32 {
	var total float32
	for _, cake := range o.Items {
		total += cakePrice(cake)
	}
	return total
}

func cakePrice(c *Cake) float32 {
	price := c.Size.Price + c.Taste.Price
	for _, d := range c.Decor {
		price += d.Price
	}
	return price * float32(c.Amount)
}

func printOrderSummary(o *Order) {
	fmt.Printf("\nЗаказ для: %s\n", o.Customer)
	fmt.Println("Торты:")
	for i, c := range o.Items {
		fmt.Printf("  %d. %s — %.2f руб\n", i+1, c.Name, cakePrice(c))
	}
	fmt.Printf("Итого: %.2f руб\n", o.TotalPrice())
}

func saveOrdersToJSON() {

}
