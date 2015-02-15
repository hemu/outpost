package event

import (
	mCard "github.com/hmuar/dominion-replay/card"
	mEvent "github.com/hmuar/dominion-replay/event"
	"testing"
)

// are new turns correctly created?
func TestTurnChange(t *testing.T) {
	hb := NewHistoryBuilder()
	if hb.getCurTurn().num != 0 {
		t.Error("Expected initial turn num to be 0")
	}
	hb.startNewTurn(1)
	if hb.getCurTurn().num != 1 {
		t.Errorf("Expected cur turn to be 1 but got %d", hb.getCurTurn().num)
	}
	if len(hb.History.Turns) != 1 {
		t.Errorf("Expected num History.Turns stored to be 1 but got %d",
			len(hb.History.Turns))
	}
	hb.startNewTurn(2)
	if hb.getCurTurn().num != 2 {
		t.Errorf("Expected cur turn to be 2 but got %d", hb.getCurTurn().num)
	}
	if len(hb.History.Turns) != 2 {
		t.Errorf("Expected num History.Turns stored to be 2 but got %d",
			len(hb.History.Turns))
	}
	hb.startNewTurn(3)
	if len(hb.History.Turns) != 3 {
		t.Errorf("Expected num History.Turns stored to be 3 but got %d",
			len(hb.History.Turns))
	}
}

// are new player turns correctly created?
func TestNewPlayerTurn(t *testing.T) {
	hb := NewHistoryBuilder()
	if hb.getCurTurn().num != 0 {
		t.Error("Expected no current turn")
	}
	hb.startNewPlayerTurn("homer", 1)
	if hb.getCurTurn().num != 1 {
		t.Errorf("Expected cur turn to be 1 but got %d", hb.getCurTurn().num)
	}
	if hb.getCurPlayerTurn().player != "homer" {
		t.Error("Expected cur player to be homer")
	}
	hb.startNewPlayerTurn("apu", 1)
	if hb.getCurTurn().num != 1 {
		t.Errorf("Expected cur turn to be 1 but got %d", hb.getCurTurn().num)
	}
	if hb.getCurPlayerTurn().player != "apu" {
		t.Error("Expected cur player to be apu")
	}

	turn1 := hb.getCurTurn()

	hb.startNewPlayerTurn("homer", 2)
	if hb.getCurTurn().num != 2 {
		t.Errorf("Expected cur turn to be 2 but got %d", hb.getCurTurn().num)
	}
	if hb.getCurPlayerTurn().player != "homer" {
		t.Error("Expected cur player to be homer")
	}
	if hb.getCurTurn().num == turn1.num {
		t.Error("Expected turn change")
	}
}

// are previous player turns correctly saved
// as new player turn is started?
func TestSavePlayerTurn(t *testing.T) {
	hb := NewHistoryBuilder()

	hb.StartPlayerTurn("homer", 1)
	hb.StartPlayerTurn("apu", 1)
	if len(hb.getCurTurn().playerTurns) != 2 {
		t.Errorf("Expected 2 stored playerTurns but got %d",
			len(hb.getCurTurn().playerTurns))
	}
	hb.StartPlayerTurn("homer", 2)
	if len(hb.History.Turns) != 2 {
		t.Errorf("Expected 2 stored turns in History.Turns but got %d",
			len(hb.History.Turns))
	}
	if len(hb.History.Turns[0].playerTurns) != 2 {
		t.Errorf("Expected 2 stored playerTurns but got %d",
			len(hb.History.Turns[0].playerTurns))
	}
}

func TestAddEvent(t *testing.T) {
	hb := NewHistoryBuilder()
	hb.StartPlayerTurn("homer", 1)
	drawCards := []mCard.CardSet{mCard.CardSet{Num: 5, Card: mCard.NewCard("Copper")}}
	hb.AddEvent("homer", mEvent.ACTION_DRAW, drawCards)
	if len(hb.getCurPlayerTurn().events) != 1 {
		t.Errorf("Expected 1 event in cur player turn but got %d",
			len(hb.getCurPlayerTurn().events))
	}
	if hb.getCurPlayerTurn().events[0].Player != "homer" {
		t.Errorf("Expected event player homer but got %v",
			hb.getCurPlayerTurn().events[0].Player)
	}
}

func TestAddEventDraw(t *testing.T) {
	hb := NewHistoryBuilder()
	hb.StartPlayerTurn("homer", 1)
	drawCards := []mCard.CardSet{mCard.CardSet{Num: 5, Card: mCard.NewCard("Copper")}}
	hb.AddEvent("homer", mEvent.ACTION_DRAW, drawCards)
	if hb.getCurPlayerTurn().events[0].Action != mEvent.ACTION_DRAW {
		t.Errorf("Expected event action %v but got %v",
			mEvent.ACTION_DRAW,
			hb.getCurPlayerTurn().events[0].Action)
	}
	if hb.getCurPlayerTurn().events[0].Cards[0].Num != 5 {
		t.Errorf("Expected event mCard num 5 but got %d",
			hb.getCurPlayerTurn().events[0].Cards[0].Num)
	}
	if hb.getCurPlayerTurn().events[0].Cards[0].Card.Name != "Copper" {
		t.Errorf("Expected event mCard Copper but got %v",
			hb.getCurPlayerTurn().events[0].Cards[0].Card.Name)
	}
}

func TestAddEventPlay(t *testing.T) {
	hb := NewHistoryBuilder()
	hb.StartPlayerTurn("homer", 1)
	playCards := []mCard.CardSet{mCard.CardSet{Num: 5, Card: mCard.NewCard("Copper")}}
	hb.AddEvent("homer", mEvent.ACTION_PLAY, playCards)
	if hb.getCurPlayerTurn().events[0].Action != mEvent.ACTION_PLAY {
		t.Errorf("Expected event action %v but got %v",
			mEvent.ACTION_PLAY,
			hb.getCurPlayerTurn().events[0].Action)
	}
	if hb.getCurPlayerTurn().events[0].Cards[0].Num != 5 {
		t.Errorf("Expected event mCard num 5 but got %d",
			hb.getCurPlayerTurn().events[0].Cards[0].Num)
	}
	if hb.getCurPlayerTurn().events[0].Cards[0].Card.Name != "Copper" {
		t.Errorf("Expected event mCard Copper but got %v",
			hb.getCurPlayerTurn().events[0].Cards[0].Card.Name)
	}
}

func TestAddEventBuy(t *testing.T) {
	hb := NewHistoryBuilder()
	hb.StartPlayerTurn("homer", 1)
	buyCards := []mCard.CardSet{mCard.CardSet{Num: 1, Card: mCard.NewCard("Gold")}}
	hb.AddEvent("homer", mEvent.ACTION_BUY, buyCards)
	if hb.getCurPlayerTurn().events[0].Action != mEvent.ACTION_BUY {
		t.Errorf("Expected event action %v but got %v",
			mEvent.ACTION_BUY,
			hb.getCurPlayerTurn().events[0].Action)
	}
	if hb.getCurPlayerTurn().events[0].Cards[0].Num != 1 {
		t.Errorf("Expected event mCard num 1 but got %d",
			hb.getCurPlayerTurn().events[0].Cards[0].Num)
	}
	if hb.getCurPlayerTurn().events[0].Cards[0].Card.Name != "Gold" {
		t.Errorf("Expected event mCard Gold but got %v",
			hb.getCurPlayerTurn().events[0].Cards[0].Card.Name)
	}
}

func TestAddEventGain(t *testing.T) {
	hb := NewHistoryBuilder()
	hb.StartPlayerTurn("homer", 1)
	buyCards := []mCard.CardSet{mCard.CardSet{Num: 1, Card: mCard.NewCard("Gold")}}
	hb.AddEvent("homer", mEvent.ACTION_GAIN, buyCards)
	if hb.getCurPlayerTurn().events[0].Action != mEvent.ACTION_GAIN {
		t.Errorf("Expected event action %v but got %v",
			mEvent.ACTION_GAIN,
			hb.getCurPlayerTurn().events[0].Action)
	}
	if hb.getCurPlayerTurn().events[0].Cards[0].Num != 1 {
		t.Errorf("Expected event mCard num 1 but got %d",
			hb.getCurPlayerTurn().events[0].Cards[0].Num)
	}
	if hb.getCurPlayerTurn().events[0].Cards[0].Card.Name != "Gold" {
		t.Errorf("Expected event mCard Gold but got %v",
			hb.getCurPlayerTurn().events[0].Cards[0].Card.Name)
	}
}

func TestAddEventDiscard(t *testing.T) {
	hb := NewHistoryBuilder()
	hb.StartPlayerTurn("homer", 1)
	buyCards := []mCard.CardSet{mCard.CardSet{Num: 5, Card: mCard.NewCard("Copper")}}
	hb.AddEvent("homer", mEvent.ACTION_DISCARD, buyCards)
	if hb.getCurPlayerTurn().events[0].Action != mEvent.ACTION_DISCARD {
		t.Errorf("Expected event action %v but got %v",
			mEvent.ACTION_DISCARD,
			hb.getCurPlayerTurn().events[0].Action)
	}
	if hb.getCurPlayerTurn().events[0].Cards[0].Num != 5 {
		t.Errorf("Expected event mCard num 5 but got %d",
			hb.getCurPlayerTurn().events[0].Cards[0].Num)
	}
	if hb.getCurPlayerTurn().events[0].Cards[0].Card.Name != "Copper" {
		t.Errorf("Expected event mCard Copper but got %v",
			hb.getCurPlayerTurn().events[0].Cards[0].Card.Name)
	}
}
