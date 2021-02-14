package main

func main() {
	cards := newDeck()
	cards.saveToFile("test")

	cardFromFile := newDeckFromFile("test§")
	cardFromFile.print()
}
