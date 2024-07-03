package main

import (
	"fmt"
	"os"
)

func displayMenu() {

	fmt.Println("Exit: 0")
	fmt.Println("Select product: product id")
	fmt.Println("Edit product discount: 4")
	fmt.Println("Checkout: 10")

	for _, product := range products {
		fmt.Println("Id", product.id)
		fmt.Println("SKU ", product.sku)
		fmt.Println("Unit Price: ", product.unitPrice)
		fmt.Println()
	}
}

func displayCheckout(basket map[string][]Product, checkoutTotal int) {
	fmt.Println("SKU		Unit Price		Special Price		Quantity")
	for _, product := range products {
		productDeal := product.unitPrice*product.discount.itemCountForDiscount + product.discount.discount
		fmt.Fprintln(os.Stdout, []any{product.sku, "		", product.unitPrice, "			", product.discount.itemCountForDiscount, "for ", productDeal, " 		 ", len(basket[product.sku])}...)
	}
	fmt.Println()
	fmt.Fprintln(os.Stdout, []any{"Total:", checkoutTotal}...)
}

func displayProductDiscount() {
	for _, product := range products {
		fmt.Println("Id", product.id)
		fmt.Println("SKU ", product.sku)
		fmt.Println("Unit Price: ", product.unitPrice)
		fmt.Println()
	}
	var productId, itemCount, discount int
	fmt.Println("Select product Id:")
	fmt.Scan(&productId)
	fmt.Println("Quantity of items to qualify for discount:")
	fmt.Scan(&itemCount)
	fmt.Println("Discount to subtract from product price:")
	fmt.Scan(&discount)

	products[productId-1].discount.itemCountForDiscount = itemCount
	products[productId-1].discount.discount = -discount
}
