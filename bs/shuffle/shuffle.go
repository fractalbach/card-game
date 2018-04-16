// package shuffle includes Fisher-Yates shuffle for integer arrays.
package shuffle

import (
	"fmt"
	"math/rand"
	"time"
)

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

// Copies an integer array, and returns a shuffled array.
func ShuffleInts(a []int) []int {
	max := len(a) - 1
	for i := max; i > 0; i-- {
		r := random.Intn(i)
		a[i], a[r] = a[r], a[i]
	}
	return a
}

// MakeInts creates an integer array,
// with integers starting with "first" and ending at "end", adding 1 to each.
// Returns an error if "last" > "first".
func MakeInts(first, last int) ([]int, error) {
	if first > last {
		return []int{}, errorFirstLast(first, last)
	}
	n := last - first + 1
	x := make([]int, n)
	for i := 0; i < n; i++ {
		x[i] = first + i
	}
	return x, nil
}

// MakeShuffledInts returns a shuffled integer array,
// with elements starting with "first" and ending at "last".
// The length of the array will be (last - first + 1)
func MakeShuffledInts(first, last int) []int {
	x, _ := MakeInts(first, last)
	return ShuffleInts(x)
}

//special error function
func errorFirstLast(first, last int) error {
	return fmt.Errorf("First:(%v) cannot be greater than Last:(%v)", first, last)
}
