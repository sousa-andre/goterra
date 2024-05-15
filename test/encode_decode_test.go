package test

import (
	"github.com/sousa-andre/goterra"
	"testing"
)

func TestEncodeDecode(t *testing.T) {
	blocks := make(chan testDeck)
	go getDeckCodesTestData(blocks)

	for deck := range blocks {
		td := goterra.NewDeck()

		for _, cg := range deck.cardGroups {
			td.AddCard(cg.Card.GetCardCode(), cg.Count)
		}
		if !cardsMatch(td.CardGroups, deck.cardGroups) {
			t.Errorf("Deck codes %s and %s don't match", td.Encode(), deck.deckCode)
		}
	}
}
