package main

import (
	"testing"
	"math/rand"
	"fmt"
)

func TestCompute(t *testing.T) {
	rand.Seed(1)

	rows := []struct {
		array  []int
		target int
		result int
	}{
		{[]int{8, 9, 1, 4, 5, 0, 6, 3, 2, 7}, 1, 2},
		{[]int{8, 9, 1, 4, 5, 0, 6, 3, 2, 7}, 10, -1},
		{[]int{8, 9, 1, 4, 5, 0, 6, 3, 2, 7}, 7, 9},
	}

	fmt.Println(factorial(10))

	for _, row := range rows {
		for i, fn := range []func([]int, int) int{
			linear,
			betterLinear,
			sentinelLinear,
			recursiveLinear,
		} {
			a := make([]int, len(row.array))
			copy(a, row.array)

			ac := fn(a, row.target)

			if ac != row.result {
				fmt.Printf("%d: %+v \n", i, ac)
				t.Fail()
			}
		}
	}
}

