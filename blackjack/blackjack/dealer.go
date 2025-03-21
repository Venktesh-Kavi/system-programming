package blackjack

import (
	"fmt"
	"github.com/google/uuid"
	"math/rand"
)

type Dealer struct {
	id        uuid.UUID
	cardDecks []CardDeck
	balance   float32
}

func NewDealer() Dealer {
	return Dealer{}
}

func (d Dealer) DealCards(players []Player) {
	fmt.Printf("Dealer: %v, starting to deal cards\n", d.id)
	shuffle(d.cardDecks)
}

func shuffle(cardDecks []CardDeck) {
	/*
		shuffle cards in each deck
		shuffle all card decks
	*/
	for _, cd := range cardDecks {
		c := cd.cards
		fmt.Println(c)
	}
}

func shuffleCards(c []Card) {
	for i := 0; i <= len(c); i++ {
		// shuffle 13 times.
		rand.Shuffle(3, func(i, j int) {
			c[i], c[j] = c[j], c[i]
		})
	}
	rand.Intn(len(c))
}
