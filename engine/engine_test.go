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
	cards = append(cards, mCard.NewCards("Copper", 3)...)
	cards = append(cards, mCard.NewCards("Estate", 2)...)

	ev := mEvent.Event{Player: "happypuppy", Action: mEvent.ACTION_DRAW, Cards: cards}
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

func TestEventEndTurn(t *testing.T) {
	state := mState.State{}
	state.SetPlayers([]string{"happypuppy", "sadpuppy"})
	state.SetTurnPlayer("happypuppy")
	eng := Engine{}

	cards := []mCard.Card{}
	cards = append(cards, mCard.NewCards("Copper", 3)...)
	cards = append(cards, mCard.NewCards("Village", 2)...)

	state.SetHand("happypuppy", cards)

}
