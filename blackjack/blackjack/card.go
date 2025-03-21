package blackjack

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
	ONE
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

func InitCards(opts ...func(cardOpts CardOpts)) []Card {
	cards := make([]Card, 52)
	for _, suit := range suits {
		for i := minRank; i < maxRank; i++ {
			card := NewCard(suit, i)
			cards[i] = card
		}
	}
	return cards
}

func GetDecks(numOfDecks int) []CardDeck {
	cardDecks := make([]CardDeck, numOfDecks)
	for i := range numOfDecks {
		cd := CardDeck{cards: InitCards()}
		cardDecks[i] = cd
	}
	return cardDecks
}
