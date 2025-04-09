package blackjack

import (
	"bufio"
	"fmt"
	"github.com/google/uuid"
	"os"
)

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
	fmt.Println("#### Starting black jack game ####")
	fmt.Println("")
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
}
