package main

import (
	"math"
)

func binary(a []int, x int) int {
	l := len(a)
	p := 0
	r := l - 1

	for p <= r {
		q := p + (r+1-p)/2
		v := a[q]

		switch {
		case v == x:
			return q
		case x < v:
			r = q - 1
		case v < x:
			p = q + 1
		}
	}

	return -1
}

func recursiveBinary(a []int, x int) int {
	l := len(a)
	p := 0
	r := l - 1

	return _recursiveBinary(a, x, p, r)
}

func _recursiveBinary(a []int, x int, p, r int) int {
	if p > r {
		return -1
	}

	q := (p + r) / 2
	v := a[q]

	switch {
	case v == x:
		return q
	case x < v:
		return _recursiveBinary(a, x, p, q-1)
	default:
		return _recursiveBinary(a, x, q+1, r)
	}
}

func selectionSort(a []int) []int {
	l := len(a)

	for i := 0; i < l; i++ {
		min := i

		for j := i; j < l; j++ {
			if a[j] < a[min] {
				min = j
			}
		}

		if min != i {
			a[i], a[min] = a[min], a[i]
		}
	}

	return a
}

func insertionSort(a []int) []int {
	l := len(a)

	for i := 1; i < l; i++ {
	INNER:
		for j := i; j >= 1; j-- {
			if a[j-1] > a[j] {
				a[j-1], a[j] = a[j], a[j-1]
			} else {
				break INNER
			}
		}
	}

	return a
}

func mergeSort(a []int) []int {
	l := len(a)

	return _mergeSort(a, 0, l-1)[:l]
}

func _mergeSort(a []int, r, p int) []int {
	if r == p {
		return []int{a[r], math.MaxInt32}
	}

	q := (p + r) / 2
	left := _mergeSort(a, r, q)
	right := _mergeSort(a, q+1, p)

	aa := make([]int, p+1-r+1)
	aa[p+1-r] = math.MaxInt32

	li := 0
	ri := 0
	for i, _ := range aa {
		if left[li] < right[ri] {
			aa[i] = left[li]
			li++
		} else {
			aa[i] = right[ri]
			ri++
		}
	}

	return aa
}

func quickSort(a []int) []int {
	l := len(a)
	_quickSort(a, 0, l-1)

	return a
}

func _quickSort(a []int, r, p int) {
	if r > p {
		return
	}

	q := partition(a, r, p)
	_quickSort(a, r, q-1)
	_quickSort(a, q+1, p)
}

func partition(a []int, r, p int) int {
	pivot := a[p]
	q := r

	for i := r; i < p; i++ {
		if a[i] <= pivot {
			a[q], a[i] = a[i], a[q]
			q++
		}
	}
	a[q], a[p] = a[p], a[q]
	return q
}

func loopMergeSort(a []int) []int {
	l := len(a)

	span := 1

	for span < l {
		twice := span * 2
		for i := 0; i < l; i += twice {
			if i+span >= l {
				continue
			}

			tail := i + twice
			if tail > l {
				tail = l
			}

			leftLength := span
			rightLength := tail - (i + span)

			left := make([]int, leftLength+1)
			right := make([]int, rightLength+1)

			left[span] = math.MaxInt32
			right[tail-(i+span)] = math.MaxInt32

			for j := 0; j < leftLength; j++ {
				left[j] = a[i+j]
			}
			for j := 0; j < rightLength; j++ {
				right[j] = a[i+span+j]
			}

			li := 0
			ri := 0

			for j := 0; j < leftLength+rightLength; j++ {
				switch {
				case left[li] < right[ri]:
					a[i+j] = left[li]
					li++
				default:
					a[i+j] = right[ri]
					ri++
				}
			}
		}

		span *= 2
	}

	return a
}

func reallySimpleSort(a []int) []int {
	bin := make([]int, 2)

	for _, n := range a {
		bin[n]++
	}

	for i := 0; i < bin[0]; i++ {
		a[i] = 0
	}
	for i := bin[0]; i < len(a); i++ {
		a[i] = 1
	}

	return a
}

var chars = []rune("0123456789abcdefghijklmnopqrstuvwxyz")
var prepared = false
var charMap = map[rune]int{}

func prepare() {
	if prepared {
		return
	}

	for i, n := range chars {
		charMap[n] = i
	}

	prepared = true
}

func countingSort(a []string) []string {
	aa := make([]string, len(a))

	indexes := generateIndexes(a, 0)

	for _, n := range a {
		j := charMap[[]rune(n)[0]]
		aa[indexes[j]] = n
		indexes[j]++
	}

	return aa
}

func generateIndexes(a []string, pos int) []int {
	prepare()
	bin := make([]int, 36)

	for _, n := range a {
		j := charMap[[]rune(n)[pos]]
		bin[j]++
	}

	indexes := make([]int, 36)
	indexes[0] = 0

	for i := 1; i < 36; i++ {
		indexes[i] = indexes[i-1] + bin[i-1]
	}

	return indexes
}

func radixSort(base []string) []string {
	aa := make([]string, len(base))

	a := base
	b := aa

	for j := 5; j >= 0; j-- {
		indexes := generateIndexes(a, j)
		for _, n := range a {
			j := charMap[[]rune(n)[j]]
			b[indexes[j]] = n
			indexes[j]++
		}
		a, b = b, a
	}

	return a
}
