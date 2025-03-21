package blackjack

import (
	"testing"
)

func TestInitCards(t *testing.T) {
	t.Run("a need to card deck with all the required cards out", func(t *testing.T) {
		cards := InitCards()
		if cards == nil {
			t.Errorf("Expected nil cards, got %v", cards)
		}
		if len(cards) != 52 {
			t.Errorf("Expected 52 cards, got %v", len(cards))
		}
	})
}

func TestGetDecks(t *testing.T) {
	subtests := []struct {
		name        string
		numOfDecks  int
		expectedLen int
	}{
		{
			name:        "given an input number of decks, I need to get the same number of decks with 52 cards",
			numOfDecks:  3,
			expectedLen: 3,
		},
		{
			name:        "given zero decks, i shouldn't get any deck back",
			numOfDecks:  0,
			expectedLen: 0,
		},
	}

	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			cd := GetDecks(subtest.numOfDecks)
			if len(cd) != subtest.expectedLen {
				t.Errorf("Expected length of decks to be %v, got %v", subtest.expectedLen, len(cd))
			}
		})
	}
}
