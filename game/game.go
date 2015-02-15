package game

import (
	mEngine "github.com/hmuar/dominion-replay/engine"
	mHistory "github.com/hmuar/dominion-replay/history"
	mState "github.com/hmuar/dominion-replay/state"
)

// turn-indexed array of game states
// stores gamestate at beginning of every turn
// turn 0 is start of game
type Game struct {
	states []mState.State
}

func NewGameBuilder() gameBuilder {
	g := Game{}
	gb := gameBuilder{game: g}
	return gb
}

// gameBuilder uses an engine and feeds it actions
// and a current state and allows engine to modify
// that state
type gameBuilder struct {
	game Game
	eng  mEngine.Engine
}

func (gb *gameBuilder) FeedHistory(h mHistory.History) {

}

func (gb *gameBuilder) registerSupply(turnNum int) {

}
