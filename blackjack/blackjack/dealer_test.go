package blackjack

import (
	"github.com/google/uuid"
	"testing"
)

func TestDealer_DealCards(t *testing.T) {
	t.Run("deal two cards to the player and dealer", func(t *testing.T) {
		d := NewDealer()
		p := Player{
			id:          uuid.New(),
			name:        "player1",
			moneyInHand: 100.0,
		}
		players := []Player{p}
		err := d.DealCards(players, 2)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if len(p.cards) != 2 {
			t.Errorf("Expected 2 cards, got %v", len(p.cards))
		}
	})
}
