package state

import (
	"testing"
)

func TestSetPlayer(t *testing.T) {
	state := State{}
	state.SetPlayers([]string{"happypuppy", "sadpuppy"})
	if len(state.Players) != 2 {
		t.Errorf("Expected 2 players but got %d", len(state.Players))
	}
	_, exists := state.getPlayer("happypuppy")
	if !exists {
		t.Error("Player happypuppy does not exist")
	}
	_, exists = state.getPlayer("sadpuppy")
	if !exists {
		t.Error("Player sadpuppy does not exist")
	}

}

func TestSetTurnPlayer(t *testing.T) {
	state := State{}
	state.SetPlayers([]string{"happypuppy", "sadpuppy"})
	state.SetTurnPlayer("sadpuppy")
	if state.TurnPlayer != "sadpuppy" {
		t.Errorf("Expected turn player sadpuppy but got %v",
			state.TurnPlayer)
	}
	_, exists := state.getPlayer(state.TurnPlayer)
	if !exists {
		t.Error("Player sadpuppy does not exist")
	}
}
