package lib

func Fibonacci(n uint64) uint64 {
	var current uint64 = 1
	var previous uint64 = 1
	var index uint64 = 0

	for index < n {
		var newPrevious = current
		current += previous
		previous = newPrevious

		index++
	}

	return current
}
