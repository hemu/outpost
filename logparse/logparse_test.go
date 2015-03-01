package logparse

import (
	mCard "github.com/hmuar/dominion-replay/card"
	mEvent "github.com/hmuar/dominion-replay/event"
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
	if len(cards) != 5 {
		t.Errorf("Expected 5 card sets but got %d", len(cards))
	}
	if cards[0].Name != "Copper" {
		t.Errorf("Expected Copper as first card but got %v", cards[0].Name)
	}
	if cards[1].Name != "Copper" {
		t.Errorf("Expected Copper as second card but got %v", cards[1].Name)
	}
	if cards[2].Name != "Estate" {
		t.Errorf("Expected Estate as third card but got %v", cards[2].Name)
	}
	if cards[3].Name != "Copper" {
		t.Errorf("Expected Copper as fourth card but got %v", cards[3].Name)
	}
	if cards[4].Name != "Copper" {
		t.Errorf("Expected Copper as fifth card but got %v", cards[4].Name)
	}
}

func TestParseSupplyCore(t *testing.T) {
	supLine := "Supply cards: Chapel, Courtyard, Haven, " +
		"Fishing Village, Village, Warehouse, " +
		"Moneylender, Monument, Navigator, Bank, " +
		"Copper, Silver, Gold, Estate, Duchy, " +
		"Province, Curse"

	cardSets := handleSupply(supLine)
	// eventually need to check correct number of cards
	// expected in supply based on card type
	// e.g. there should be 8 provinces
	expected := [17]mCard.CardSet{
		mCard.CardSet{Num: 10, Card: mCard.NewCard("Chapel")},
		mCard.CardSet{Num: 10, Card: mCard.NewCard("Courtyard")},
		mCard.CardSet{Num: 10, Card: mCard.NewCard("Haven")},
		mCard.CardSet{Num: 10, Card: mCard.NewCard("Fishing Village")},
		mCard.CardSet{Num: 10, Card: mCard.NewCard("Village")},
		mCard.CardSet{Num: 10, Card: mCard.NewCard("Warehouse")},
		mCard.CardSet{Num: 10, Card: mCard.NewCard("Moneylender")},
		mCard.CardSet{Num: 10, Card: mCard.NewCard("Monument")},
		mCard.CardSet{Num: 10, Card: mCard.NewCard("Navigator")},
		mCard.CardSet{Num: 10, Card: mCard.NewCard("Bank")},
		mCard.CardSet{Num: 10, Card: mCard.NewCard("Copper")},
		mCard.CardSet{Num: 10, Card: mCard.NewCard("Silver")},
		mCard.CardSet{Num: 10, Card: mCard.NewCard("Gold")},
		mCard.CardSet{Num: 10, Card: mCard.NewCard("Estate")},
		mCard.CardSet{Num: 10, Card: mCard.NewCard("Duchy")},
		mCard.CardSet{Num: 10, Card: mCard.NewCard("Province")},
		mCard.CardSet{Num: 10, Card: mCard.NewCard("Curse")},
	}
	if len(cardSets) != len(expected) {
		t.Error("Got unexpected supply cardSets")
	}
	for i := range cardSets {
		if cardSets[i].Card.Name != expected[i].Card.Name {
			t.Errorf("Expected %v but got %v", expected[i].Card.Name, cardSets[i].Card.Name)
		}
		if cardSets[i].Num != expected[i].Num {
			t.Errorf("Expected %d but got %d", expected[i].Num, cardSets[i].Num)
		}
	}
}

func TestParseSupply(t *testing.T) {
	history := ParseLog("test/testlogs/testlog_gamesetup.txt")
	gameSupplyCards := history.Supply
	expected := [17]mCard.CardSet{
		mCard.CardSet{Num: 1, Card: mCard.NewCard("Chapel")},
		mCard.CardSet{Num: 1, Card: mCard.NewCard("Courtyard")},
		mCard.CardSet{Num: 1, Card: mCard.NewCard("Haven")},
		mCard.CardSet{Num: 1, Card: mCard.NewCard("Fishing Village")},
		mCard.CardSet{Num: 1, Card: mCard.NewCard("Village")},
		mCard.CardSet{Num: 1, Card: mCard.NewCard("Warehouse")},
		mCard.CardSet{Num: 1, Card: mCard.NewCard("Moneylender")},
		mCard.CardSet{Num: 1, Card: mCard.NewCard("Monument")},
		mCard.CardSet{Num: 1, Card: mCard.NewCard("Navigator")},
		mCard.CardSet{Num: 1, Card: mCard.NewCard("Bank")},
		mCard.CardSet{Num: 1, Card: mCard.NewCard("Copper")},
		mCard.CardSet{Num: 1, Card: mCard.NewCard("Silver")},
		mCard.CardSet{Num: 1, Card: mCard.NewCard("Gold")},
		mCard.CardSet{Num: 1, Card: mCard.NewCard("Estate")},
		mCard.CardSet{Num: 1, Card: mCard.NewCard("Duchy")},
		mCard.CardSet{Num: 1, Card: mCard.NewCard("Province")},
		mCard.CardSet{Num: 1, Card: mCard.NewCard("Curse")},
	}
	if len(gameSupplyCards) != len(expected) {
		t.Error("Got unexpected supply mCards")
	}
	for i := range gameSupplyCards {
		if gameSupplyCards[i].Card.Name != expected[i].Card.Name {
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
	if history.Turns[0].GetPlayerEvents(0)[0].Action != mEvent.ACTION_DRAW {
		t.Errorf("Expected first player action draw but got %v",
			history.Turns[0].GetPlayerEvents(0)[0].Action)
	}
	if history.Turns[0].GetPlayerEvents(1)[0].Action != mEvent.ACTION_DRAW {
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
	if firstTurnEvents[0].Action != mEvent.ACTION_PLAY {
		t.Errorf("Expected first event action to be %v but got %v",
			mEvent.ACTION_PLAY, firstTurnEvents[0].Action)
	}
	for i := 0; i < 3; i++ {
		if firstTurnEvents[0].Cards[i].Name != "Copper" {
			t.Errorf("Expected first event card #%d to be Copper but got %v",
				i, firstTurnEvents[0].Cards[i].Name)
		}
	}
}

func TestHandlePlaceOnDeckCore(t *testing.T) {
	text := "stanleygoodspeed - places Estate on top of deck"
	player, cards := handlePlaceOnDeck(text)
	if player != "stanleygoodspeed" {
		t.Errorf("Expected player stanleygoodspeed but got %v", player)
	}
	if len(cards) != 1 {
		t.Errorf("Expected 1 card but got %d", len(cards))
	}
	card := cards[0]
	if card.Name != "Estate" {
		t.Errorf("Expected Estate but got %v", card.Name)
	}
}

func TestLookAtCore(t *testing.T) {
	text := "stanleygoodspeed - looks at Estate, Copper, Warehouse, Copper, Estate"
	player, cards := handleLookAt(text)
	if player != "stanleygoodspeed" {
		t.Errorf("Expected player stanleygoodspeed but got %v", player)
	}
	if len(cards) != 5 {
		t.Errorf("Expected 5 cardsets but got %d", len(cards))
	}
	if len(cards) == 5 {
		if cards[0].Name != "Estate" {
			t.Errorf("Expected Estate but got %v", cards[0].Name)
		}
		if cards[1].Name != "Copper" {
			t.Errorf("Expected Copper but got %v", cards[1].Name)
		}
		if cards[2].Name != "Warehouse" {
			t.Errorf("Expected Warehouse but got %v", cards[2].Name)
		}
		if cards[3].Name != "Copper" {
			t.Errorf("Expected Copper but got %v", cards[3].Name)
		}
		if cards[4].Name != "Estate" {
			t.Errorf("Expected Estate but got %v", cards[4].Name)
		}
	}
}

func TestPlayerOrder(t *testing.T) {
	history := ParseLog("test/testlogs/testlog_turn0.txt")
	if len(history.PlayerOrder) != 2 {
		t.Errorf("Expected 2 player order entries but got %d",
			len(history.PlayerOrder))
	} else {
		if history.PlayerOrder[0] != "stanleygoodspeed" {
			t.Errorf("Expected first player stanleygoodspeed but got %v",
				history.PlayerOrder[0])
		}
		if history.PlayerOrder[1] != "sadpuppyfarm" {
			t.Errorf("Expected second player sadpuppyfarm but got %v",
				history.PlayerOrder[1])
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
