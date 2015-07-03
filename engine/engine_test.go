package engine

import (
	mCard "github.com/hmuar/dominion-replay/card"
	mEvent "github.com/hmuar/dominion-replay/event"
	mState "github.com/hmuar/dominion-replay/state"
	"testing"
)

func TestAddSupply(t *testing.T) {
	state := mState.State{}
	eng := Engine{}
	supp := []mCard.CardSet{}
	supp = append(supp, mCard.CardSet{Num: 10, Card: mCard.NewCard("Chapel")})
	supp = append(supp, mCard.CardSet{Num: 6, Card: mCard.NewCard("Courtyard")})
	supp = append(supp, mCard.CardSet{Num: 3, Card: mCard.NewCard("Haven")})
	// supp = append(supp, mCard.CardSet{Num: 10, Card: mCard.NewCard("FishingVillage")})
	// supp = append(supp, mCard.CardSet{Num: 10, Card: mCard.NewCard("Village")})
	// supp = append(supp, mCard.CardSet{Num: 10, Card: mCard.NewCard("Warehouse")})
	// supp = append(supp, mCard.CardSet{Num: 10, Card: mCard.NewCard("Moneylender")})
	// supp = append(supp, mCard.CardSet{Num: 10, Card: mCard.NewCard("Monument")})
	// supp = append(supp, mCard.CardSet{Num: 10, Card: mCard.NewCard("Navigator")})
	// supp = append(supp, mCard.CardSet{Num: 10, Card: mCard.NewCard("Bank")})
	// supp = append(supp, mCard.CardSet{Num: 10, Card: mCard.NewCard("Copper")})
	// supp = append(supp, mCard.CardSet{Num: 10, Card: mCard.NewCard("Silver")})
	// supp = append(supp, mCard.CardSet{Num: 10, Card: mCard.NewCard("Gold")})
	// supp = append(supp, mCard.CardSet{Num: 10, Card: mCard.NewCard("Estate")})
	// supp = append(supp, mCard.CardSet{Num: 10, Card: mCard.NewCard("Duchy")})
	// supp = append(supp, mCard.CardSet{Num: 10, Card: mCard.NewCard("Province")})
	// supp = append(supp, mCard.CardSet{Num: 10, Card: mCard.NewCard("Curse")})
	if len(state.Board.Supply) != 0 {
		t.Errorf("Expected initial len of supply cards to be 0 but got %d",
			len(state.Board.Supply))
	}

	eng.SetSupply(supp, &state)

	if len(state.Board.Supply) != 3 {
		t.Errorf("Expected initial len of supply cards to be 17 but got %d",
			len(state.Board.Supply))
	}
	chapelSet := state.Board.Supply["Chapel"]
	if (*chapelSet).Card.Name != "Chapel" {
		t.Errorf("Expected Chapel supply card but got %v",
			(*chapelSet).Card.Name)
	}
	if (*chapelSet).Num != 10 {
		t.Errorf("Expected 10 Chapel supply cards but got %d",
			(*chapelSet).Num)
	}
	courtyardSet := state.Board.Supply["Courtyard"]
	if (*courtyardSet).Card.Name != "Courtyard" {
		t.Errorf("Expected Courtyard supply card but got %v",
			(*courtyardSet).Card.Name)
	}
	if (*courtyardSet).Num != 6 {
		t.Errorf("Expected 6 Courtyard supply cards but got %d",
			(*courtyardSet).Num)
	}
	havenSet := state.Board.Supply["Haven"]
	if (*havenSet).Card.Name != "Haven" {
		t.Errorf("Expected Haven supply card but got %v",
			(*havenSet).Card.Name)
	}
	if (*havenSet).Num != 3 {
		t.Errorf("Expected 3 Haven supply cards but got %d",
			(*havenSet).Num)
	}
}

func TestEventDraw(t *testing.T) {
	state := mState.State{}
	state.SetPlayers([]string{"happypuppy", "sadpuppy"})
	eng := Engine{}
	cards := []mCard.Card{}
	cards = append(cards, mCard.NewCards("Copper", 7)...)
	cards = append(cards, mCard.NewCards("Estate", 3)...)
	state.Players["happypuppy"].Draw = cards
	drawCards := []mCard.Card{}
	drawCards = append(drawCards, mCard.NewCards("Copper", 3)...)
	drawCards = append(drawCards, mCard.NewCards("Estate", 2)...)

	ev := mEvent.Event{Player: "happypuppy", Action: mEvent.ACTION_DRAW, Cards: drawCards}
	_ = eng.RegisterEvent(ev, &state)

	if len(state.GetHand("happypuppy")) != 5 {
		t.Errorf("Expected 5 cards in player hand but got %d",
			len(state.GetHand("happypuppy")))
	}

	hand := state.GetHand("happypuppy")
	if len(hand) == 0 {
		t.Errorf("No hand found for player happypuppy")
	} else {
		for i := 0; i < 3; i++ {
			if hand[i].Name != cards[i].Name {
				t.Errorf("Expected %v as %dth hand card but got %v",
					cards[i].Name, i, hand[i].Name)
			}
		}
	}

	playerDraw := state.GetDraw("happypuppy")
	if len(playerDraw) != 5 {
		t.Errorf("Expected 5 cards in player hand but got %d",
			len(playerDraw))
	}
	numDrawCopper := 0
	numDrawEstate := 0
	for _, card := range playerDraw {
		if card.Name == "Copper" {
			numDrawCopper += 1
		} else if card.Name == "Estate" {
			numDrawEstate += 1
		}
	}
	if numDrawCopper != 4 {
		t.Errorf("Expected 4 Coppers in player draw but got %d", numDrawCopper)
	}
	if numDrawEstate != 1 {
		t.Errorf("Expected 1 Estate in player draw but got %d", numDrawEstate)
	}

}

func TestMultiPlayerDraw(t *testing.T) {
	state := mState.State{}
	state.SetPlayers([]string{"happypuppy", "sadpuppy"})
	eng := Engine{}
	cards := mCard.NewCards("Copper", 7)
	cards = append(cards, mCard.NewCards("Estate", 3)...)
	state.SetDraw("happypuppy", cards)

	cards2 := mCard.NewCards("Copper", 7)
	cards2 = append(cards2, mCard.NewCards("Estate", 3)...)
	state.SetDraw("sadpuppy", cards2)

	numC := 0
	numE := 0
	for _, card := range state.GetDraw("happypuppy") {
		if card.Name == "Copper" {
			numC += 1
		} else if card.Name == "Estate" {
			numE += 1
		}
	}
	if numC != 7 {
		t.Errorf("Expected 7 Coppers in player draw but got %d", numC)
	}
	if numE != 3 {
		t.Errorf("Expected 3 Estate in player draw but got %d", numE)
	}

	numC = 0
	numE = 0
	for _, card := range state.GetDraw("sadpuppy") {
		if card.Name == "Copper" {
			numC += 1
		} else if card.Name == "Estate" {
			numE += 1
		}
	}
	if numC != 7 {
		t.Errorf("Expected 7 Coppers in player draw but got %d", numC)
	}
	if numE != 3 {
		t.Errorf("Expected 3 Estate in player draw but got %d", numE)
	}

	drawCards := mCard.NewCards("Copper", 3)
	drawCards = append(drawCards, mCard.NewCards("Estate", 2)...)

	ev := mEvent.Event{Player: "happypuppy", Action: mEvent.ACTION_DRAW, Cards: drawCards}
	_ = eng.RegisterEvent(ev, &state)

	numC = 0
	numE = 0
	for _, card := range state.GetDraw("sadpuppy") {
		if card.Name == "Copper" {
			numC += 1
		} else if card.Name == "Estate" {
			numE += 1
		}
	}
	if numC != 7 {
		t.Errorf("Expected 7 Coppers in player draw but got %d", numC)
	}
	if numE != 3 {
		t.Errorf("Expected 3 Estate in player draw but got %d", numE)
	}

	if len(state.GetHand("happypuppy")) != 5 {
		t.Errorf("Expected 5 cards in player hand but got %d",
			len(state.GetHand("happypuppy")))
	}

	hand := state.GetHand("happypuppy")
	if len(hand) == 0 {
		t.Errorf("No hand found for player happypuppy")
	} else {
		for i := 0; i < 3; i++ {
			if hand[i].Name != cards[i].Name {
				t.Errorf("Expected %v as %dth hand card but got %v",
					cards[i].Name, i, hand[i].Name)
			}
		}
	}

	playCards := mCard.NewCards("Copper", 3)
	ev = mEvent.Event{Player: "happypuppy", Action: mEvent.ACTION_PLAY, Cards: playCards}
	_ = eng.RegisterEvent(ev, &state)

	numC = 0
	numE = 0
	for _, card := range state.GetDraw("sadpuppy") {
		if card.Name == "Copper" {
			numC += 1
		} else if card.Name == "Estate" {
			numE += 1
		}
	}
	if numC != 7 {
		t.Errorf("Expected 7 Coppers in player draw but got %d", numC)
	}
	if numE != 3 {
		t.Errorf("Expected 3 Estate in player draw but got %d", numE)
	}

	ev = mEvent.Event{Player: "happypuppy", Action: mEvent.ACTION_END_TURN,
		Cards: []mCard.Card{}}
	_ = eng.RegisterEvent(ev, &state)

	numC = 0
	numE = 0
	for _, card := range state.GetDraw("sadpuppy") {
		if card.Name == "Copper" {
			numC += 1
		} else if card.Name == "Estate" {
			numE += 1
		}
	}
	if numC != 7 {
		t.Errorf("Expected 7 Coppers in player draw but got %d", numC)
	}
	if numE != 3 {
		t.Errorf("Expected 3 Estate in player draw but got %d", numE)
	}

	drawCards2 := mCard.NewCards("Copper", 2)
	drawCards2 = append(drawCards2, mCard.NewCards("Estate", 3)...)
	ev = mEvent.Event{Player: "sadpuppy", Action: mEvent.ACTION_DRAW, Cards: drawCards2}
	_ = eng.RegisterEvent(ev, &state)
	playCards2 := mCard.NewCards("Copper", 2)
	ev = mEvent.Event{Player: "sadpuppy", Action: mEvent.ACTION_PLAY, Cards: playCards2}
	_ = eng.RegisterEvent(ev, &state)

	playerDraw := state.GetDraw("happypuppy")
	if len(playerDraw) != 5 {
		t.Errorf("Expected 5 cards in player hand but got %d",
			len(playerDraw))
	}
	numDrawCopper := 0
	numDrawEstate := 0
	for _, card := range playerDraw {
		if card.Name == "Copper" {
			numDrawCopper += 1
		} else if card.Name == "Estate" {
			numDrawEstate += 1
		}
	}
	if numDrawCopper != 4 {
		t.Errorf("Expected 4 Coppers in player draw but got %d", numDrawCopper)
	}
	if numDrawEstate != 1 {
		t.Errorf("Expected 1 Estate in player draw but got %d", numDrawEstate)
	}

}

func TestEventPlay(t *testing.T) {
	state := mState.State{}
	state.SetPlayers([]string{"happypuppy", "sadpuppy"})
	eng := Engine{}
	cards := mCard.NewCards("Copper", 3)
	ev := mEvent.Event{Player: "happypuppy", Action: mEvent.ACTION_PLAY, Cards: cards}
	_ = eng.RegisterEvent(ev, &state)

	play := state.GetPlay()
	if len(play) == 0 {
		t.Errorf("No play cards found")
	} else {
		for i := 0; i < len(cards); i++ {
			if play[i].Name != cards[i].Name {
				t.Errorf("Expected %v as %dth play card but got %v",
					cards[i].Name, i, play[i].Name)
			}
		}
	}
}

func TestEventDrawAndPlay(t *testing.T) {
	state := mState.State{}
	state.SetPlayers([]string{"happypuppy", "sadpuppy"})
	state.SetTurnPlayer("happypuppy")
	eng := Engine{}
	cards := []mCard.Card{}
	cards = append(cards, mCard.NewCards("Copper", 3)...)
	cards = append(cards, mCard.NewCards("Estate", 2)...)
	ev := mEvent.Event{Player: "happypuppy", Action: mEvent.ACTION_DRAW, Cards: cards}
	_ = eng.RegisterEvent(ev, &state)

	hand := state.GetHand("happypuppy")
	if len(hand) == 0 {
		t.Errorf("No hand found for player happypuppy")
	} else {
		for i := 0; i < 3; i++ {
			if hand[i].Name != cards[i].Name {
				t.Errorf("Expected %v as %dth hand card but got %v",
					cards[i].Name, i, hand[i].Name)
			}
		}
	}

	playCards := mCard.NewCards("Copper", 3)
	ev = mEvent.Event{Player: "happypuppy", Action: mEvent.ACTION_PLAY, Cards: playCards}
	_ = eng.RegisterEvent(ev, &state)

	play := state.GetPlay()
	if len(play) == 0 {
		t.Errorf("No play cards found")
	} else {
		for i := 0; i < len(playCards); i++ {
			if play[i].Name != playCards[i].Name {
				t.Errorf("Expected %v as %dth play card but got %v",
					playCards[i].Name, i, play[i].Name)
			}
		}
	}

	remainHand := state.GetHand("happypuppy")
	if len(remainHand) != 2 {
		t.Errorf("Expected 2 cards in hand but got %d", len(remainHand))
	} else {
		for i := 0; i < 2; i++ {
			if remainHand[i].Name != "Estate" {
				t.Errorf("Expected %v as %dth hand card but got %v",
					"Estate", i, remainHand[i].Name)
			}
		}
	}
}

func TestAddStats(t *testing.T) {
	state := mState.State{}
	state.SetPlayers([]string{"happypuppy", "sadpuppy"})
	state.SetTurnPlayer("happypuppy")
	eng := Engine{}
	cards := []mCard.Card{}
	cards = append(cards, mCard.NewCards("Copper", 3)...)
	cards = append(cards, mCard.NewCards("Village", 2)...)
	ev := mEvent.Event{Player: "happypuppy", Action: mEvent.ACTION_DRAW, Cards: cards}
	_ = eng.RegisterEvent(ev, &state)

	hand := state.GetHand("happypuppy")
	if len(hand) == 0 {
		t.Errorf("No hand found for player happypuppy")
	} else {
		for i := 0; i < 3; i++ {
			if hand[i].Name != cards[i].Name {
				t.Errorf("Expected %v as %dth hand card but got %v",
					cards[i].Name, i, hand[i].Name)
			}
		}
	}

	origStats := state.GetPlayerStats("happypuppy")

	// test playing an action card
	playCards := mCard.NewCards("Village", 1)
	ev = mEvent.Event{Player: "happypuppy", Action: mEvent.ACTION_PLAY, Cards: playCards}
	_ = eng.RegisterEvent(ev, &state)

	afterVillageStats := state.GetPlayerStats("happypuppy")

	if afterVillageStats[0] != origStats[0]+1 {
		t.Errorf("Expected extra action for num player actions")
	}
	if afterVillageStats[1] != origStats[1] {
		t.Errorf("Expected no change in num player buys")
	}
	if afterVillageStats[2] != origStats[2] {
		t.Errorf("Expected no change in num player coin")
	}
	if afterVillageStats[3] != origStats[3] {
		t.Errorf("Expected no change in num player victory")
	}

	// test playing coin cards
	playCards = mCard.NewCards("Copper", 3)
	ev = mEvent.Event{Player: "happypuppy", Action: mEvent.ACTION_PLAY, Cards: playCards}
	_ = eng.RegisterEvent(ev, &state)

	afterCopperStats := state.GetPlayerStats("happypuppy")

	if afterCopperStats[0] != afterVillageStats[0] {
		t.Errorf("Expected no change in num player actions")
	}
	if afterCopperStats[1] != afterVillageStats[1] {
		t.Errorf("Expected no change in num player buys")
	}
	if afterCopperStats[2] != afterVillageStats[2]+3 {
		t.Errorf("Expected a change of 3 coins in num player coin")
	}
	if afterCopperStats[3] != afterVillageStats[3] {
		t.Errorf("Expected no change in num player victory")
	}
}

func TestEventBuy(t *testing.T) {
	state := mState.State{}
	state.SetPlayers([]string{"happypuppy", "sadpuppy"})
	state.SetTurnPlayer("happypuppy")

	supp := []mCard.CardSet{}
	supp = append(supp, mCard.CardSet{Num: 10, Card: mCard.NewCard("Chapel")})
	supp = append(supp, mCard.CardSet{Num: 10, Card: mCard.NewCard("Courtyard")})

	eng := Engine{}
	eng.SetSupply(supp, &state)

	cards := []mCard.Card{}
	cards = append(cards, mCard.NewCards("Copper", 3)...)
	cards = append(cards, mCard.NewCards("Village", 2)...)

	state.SetDiscard("happypuppy", cards)

	if state.Board.Supply["Chapel"].Num != 10 {
		t.Errorf("Expected 10 chapel supply cards but got %d",
			state.Board.Supply["Chapel"].Num)
	}

	buyCards := mCard.NewCards("Chapel", 1)

	ev := mEvent.Event{Player: "happypuppy", Action: mEvent.ACTION_BUY, Cards: buyCards}
	_ = eng.RegisterEvent(ev, &state)

	if state.Board.Supply["Chapel"].Num != 9 {
		t.Errorf("Expected 9 chapel supply cards but got %d",
			state.Board.Supply["Chapel"].Num)
	}

}

func TestEventGain(t *testing.T) {
	state := mState.State{}
	state.SetPlayers([]string{"happypuppy", "sadpuppy"})
	state.SetTurnPlayer("happypuppy")
	eng := Engine{}

	cards := []mCard.Card{}
	cards = append(cards, mCard.NewCards("Copper", 3)...)
	cards = append(cards, mCard.NewCards("Village", 2)...)

	state.SetDiscard("happypuppy", cards)

	discard := state.GetDiscard("happypuppy")

	if len(discard) != 5 {
		t.Errorf("Expected 5 discard cards but got %d", len(discard))
	}

	gainCards := mCard.NewCards("Chapel", 1)

	ev := mEvent.Event{Player: "happypuppy", Action: mEvent.ACTION_GAIN, Cards: gainCards}
	_ = eng.RegisterEvent(ev, &state)

	discard = state.GetDiscard("happypuppy")

	if len(discard) != 6 {
		t.Errorf("Expected 6 discard cards but got %d", len(discard))
	}

	//TODO: check that card is on top of discard
	topCard := discard[len(discard)-1]
	if topCard.Name != "Chapel" {
		t.Errorf("Expected top discard card to be Chapel but got %v", topCard.Name)
	}

}

func TestEventDiscard(t *testing.T) {
	state := mState.State{}
	state.SetPlayers([]string{"happypuppy", "sadpuppy"})
	state.SetTurnPlayer("happypuppy")
	eng := Engine{}

	cards := []mCard.Card{}
	cards = append(cards, mCard.NewCards("Copper", 3)...)
	cards = append(cards, mCard.NewCards("Village", 2)...)

	state.SetHand("happypuppy", cards)

	discard := state.GetDiscard("happypuppy")

	if len(discard) != 0 {
		t.Errorf("Expected 0 discard cards but got %d", len(discard))
	}

	discardCards := mCard.NewCards("Village", 1)

	ev := mEvent.Event{Player: "happypuppy", Action: mEvent.ACTION_DISCARD, Cards: discardCards}
	_ = eng.RegisterEvent(ev, &state)

	discard = state.GetDiscard("happypuppy")

	if len(discard) != 1 {
		t.Errorf("Expected 1 discard card but got %d", len(discard))
	}

	//TODO: check that card is on top of discard
	topCard := discard[len(discard)-1]
	if topCard.Name != "Village" {
		t.Errorf("Expected top discard card to be Village but got %v", topCard.Name)
	}

}

// add board play and player hand to discard
// clear board play
// clear player hand
// player action = 0
// player buy = 0
// player coin = 0
func TestEventEndTurnThenStart(t *testing.T) {
	state := mState.State{}
	state.SetPlayers([]string{"happypuppy", "sadpuppy"})
	state.SetTurnPlayer("happypuppy")
	eng := Engine{}
	cards := []mCard.Card{}
	cards = append(cards, mCard.NewCards("Copper", 3)...)
	cards = append(cards, mCard.NewCards("Estate", 2)...)
	ev := mEvent.Event{Player: "happypuppy", Action: mEvent.ACTION_DRAW, Cards: cards}
	_ = eng.RegisterEvent(ev, &state)

	hand := state.GetHand("happypuppy")
	if len(hand) == 0 {
		t.Errorf("No hand found for player happypuppy")
	} else {
		for i := 0; i < 3; i++ {
			if hand[i].Name != cards[i].Name {
				t.Errorf("Expected %v as %dth hand card but got %v",
					cards[i].Name, i, hand[i].Name)
			}
		}
	}

	playCards := mCard.NewCards("Copper", 3)
	ev = mEvent.Event{Player: "happypuppy", Action: mEvent.ACTION_PLAY, Cards: playCards}
	_ = eng.RegisterEvent(ev, &state)

	play := state.GetPlay()
	if len(play) == 0 {
		t.Errorf("No play cards found")
	} else {
		for i := 0; i < len(playCards); i++ {
			if play[i].Name != playCards[i].Name {
				t.Errorf("Expected %v as %dth play card but got %v",
					playCards[i].Name, i, play[i].Name)
			}
		}
	}

	remainHand := state.GetHand("happypuppy")
	if len(remainHand) != 2 {
		t.Errorf("Expected 2 cards in hand but got %d", len(remainHand))
	} else {
		for i := 0; i < 2; i++ {
			if remainHand[i].Name != "Estate" {
				t.Errorf("Expected %v as %dth hand card but got %v",
					"Estate", i, remainHand[i].Name)
			}
		}
	}

	ev = mEvent.Event{Player: "happypuppy", Action: mEvent.ACTION_END_TURN, Cards: []mCard.Card{}}
	_ = eng.RegisterEvent(ev, &state)

	endHand := state.GetHand("happypuppy")
	endPlay := state.GetPlay()
	endStats := state.GetPlayerStats("happypuppy")

	if len(endHand) != 0 {
		t.Errorf("Expected 0 cards in hand but got %d", len(endHand))
	}
	if len(endPlay) != 0 {
		t.Errorf("Expected 0 cards in play but got %d", len(endPlay))
	}
	if endStats[0] != 0 {
		t.Errorf("Expected 0 actions for player happypuppy but got %d", endStats[0])
	}
	if endStats[1] != 0 {
		t.Errorf("Expected 0 buys for player happypuppy but got %d", endStats[1])
	}
	if endStats[2] != 0 {
		t.Errorf("Expected 0 coins for player happypuppy but got %d", endStats[2])
	}

	nextPlayerStats := state.GetPlayerStats("sadpuppy")

	if nextPlayerStats[0] != 0 {
		t.Errorf("Expected 0 actions for player sadpuppy but got %d", nextPlayerStats[0])
	}
	if nextPlayerStats[1] != 0 {
		t.Errorf("Expected 0 buys for player sadpuppy but got %d", nextPlayerStats[1])
	}
	if nextPlayerStats[2] != 0 {
		t.Errorf("Expected 0 coins for player sadpuppy but got %d", nextPlayerStats[2])
	}

	eng.RegisterPlayerTurnStart("sadpuppy", &state)

	drawCards2 := []mCard.Card{}
	drawCards2 = append(drawCards2, mCard.NewCards("Copper", 2)...)
	drawCards2 = append(drawCards2, mCard.NewCards("Estate", 3)...)
	ev = mEvent.Event{Player: "sadpuppy", Action: mEvent.ACTION_DRAW, Cards: drawCards2}
	_ = eng.RegisterEvent(ev, &state)

	nextPlayerStats = state.GetPlayerStats("sadpuppy")

	if nextPlayerStats[0] != 1 {
		t.Errorf("Expected 1 action for player sadpuppy but got %d", nextPlayerStats[0])
	}
	if nextPlayerStats[1] != 1 {
		t.Errorf("Expected 1 buy for player sadpuppy but got %d", nextPlayerStats[1])
	}
	if nextPlayerStats[2] != 0 {
		t.Errorf("Expected 0 coins for player sadpuppy but got %d", nextPlayerStats[2])
	}

}
