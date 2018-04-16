package shuffle

import (
	"testing"
)

func TestMakeInts(t *testing.T) {
	a, err := MakeInts(1, 10)
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
	t.Log(a)
	a, err = MakeInts(-1, -5)
	if err == nil {
		t.Error("there should have been an error", a)
	}
	t.Log(a, err)
}

func TestBasic(t *testing.T) {
	a, _ := MakeInts(1, 10)
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
	a, _ := MakeInts(1, 10)
	for i := 0; i < 10; i++ {
		a = ShuffleInts(a)
		t.Log(a)
	}
}
