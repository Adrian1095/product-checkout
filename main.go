package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		basket        = make(map[string][]Product)
		checkoutTotal = 0
	)

	displayMenu()

	var option int

	for {
		fmt.Scan(&option)
		if option == 0 {
			break
		}

		if option == 10 {
			displayCheckout(basket, checkoutTotal)
			fmt.Println("press Enter to navigate to menu...")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
			displayMenu()
		}

		if option >= 1 && option <= 3 {
			product := addProductToBasket(option)
			basket[product.sku] = append(basket[product.sku], product)

			if addDiscount(basket[product.sku], product.discount.itemCountForDiscount) {
				discountPrice := product.unitPrice + product.discount.discount
				checkoutTotal = calculateCheckoutTotal(discountPrice, checkoutTotal)
			} else {
				checkoutTotal = calculateCheckoutTotal(product.unitPrice, checkoutTotal)
			}
			fmt.Println("Product added to basket")
		}

		if option == 4 {
			displayProductDiscount()
			fmt.Println("press Enter to navigate to menu...")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
			displayMenu()
		}

	}
}
