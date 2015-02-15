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
	supp = append(supp, mCard.CardSet{Num: 10, Card: mCard.NewCard("Courtyard")})
	supp = append(supp, mCard.CardSet{Num: 10, Card: mCard.NewCard("Haven")})
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
}

func TestRegisterEventDraw(t *testing.T) {
	state := mState.State{}
	state.SetPlayers([]string{"happypuppy", "sadpuppy"})
	eng := Engine{}
	cards := []mCard.CardSet{}
	cards = append(cards, mCard.CardSet{Num: 3, Card: mCard.NewCard("Copper")})
	cards = append(cards, mCard.CardSet{Num: 2, Card: mCard.NewCard("Estate")})
	ev := mEvent.Event{Player: "happypuppy", Action: mEvent.ACTION_DRAW, Cards: cards}
	eng.RegisterEvent(ev, &state)

	if len(state.GetHand("happypuppy")) != 2 {
		t.Errorf("Expected 2 card sets in player hand but got %d",
			len(state.GetHand("happypuppy")))
	}

	fhand := state.GetHand("happypuppy")
	if len(fhand) == 0 {
		t.Errorf("No hand found for player happypuppy")
	} else {
		fc := fhand[0]
		if fc.Card.Name != "Copper" || fc.Num != 3 {
			t.Errorf("Expected 3 Coppers but got %d %v", fc.Num, fc.Card.Name)
		}
	}

	if len(fhand) == 0 {
		t.Errorf("No hand found for player happypuppy")
	} else {
		sc := fhand[1]
		if sc.Card.Name != "Estate" || sc.Num != 2 {
			t.Errorf("Expected 2 Estates but got %d %v", sc.Num, sc.Card.Name)
		}
	}

}
