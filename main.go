package main

func main() {

	cards := newDeck()
	//cards.print()
	cards.shuffle()
	cards.print()
	//for i, card := range cards {
	//	fmt.Println(i, "--->", card)
	//}

	//fmt.Println(cards.toString())
	//cards.saveToFile("testSave.txt")
	//hand, remainingCards := deal(cards, 5)
	//hand.print()
	//remainingCards.print()
}
