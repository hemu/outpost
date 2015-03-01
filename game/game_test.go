package game

import (
	// "fmt"
	mParse "github.com/hmuar/dominion-replay/logparse"
	"testing"
)

func TestInitHistory(t *testing.T) {
	history := mParse.ParseLog("../logparse/test/testlogs/testlog_turn0.txt")
	gb := NewGameBuilder()
	gb.registerInitHistory(history)

	// test supply
	if len(gb.state.Board.Supply) != 17 {
		t.Errorf("Expected 17 supply cards but got %d",
			len(gb.state.Board.Supply))
	}
	// test init players
	stanPlayerExists := false
	sadPlayerExists := false
	for player, _ := range gb.state.Players {
		if player == "stanleygoodspeed" && !stanPlayerExists {
			stanPlayerExists = true
		} else if player == "sadpuppyfarm" && !sadPlayerExists {
			sadPlayerExists = true
		} else {
			t.Errorf("Found unexpected player %v", player)
		}
	}
	if !(stanPlayerExists && sadPlayerExists) {
		t.Error("Expected only players stanleygoodspeed and sadpuppyfarm")
	}
	// test init draw pile
	for _, playerName := range [2]string{"stanleygoodspeed", "sadpuppyfarm"} {
		numCopper := 0
		numEstate := 0
		for _, card := range gb.state.Players[playerName].Draw {
			if card.Name == "Copper" {
				numCopper += 1
			} else if card.Name == "Estate" {
				numEstate += 1
			} else {
				t.Errorf("Found unexpected card %v for player %v",
					card.Name, playerName)
			}
		}
		if numCopper != 7 {
			t.Errorf("Expected 7 Copper but got %d for player %v",
				numCopper, playerName)
		}
		if numEstate != 3 {
			t.Errorf("Expected 3 Estates but got %d for player %v",
				numEstate, playerName)
		}
	}
}

func TestHistorySandbox(t *testing.T) {
	history := mParse.ParseLog("../logparse/test/testlogs/testlog_turns.txt")
	gb := NewGameBuilder()
	gb.FeedHistory(history)
}
