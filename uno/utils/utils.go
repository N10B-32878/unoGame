package utils

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Card struct {
	Color	string
	Number	int
}

//main logic function, magic happens here
func RunLogic(){
	deck := GenerateCardDeck()
	deck = ShuffleDeck(deck)

	//TODO get number of players from user
	numberOfPlayers := 4

	err, playersCards := dealCards(numberOfPlayers,deck)
	if err != nil {
		fmt.Println("Houston we have a problem!", err)
	} else {
		deck, pileCard := GiveCardFromDeck(deck, nil)
		fmt.Println("First on pile: ", pileCard)

		gameLoop(deck, playersCards,numberOfPlayers, pileCard)
	}
}

func gameLoop(deck []Card, playersCards [][]Card, numberOfPlayers int, pileCard Card){
	for true{
		for i:=0;i<numberOfPlayers;i++{
			printHand(playersCards, i)
			place, cardNumber := canPlaceACardOnTop(playersCards[i],pileCard)
			if place{
				playersCards[i], pileCard = GiveCardFromDeck(playersCards[i], cardNumber)
				fmt.Println("Player ", i, " placed card ", pileCard, " on pile")
			} else {
				deck, playersCards = assignCardToPlayer(i, playersCards, deck)
				fmt.Println("Player ", i, " took new card")
			}
		}
		if SomeoneWon(playersCards){
			break
		}
	}
}

//check if player can put a card on pile or has to draw
func canPlaceACardOnTop(playerHand []Card, currentCard Card) (bool, *int) {
	var retInt *int
	retBool := false

	if len(playerHand) < 1{
		retBool = false
	} else {
		for key, val := range playerHand {
			if val.Color == currentCard.Color || val.Number == currentCard.Number {
				retBool = true
				retInt = &key
				break
			}
		}
	}

	return retBool, retInt
}

//generate card deck according to uno rules, with exception of special and wilds cards (only number cards)
func GenerateCardDeck() []Card{
	var retDeck []Card

	colors := [4]string{"red", "yellow", "green", "blue"}

	for _, val := range colors{
		for i:=0;i<10;i++{
			if i > 0 && i < 10{
				retDeck = assignCardValueAndAddToDeck(retDeck, val, i)
			}
			retDeck = assignCardValueAndAddToDeck(retDeck, val, i)
		}
	}

	return retDeck
}

//create card based on number and color, then add it to deck
func assignCardValueAndAddToDeck(deck []Card, color string, number int) []Card{
	tmpCard := Card{color, number}
	deck = append(deck,tmpCard)

	return deck
}

//shuffle deck
func ShuffleDeck(deck []Card) []Card{
	tmpDeck := deck

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(tmpDeck), func(i, j int) { tmpDeck[i], tmpDeck[j] = tmpDeck[j], tmpDeck[i] })

	return tmpDeck
}

//deal 7 cards to each player
func dealCards(numberOfPlayers int, deck []Card) (error, [][]Card){
	var playersCards [][]Card

	if numberOfPlayers < 2 || numberOfPlayers > 10 {
		err := errors.New("inappropriate number of players")
		return err, playersCards
	}

	shuffledDeck := ShuffleDeck(deck)

	//number of cards for
	for i := 0; i < 7; i++{
		//number of players for
		for j := 0; j < numberOfPlayers; j++{
			//assign card to given player
			shuffledDeck, playersCards = assignCardToPlayer(j, playersCards, shuffledDeck)
		}
	}

	return nil, playersCards
}

//get first card from deck, then return it and deck without this card
func GiveCardFromDeck(deck []Card, position *int) ([]Card, Card){
	var retCard Card
	if position == nil {
		retCard = deck[0]
		deck = append(deck[:0], deck[1:]...)
	} else {
		retCard = deck[*position]
		deck = append(deck[:*position], deck[*position+1:]...)
	}

	return deck, retCard
}

//assign card to given player
func assignCardToPlayer(playerNumber int, playersCards [][]Card, deck []Card) ([]Card, [][]Card){
	var (
		tmpCard Card
		tmpDeck []Card
	)

	deck, tmpCard = GiveCardFromDeck(deck, nil)
	if len(playersCards) < playerNumber+1{
		tmpDeck	= append(tmpDeck, tmpCard)
		playersCards = append(playersCards, tmpDeck)
	} else {
		playersCards[playerNumber] = append(playersCards[playerNumber], tmpCard)
	}



	return deck, playersCards
}

func printHand(cards [][]Card, player int){
	fmt.Println("Player number ", player, " cards are: ")
	fmt.Println(cards[player])
}

func SomeoneWon(playersCards [][]Card) bool {
	winner := false
	for key, playerHand := range playersCards{
		if len(playerHand) < 1{
			winner = true
			fmt.Println("Won player ", key, "!")
		}
	}
	return winner
}