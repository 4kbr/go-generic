package main

import (
	"fmt"
	"testing"
)

type Bag[T any] []T

func PrintBag[T any](bag Bag[T]) {
	for _, value := range bag {
		fmt.Println(value)
	}
}

func TestBagString(t *testing.T) {
	names := Bag[string]{"Nama", "Mantap"}
	PrintBag(names)
}
func TestBagInt(t *testing.T) {
	numbers := Bag[int]{12, 13, 14}
	PrintBag(numbers)
}
