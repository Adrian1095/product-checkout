package main

import (
	"log"
)

type Discount struct {
	itemCountForDiscount int
	discount             int
}

type Product struct {
	id        int
	sku       string
	unitPrice int
	discount  Discount
}

var products = []Product{
	{id: 1, sku: "A", unitPrice: 20, discount: Discount{itemCountForDiscount: 3, discount: -5}},
	{id: 2, sku: "B", unitPrice: 15, discount: Discount{itemCountForDiscount: 2, discount: -10}},
	{id: 3, sku: "C", unitPrice: 50, discount: Discount{itemCountForDiscount: 3, discount: -30}},
}

func addProductToBasket(option int) Product {
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

func addDiscount(products []Product, itemCountForDeal int) bool {
	productCount := len(products)

	return productCount%itemCountForDeal == 0
}

// func editProductDiscount(product Product, discount Discount) {
// 	product.discount = discount
// }
