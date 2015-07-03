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

// Does an end turn event get correctly added?
func TestEndTurnEvent(t *testing.T) {
	hb := NewHistoryBuilder()

	hb.StartPlayerTurn("homer", 1)
	drawCards := mCard.NewCards("Copper", 3)
	drawCards = append(drawCards, mCard.NewCards("Estate", 2)...)
	hb.AddEvent("homer", mEvent.ACTION_DRAW, drawCards)
	playCards := mCard.NewCards("Copper", 3)
	hb.AddEvent("homer", mEvent.ACTION_PLAY, playCards)

	hb.StartPlayerTurn("apu", 1)
	drawCards = append(drawCards, mCard.NewCards("Estate", 2)...)
	hb.AddEvent("apu", mEvent.ACTION_DRAW, drawCards)

	turn1 := hb.getCurTurn()
	playerEvents := turn1.GetPlayerEvents(0)

	// first event should be draw
	firstEv := playerEvents[0]
	if firstEv.Action != mEvent.ACTION_DRAW {
		t.Errorf("Expected first player event to be '%v' but got '%v'",
			mEvent.ACTION_DRAW, firstEv.Action)
	}
	// last event should be player end turn
	lastEv := playerEvents[len(playerEvents)-1]
	if lastEv.Action != mEvent.ACTION_END_TURN {
		t.Errorf("Expected last player event to be '%v' but got '%v'",
			mEvent.ACTION_END_TURN, lastEv.Action)
	}
	if lastEv.Player != "homer" {
		t.Errorf("Expected last event player to be homer but got %v",
			lastEv.Player)
	}
	otherPlayerEvents := turn1.GetPlayerEvents(1)
	// last event should not be player end turn yet
	lastEv = otherPlayerEvents[len(otherPlayerEvents)-1]
	if lastEv.Action == mEvent.ACTION_END_TURN {
		t.Errorf("Did not expect last player event to be '%v' but got '%v'",
			mEvent.ACTION_END_TURN, lastEv.Action)
	}
	if lastEv.Player != "apu" {
		t.Errorf("Expected last event player to be apu but got %v",
			lastEv.Player)
	}

	hb.StartPlayerTurn("homer", 2)
	drawCards = mCard.NewCards("Copper", 3)
	drawCards = append(drawCards, mCard.NewCards("Estate", 2)...)
	hb.AddEvent("homer", mEvent.ACTION_DRAW, drawCards)
	playCards = mCard.NewCards("Copper", 3)
	hb.AddEvent("homer", mEvent.ACTION_PLAY, playCards)

	hb.StartPlayerTurn("apu", 2)
	drawCards = append(drawCards, mCard.NewCards("Estate", 2)...)
	hb.AddEvent("apu", mEvent.ACTION_DRAW, drawCards)

	turn2 := hb.getCurTurn()
	playerEvents = turn2.GetPlayerEvents(0)

	// first event should be draw
	firstEv = playerEvents[0]
	if firstEv.Action != mEvent.ACTION_DRAW {
		t.Errorf("Expected first player event to be '%v' but got '%v'",
			mEvent.ACTION_DRAW, firstEv.Action)
	}
	// last event should be player end turn
	lastEv = playerEvents[len(playerEvents)-1]
	if lastEv.Action != mEvent.ACTION_END_TURN {
		t.Errorf("Expected last player event to be '%v' but got '%v'",
			mEvent.ACTION_END_TURN, lastEv.Action)
	}
	if lastEv.Player != "homer" {
		t.Errorf("Expected last event player to be homer but got %v",
			lastEv.Player)
	}

	otherPlayerEvents = turn2.GetPlayerEvents(1)
	// last event should not be player end turn yet
	lastEv = otherPlayerEvents[len(otherPlayerEvents)-1]
	if lastEv.Action == mEvent.ACTION_END_TURN {
		t.Errorf("Did not expect last player event to be '%v' but got '%v'",
			mEvent.ACTION_END_TURN, lastEv.Action)
	}
	if lastEv.Player != "apu" {
		t.Errorf("Expected last event player to be apu but got %v",
			lastEv.Player)
	}

}

func TestAddEvent(t *testing.T) {
	hb := NewHistoryBuilder()
	hb.StartPlayerTurn("homer", 1)
	drawCards := mCard.NewCards("Copper", 5)
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
	drawCards := mCard.NewCards("Copper", 5)
	hb.AddEvent("homer", mEvent.ACTION_DRAW, drawCards)
	if hb.getCurPlayerTurn().events[0].Action != mEvent.ACTION_DRAW {
		t.Errorf("Expected event action %v but got %v",
			mEvent.ACTION_DRAW,
			hb.getCurPlayerTurn().events[0].Action)
	}
	for i, card := range drawCards {
		if hb.getCurPlayerTurn().events[0].Cards[i].Name != card.Name {
			t.Errorf("Expected event card %v but got %v",
				card.Name,
				hb.getCurPlayerTurn().events[0].Cards[i].Name)
		}
	}
}

func TestAddEventPlay(t *testing.T) {
	hb := NewHistoryBuilder()
	hb.StartPlayerTurn("homer", 1)
	playCards := mCard.NewCards("Copper", 5)
	hb.AddEvent("homer", mEvent.ACTION_PLAY, playCards)
	if hb.getCurPlayerTurn().events[0].Action != mEvent.ACTION_PLAY {
		t.Errorf("Expected event action %v but got %v",
			mEvent.ACTION_PLAY,
			hb.getCurPlayerTurn().events[0].Action)
	}
	for i, card := range playCards {
		if hb.getCurPlayerTurn().events[0].Cards[i].Name != card.Name {
			t.Errorf("Expected event card %v but got %v",
				card.Name,
				hb.getCurPlayerTurn().events[0].Cards[i].Name)
		}
	}
}

func TestAddEventBuy(t *testing.T) {
	hb := NewHistoryBuilder()
	hb.StartPlayerTurn("homer", 1)
	buyCards := []mCard.Card{mCard.NewCard("Gold")}
	hb.AddEvent("homer", mEvent.ACTION_BUY, buyCards)
	if hb.getCurPlayerTurn().events[0].Action != mEvent.ACTION_BUY {
		t.Errorf("Expected event action %v but got %v",
			mEvent.ACTION_BUY,
			hb.getCurPlayerTurn().events[0].Action)
	}
	if hb.getCurPlayerTurn().events[0].Cards[0].Name != "Gold" {
		t.Errorf("Expected event card Gold but got %v",
			hb.getCurPlayerTurn().events[0].Cards[0].Name)
	}
}

func TestAddEventGain(t *testing.T) {
	hb := NewHistoryBuilder()
	hb.StartPlayerTurn("homer", 1)
	buyCards := []mCard.Card{mCard.NewCard("Gold")}
	hb.AddEvent("homer", mEvent.ACTION_GAIN, buyCards)
	if hb.getCurPlayerTurn().events[0].Action != mEvent.ACTION_GAIN {
		t.Errorf("Expected event action %v but got %v",
			mEvent.ACTION_GAIN,
			hb.getCurPlayerTurn().events[0].Action)
	}
	if hb.getCurPlayerTurn().events[0].Cards[0].Name != "Gold" {
		t.Errorf("Expected event card Gold but got %v",
			hb.getCurPlayerTurn().events[0].Cards[0].Name)
	}
}

func TestAddEventDiscard(t *testing.T) {
	hb := NewHistoryBuilder()
	hb.StartPlayerTurn("homer", 1)
	discardCards := []mCard.Card{mCard.NewCard("Copper")}
	hb.AddEvent("homer", mEvent.ACTION_DISCARD, discardCards)
	if hb.getCurPlayerTurn().events[0].Action != mEvent.ACTION_DISCARD {
		t.Errorf("Expected event action %v but got %v",
			mEvent.ACTION_DISCARD,
			hb.getCurPlayerTurn().events[0].Action)
	}
	if hb.getCurPlayerTurn().events[0].Cards[0].Name != "Copper" {
		t.Errorf("Expected event card Copper but got %v",
			hb.getCurPlayerTurn().events[0].Cards[0].Name)
	}
}
