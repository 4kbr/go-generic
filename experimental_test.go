package main

import (
	"maps"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
)

func ExperimentalMin[T constraints.Ordered](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestExperimentalMin(t *testing.T) {
	assert.Equal(t, 100, ExperimentalMin(100, 200))
	assert.Equal(t, 100.0, ExperimentalMin[float64](100, 200))
}

func TestExperimentalMaps(t *testing.T) {
	first := map[string]string{
		"Name": "Oke",
	}
	second := map[string]string{
		"Name": "Oke",
	}

	assert.True(t, maps.Equal(first, second))
}

func TestExperimentalSlice(t *testing.T) {
	first := []string{"Oke"}
	second := []string{"Oke"}

	assert.True(t, slices.Equal(first, second))
}
