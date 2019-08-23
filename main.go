package main

func main() {
	//var card string = "Joker"
	//card := "Joker"
	//card = "Five"
	//card := newCard()
	//fmt.Println(card)
	//cards := []string{"Joker", newCard()}
	//cards := deck{"Joker", newCard()}
	//cards = append(cards, "six of spades")

	//cards := newDeck()

	cards := newDeckFromFile("tes_cards")
	cards.shuffle()

	hand, remainingCard := deal(cards, 5)

	hand.print()
	remainingCard.print()

	//cards.print()

	//save to file
	//cards.toFile("tes_cards")

	//cards.print()
	/*
		for i, card := range cards {
			fmt.Println(i, card)
		}
	*/
}

func newCard() string {
	return "Six Secop"
}

// array in go
// array
// Fixed length list of things
// slice
// an array that can grow or shrink
