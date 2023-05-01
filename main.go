package main

func main() {

	cards := newDeckFromFile("testSave.txt")
	//for i, card := range cards {
	//	fmt.Println(i, "--->", card)
	//}
	cards.print()
	//fmt.Println(cards.toString())
	//cards.saveToFile("testSave.txt")
	//hand, remainingCards := deal(cards, 5)
	//hand.print()
	//remainingCards.print()
}
