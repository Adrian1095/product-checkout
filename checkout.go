package main

import "fmt"

func displayProductMenu() {
	products := getProductList()

	for _, product := range products {
		fmt.Println("Sku: ", product.sku)
		fmt.Println("Unit Price: ", product.unitPrice)
		fmt.Println()
	}
}
