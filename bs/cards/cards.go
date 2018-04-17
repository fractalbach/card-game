package cards

import (
	"fmt"
)

const (
	number_of_cards_in_deck = 52
	cards_per_suit          = 13
	unicode_facedown_card   = 127136
)

func rank(n int) int {
	return n%cards_per_suit + 1
}

func suit(n int) int {
	return n / cards_per_suit
}

func getSuitRank(n int) (int, int) {
	return suit(n), rank(n)
}

// SprintRank returns a string of the rank, which is often
// simply the string form of the integer, but can be ace, king, queen, or jack
// for certain special values.
func SprintRank(rank int) string {
	s := fmt.Sprint(rank)
	switch rank {
	case 1:
		s = "Ace"
	case 13:
		s = "King"
	case 12:
		s = "Queen"
	case 11:
		s = "Jack"
	}
	return s
}

// SprintSuit returns the string form of the suit.  The valid inputs are
// between 0-3 inclusively.  Any other values will simply be passed through
// and convert into the string form of the integer given.
func SprintSuit(suit int) string {
	s := fmt.Sprint(suit)
	switch suit {
	case 0:
		s = "Spades"
	case 1:
		s = "Diamonds"
	case 2:
		s = "Hearts"
	case 3:
		s = "Clubs"
	}
	return s
}

// SprintCard returns a string of the given card in the form of
// "Queen of Hearts" or "Ace of Spades".
func SprintCard(n int) string {
	suit, rank := getSuitRank(n)
	return SprintRank(rank) + " of " + SprintSuit(suit)
}

// ToUnicode returns the Unicode integer value of the designated card.
// If the card is out of bounds, then ToUnicode will return an error, and
// the Unicode of the face-down card
func ToUnicodeInt(suit, rank int) (int, error) {
	base := unicode_facedown_card
	if suit < 0 || suit > 3 {
		return base, inputError(0, 3, suit)
	}
	if rank < 1 || rank > 13 {
		return base, inputError(1, 13, rank)
	}
	val := base + (suit * 16) + rank
	return val, nil
}

// ToUnicodeString returns the special Playing-Card Unicode character of
// the given rank and suit, or returns an error if the rank or suit are
// out of bounds.
func ToUnicodeString(suit, rank int) (string, error) {
	val, err := ToUnicodeInt(suit, rank)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%c", val), nil
}

func inputError(min, max, got interface{}) error {
	return fmt.Errorf("Out of Bounds: Got:(%v), Want between:(%v, %v) ", got, min, max)
}
