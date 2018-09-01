package main

import (
	"testing"
	"math/rand"
	"fmt"
	"sort"
)

func TestCompute(t *testing.T) {
	rand.Seed(1)

	rows := []struct {
		array  []int
		target int
		result int
	}{
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 0, 0},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 11, -1},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10, 10},
	}

	for _, row := range rows {
		for i, fn := range []func([]int, int) int{
			binary,
			recursiveBinary,
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

func generateArray(l int) []int {
	a := make([]int, l)

	for i, _ := range a {
		a[i] = rand.Intn(1000000)
	}

	return a
}

func clone(base []int) []int {
	a := make([]int, len(base))
	copy(a, base)

	return a
}

func clone2(base []string) []string {
	a := make([]string, len(base))
	copy(a, base)

	return a
}

func eq(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, an := range a {
		if an != b[i] {
			return false
		}
	}
	return true
}

func eq2(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, an := range a {
		if an != b[i] {
			return false
		}
	}
	return true
}


func TestCompute2(t *testing.T) {
	rows := [][]int{
		generateArray(100),
		generateArray(200),
		generateArray(15),
		generateArray(10),
		generateArray(9),
		generateArray(8),
		generateArray(7),
	}

	for _, row := range rows {
		ac := clone(row)
		sort.Ints(ac)

		for i, fn := range []func([]int) []int{
			selectionSort,
			insertionSort,
			mergeSort,
			quickSort,
			loopMergeSort,
		} {
			a := clone(row)
			ex := fn(a)

			if !eq(ac, ex) {
				fmt.Printf("%d: %+v %+v \n", i, ac, ex)
				t.Fail()
			}
		}
	}
}

func generateArray2(l int) []int {
	a := make([]int, l)

	for i, _ := range a {
		a[i] = rand.Intn(2)
	}

	return a
}

func TestCompute3(t *testing.T) {
	rows := [][]int{
		generateArray2(100),
	}

	for _, row := range rows {
		ac := clone(row)
		sort.Ints(ac)

		for i, fn := range []func([]int) []int{
			reallySimpleSort,
		} {
			a := clone(row)
			ex := fn(a)

			if !eq(ac, ex) {
				fmt.Printf("%d: %+v %+v \n", i, ac, ex)
				t.Fail()
			}
		}
	}
}

func randomChar() string {
	return string(chars[rand.Intn(36)])
}

func generateArray3(l int) []string {
	a := make([]string, l)

	for i, _ := range a {
		a[i] = randomChar()
	}

	return a
}

func generateArray4(l int) []string {
	a := make([]string, l)

	for i, _ := range a {
		for j := 0; j < 6; j++ {
			a[i] += randomChar()
		}
	}

	return a
}

func TestCompute4(t *testing.T) {
	rows := [][]string{
		generateArray3(100),
		generateArray3(200),
		generateArray3(15),
		generateArray3(10),
	}

	for _, row := range rows {
		ac := clone2(row)
		sort.Strings(ac)

		for i, fn := range []func([]string) []string{
			countingSort,
		} {
			a := clone2(row)
			ex := fn(a)

			if !eq2(ac, ex) {
				fmt.Printf("%d: %+v %+v \n", i, ac, ex)
				t.Fail()
			}
		}
	}
}

func TestCompute5(t *testing.T) {
	rows := [][]string{
		generateArray4(100),
		generateArray4(200),
		generateArray4(15),
		generateArray4(10),
	}

	for _, row := range rows {
		ac := clone2(row)
		sort.Strings(ac)

		for i, fn := range []func([]string) []string{
			radixSort,
		} {
			a := clone2(row)
			ex := fn(a)

			if !eq2(ac, ex) {
				fmt.Printf("%d: %+v %+v \n", i, ac, ex)
				t.Fail()
			}
		}
	}
}
