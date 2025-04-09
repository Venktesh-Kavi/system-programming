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
func TestSorting(t *testing.T) {
	t.Run("default sorting", func(t *testing.T) {
		cards := InitCards(DefaultSort)
		if cards[0].rank != ACE {
			t.Errorf("Expected first card to be ACE, got %v", cards[0].rank)
		}
	})
	t.Run("custom sorting", func(t *testing.T) {
		cs := func(cards []Card) func(i, j int) bool {
			return func(i, j int) bool {
				return absRank(cards[i]) > absRank(cards[j])
			}
		}
		cards := InitCards(CustomSort(cs))
		if cards[0].rank != KING {
			t.Errorf("Expected first card to be KING, got %v", cards[0].rank)
		}
	})
}

func TestDefaultShuffle(t *testing.T) {
	t.Run("shuffle the cards", func(t *testing.T) {
		cards := InitCards(DefaultShuffle)
		if cards[0].suit == 0 && cards[0].rank == 0 {
			t.Errorf("Expected cards to be shuffled, got %v", cards)
		}
	})
}

func TestFlattenCardDeck(t *testing.T) {
	t.Run("flatten the card decks", func(t *testing.T) {
		cd := GetDecks(3)
		cards := FlattenCardDeck(cd)
		if len(cards) != 156 {
			t.Errorf("Expected 156 cards, got %v", len(cards))
		}

		if cards[13].rank != ACE && cards[13].suit == DIAMOND {
			t.Errorf("Expected first card in the next suite to be ACE, got %v", cards[14].rank)
		}
	})
}
