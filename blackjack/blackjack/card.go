//go:generate stringer -type=Suit,Rank
package blackjack

import (
	"sort"
)

type Rank uint8
type Suit uint8

const (
	HEART Suit = iota
	DIAMOND
	SPADE
	CLUB
)

const (
	_ Rank = iota
	ACE
	TWO
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	TEN
	JACK
	QUEEN
	KING
)

const (
	minRank = ACE
	maxRank = KING
)

var suits = [...]Suit{HEART, DIAMOND, SPADE, CLUB}

type Opts struct {
	Shuffle bool
}

type Card struct {
	suit Suit
	rank Rank
}

type CardDeck struct {
	cards []Card
}

type CardOpts struct {
	Shuffle    bool
	Sort       bool
	JokerCount int
}

func NewCard(suit Suit, rank Rank) Card {
	return Card{suit: suit, rank: rank}
}

// InitCards functional options, provide options as functions.
func InitCards(opts ...func(cards []Card) []Card) []Card {
	cards := make([]Card, 52)
	index := 0 // Keep a separate counter for index
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ { // Corrected condition to include maxRank
			card := NewCard(suit, rank)
			cards[index] = card // Use separate counter
			index++
		}
	}
	for _, opt := range opts {
		cards = opt(cards)
	}
	return cards
}

func CustomSort(less func(cards []Card) func(i, j int) bool) func(cards []Card) []Card {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

func Less(card []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(card[i]) < absRank(card[j])
	}
}

func absRank(c Card) int {
	return (int(c.suit) * 13) + int(c.rank)
}

func GetDecks(numOfDecks int) []CardDeck {
	cardDecks := make([]CardDeck, numOfDecks)
	for i := range numOfDecks {
		cd := CardDeck{cards: InitCards()}
		cardDecks[i] = cd
	}
	return cardDecks
}
