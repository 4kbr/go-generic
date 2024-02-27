package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func FindMin[T interface{ int | int64 | float64 }](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestFindMin(t *testing.T) {
	assert.Equal(t, 10, FindMin(10, 20))
	assert.Equal(t, int64(10), FindMin[int64](int64(10), int64(20)))
	assert.Equal(t, 10.0, FindMin(10.0, 20.0))
}

func GetFirst[T []E, E any](slice T) E {
	first := slice[0]
	return first
}

func TestGetFirst(t *testing.T) {
	names := []string{
		"Aldi", "Uda", "Bad",
	}
	// fisrt:=GetFirst(names) //bisa
	fisrt := GetFirst[[]string, string](names) //bisa
	assert.Equal(t, "Aldi", fisrt)
}
