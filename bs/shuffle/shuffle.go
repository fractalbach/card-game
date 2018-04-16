// package shuffle includes Fisher-Yates shuffle for integer arrays.
package shuffle

import (
	"fmt"
	"math/rand"
	"time"
)

var random *rand.Rand

// Initialize the pseudo random number generator using the time,
// but only do it once.
func init() {
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// Copies an integer array, and returns a shuffled array.
// Uses the Fisher-Yates shuffling algorithm.
//
func ShuffleInts(a []int) []int {
	max := len(a) - 1
	for i := max; i > 0; i-- {
		r := random.Intn(i)
		a[i], a[r] = a[r], a[i]
	}
	return a
}

// MakeInts creates an integer array,
// where each integer is (+/-)1 from the previous one.
// If "first" < "last", then the difference between each is +1.
// If "first" > "last", then the difference is -1.
// Otherwise, an empty array is returned.
//
func MakeInts(first, last int) []int {

	var (
		difference int
		length     int
	)

	switch {
	case first > last:
		difference = -1
		length = first - last + 1

	case first < last:
		difference = 1
		length = last - first + 1

	default:
		return []int{}
	}

	x := make([]int, length)
	for i := 0; i < length; i++ {
		x[i] = first + (i * difference)
	}
	return x
}

// MakeShuffledInts returns a shuffled integer array,
// with elements starting with "first" and ending at "last".
// The length of the array will be (last - first + 1)
func MakeShuffledInts(first, last int) []int {
	x := MakeInts(first, last)
	return ShuffleInts(x)
}

//special error function
func errorFirstLast(first, last int) error {
	return fmt.Errorf("First:(%v) cannot be greater than Last:(%v)", first, last)
}
