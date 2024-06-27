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
