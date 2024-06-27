package main

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
