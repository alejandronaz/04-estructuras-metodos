package main

import "fmt"

type Product struct {
	ID          string
	Name        string
	Price       float64
	Description string
	Category    string
}

var products = []Product{
	{
		ID:          "1",
		Name:        "Laptop",
		Price:       1000,
		Description: "MacBook Pro",
		Category:    "Technology",
	},
	{
		ID:          "2",
		Name:        "Mouse",
		Price:       20,
		Description: "Logitech",
		Category:    "Technology",
	},
	{
		ID:          "3",
		Name:        "Keyboard",
		Price:       40,
		Description: "Logitech",
		Category:    "Technology",
	},
	{
		ID:          "4",
		Name:        "Monitor",
		Price:       200,
		Description: "Samsung",
		Category:    "Technology",
	},
}

func (prod Product) Save() {
	products = append(products, prod)
}

func (prod Product) GetAll() {
	for _, prod := range products {
		fmt.Println(prod)
	}
}

func getById(idProd int) Product {
	return products[idProd]
}

func main() {
	prod := Product{
		ID:          "5",
		Name:        "Keyboard",
		Price:       40,
		Description: "Logitech",
		Category:    "Technology",
	}
	prod.Save()
	prod.GetAll()
	fmt.Println(getById(2))
}
