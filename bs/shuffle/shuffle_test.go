package shuffle

import (
	"testing"
)

func TestMakeInts(t *testing.T) {
	a := MakeInts(1, 10)
	t.Log(a)
	a = MakeInts(-1, -5)
	t.Log(a)
}

func TestBasic(t *testing.T) {
	a := MakeInts(1, 10)
	t.Log(a)
	a = ShuffleInts(a)
	t.Log(a)
}

func TestMakers(t *testing.T) {
	for i := 1; i < 20; i += 5 {
		a := MakeShuffledInts(1, i)
		t.Log(a)
	}
}

func TestRandomness(t *testing.T) {
	a := MakeInts(1, 10)
	for i := 0; i < 10; i++ {
		a = ShuffleInts(a)
		t.Log(a)
	}
}
