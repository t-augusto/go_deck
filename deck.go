package main

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"strings"
)

var globalDecks []Deck

type Card struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Code  string `json:"code"`
}

type Deck struct {
	Cards    []Card `json:"cards"`
	Id       string `json:"id"`
	Shuffled bool   `json:"shuffled"`
}

func CreateDeck(shuffle bool, partial []string) (deck Deck) {
	fmt.Printf("creating new deck \n")

	//Default cards
	types := []string{"ACE", "TWO", "THREE", "FOUR", "FIVE", "SIX", "SEVEN",
		"EIGHT", "NINE", "TEN", "JACK", "QUEEN", "KING"}

	suits := []string{"SPADES", "DIAMONDS", "CLUBS", "HEARTS"}

	// Non shuffled deck
	for i := 0; i < len(types); i++ {
		for n := 0; n < len(suits); n++ {
			card := Card{
				Value: types[i],
				Suit:  suits[n],
				Code:  types[i][0:1] + suits[n][0:1],
			}
			deck.Cards = append(deck.Cards, card)
		}
	}
	deck.Id = uuid.NewString()
	deck.Shuffled = false
	if len(partial) > 0 {
		//args from query are in the first index, following the challenges parameters
		// partial[0] would return something like so: "AS,AD"
		deck.Cards = DrawSpecific(deck, partial[0])
	}
	if shuffle {
		deck = Shuffle(deck)
	}
	globalDecks = append(globalDecks, deck)
	return
}

func Shuffle(d Deck) Deck {
	d.Shuffled = true
	for i := 1; i < len(d.Cards); i++ {
		r := rand.Intn(i + 1)
		if i != r {
			d.Cards[r], d.Cards[i] = d.Cards[i], d.Cards[r]
		}
	}
	return d
}

func Draw(d Deck, amount int) []Card {
	for i, deck := range globalDecks {
		if d.Id == deck.Id {
			globalDecks[i].Cards = deck.Cards[amount:]
		}
	}
	return d.Cards[:amount]
}

func DrawSpecific(baseDeck Deck, codes string) (withdraw []Card) {
	splitCodes := strings.Split(codes, ",")
	codeMap := make(map[string]int)
	for i, card := range baseDeck.Cards {
		codeMap[card.Code] = i
	}
	for _, code := range splitCodes {
		if val, ok := codeMap[code]; ok {
			withdraw = append(withdraw, baseDeck.Cards[val])
		}
	}
	return
}

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

func DeckExists(id string) (deck Deck, err error) {
	err = errors.New("Invalid Deck UUID")
	if IsValidUUID(id) {
		for _, globalDeck := range globalDecks {
			if id == globalDeck.Id {
				return globalDeck, nil
			}
		}
	}
	return deck, err

}
