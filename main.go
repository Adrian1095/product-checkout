package main

import (
	"fmt"
)

func main() {
	var (
		basket        = make(map[string][]Product)
		checkoutTotal = 0
	)
	products := getProductList()
	displayMenu(products)

	var option int

	for {
		fmt.Scan(&option)
		if option == 0 {
			break
		}

		product := addProductToBasket(option)
		basket[product.sku] = append(basket[product.sku], product)
		checkoutTotal = calculateCheckoutTotal(product.unitPrice, checkoutTotal)
		fmt.Println(checkoutTotal)
	}
}
