package blackjack

import "github.com/google/uuid"

type Game struct {
	dealerId uuid.UUID
	players  []Player
	round    Round
}

type Round struct {
	roundNo      int
	betPerPlayer map[uuid.UUID]float32
}

func Start() {

}
