package main

import "fmt"
import "goapp/Model"

func main() {
  card := Model.Card{
		Rank: Model.TWO, 
		Suit: Model.DIAMOND,
	}
	fmt.Printf("%v", card.Display())
}
