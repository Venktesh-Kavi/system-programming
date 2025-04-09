package blackjack

import (
	"fmt"
	"github.com/google/uuid"
)

type Dealer struct {
	id        uuid.UUID
	cardDecks []CardDeck
	balance   float32
}

func NewDealer() Dealer {
	return Dealer{
		id:        uuid.New(),
		cardDecks: NewCardDeck(),
		balance:   0.0,
	}
}

func NewCardDeck() []CardDeck {
	return []CardDeck{{cards: InitCards()}}
}

func (d Dealer) DealCards(players []Player, noOfCards int) error {
	fmt.Printf("Dealer: %v, starting to deal cards\n", d.id)
	if players == nil {
		return fmt.Errorf("player is nil, ensure to create a player")
	}
	// deal 2 cards initially
	cs := FlattenCardDeck(d.cardDecks)
	for i := 0; i < noOfCards; i++ {
		for _, player := range players {
			player.cards = append(player.cards, cs[i])
		}
	}
	fmt.Println(players)
	return nil
}
