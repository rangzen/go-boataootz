package functional

// Map apply a function on a slice.
// The function can have two different types for input and output.
func Map[A, B any](collection []A, f func(A) B) []B {
	result := make([]B, len(collection))
	for i, c := range collection {
		result[i] = f(c)
	}
	return result
}

// Reduce apply accumulator to a slice.
// The accumulator can use two different type, one for input, one for accumulated value and output.
func Reduce[A, B any](collection []A, accumulator func(B, A) B, initialValue B) B {
	result := initialValue
	for _, c := range collection {
		result = accumulator(result, c)
	}
	return result
}
