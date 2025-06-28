package blackjack

import "fmt"

func ExampleDefaultSort() {
	cards := InitCards(DefaultSort)
	fmt.Println(cards[0].suit, cards[0].rank)
	// Output:
	// HEART ACE
}

func ExampleCustomSort() {
	// the custom sorting function requires a card object to get absRank. How to do it?
	cs := func(cards []Card) func(i, j int) bool {
		return func(i, j int) bool {
			return absRank(cards[i]) > absRank(cards[j])
		}
	}
	cards := InitCards(CustomSort(cs))
	fmt.Println(cards[0].suit, cards[0].rank)
	// Output:
	// CLUB KING
}
