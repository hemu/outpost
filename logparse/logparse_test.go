package logparse

import (
	mCard "github.com/hmuar/dominion-replay/card"
	mGame "github.com/hmuar/dominion-replay/game"
	"testing"
)

func TestParseAction(t *testing.T) {
	text := "stanleygoodspeed - plays 3 Copper"
	player, cardsText, err := parseActionWithCards(text, "plays")
	if err != nil {
		t.Errorf("Parse error")
	}
	if player != "stanleygoodspeed" {
		t.Errorf("Expected player stanleygoodspeed but got '%v'", player)
	}
	if cardsText != "3 Copper" {
		t.Errorf("Expected cardsText 'plays 3 Copper' but got '%v'", cardsText)
	}

	text = "sadpuppyfarm - draws Copper, Copper, Copper, Copper, Estate"
	player, cardsText, err = parseActionWithCards(text, "draws")
	if err != nil {
		t.Errorf("Parse error")
	}
	if player != "sadpuppyfarm" {
		t.Errorf("Expected player sadpuppyfarm but got '%v'", player)
	}
	if cardsText != "Copper, Copper, Copper, Copper, Estate" {
		t.Errorf("Expected cardsText 'Copper, Copper, Copper, Copper, Estate' but got '%v'", cardsText)
	}
}

func TestParseCards(t *testing.T) {
	text := "stanleygoodspeed - draws Copper, Copper, Estate, Copper, Copper"
	_, cardsText, err := parseActionWithCards(text, "draws")
	if err != nil {
		t.Errorf("Parse error")
	}
	cards := parseCards(cardsText)
	if len(cards) != 2 {
		t.Errorf("Expected 2 card sets but got %d", len(cards))
	}
	if cards[0].Num != 4 {
		t.Errorf("Expected 4 cards for first set but got %d", cards[0].Num)
	}
	if cards[0].Card.Name != "Copper" {
		t.Errorf("Expected Copper for first set but got %v", cards[0].Card.Name)
	}
	if cards[1].Num != 1 {
		t.Errorf("Expected 1 cardc for second set but got %d", cards[1].Num)
	}
	if cards[1].Card.Name != "Estate" {
		t.Errorf("Expected Estate for second set but got %v", cards[1].Card.Name)
	}
}

func TestParseSupplyCore(t *testing.T) {
	supLine := "Supply cards: Chapel, Courtyard, Haven, " +
		"Fishing Village, Village, Warehouse, " +
		"Moneylender, Monument, Navigator, Bank, " +
		"Copper, Silver, Gold, Estate, Duchy, " +
		"Province, Curse"

	cards := handleSupply(supLine)
	expected := [17]mCard.CardSet{
		mCard.CardSet{Num: 1, Card: mCard.CardFactory["Chapel"]},
		mCard.CardSet{Num: 1, Card: mCard.CardFactory["Courtyard"]},
		mCard.CardSet{Num: 1, Card: mCard.CardFactory["Haven"]},
		mCard.CardSet{Num: 1, Card: mCard.CardFactory["Fishing Village"]},
		mCard.CardSet{Num: 1, Card: mCard.CardFactory["Village"]},
		mCard.CardSet{Num: 1, Card: mCard.CardFactory["Warehouse"]},
		mCard.CardSet{Num: 1, Card: mCard.CardFactory["Moneylender"]},
		mCard.CardSet{Num: 1, Card: mCard.CardFactory["Monument"]},
		mCard.CardSet{Num: 1, Card: mCard.CardFactory["Navigator"]},
		mCard.CardSet{Num: 1, Card: mCard.CardFactory["Bank"]},
		mCard.CardSet{Num: 1, Card: mCard.CardFactory["Copper"]},
		mCard.CardSet{Num: 1, Card: mCard.CardFactory["Silver"]},
		mCard.CardSet{Num: 1, Card: mCard.CardFactory["Gold"]},
		mCard.CardSet{Num: 1, Card: mCard.CardFactory["Estate"]},
		mCard.CardSet{Num: 1, Card: mCard.CardFactory["Duchy"]},
		mCard.CardSet{Num: 1, Card: mCard.CardFactory["Province"]},
		mCard.CardSet{Num: 1, Card: mCard.CardFactory["Curse"]},
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

func TestParseSupply(t *testing.T) {
	game := ParseLog("test/testlogs/testlog_gamesetup.txt")
	gameSupplyCards := game.Supply
	expected := [17]mCard.CardSet{
		mCard.CardSet{Num: 1, Card: mCard.CardFactory["Chapel"]},
		mCard.CardSet{Num: 1, Card: mCard.CardFactory["Courtyard"]},
		mCard.CardSet{Num: 1, Card: mCard.CardFactory["Haven"]},
		mCard.CardSet{Num: 1, Card: mCard.CardFactory["Fishing Village"]},
		mCard.CardSet{Num: 1, Card: mCard.CardFactory["Village"]},
		mCard.CardSet{Num: 1, Card: mCard.CardFactory["Warehouse"]},
		mCard.CardSet{Num: 1, Card: mCard.CardFactory["Moneylender"]},
		mCard.CardSet{Num: 1, Card: mCard.CardFactory["Monument"]},
		mCard.CardSet{Num: 1, Card: mCard.CardFactory["Navigator"]},
		mCard.CardSet{Num: 1, Card: mCard.CardFactory["Bank"]},
		mCard.CardSet{Num: 1, Card: mCard.CardFactory["Copper"]},
		mCard.CardSet{Num: 1, Card: mCard.CardFactory["Silver"]},
		mCard.CardSet{Num: 1, Card: mCard.CardFactory["Gold"]},
		mCard.CardSet{Num: 1, Card: mCard.CardFactory["Estate"]},
		mCard.CardSet{Num: 1, Card: mCard.CardFactory["Duchy"]},
		mCard.CardSet{Num: 1, Card: mCard.CardFactory["Province"]},
		mCard.CardSet{Num: 1, Card: mCard.CardFactory["Curse"]},
	}
	if len(gameSupplyCards) != len(expected) {
		t.Error("Got unexpected supply mCards")
	}
	for i := range gameSupplyCards {
		if gameSupplyCards[i] != expected[i] {
			t.Error("Expected " + expected[i].Card.Name + " but got " + gameSupplyCards[i].Card.Name)
		}
	}
}

func TestTurnCore(t *testing.T) {
	turnLine := "---------- sadpuppyfarm: turn 3 ----------"
	player, turnNum := handleTurn(turnLine)
	if player != "sadpuppyfarm" {
		t.Error("Expected player saddpuppyfarm but got " + player)
	}
	if turnNum != 3 {
		t.Errorf("Expected turn 3 but got %d", turnNum)
	}
}

func TestTurn(t *testing.T) {
	game := ParseLog("test/testlogs/testlog_turns.txt")
	if len(game.Turns) != 5 {
		t.Errorf("Expected 2 turns but got %d", len(game.Turns))
	}
	if game.Turns[0].GetNumPlayerTurns() != 2 {
		t.Errorf("Expected 2 player turns but got %d",
			game.Turns[0].GetNumPlayerTurns())
	}
	if game.Turns[4].GetNumPlayerTurns() != 1 {
		t.Errorf("Expected 1 player turn for turn 5 but got %d",
			game.Turns[4].GetNumPlayerTurns())
	}
}

func TestEvent(t *testing.T) {
	game := ParseLog("test/testlogs/testlog_turns.txt")
	firstTurnEvents := game.Turns[0].GetPlayerEvents(0)
	if firstTurnEvents[0].Player != "stanleygoodspeed" {
		t.Errorf("Expected first event player to be stanleygoodspeed but got %v",
			firstTurnEvents[0].Player)
	}
	if firstTurnEvents[0].Action != mGame.ACTION_PLAY {
		t.Errorf("Expected first event action to be %v but got %v",
			mGame.ACTION_PLAY, firstTurnEvents[0].Action)
	}
	if firstTurnEvents[0].Cards[0].Num != 3 {
		t.Errorf("Expected first event cards num to be 3 but got %d",
			mGame.ACTION_PLAY, firstTurnEvents[0].Cards[0].Num)
	}
	if firstTurnEvents[0].Cards[0].Card.Name != "Copper" {
		t.Errorf("Expected first event cards to be Copper but got %v",
			mGame.ACTION_PLAY, firstTurnEvents[0].Cards[0].Card.Name)
	}
}

func TestFullTurnGame(t *testing.T) {
	game := ParseLog("test/testlogs/testlog_turns.txt")
	game.PrintGame()
}

// func TestDrawAction(t *testing.T) {
// 	log.Println("-- TestDrawAction --")
// 	game := ParseLog("test/testlogs/testlog_turns.txt")
// 	if len(game.Turns) != 2 {
// 		t.Errorf("Expected 2 turns but got %d", len(game.Turns))
// 	}
// }

// func TestGameSetup(t *testing.T) {
// 	log.Println("-- TestGameSetup --")
// 	// parsedGame := ParseLog("test/testlogs/testlog_gamesetup.txt")
// }
