package main

import (
	"log"
)

type Discount struct {
	itemCountForDiscount int
	discount             int
}

type Product struct {
	sku       string
	unitPrice int
	discount  Discount
}

func getProductList() []Product {
	return []Product{
		{sku: "A", unitPrice: 20, discount: Discount{itemCountForDiscount: 3, discount: -5}},
		{sku: "B", unitPrice: 15, discount: Discount{itemCountForDiscount: 2, discount: -10}},
		{sku: "C", unitPrice: 50, discount: Discount{itemCountForDiscount: 3, discount: -30}},
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

func addDiscount(products []Product, itemCountForDeal int) bool {
	productCount := len(products)

	return productCount%itemCountForDeal == 0
}
