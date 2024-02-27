package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func IsSame[T comparable](param1, param2 T) bool {
	if param1 == param2 {
		return true
	}
	return false
}

func TestIsSame(t *testing.T) {
	assert.True(t, IsSame[string]("Oke", "Oke"))
	assert.True(t, IsSame[int](100, 100))
	assert.True(t, IsSame[bool](true, true))
}
