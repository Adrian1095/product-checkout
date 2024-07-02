package main

import "fmt"

func displayMenu(products []Product) {
	fmt.Println("Exit: 0")
	fmt.Println("Enter SKU to add to checkout: SKU value")

	for _, product := range products {
		fmt.Println("SKU ", product.sku)
		fmt.Println("Unit Price: ", product.unitPrice)
		fmt.Println()
	}
}
