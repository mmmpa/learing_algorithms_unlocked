package main

func linear(array []int, target int) int {
	l := len(array)
	result := -1

	for i := 0; i < l; i++ {
		if array[i] == target {
			result = i
		}
	}

	return result
}

func betterLinear(array []int, target int) int {
	l := len(array)

	for i := 0; i < l; i++ {
		if array[i] == target {
			return i
		}
	}

	return -1
}

func sentinelLinear(array []int, target int) int {
	l := len(array)

	last := array[l-1]
	array[l-1] = target

	i := 0
	for {
		if array[i] == target {
			if i < l-1 {
				return i
			} else if last == target {
				return i
			} else {
				return -1
			}
		}
		i++
	}
}

func recursiveLinear(array []int, target int) int {
	return _recursiveLinear(array, target, 0)
}

func _recursiveLinear(array []int, target int, i int) int {
	l := len(array)

	if i == l {
		return -1
	}

	if array[i] == target {
		return i
	}

	return _recursiveLinear(array, target, i+1)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}
