package game

import (
	mCard "github.com/hmuar/dominion-replay/card"
	"testing"
)

// are new turns correctly created?
func TestTurnChange(t *testing.T) {
	gb := NewGameBuilder()
	if gb.getCurTurn().num != 0 {
		t.Error("Expected initial turn num to be 0")
	}
	gb.startNewTurn(1)
	if gb.getCurTurn().num != 1 {
		t.Errorf("Expected cur turn to be 1 but got %d", gb.getCurTurn().num)
	}
	if len(gb.Game.Turns) != 1 {
		t.Errorf("Expected num Game.Turns stored to be 1 but got %d",
			len(gb.Game.Turns))
	}
	gb.startNewTurn(2)
	if gb.getCurTurn().num != 2 {
		t.Errorf("Expected cur turn to be 2 but got %d", gb.getCurTurn().num)
	}
	if len(gb.Game.Turns) != 2 {
		t.Errorf("Expected num Game.Turns stored to be 2 but got %d",
			len(gb.Game.Turns))
	}
	gb.startNewTurn(3)
	if len(gb.Game.Turns) != 3 {
		t.Errorf("Expected num Game.Turns stored to be 3 but got %d",
			len(gb.Game.Turns))
	}
}

// are new player turns correctly created?
func TestNewPlayerTurn(t *testing.T) {
	gb := NewGameBuilder()
	if gb.getCurTurn().num != 0 {
		t.Error("Expected no current turn")
	}
	gb.startNewPlayerTurn("homer", 1)
	if gb.getCurTurn().num != 1 {
		t.Errorf("Expected cur turn to be 1 but got %d", gb.getCurTurn().num)
	}
	if gb.getCurPlayerTurn().player != "homer" {
		t.Error("Expected cur player to be homer")
	}
	gb.startNewPlayerTurn("apu", 1)
	if gb.getCurTurn().num != 1 {
		t.Errorf("Expected cur turn to be 1 but got %d", gb.getCurTurn().num)
	}
	if gb.getCurPlayerTurn().player != "apu" {
		t.Error("Expected cur player to be apu")
	}

	turn1 := gb.getCurTurn()

	gb.startNewPlayerTurn("homer", 2)
	if gb.getCurTurn().num != 2 {
		t.Errorf("Expected cur turn to be 2 but got %d", gb.getCurTurn().num)
	}
	if gb.getCurPlayerTurn().player != "homer" {
		t.Error("Expected cur player to be homer")
	}
	if gb.getCurTurn().num == turn1.num {
		t.Error("Expected turn change")
	}
}

// are previous player turns correctly saved
// as new player turn is started?
func TestSavePlayerTurn(t *testing.T) {
	gb := NewGameBuilder()

	gb.StartPlayerTurn("homer", 1)
	gb.StartPlayerTurn("apu", 1)
	if len(gb.getCurTurn().playerTurns) != 2 {
		t.Errorf("Expected 2 stored playerTurns but got %d",
			len(gb.getCurTurn().playerTurns))
	}
	gb.StartPlayerTurn("homer", 2)
	if len(gb.Game.Turns) != 2 {
		t.Errorf("Expected 2 stored turns in Game.Turns but got %d",
			len(gb.Game.Turns))
	}
	if len(gb.Game.Turns[0].playerTurns) != 2 {
		t.Errorf("Expected 2 stored playerTurns but got %d",
			len(gb.Game.Turns[0].playerTurns))
	}
}

func TestAddEvent(t *testing.T) {
	gb := NewGameBuilder()
	gb.StartPlayerTurn("homer", 1)
	drawCards := []mCard.CardSet{mCard.CardSet{Num: 5, Card: mCard.CardFactory["Copper"]}}
	gb.AddEvent("homer", ACTION_DRAW, drawCards)
	if len(gb.getCurPlayerTurn().events) != 1 {
		t.Errorf("Expected 1 event in cur player turn but got %d",
			len(gb.getCurPlayerTurn().events))
	}
	if gb.getCurPlayerTurn().events[0].player != "homer" {
		t.Errorf("Expected event player homer but got %v",
			gb.getCurPlayerTurn().events[0].player)
	}
}

func TestAddEventDraw(t *testing.T) {
	gb := NewGameBuilder()
	gb.StartPlayerTurn("homer", 1)
	drawCards := []mCard.CardSet{mCard.CardSet{Num: 5, Card: mCard.CardFactory["Copper"]}}
	gb.AddEvent("homer", ACTION_DRAW, drawCards)
	if gb.getCurPlayerTurn().events[0].action != ACTION_DRAW {
		t.Errorf("Expected event action %v but got %v",
			ACTION_DRAW,
			gb.getCurPlayerTurn().events[0].action)
	}
	if gb.getCurPlayerTurn().events[0].cards[0].Num != 5 {
		t.Errorf("Expected event mCard num 5 but got %d",
			gb.getCurPlayerTurn().events[0].cards[0].Num)
	}
	if gb.getCurPlayerTurn().events[0].cards[0].Card.Name != "Copper" {
		t.Errorf("Expected event mCard Copper but got %v",
			gb.getCurPlayerTurn().events[0].cards[0].Card.Name)
	}
}

func TestAddEventPlay(t *testing.T) {
	gb := NewGameBuilder()
	gb.StartPlayerTurn("homer", 1)
	playCards := []mCard.CardSet{mCard.CardSet{Num: 5, Card: mCard.CardFactory["Copper"]}}
	gb.AddEvent("homer", ACTION_PLAY, playCards)
	if gb.getCurPlayerTurn().events[0].action != ACTION_PLAY {
		t.Errorf("Expected event action %v but got %v",
			ACTION_PLAY,
			gb.getCurPlayerTurn().events[0].action)
	}
	if gb.getCurPlayerTurn().events[0].cards[0].Num != 5 {
		t.Errorf("Expected event mCard num 5 but got %d",
			gb.getCurPlayerTurn().events[0].cards[0].Num)
	}
	if gb.getCurPlayerTurn().events[0].cards[0].Card.Name != "Copper" {
		t.Errorf("Expected event mCard Copper but got %v",
			gb.getCurPlayerTurn().events[0].cards[0].Card.Name)
	}
}

func TestAddEventBuy(t *testing.T) {
	gb := NewGameBuilder()
	gb.StartPlayerTurn("homer", 1)
	buyCards := []mCard.CardSet{mCard.CardSet{Num: 1, Card: mCard.CardFactory["Gold"]}}
	gb.AddEvent("homer", ACTION_BUY, buyCards)
	if gb.getCurPlayerTurn().events[0].action != ACTION_BUY {
		t.Errorf("Expected event action %v but got %v",
			ACTION_BUY,
			gb.getCurPlayerTurn().events[0].action)
	}
	if gb.getCurPlayerTurn().events[0].cards[0].Num != 1 {
		t.Errorf("Expected event mCard num 1 but got %d",
			gb.getCurPlayerTurn().events[0].cards[0].Num)
	}
	if gb.getCurPlayerTurn().events[0].cards[0].Card.Name != "Gold" {
		t.Errorf("Expected event mCard Gold but got %v",
			gb.getCurPlayerTurn().events[0].cards[0].Card.Name)
	}
}

func TestAddEventGain(t *testing.T) {
	gb := NewGameBuilder()
	gb.StartPlayerTurn("homer", 1)
	buyCards := []mCard.CardSet{mCard.CardSet{Num: 1, Card: mCard.CardFactory["Gold"]}}
	gb.AddEvent("homer", ACTION_GAIN, buyCards)
	if gb.getCurPlayerTurn().events[0].action != ACTION_GAIN {
		t.Errorf("Expected event action %v but got %v",
			ACTION_GAIN,
			gb.getCurPlayerTurn().events[0].action)
	}
	if gb.getCurPlayerTurn().events[0].cards[0].Num != 1 {
		t.Errorf("Expected event mCard num 1 but got %d",
			gb.getCurPlayerTurn().events[0].cards[0].Num)
	}
	if gb.getCurPlayerTurn().events[0].cards[0].Card.Name != "Gold" {
		t.Errorf("Expected event mCard Gold but got %v",
			gb.getCurPlayerTurn().events[0].cards[0].Card.Name)
	}
}

func TestAddEventDiscard(t *testing.T) {
	gb := NewGameBuilder()
	gb.StartPlayerTurn("homer", 1)
	buyCards := []mCard.CardSet{mCard.CardSet{Num: 5, Card: mCard.CardFactory["Copper"]}}
	gb.AddEvent("homer", ACTION_DISCARD, buyCards)
	if gb.getCurPlayerTurn().events[0].action != ACTION_DISCARD {
		t.Errorf("Expected event action %v but got %v",
			ACTION_DISCARD,
			gb.getCurPlayerTurn().events[0].action)
	}
	if gb.getCurPlayerTurn().events[0].cards[0].Num != 5 {
		t.Errorf("Expected event mCard num 5 but got %d",
			gb.getCurPlayerTurn().events[0].cards[0].Num)
	}
	if gb.getCurPlayerTurn().events[0].cards[0].Card.Name != "Copper" {
		t.Errorf("Expected event mCard Copper but got %v",
			gb.getCurPlayerTurn().events[0].cards[0].Card.Name)
	}
}
