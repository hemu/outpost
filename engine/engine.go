package engine

import (
	// mCard "github.com/hmuar/dominion-replay/card"
	mGame "github.com/hmuar/dominion-replay/game"
)

// turn indexed array of game states
// stores gamestate at beginning of every turn
// turn 0 is start of game
type gameEngine struct {
	game mGame.Game
	// states []mGameState.GameState
}

// factory func
func NewGameEngine() gameEngine {
	e := gameEngine{}
	return e
}

func (e *gameEngine) FeedGame(g mGame.Game) {

}
