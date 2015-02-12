package logparse

import (
	mCard "github.com/hmuar/dominion-replay/card"
	mHistory "github.com/hmuar/dominion-replay/history"
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
	history := ParseLog("test/testlogs/testlog_gamesetup.txt")
	gameSupplyCards := history.Supply
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

func TestTurn0(t *testing.T) {
	history := ParseLog("test/testlogs/testlog_turn0.txt")
	if history.Turns[0].GetNumPlayerTurns() != 2 {
		t.Errorf("Expected 2 player turns for turn 0 but got %d",
			history.Turns[0].GetNumPlayerTurns())
	}
	if len(history.Turns[0].GetPlayerEvents(0)) != 1 {
		t.Errorf("Expected 1 player event in turn 1 but got %d",
			len(history.Turns[0].GetPlayerEvents(0)))
	}
	if history.Turns[0].GetPlayerEvents(0)[0].Action != mHistory.ACTION_DRAW {
		t.Errorf("Expected first player action draw but got %v",
			history.Turns[0].GetPlayerEvents(0)[0].Action)
	}
	if history.Turns[0].GetPlayerEvents(1)[0].Action != mHistory.ACTION_DRAW {
		t.Errorf("Expected sec player action draw but got %v",
			history.Turns[0].GetPlayerEvents(0)[0].Action)
	}
}

func TestTurn(t *testing.T) {
	history := ParseLog("test/testlogs/testlog_turns.txt")
	if len(history.Turns) != 6 {
		t.Errorf("Expected 6 turns but got %d", len(history.Turns))
	}
	if history.Turns[1].GetNumPlayerTurns() != 2 {
		t.Errorf("Expected 2 player turns but got %d",
			history.Turns[1].GetNumPlayerTurns())
	}
	if history.Turns[5].GetNumPlayerTurns() != 1 {
		t.Errorf("Expected 1 player turn for turn 5 but got %d",
			history.Turns[5].GetNumPlayerTurns())
	}
}

func TestEvent(t *testing.T) {
	history := ParseLog("test/testlogs/testlog_turns.txt")
	firstTurnEvents := history.Turns[1].GetPlayerEvents(0)
	if firstTurnEvents[0].Player != "stanleygoodspeed" {
		t.Errorf("Expected first event player to be stanleygoodspeed but got %v",
			firstTurnEvents[0].Player)
	}
	if firstTurnEvents[0].Action != mHistory.ACTION_PLAY {
		t.Errorf("Expected first event action to be %v but got %v",
			mHistory.ACTION_PLAY, firstTurnEvents[0].Action)
	}
	if firstTurnEvents[0].Cards[0].Num != 3 {
		t.Errorf("Expected first event cards num to be 3 but got %d",
			mHistory.ACTION_PLAY, firstTurnEvents[0].Cards[0].Num)
	}
	if firstTurnEvents[0].Cards[0].Card.Name != "Copper" {
		t.Errorf("Expected first event cards to be Copper but got %v",
			mHistory.ACTION_PLAY, firstTurnEvents[0].Cards[0].Card.Name)
	}
}

func TestHandlePlaceOnDeckCore(t *testing.T) {
	text := "stanleygoodspeed - places Estate on top of deck"
	player, cardSets := handlePlaceOnDeck(text)
	if player != "stanleygoodspeed" {
		t.Errorf("Expected player stanleygoodspeed but got %v", player)
	}
	if len(cardSets) != 1 {
		t.Errorf("Expected 1 cardset but got %d", len(cardSets))
	}
	cardSet := cardSets[0]
	if cardSet.Num != 1 {
		t.Errorf("Expected 1 card but got %d", cardSet.Num)
	}
	if cardSet.Card.Name != "Estate" {
		t.Errorf("Expected Estate but got %v", cardSet.Card.Name)
	}
}

func TestLookAtCore(t *testing.T) {
	text := "stanleygoodspeed - looks at Estate, Copper, Warehouse, Copper, Estate"
	player, cardSets := handleLookAt(text)
	if player != "stanleygoodspeed" {
		t.Errorf("Expected player stanleygoodspeed but got %v", player)
	}
	if len(cardSets) != 5 {
		t.Errorf("Expected 5 cardsets but got %d", len(cardSets))
	}
	for _, cardSet := range cardSets {
		if cardSet.Num != 1 {
			t.Errorf("Expected 1 card but got %d", cardSet.Num)
		}
	}
	if len(cardSets) == 5 {
		if cardSets[0].Card.Name != "Estate" {
			t.Errorf("Expected Estate but got %v", cardSets[0].Card.Name)
		}
		if cardSets[1].Card.Name != "Copper" {
			t.Errorf("Expected Copper but got %v", cardSets[1].Card.Name)
		}
		if cardSets[2].Card.Name != "Warehouse" {
			t.Errorf("Expected Warehouse but got %v", cardSets[2].Card.Name)
		}
		if cardSets[3].Card.Name != "Copper" {
			t.Errorf("Expected Copper but got %v", cardSets[3].Card.Name)
		}
		if cardSets[4].Card.Name != "Estate" {
			t.Errorf("Expected Estate but got %v", cardSets[4].Card.Name)
		}
	}
}

func TestFullTurnGame(t *testing.T) {
	// history := ParseLog("test/testlogs/testlog_turns.txt")
	// history.PrintGame()
}

// func TestDrawAction(t *testing.T) {
// 	log.Println("-- TestDrawAction --")
// 	history := ParseLog("test/testlogs/testlog_turns.txt")
// 	if len(history.Turns) != 2 {
// 		t.Errorf("Expected 2 turns but got %d", len(history.Turns))
// 	}
// }

// func TestGameSetup(t *testing.T) {
// 	log.Println("-- TestGameSetup --")
// 	// parsedGame := ParseLog("test/testlogs/testlog_gamesetup.txt")
// }
