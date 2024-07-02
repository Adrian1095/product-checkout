package main

import (
	"reflect"
	"testing"
)

func TestGetProductList(t *testing.T) {
	expectedProducts := []Product{
		{sku: "A", unitPrice: 20, discount: Discount{itemCountForDiscount: 3, discount: -5}},
		{sku: "B", unitPrice: 15, discount: Discount{itemCountForDiscount: 2, discount: -10}},
		{sku: "C", unitPrice: 50, discount: Discount{itemCountForDiscount: 3, discount: -30}},
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
			{sku: "A", unitPrice: 20, discount: Discount{itemCountForDiscount: 3, discount: -5}},
			{sku: "A", unitPrice: 20, discount: Discount{itemCountForDiscount: 3, discount: -5}},
		},
		"B": {
			{sku: "B", unitPrice: 15, discount: Discount{itemCountForDiscount: 2, discount: -10}},
		},
		"C": {
			{sku: "C", unitPrice: 50, discount: Discount{itemCountForDiscount: 3, discount: -30}},
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

func TestCalculateCheckoutTotal(t *testing.T) {
	checkoutTotal := 10
	expectedTotal := 15

	checkoutTotal = calculateCheckoutTotal(5, 10)

	if expectedTotal != checkoutTotal {
		t.Fatalf(`Expected total does not equal checkoutTotal
		expectedTotal: %v,
		checkoutTotal: %v`,
			expectedTotal, checkoutTotal)
	}
}

func TestAddDiscount(t *testing.T) {
	var products = []Product{
		{sku: "A", unitPrice: 20},
		{sku: "A", unitPrice: 20},
	}

	tests := map[string]struct {
		products         []Product
		itemCountForDeal int
		result           bool
	}{
		"return true": {
			products:         products,
			itemCountForDeal: 2,
			result:           true,
		},
		"return false": {
			products:         products,
			itemCountForDeal: 3,
			result:           false,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got, expected := addDiscount(test.products, test.itemCountForDeal), test.result; got != expected {
				t.Fatalf("addDiscount(%v,%v) returned %v; expected %v", test.products, test.itemCountForDeal, got, expected)
			}
		})
	}

}
