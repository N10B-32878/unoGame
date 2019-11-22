package utils_test

import (
	"testing"
	"uno/utils"
)

func TestDeckGenerator(t *testing.T) {
	//test deck generator
	deck := utils.GenerateCardDeck()
	if len(deck) != 76{
		t.FailNow()
	}
}

//GoLang is pass-by-reference so this test will always fail. TODO improve test
/*func TestDeckShuffle(t *testing.T) {
	//test deck shuffle
	deck := utils.GenerateCardDeck()
	tmpDeck := utils.ShuffleDeck(deck)


	for key, val := range deck{
		if val == tmpDeck[key] {
			t.FailNow()
		}
	}
}*/

func TestCardDraw(t *testing.T){
	deck := utils.GenerateCardDeck()
	tmpDeck, tmpCard := utils.GiveCardFromDeck(deck)

	for _, val := range tmpDeck{
		if val == tmpCard{
			t.FailNow()
		}
	}
}