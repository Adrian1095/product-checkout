package main

import (
	"reflect"
	"testing"
)

func TestGetProductList(t *testing.T) {
	expectedProducts := []Product{
		{sku: "A", unitPrice: 20},
		{sku: "B", unitPrice: 15},
		{sku: "C", unitPrice: 50},
	}

	actualProducts := getProductList()

	if !reflect.DeepEqual(expectedProducts, actualProducts) {
		t.Fatalf(`Expected products list does not match actual products list.
		Expected: %v
		Actual: %v`,
			expectedProducts, actualProducts)
	}
}

func TestAddProductToBasket(t *testing.T) {
	basket := make(map[string][]Product)

	var expectedBasket = map[string][]Product{
		"A": {
			{sku: "A", unitPrice: 20},
			{sku: "A", unitPrice: 20},
		},
		"B": {
			{sku: "B", unitPrice: 15},
		},
		"C": {
			{sku: "C", unitPrice: 50},
		},
	}

	product := addProductToBasket(1)
	basket[product.sku] = append(basket[product.sku], product)

	product = addProductToBasket(2)
	basket[product.sku] = append(basket[product.sku], product)

	product = addProductToBasket(3)
	basket[product.sku] = append(basket[product.sku], product)

	product = addProductToBasket(1)
	basket[product.sku] = append(basket[product.sku], product)

	if !reflect.DeepEqual(expectedBasket, basket) {
		t.Fatalf(`Expected products list does not match actual products list.
		Expected: %v
		Actual: %v`,
			expectedBasket, basket)
	}
}
