package main

import (
	"log"
)

type Product struct {
	sku       string
	unitPrice int
}

func getProductList() []Product {
	return []Product{
		{sku: "A", unitPrice: 20},
		{sku: "B", unitPrice: 15},
		{sku: "C", unitPrice: 50},
	}
}

func addProductToBasket(option int) Product {
	products := getProductList()
	defer func() {
		if err := recover(); err != nil {
			log.Println("Invalid item cannot be added to basket.")
		}
	}()

	product := products[option-1]

	return product
}

func calculateCheckoutTotal(productPrice int, checkoutTotal int) int {
	checkoutTotal += productPrice
	return checkoutTotal
}
