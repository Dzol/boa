package main

import (
	"testing"
	"testing/quick"
)

func TestMinimumFirst(t *testing.T) {
	f := func(x []int) bool {
		if len(x) == 0 {
			return true
		}
		bubbleSort(x)
		return minimum(x...) == x[0]
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMaximumLast(t *testing.T) {
	f := func(x []int) bool {
		if len(x) == 0 {
			return true
		}
		bubbleSort(x)
		return maximum(x...) == x[len(x)-1]
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func minimum(i ...int) int {
	minimum := i[0]
	for _, v := range i {
		if v < minimum {
			minimum = v
		}
	}
	return minimum
}

func maximum(i ...int) int {
	maximum := i[0]
	for _, v := range i {
		if v > maximum {
			maximum = v
		}
	}
	return maximum
}
