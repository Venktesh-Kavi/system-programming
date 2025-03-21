package blackjack

import "github.com/google/uuid"

type Player struct {
	id          uuid.UUID
	name        string
	moneyInHand float32
	cards       []Card
}
