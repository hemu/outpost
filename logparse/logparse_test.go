package logparse

import (
	"fmt"
	"github.com/hmuar/dominion-replay/card"
	"testing"
)

func TestParseSupply(t *testing.T) {
	supLine := "Supply cards: Chapel, Courtyard, Haven, " +
		"Fishing Village, Village, Warehouse, " +
		"Moneylender, Monument, Navigator, Bank, " +
		"Copper, Silver, Gold, Estate, Duchy, " +
		"Province, Curse"

	cards := handleSupply(supLine)
	expected := [17]card.CardSet{
		card.CardSet{Num: 1, Card: card.CardFactory["Chapel"]},
		card.CardSet{Num: 1, Card: card.CardFactory["Courtyard"]},
		card.CardSet{Num: 1, Card: card.CardFactory["Haven"]},
		card.CardSet{Num: 1, Card: card.CardFactory["Fishing Village"]},
		card.CardSet{Num: 1, Card: card.CardFactory["Village"]},
		card.CardSet{Num: 1, Card: card.CardFactory["Warehouse"]},
		card.CardSet{Num: 1, Card: card.CardFactory["Moneylender"]},
		card.CardSet{Num: 1, Card: card.CardFactory["Monument"]},
		card.CardSet{Num: 1, Card: card.CardFactory["Navigator"]},
		card.CardSet{Num: 1, Card: card.CardFactory["Bank"]},
		card.CardSet{Num: 1, Card: card.CardFactory["Copper"]},
		card.CardSet{Num: 1, Card: card.CardFactory["Silver"]},
		card.CardSet{Num: 1, Card: card.CardFactory["Gold"]},
		card.CardSet{Num: 1, Card: card.CardFactory["Estate"]},
		card.CardSet{Num: 1, Card: card.CardFactory["Duchy"]},
		card.CardSet{Num: 1, Card: card.CardFactory["Province"]},
		card.CardSet{Num: 1, Card: card.CardFactory["Curse"]},
	}
	if len(cards) != len(expected) {
		t.Error("Got unexpected supply cards")
	}
	for i := range cards {
		if cards[i] != expected[i] {
			t.Error("Expected " + expected[i].Card.Name + " but got " + cards[i].Card.Name)
		}
	}
}

func TestGameSetup(t *testing.T) {
	parsedGame := ParseLog("test/testlogs/testlog_gamesetup.txt")
	fmt.Println(parsedGame)
}
