package cards

import (
	"testing"
)

var cards = [][]string{
	[]string{"🂡", "🂢", "🂣", "🂤", "🂥", "🂦", "🂧", "🂨", "🂩", "🂪", "🂫", "🂬", "🂭", "🂮"},
	[]string{"🂱", "🂲", "🂳", "🂴", "🂵", "🂶", "🂷", "🂸", "🂹", "🂺", "🂻", "🂼", "🂽", "🂾"},
	[]string{"🃁", "🃂", "🃃", "🃄", "🃅", "🃆", "🃇", "🃈", "🃉", "🃊", "🃋", "🃌", "🃍", "🃎"},
	[]string{"🃑", "🃒", "🃓", "🃔", "🃕", "🃖", "🃗", "🃘", "🃙", "🃚", "🃛", "🃜", "🃝", "🃞"},
}

func TestUnicode(t *testing.T) {
	for i := 0; i < 52; i++ {
		suit, rank := getSuitRank(i)
		s, err := ToUnicodeString(suit, rank)
		if err != nil {
			t.Error(err)
		}
		expected := cards[suit][rank-1]
		if s != expected {
			t.Errorf("Got:(%v), Expected:(%v)", s, expected)
		} else {
			t.Log(s)
		}
	}
}

func TestStrings(t *testing.T) {
	for i := 0; i < 52; i++ {
		suit, rank := getSuitRank(i)
		u, _ := ToUnicodeString(suit, rank)
		s := SprintCard(i)
		t.Logf("Card Number: (%v): The (%v). Unicode:(%v).", i, s, u)
	}
}
